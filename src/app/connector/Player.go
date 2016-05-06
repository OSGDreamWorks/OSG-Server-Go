package connector
import (
    "protobuf"
    "component/server"
    "component/db"
    "common/logger"
)

type Player struct {
    *protobuf.PlayerBaseInfo
    conn      server.RpcConn
}

func (p *Player) Save() {
    db.Write("PlayerBaseInfo", p.Uid, p.PlayerBaseInfo)
}

func (p *Player) OnQuit() {

    logger.Info("OnQuit Begin")

    //pConn
    if p.conn != nil {
        p.conn.Lock()
        defer func() {
            p.conn.Unlock()
            logger.Info("OnQuit End")
        }()
    }

    p.Save()
}