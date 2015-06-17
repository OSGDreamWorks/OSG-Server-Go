--加载protobuf模块
local CSPacket_pb = import("CSPacket_pb")
local SCPacket_pb = import("SCPacket_pb")
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

    class.loginServer = Server:new()
    class.loginServer:Register(class)

    class.loginServer:ListenAndServe(cfg.TcpHost, cfg.HttpHost)

end

function GameServer:CS_CheckSession(conn, buf)

    local checkSession = CSPacket_pb.CS_CheckSession()
    checkSession:ParseFromString(buf)

    local checkSessionResult = SCPacket_pb.SC_CheckSessionResult()
    checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.SERVERERROR
    checkSessionResult.server_time = os.time()

    if string.len(checkSession.uid) > 0 then
        local sid, err = self.mainCache:Do("GET", checkSession.uid)
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