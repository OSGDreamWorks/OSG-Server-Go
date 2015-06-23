--加载protobuf模块
local CSPacket_pb = import("CSPacket_pb")
local SCPacket_pb = import("SCPacket_pb")
local SLPacket_pb = import("SLPacket_pb")
local LSPacket_pb = import("LSPacket_pb")

local XShare_Logic_pb = import("XShare_Logic_pb")

local db = import("db")
local logger = import("logger")
local mvc = import("mvc")

local Player = import(".Player")

local GameServer = class("GameServer", mvc.AppBase)

function GameServer:ctor(appName)
    GameServer.super.ctor(self, appName)
end

function GameServer:CreateServices(cfg)

    local class = self.class

    --初始化DB
    db.Init()

    --初始化Cache
    class.mainCache = CachePool:new("etc/maincache.json")

    --
    local loginCfg = common.config.ReadConfig("etc/loginserver.json")
    class.loginServer = RpcClient.new(loginCfg.LoginHost)

    class.gameServer = Server:new()
    class.gameServer:Register(class)

    class.gameServer:ListenAndServe(cfg.TcpHost, cfg.HttpHost)

    local updatePlayerCount = function()
        local rpcCall = SLPacket_pb.SL_UpdatePlayerCount()
        rpcCall.ServerId = 1
        rpcCall.PlayerCount = 0
        rpcCall.TcpServerIp = cfg.TcpHost
        rpcCall.HttpServerIp = cfg.HttpHost
        logger.Debug("updatePlayerCount : %d, %d, %s, %s", rpcCall.ServerId, rpcCall.PlayerCount, rpcCall.TcpServerIp, rpcCall.HttpServerIp)
        if class.loginServer ~= nil then
            local rep = class.loginServer:Call("LoginRpcServer.SL_UpdatePlayerCount", rpcCall:SerializeToString(), "")
            local rpcResult = LSPacket_pb.LS_UpdatePlayerCountResult()
            rpcResult:ParseFromString(rep)
            logger.Debug("server_time %d", rpcResult.server_time)
        end
    end

    common.SetInterval("updatePlayerCount", 5, updatePlayerCount)

end

function GameServer:CS_CheckSession(conn, buf)

    local checkSession = CSPacket_pb.CS_CheckSession()
    checkSession:ParseFromString(buf)

    local checkSessionResult = SCPacket_pb.SC_CheckSessionResult()
    checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.SERVERERROR
    checkSessionResult.server_time = os.time()

    if string.len(checkSession.uid) > 0 then
        local sid, err = self.mainCache:Do("GET", "SessionKey_" .. checkSession.uid)
        if string.len(err) == 0  and sid == checkSession.sessionKey then
            checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.OK
            --登陆成功
            local info_buf, result, err = db.Query("PlayerBaseInfo", checkSession.uid, "")
            if result == false then
                local playerBaseInfo = XShare_Logic_pb.PlayerBaseInfo()
                playerBaseInfo.uid = checkSession.uid
                playerBaseInfo.stat.name = "name"
                playerBaseInfo.stat.level = 1
                info_buf = playerBaseInfo:SerializeToString()
                result, err = db.Write("PlayerBaseInfo",checkSession.uid,info_buf)
            end

            if string.len(err) == 0 and string.len(info_buf) > 0 then

                local playerBaseInfo = XShare_Logic_pb.PlayerBaseInfo()
                playerBaseInfo:ParseFromString(info_buf)
                local player = Player.new({info = playerBaseInfo})
                --self:addPlayer(conn:GetId(), player)

            else
                --查询或创建角色失败
                checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.SERVERERROR
            end

        else
            checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.AUTH_FAILED
        end
    end

    conn:WriteObj("protobuf.SC_CheckSessionResult", checkSessionResult:SerializeToString())

end

function GameServer:CS_Ping(conn, buf)

    local pingResult = SCPacket_pb.SC_PingResult()
    pingResult.server_time = os.time()
    conn:WriteObj("protobuf.SC_PingResult", pingResult:SerializeToString())

end

return GameServer