package connector

import (
    "component/server"
    "protobuf"
    "math/rand"
    "time"
    "common/logger"
)

func checkEnemyIn(from *protobuf.Transform, to *protobuf.Transform) bool {

    if (from.GetPosition().GetX() != to.GetPosition().GetX()) || (from.GetPosition().GetY() != to.GetPosition().GetY()) {

        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        n := r.Int31n(5)
        if n == 1 {
            return true
        }

    }

    return false
}

func (self *Connector) UpdatePlayerInfo(conn server.RpcConn, info protobuf.PlayerBaseInfo) error {

    self.l.RLock()
    p, exist := self.players[conn.GetId()]
    self.l.RUnlock()
    if !exist {
        return nil
    }

    if info.GetTransform() != nil  {
        if checkEnemyIn(info.GetTransform(), p.GetTransform()) {
            self.FsMgr.Call("FightServer.StartBattle", p.PlayerBaseInfo)
        }
        p.SetTransform(info.GetTransform())
    }

    WriteResult(conn, p.PlayerBaseInfo)

    return nil
}

func (self *Connector) BattleInfo(conn server.RpcConn, info protobuf.BattleInfo) error {

    logger.Debug("BattleInfo")

    for _, p := range info.Partner {

        self.l.RLock()
        p, exist :=  self.playersbyid[p.GetUid()]
        self.l.RUnlock()
        if !exist {
            return nil
        }

        WriteResult(p.conn, info)
        logger.Debug("send BattleInfo to player %v", p)
    }

    return nil
}