--¼ÓÔØprotobufÄ£¿é
local XShare_Logic_pb = require("XShare_Logic_pb")

local Player = XShare_Logic_pb.PlayerInfo()

function Player:new()
    self.uid = ""
    self.stat = ""
end