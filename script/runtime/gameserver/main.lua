require("script.runtime.preload")
--����ʹ�õ���
local config = require("script.common.config")
local common = require("script.common.define")
local logger = require("script.common.logger")

local GameServer = require("script.app.gameserver.GameServer")

local cfg = config.ReadConfig("etc/gameserver.json")

local GameServices = GameServer.new("GameServer")
GameServices:CreateServices(cfg)

common.WatchSystemSignal()
