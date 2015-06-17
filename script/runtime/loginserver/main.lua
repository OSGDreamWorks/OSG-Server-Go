require("script.runtime.preload")

--����ʹ�õ���
local common = import("common")

local LoginServer = import("script.app.loginserver.LoginServer")

local cfg = common.config.ReadConfig("etc/loginserver.json")

local loginServices = LoginServer.new("LoginServer")
loginServices:CreateServices(cfg)

common.WatchSystemSignal()
