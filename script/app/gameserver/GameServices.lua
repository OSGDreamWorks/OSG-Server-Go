--¼ÓÔØprotobufÄ£¿é
local CSPacket_pb = require("CSPacket_pb")
local SCPacket_pb = require("SCPacket_pb")

local config = require("script.common.config")
local logger = require("script.common.logger")

local GameServices = {}

GameServices.name = "GameServer"

function GameServices:CreateServices(cfg)

    self.mainCache = CachePool:new("etc/maincache.json")

    self.loginServer = Server:new()
    self.loginServer:Register(self)

    self.loginServer:ListenAndServe(cfg.TcpHost, cfg.HttpHost)

end

function GameServices:CS_CheckSession(conn, buf)

    local checkSession = CSPacket_pb.CS_CheckSession()
    checkSession:ParseFromString(buf)

    local checkSessionResult = SCPacket_pb.SC_CheckSessionResult()
    checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.SERVERERROR
    checkSessionResult.server_time = os.time()

    if string.len(checkSession.uid) > 0 then
        local sid, err = self.mainCache:Do("GET", checkSession.uid)
        if string.len(err) == 0  and sid == checkSession.sessionKey then
            checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.OK
        else
            checkSessionResult.result = SCPacket_pb.SC_CheckSessionResult.AUTH_FAILED
        end
    end

    conn:WriteObj("protobuf.SC_CheckSessionResult", checkSessionResult:SerializeToString())

end

function GameServices:CS_Ping(conn, buf)

    local pingResult = SCPacket_pb.SC_PingResult()
    pingResult.server_time = os.time()
    conn:WriteObj("protobuf.SC_PingResult", pingResult:SerializeToString())

end

return GameServices