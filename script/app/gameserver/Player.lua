--����protobufģ��
local XShare_Logic_pb = import("XShare_Logic_pb")

local logger = import("logger")
local mvc = import("mvc")

local Player = class("Player", mvc.ModelBase)

-- ��������
Player.schema = clone(mvc.ModelBase.schema)
Player.schema["info"]       = {"table", XShare_Logic_pb.PlayerBaseInfo()}   -- ��Χ���� 0Ϊû��

function Player:ctor(properties, events, callbacks)
    Player.super.ctor(self, properties)
end

function Player:Save()
end

function Player:OnQuit()
end

return Player