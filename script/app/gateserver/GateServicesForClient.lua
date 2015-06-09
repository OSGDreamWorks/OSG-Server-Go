local logger = require("script.common.logger")

local GateServicesForClient = {}

GateServicesForClient.name = "GateServer"

function GateServicesForClient:CreateServices(cfg)
    self.rpcServer = Server:new()
    logger.Dump(self)
    self.rpcServer:Register(self)
    self.rpcServer:ListenAndServe(cfg.TcpHostForClient, cfg.HttpHostForClient)
end

function GateServicesForClient:TestPrint(str, ok)
    logger.Debug("hello")
end

function GateServicesForClient:TestPrint2(str, ok)
    logger.Debug("hello2 %s", self)
end

return GateServicesForClient