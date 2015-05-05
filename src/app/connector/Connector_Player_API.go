package connector

import (
    "component/server"
    "protobuf"
    "common/logger"
)

func (self *Connector) UpdatePlayerInfo(conn server.RpcConn, info protobuf.PlayerBaseInfo) error {

    self.l.RLock()
    p, exist := self.players[conn.GetId()]
    self.l.RUnlock()
    if !exist {
        return nil
    }

    if info.GetTransform() != nil  {
        p.SetTransform(info.GetTransform())
    }

    WriteResult(conn, p.PlayerBaseInfo)

    return nil
}