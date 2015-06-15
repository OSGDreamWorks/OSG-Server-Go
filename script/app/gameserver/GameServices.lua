--¼ÓÔØprotobufÄ£¿é
local CSPacket_pb = require("CSPacket_pb")
local SCPacket_pb = require("SCPacket_pb")

local config = require("script.common.config")
local logger = require("script.common.logger")

local GameServices = {}

GameServices.name = "GameServer"

function GameServices:CreateServices(cfg)

    self.loginServer = Server:new()
    self.loginServer:Register(self)

    self.loginServer:ListenAndServe(cfg.TcpHost, cfg.HttpHost)

end

function GameServices:CS_CheckSession(conn, buf)

    local checkSession = CSPacket_pb.CS_CheckSession()
    checkSession:ParseFromString(buf)

    logger.Debug(checkSession.uid)
    logger.Debug(checkSession.sessionKey)

    local checkSessionResult = SCPacket_pb.SC_CheckSessionResult()
    checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.OK
    checkSessionResult.server_time = os.time()

    conn:WriteObj("protobuf.SC_CheckSessionResult", checkSessionResult:SerializeToString())

end

function GameServices:CS_Ping(conn, buf)

    local pingResult = SCPacket_pb.SC_PingResult()
    pingResult.server_time = os.time()
    conn:WriteObj("protobuf.SC_PingResult", pingResult:SerializeToString())

end

return GameServices