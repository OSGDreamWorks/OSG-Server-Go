require("script.runtime.preload")
--引用使用的类
local config = require("script.common.config")
local common = require("script.common.define")

local LoginServer = require("script.app.loginserver.LoginServer")

local cfg = config.ReadConfig("etc/loginserver.json")

local loginServices = LoginServer.new("LoginServer")
loginServices:CreateServices(cfg)

common.WatchSystemSignal()
