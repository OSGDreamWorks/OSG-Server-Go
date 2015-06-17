--¼ÓÔØprotobufÄ£¿é
local CLPacket_pb = import("CLPacket_pb")
local LCPacket_pb = import("LCPacket_pb")
local LAPacket_pb = import("LAPacket_pb")
local ALPacket_pb = import("ALPacket_pb")

local common = import("common")
local logger = import("logger")
local mvc = import("mvc")

local LoginServer = class("LoginServer", mvc.AppBase)

function LoginServer:ctor(appName)
    LoginServer.super.ctor(self, appName)
end

function LoginServer:CreateServices(cfg)

    local class = self.class

    class.loginServer = Server:new()
    class.loginServer:Register(class)

    local authCfg = common.config.ReadConfig("etc/authserver.json")
    class.authServer = RpcClient.new(authCfg.AuthHost)

    local gameCfg = common.config.ReadConfig("etc/gameserver.json")
    class.stableTcpServer = gameCfg.TcpHost
    class.stableHttpServer = gameCfg.HttpHost

    class.loginServer:ListenAndServe(cfg.TcpHostForClient, cfg.HttpHostForClient)

end

function LoginServer:CL_CheckAccount(conn, buf)

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

        if conn:IsWebConn() then
            checkAccountResult.gameServerIp = self.stableHttpServer
        else
            checkAccountResult.gameServerIp = self.stableTcpServer
        end

    end

    conn:WriteObj("protobuf.LC_CheckAccountResult", checkAccountResult:SerializeToString())

    return 0

end

return LoginServer