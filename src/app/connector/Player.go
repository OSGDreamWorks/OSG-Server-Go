package connector
import (
    "common/protobuf"
    "component/server"
)

type Player struct {
    *protobuf.PlayerBaseInfo
    conn      server.RpcConn
}

func (p *Player) OnQuit() {

}