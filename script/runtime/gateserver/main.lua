require("script.runtime.preload")
--引用使用的类
local config = require("script.common.config")
local common = require("script.common.define")

local gateServicesForServer = require("script.app.gateserver.GateServicesForServer")
local gateServicesForClient = require("script.app.gateserver.GateServicesForClient")

local cfg = config.ReadConfig("etc/gateserver.json")

gateServicesForServer:CreateServices(cfg)
gateServicesForClient:CreateServices(cfg)

common.WatchSystemSignal()
