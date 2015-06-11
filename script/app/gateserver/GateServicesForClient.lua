local logger = require("script.common.logger")
local common = require("script.common.define")

local GateServicesForClient = {}

GateServicesForClient.name = "GateServer"

function GateServicesForClient:CreateServices(cfg)
    self.rpcServer = Server:new()
    logger.Dump(_G)
    self.rpcServer:Register(self)
    self.rpcServer:ListenAndServe(cfg.TcpHostForClient, cfg.HttpHostForClient)
end

function GateServicesForClient:TestPrint(conn, buf)
    logger.Debug("hello")
end

function GateServicesForClient:TestPrint2(conn, buf)

    local msg_pb = require("msg_pb")

    local chatdata = msg_pb.Login()
    chatdata:ParseFromString(buf)
    logger.DumpString(buf)
    logger.Debug(chatdata.account)
    logger.Debug(chatdata.password)
    logger.Debug("%d",chatdata.create_time)

    local loginResultdata = msg_pb.LoginResult()
    loginResultdata.result = msg_pb.LoginResult.OK
    loginResultdata.server_time = os.time()
    loginResultdata.sessionKey = "test"
    loginResultdata.uid = chatdata.account
    conn:WriteObj("protobuf.LoginResult", loginResultdata:SerializeToString())

    return 0

end

return GateServicesForClient