-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
module('PB_PacketServerDefine_pb')


local LA_PROTOCOL = protobuf.EnumDescriptor();
local LA_PROTOCOL_ELA_PACKETBEGIN_ENUM = protobuf.EnumValueDescriptor();
local LA_PROTOCOL_ELA_CONNECTED_ENUM = protobuf.EnumValueDescriptor();
local LA_PROTOCOL_ELA_DISCONNECTED_ENUM = protobuf.EnumValueDescriptor();
local LA_PROTOCOL_ELA_CHECKACCOUNT_ENUM = protobuf.EnumValueDescriptor();
local LA_PROTOCOL_ELA_PACKETEND_ENUM = protobuf.EnumValueDescriptor();
local AL_PROTOCOL = protobuf.EnumDescriptor();
local AL_PROTOCOL_EAL_PACKETBEGIN_ENUM = protobuf.EnumValueDescriptor();
local AL_PROTOCOL_EAL_CONNECTED_ENUM = protobuf.EnumValueDescriptor();
local AL_PROTOCOL_EAL_DISCONNECTED_ENUM = protobuf.EnumValueDescriptor();
local AL_PROTOCOL_EAL_CHECKACCOUNTRESULT_ENUM = protobuf.EnumValueDescriptor();
local AL_PROTOCOL_EAL_PACKETEND_ENUM = protobuf.EnumValueDescriptor();
local LS_PROTOCOL = protobuf.EnumDescriptor();
local LS_PROTOCOL_ELS_PACKETBEGIN_ENUM = protobuf.EnumValueDescriptor();
local LS_PROTOCOL_ELS_CONNECTED_ENUM = protobuf.EnumValueDescriptor();
local LS_PROTOCOL_ELS_DISCONNECTED_ENUM = protobuf.EnumValueDescriptor();
local LS_PROTOCOL_ELS_UPDATEPLAYERCOUNTRESULT_ENUM = protobuf.EnumValueDescriptor();
local LS_PROTOCOL_ESLS_PACKETEND_ENUM = protobuf.EnumValueDescriptor();
local SL_PROTOCOL = protobuf.EnumDescriptor();
local SL_PROTOCOL_ESL_PACKETBEGIN_ENUM = protobuf.EnumValueDescriptor();
local SL_PROTOCOL_ESL_CONNECTED_ENUM = protobuf.EnumValueDescriptor();
local SL_PROTOCOL_ESL_DISCONNECTED_ENUM = protobuf.EnumValueDescriptor();
local SL_PROTOCOL_ESL_UPDATEPLAYERCOUNT_ENUM = protobuf.EnumValueDescriptor();
local SL_PROTOCOL_ESL_PACKETEND_ENUM = protobuf.EnumValueDescriptor();

