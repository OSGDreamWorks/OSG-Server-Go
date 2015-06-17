require("script.runtime.preload")
--����ʹ�õ���
local common = import("common")

local GameServer = import("script.app.gameserver.GameServer")

local cfg = common.config.ReadConfig("etc/gameserver.json")

local GameServices = GameServer.new("GameServer")
GameServices:CreateServices(cfg)

common.WatchSystemSignal()
