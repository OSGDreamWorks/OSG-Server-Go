--预加载protobuf模块
local PB_PacketCommon_pb = require("script.protobuf.PB_PacketCommon_pb")
local PB_PacketDefine_pb = require("script.protobuf.PB_PacketDefine_pb")
local PB_PacketServerDefine_pb = require("script.protobuf.PB_PacketServerDefine_pb")
local XShare_Logic_pb = require("script.protobuf.XShare_Logic_pb")
local XShare_Server_pb = require("script.protobuf.XShare_Server_pb")
local CLPacket_pb = require("script.protobuf.CLPacket_pb")
local LCPacket_pb = require("script.protobuf.LCPacket_pb")
local CSPacket_pb = require("script.protobuf.CSPacket_pb")
local SCPacket_pb = require("script.protobuf.SCPacket_pb")
local LAPacket_pb = require("script.protobuf.LAPacket_pb")
local LAPacket_pb = require("script.protobuf.ALPacket_pb")
local SLPacket_pb = require("script.protobuf.SLPacket_pb")
local SLPacket_pb = require("script.protobuf.LSPacket_pb")
--加载osg模块
local osg = require("osg")
--引用使用的类
local config = require("script.common.config")
local common = require("script.common.define")

local gateServicesForServer = require("script.app.gateserver.GateServicesForServer")
local gateServicesForClient = require("script.app.gateserver.GateServicesForClient")

local cfg = config.ReadConfig("etc/gateserver.json")

gateServicesForServer:CreateServices(cfg)
gateServicesForClient:CreateServices(cfg)

common.WatchSystemSignal()
