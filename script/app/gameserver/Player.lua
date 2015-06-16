--����protobufģ��
local XShare_Logic_pb = require("XShare_Logic_pb")

local logger = require("script.common.logger")

local Player = class("Player", osg.mvc.ModelBase)

-- ��������
Player.schema = clone(osg.mvc.ModelBase.schema)
Player.schema["info"]       = {"table", XShare_Logic_pb.PlayerInfo()}   -- ��Χ���� 0Ϊû��

function Player:ctor(properties, events, callbacks)
    Player.super.ctor(self, properties)
    self.info_.uid = "test"
end

return Player