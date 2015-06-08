local logger = require("script.common.logger")
local osg = require("osg")

local conn = ProtoBufConn:new()
local svc = Server:new()

logger.Dump(_G, conn, svc, p)

conn:SetResultServer("test")