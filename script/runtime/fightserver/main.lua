require("script.runtime.preload")
--引用使用的类
local common = import("common")

local FightServer = import("script.app.fightserver.FightServer")

local cfg = common.config.ReadConfig(svrConfigFile)

local FightServices = FightServer.new("FightServer")
FightServices:CreateServices(cfg,fightServerId)

common.WatchSystemSignal()
