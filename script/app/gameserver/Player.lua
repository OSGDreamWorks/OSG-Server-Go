--加载protobuf模块
local XShare_Logic_pb = require("XShare_Logic_pb")

local logger = require("script.common.logger")

local Player = class("Player", osg.mvc.ModelBase)

-- 定义属性
Player.schema = clone(osg.mvc.ModelBase.schema)
Player.schema["info"]       = {"table", XShare_Logic_pb.PlayerInfo()}   -- 周围索引 0为没有

function Player:ctor(properties, events, callbacks)
    Player.super.ctor(self, properties)
    self.info_.uid = "test"
end

return Player