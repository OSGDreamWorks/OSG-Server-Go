--¼ÓÔØprotobufÄ£¿é
local CLPacket_pb = require("CLPacket_pb")
local LCPacket_pb = require("LCPacket_pb")
local LAPacket_pb = require("LAPacket_pb")
local ALPacket_pb = require("ALPacket_pb")

local config = require("script.common.config")
local logger = require("script.common.logger")

local GateServicesForClient = {}

GateServicesForClient.name = "LoginServer"

function GateServicesForClient:CreateServices(cfg)

    self.loginServer = Server:new()
    self.loginServer:Register(self)

    local authCfg = config.ReadConfig("etc/authserver.json")
    self.authServer = RpcClient:new(authCfg.AuthHost)

    local gameCfg = config.ReadConfig("etc/gameserver.json")
    self.stableTcpServer = gameCfg.TcpHost
    self.stableHttpServer = gameCfg.HttpHost

    self.loginServer:ListenAndServe(cfg.TcpHostForClient, cfg.HttpHostForClient)

end

function GateServicesForClient:CL_CheckAccount(conn, buf)

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

return GateServicesForClient