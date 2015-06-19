--加载protobuf模块
local CLPacket_pb = import("CLPacket_pb")
local LCPacket_pb = import("LCPacket_pb")
local ALPacket_pb = import("ALPacket_pb")
local LAPacket_pb = import("LAPacket_pb")
local SLPacket_pb = import("SLPacket_pb")
local LSPacket_pb = import("LSPacket_pb")

local common = import("common")
local logger = import("logger")
local mvc = import("mvc")

local LoginServerForClient = class("LoginServerForClient")

LoginServerForClient.loginServer = nil

function LoginServerForClient:CL_CheckAccount(conn, buf)

    local checkAccount = CLPacket_pb.CL_CheckAccount()
    checkAccount:ParseFromString(buf)

    local rpcCall = LAPacket_pb.LA_CheckAccount()
    rpcCall.account = checkAccount.account
    rpcCall.password = checkAccount.password

    local rep = self.authServer:Call("AuthServer.LA_CheckAccount", rpcCall:SerializeToString(), "", "LA_CheckAccount", "AL_CheckAccountResult")

    local rpcResult = ALPacket_pb.AL_CheckAccountResult()

    local checkAccountResult = LCPacket_pb.LC_CheckAccountResult()
    checkAccountResult.result = LCPacket_pb.LC_CheckAccountResult.SERVERERROR
    checkAccountResult.server_time = os.time()
    checkAccountResult.sessionKey = ""
    checkAccountResult.uid = ""

    if rep ~= nil then

        rpcResult:ParseFromString(rep)

        checkAccountResult.result = rpcResult.result
        checkAccountResult.server_time = rpcResult.server_time
        checkAccountResult.sessionKey = rpcResult.sessionKey
        checkAccountResult.uid = rpcResult.uid

        if self.loginServer ~= nil then
            if conn:IsWebConn() then
                checkAccountResult.gameServerIp = self.loginServer.stableHttpServer
            else
                checkAccountResult.gameServerIp = self.loginServer.stableTcpServer
            end
        end

    end

    conn:WriteObj("protobuf.LC_CheckAccountResult", checkAccountResult:SerializeToString())

    return 0

end

local LoginServerForGameServer = class("LoginServerForGameServer")

LoginServerForGameServer.loginServer = nil

function LoginServerForGameServer:SL_UpdatePlayerCount(req, ret)

    local upCount = SLPacket_pb.SL_UpdatePlayerCount()
    upCount:ParseFromString(req)

    if self.loginServer ~= nil then
        if string.len(upCount.TcpServerIp) > 0  and string.len(upCount.HttpServerIp) > 0  and upCount.PlayerCount < self.loginServer.uConnId then
            self.loginServer.uConnId = upCount.PlayerCount
            self.loginServer.stableTcpServer = upCount.TcpServerIp
            self.loginServer.stableHttpServer = upCount.HttpServerIp
        end
    end

    local result = LSPacket_pb.LS_UpdatePlayerCountResult()
    result.result = LSPacket_pb.LS_UpdatePlayerCountResult.OK
    result.server_time = os.time()
    local ret = result:SerializeToString()

    logger.Debug("SL_UpdatePlayerCount : %d, %s, %s", upCount.PlayerCount, upCount.TcpServerIp, upCount.HttpServerIp)

    return ret
end

local LoginServer = class("LoginServer", mvc.AppBase)

function LoginServer:ctor(appName)
    LoginServer.super.ctor(self, appName)
end

function LoginServer:CreateServices(cfg)

    self.loginServer = Server.new()
    LoginServerForClient.__cname = "LoginServer"
    self.loginServer:Register(LoginServerForClient)
    LoginServerForClient.loginServer = self


    self.rpcServer = RpcServer.new()
    LoginServerForGameServer.__cname = "LoginServer"
    self.rpcServer:Register(LoginServerForGameServer)
    LoginServerForGameServer.loginServer = self

    local authCfg = common.config.ReadConfig("etc/authserver.json")
    LoginServerForClient.authServer = RpcClient.new(authCfg.AuthHost)

    --设置玩家数量，记录最少玩家数量的服务器ip地址
    local gameCfg = common.config.ReadConfig("etc/gameserver.json")
    self.stableTcpServer = gameCfg.TcpHost
    self.stableHttpServer = gameCfg.HttpHost
    self.uConnId = 2^53

    self.loginServer:ListenAndServe(cfg.TcpHostForClient, cfg.HttpHostForClient)
    self.rpcServer:ListenAndServe(cfg.LoginHost)

end

return LoginServer