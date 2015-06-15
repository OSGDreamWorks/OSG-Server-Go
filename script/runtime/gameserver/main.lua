require("script.runtime.preload")
--引用使用的类
local config = require("script.common.config")
local common = require("script.common.define")

local GameServices = require("script.app.gameserver.GameServices")

local cfg = config.ReadConfig("etc/gameserver.json")

GameServices:CreateServices(cfg)

common.WatchSystemSignal()
