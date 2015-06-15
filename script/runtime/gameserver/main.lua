require("script.runtime.preload")
--����ʹ�õ���
local config = require("script.common.config")
local common = require("script.common.define")

local GameServices = require("script.app.gameserver.GameServices")

local cfg = config.ReadConfig("etc/gameserver.json")

GameServices:CreateServices(cfg)

common.WatchSystemSignal()
