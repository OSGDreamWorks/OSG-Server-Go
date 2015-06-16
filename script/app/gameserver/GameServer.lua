--¼ÓÔØprotobufÄ£¿é
local CSPacket_pb = require("CSPacket_pb")
local SCPacket_pb = require("SCPacket_pb")

local config = require("script.common.config")
local logger = require("script.common.logger")

local Player = import(".Player")

local GameServer = class("GameServer", osg.mvc.AppBase)

function GameServer:ctor(appName)
    GameServer.super.ctor(self, appName)
end

function GameServer:CreateServices(cfg)

    local class = self.class

    class.mainCache = CachePool:new("etc/maincache.json")

    class.loginServer = Server:new()
    class.loginServer:Register(class)

    class.loginServer:ListenAndServe(cfg.TcpHost, cfg.HttpHost)

    local p = Player.new()
    --logger.Dump(p.info_.uid)

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