LA_PROTOCOL_ELA_PACKETBEGIN_ENUM.name = "eLA_PacketBegin"
LA_PROTOCOL_ELA_PACKETBEGIN_ENUM.index = 0
LA_PROTOCOL_ELA_PACKETBEGIN_ENUM.number = 11000
LA_PROTOCOL_ELA_CONNECTED_ENUM.name = "eLA_Connected"
LA_PROTOCOL_ELA_CONNECTED_ENUM.index = 1
LA_PROTOCOL_ELA_CONNECTED_ENUM.number = 11000
LA_PROTOCOL_ELA_DISCONNECTED_ENUM.name = "eLA_Disconnected"
LA_PROTOCOL_ELA_DISCONNECTED_ENUM.index = 2
LA_PROTOCOL_ELA_DISCONNECTED_ENUM.number = 11001
LA_PROTOCOL_ELA_CHECKACCOUNT_ENUM.name = "eLA_CheckAccount"
LA_PROTOCOL_ELA_CHECKACCOUNT_ENUM.index = 3
LA_PROTOCOL_ELA_CHECKACCOUNT_ENUM.number = 11002
LA_PROTOCOL_ELA_PACKETEND_ENUM.name = "eLA_PacketEnd"
LA_PROTOCOL_ELA_PACKETEND_ENUM.index = 4
LA_PROTOCOL_ELA_PACKETEND_ENUM.number = 12000
LA_PROTOCOL.name = "LA_Protocol"
LA_PROTOCOL.full_name = ".protobuf.LA_Protocol"
LA_PROTOCOL.values = {LA_PROTOCOL_ELA_PACKETBEGIN_ENUM,LA_PROTOCOL_ELA_CONNECTED_ENUM,LA_PROTOCOL_ELA_DISCONNECTED_ENUM,LA_PROTOCOL_ELA_CHECKACCOUNT_ENUM,LA_PROTOCOL_ELA_PACKETEND_ENUM}
AL_PROTOCOL_EAL_PACKETBEGIN_ENUM.name = "eAL_PacketBegin"
AL_PROTOCOL_EAL_PACKETBEGIN_ENUM.index = 0
AL_PROTOCOL_EAL_PACKETBEGIN_ENUM.number = 12000
AL_PROTOCOL_EAL_CONNECTED_ENUM.name = "eAL_Connected"
AL_PROTOCOL_EAL_CONNECTED_ENUM.index = 1
AL_PROTOCOL_EAL_CONNECTED_ENUM.number = 12000
AL_PROTOCOL_EAL_DISCONNECTED_ENUM.name = "eAL_Disconnected"
AL_PROTOCOL_EAL_DISCONNECTED_ENUM.index = 2
AL_PROTOCOL_EAL_DISCONNECTED_ENUM.number = 12001
AL_PROTOCOL_EAL_CHECKACCOUNTRESULT_ENUM.name = "eAL_CheckAccountResult"
AL_PROTOCOL_EAL_CHECKACCOUNTRESULT_ENUM.index = 3
AL_PROTOCOL_EAL_CHECKACCOUNTRESULT_ENUM.number = 12002
AL_PROTOCOL_EAL_PACKETEND_ENUM.name = "eAL_PacketEnd"
AL_PROTOCOL_EAL_PACKETEND_ENUM.index = 4
AL_PROTOCOL_EAL_PACKETEND_ENUM.number = 13000
AL_PROTOCOL.name = "AL_Protocol"
AL_PROTOCOL.full_name = ".protobuf.AL_Protocol"
AL_PROTOCOL.values = {AL_PROTOCOL_EAL_PACKETBEGIN_ENUM,AL_PROTOCOL_EAL_CONNECTED_ENUM,AL_PROTOCOL_EAL_DISCONNECTED_ENUM,AL_PROTOCOL_EAL_CHECKACCOUNTRESULT_ENUM,AL_PROTOCOL_EAL_PACKETEND_ENUM}
LS_PROTOCOL_ELS_PACKETBEGIN_ENUM.name = "eLS_PacketBegin"
LS_PROTOCOL_ELS_PACKETBEGIN_ENUM.index = 0
LS_PROTOCOL_ELS_PACKETBEGIN_ENUM.number = 13000
LS_PROTOCOL_ELS_CONNECTED_ENUM.name = "eLS_Connected"
LS_PROTOCOL_ELS_CONNECTED_ENUM.index = 1
LS_PROTOCOL_ELS_CONNECTED_ENUM.number = 13000
LS_PROTOCOL_ELS_DISCONNECTED_ENUM.name = "eLS_Disconnected"
LS_PROTOCOL_ELS_DISCONNECTED_ENUM.index = 2
LS_PROTOCOL_ELS_DISCONNECTED_ENUM.number = 13001
LS_PROTOCOL_ELS_UPDATEPLAYERCOUNTRESULT_ENUM.name = "eLS_UpdatePlayerCountResult"
LS_PROTOCOL_ELS_UPDATEPLAYERCOUNTRESULT_ENUM.index = 3
LS_PROTOCOL_ELS_UPDATEPLAYERCOUNTRESULT_ENUM.number = 13002
LS_PROTOCOL_ESLS_PACKETEND_ENUM.name = "eSLS_PacketEnd"
LS_PROTOCOL_ESLS_PACKETEND_ENUM.index = 4
LS_PROTOCOL_ESLS_PACKETEND_ENUM.number = 14000
LS_PROTOCOL.name = "LS_Protocol"
LS_PROTOCOL.full_name = ".protobuf.LS_Protocol"
LS_PROTOCOL.values = {LS_PROTOCOL_ELS_PACKETBEGIN_ENUM,LS_PROTOCOL_ELS_CONNECTED_ENUM,LS_PROTOCOL_ELS_DISCONNECTED_ENUM,LS_PROTOCOL_ELS_UPDATEPLAYERCOUNTRESULT_ENUM,LS_PROTOCOL_ESLS_PACKETEND_ENUM}
SL_PROTOCOL_ESL_PACKETBEGIN_ENUM.name = "eSL_PacketBegin"
SL_PROTOCOL_ESL_PACKETBEGIN_ENUM.index = 0
SL_PROTOCOL_ESL_PACKETBEGIN_ENUM.number = 14000
SL_PROTOCOL_ESL_CONNECTED_ENUM.name = "eSL_Connected"
SL_PROTOCOL_ESL_CONNECTED_ENUM.index = 1
SL_PROTOCOL_ESL_CONNECTED_ENUM.number = 14000
SL_PROTOCOL_ESL_DISCONNECTED_ENUM.name = "eSL_Disconnected"
SL_PROTOCOL_ESL_DISCONNECTED_ENUM.index = 2
SL_PROTOCOL_ESL_DISCONNECTED_ENUM.number = 14001
SL_PROTOCOL_ESL_UPDATEPLAYERCOUNT_ENUM.name = "eSL_UpdatePlayerCount"
SL_PROTOCOL_ESL_UPDATEPLAYERCOUNT_ENUM.index = 3
SL_PROTOCOL_ESL_UPDATEPLAYERCOUNT_ENUM.number = 14002
SL_PROTOCOL_ESL_PACKETEND_ENUM.name = "eSL_PacketEnd"
SL_PROTOCOL_ESL_PACKETEND_ENUM.index = 4
SL_PROTOCOL_ESL_PACKETEND_ENUM.number = 20000
SL_PROTOCOL.name = "SL_Protocol"
SL_PROTOCOL.full_name = ".protobuf.SL_Protocol"
SL_PROTOCOL.values = {SL_PROTOCOL_ESL_PACKETBEGIN_ENUM,SL_PROTOCOL_ESL_CONNECTED_ENUM,SL_PROTOCOL_ESL_DISCONNECTED_ENUM,SL_PROTOCOL_ESL_UPDATEPLAYERCOUNT_ENUM,SL_PROTOCOL_ESL_PACKETEND_ENUM}

eAL_CheckAccountResult = 12002
eAL_Connected = 12000
eAL_Disconnected = 12001
eAL_PacketBegin = 12000
eAL_PacketEnd = 13000
eLA_CheckAccount = 11002
eLA_Connected = 11000
eLA_Disconnected = 11001
eLA_PacketBegin = 11000
eLA_PacketEnd = 12000
eLS_Connected = 13000
eLS_Disconnected = 13001
eLS_PacketBegin = 13000
eLS_UpdatePlayerCountResult = 13002
eSLS_PacketEnd = 14000
eSL_Connected = 14000
eSL_Disconnected = 14001
eSL_PacketBegin = 14000
eSL_PacketEnd = 20000
eSL_UpdatePlayerCount = 14002
