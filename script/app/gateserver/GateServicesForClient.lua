--¼ÓÔØprotobufÄ£¿é
local CLPacket_pb = require("CLPacket_pb")
local LCPacket_pb = require("LCPacket_pb")

local logger = require("script.common.logger")
local common = require("script.common.define")

local GateServicesForClient = {}

GateServicesForClient.name = "LoginServer"

function GateServicesForClient:CreateServices(cfg)
    self.rpcServer = Server:new()
    self.rpcServer:Register(self)
    self.rpcServer:ListenAndServe(cfg.TcpHostForClient, cfg.HttpHostForClient)
end

function GateServicesForClient:CL_CheckAccount(conn, buf)

    local checkAccount = CLPacket_pb.CL_CheckAccount()
    checkAccount:ParseFromString(buf)
    logger.Debug(checkAccount.account)
    logger.Debug(checkAccount.password)

    local checkAccountResult = LCPacket_pb.LC_CheckAccountResult()
    checkAccountResult.result = LCPacket_pb.LC_CheckAccountResult.OK
    checkAccountResult.server_time = os.time()
    checkAccountResult.sessionKey = "test"
    checkAccountResult.uid = checkAccount.account
    conn:WriteObj("protobuf.LC_CheckAccountResult", checkAccountResult:SerializeToString())

    return 0

end

return GateServicesForClient