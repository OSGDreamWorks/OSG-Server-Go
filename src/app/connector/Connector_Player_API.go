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
        n := r.Int31n(100)
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
            //self.FsMgr.Call("FightServer.StartBattle", p.PlayerBaseInfo)
        }
        p.SetTransform(info.GetTransform())
    }

    WriteResult(conn, p.PlayerBaseInfo)

    return nil
}

func (self *Connector) BattleTest(conn server.RpcConn, test protobuf.BattleTest) error {

    self.l.RLock()
    p, exist := self.players[conn.GetId()]
    self.l.RUnlock()
    if !exist {
        return nil
    }

    test.SetPlayer(p.PlayerBaseInfo)

    self.FsMgr.Call("FightServer.StartBattleTest", test)

    return nil
}

func (self *Connector) BattleInfo(conn server.RpcConn, info protobuf.BattleInfo) error {

    logger.Debug("BattleInfo")

    for _, p := range info.GetPartner() {

        self.l.RLock()
        p, exist :=  self.playersbyid[p.GetUid()]
        self.l.RUnlock()
        if !exist {
            return nil
        }

        WriteResult(p.conn, &info)
        logger.Debug("send BattleInfo to player %v", p)
    }

    return nil
}

func (self *Connector) BattleAttackQueue(conn server.RpcConn, queue protobuf.BattleAttackQueue) error {
    self.FsMgr.Call("FightServer.CalculateBattleResult", &queue)
    return nil
}

func (self *Connector) NotifyBattleStart(conn server.RpcConn, notify protobuf.NotifyBattleStart) error {

    logger.Debug("NotifyBattleStart")

    for _, p := range notify.GetPartner() {

        self.l.RLock()
        p, exist :=  self.playersbyid[p.GetUid()]
        self.l.RUnlock()
        if !exist {
            return nil
        }

        WriteResult(p.conn, &notify)
        logger.Debug("send NotifyBattleStart to player %v", p)
    }

    return nil
}

func (self *Connector) NotifyBattleEnd(conn server.RpcConn, notify protobuf.NotifyBattleEnd) error {

    logger.Debug("NotifyBattleEnd")

    self.l.RLock()
    p, exist :=  self.playersbyid[notify.GetPlayerlid()]
    self.l.RUnlock()
    if !exist {
        return nil
    }

    WriteResult(p.conn, notify)
    logger.Debug("send NotifyBattleEnd to player %v", p)

    return nil
}