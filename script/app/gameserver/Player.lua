--加载protobuf模块
local XShare_Logic_pb = import("XShare_Logic_pb")

local logger = import("logger")
local mvc = import("mvc")

local Player = class("Player", mvc.ModelBase)

-- 定义属性
Player.schema = clone(mvc.ModelBase.schema)
Player.schema["info"]       = {"table", XShare_Logic_pb.PlayerBaseInfo()}   -- 周围索引 0为没有

function Player:ctor(properties, events, callbacks)
    Player.super.ctor(self, properties)
end

function Player:Save()
end

function Player:OnQuit()
end

return Player