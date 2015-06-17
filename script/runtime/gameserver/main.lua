require("script.runtime.preload")
--引用使用的类
local common = import("common")

local GameServer = import("script.app.gameserver.GameServer")

local cfg = common.config.ReadConfig("etc/gameserver.json")

local GameServices = GameServer.new("GameServer")
GameServices:CreateServices(cfg)

common.WatchSystemSignal()
