--加载osg模块
require("osg")
--加载lua扩展函数
require("script.runtime.extern")

--预加载protobuf模块
import("script.protobuf.PB_PacketCommon_pb")
import("script.protobuf.PB_PacketDefine_pb")
import("script.protobuf.PB_PacketServerDefine_pb")
import("script.protobuf.XShare_Logic_pb")
import("script.protobuf.XShare_Server_pb")
import("script.protobuf.CLPacket_pb")
import("script.protobuf.LCPacket_pb")
import("script.protobuf.CSPacket_pb")
import("script.protobuf.SCPacket_pb")
import("script.protobuf.LAPacket_pb")
import("script.protobuf.ALPacket_pb")
import("script.protobuf.SLPacket_pb")
import("script.protobuf.LSPacket_pb")

--预加载其他模块
import("script.common.init")
import("script.component.db.init")
import("script.component.mvc.init")
