require("script.runtime.preload")
--����ʹ�õ���
local config = require("script.common.config")
local common = require("script.common.define")

local LoginServer = require("script.app.loginserver.LoginServer")

local cfg = config.ReadConfig("etc/loginserver.json")

local loginServices = LoginServer.new("LoginServer")
loginServices:CreateServices(cfg)

common.WatchSystemSignal()
