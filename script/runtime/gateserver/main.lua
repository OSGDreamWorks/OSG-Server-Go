local osg = require("osg")
local msg_pb = require("script.protobuf.msg_pb")
local logger = require("script.common.logger")
local config = require("script.common.config")
local common = require("script.common.define")

local gateServicesForServer = require("script.app.gateserver.GateServicesForServer")
local gateServicesForClient = require("script.app.gateserver.GateServicesForClient")

local cfg = config.ReadConfig("etc/gateserver.json")

gateServicesForServer:CreateServices(cfg)
gateServicesForClient:CreateServices(cfg)

common.WatchSystemSignal()
