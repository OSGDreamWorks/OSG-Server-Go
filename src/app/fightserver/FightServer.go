package fightserver

import (
    "common/config"
    "component/server"
    "common/logger"
    "net"
    "protobuf"
    "runtime/debug"
    "time"
    "sync"
    "sync/atomic"
    "common"
    "fmt"
)

type FightServer struct {
    id              uint64
    rpcServer    *server.Server
    battles         map[string]*Battle
    l               sync.RWMutex
}

var fightServer *FightServer

func ConvertPlayerToCreature(p *protobuf.PlayerBaseInfo) *protobuf.CreatureBaseInfo {
    creature  := &protobuf.CreatureBaseInfo{}
    creature.SetUid(p.GetUid())
    creature.SetStat(p.GetStat())
    creature.SetBp(p.GetBp())
    creature.SetProperty(p.GetProperty())
    creature.SetRevise(p.GetRevise())
    creature.SetAgainst(p.GetAgainst())
    return  creature
}

func RandomCreature(uid string) *protobuf.CreatureBaseInfo {
    creature  := &protobuf.CreatureBaseInfo{}
    creature.SetUid(uid)
    stat := &protobuf.StatusInfo{}
    stat.SetName("daocaoren")
    stat.SetLevel(1)
    stat.SetExperience(0)
    stat.SetHP(106)
    stat.SetMP(109)
    stat.SetRage(0)
    stat.SetGender(0)
    stat.SetModelid(0)
    creature.SetStat(stat)
    bp := &protobuf.PropertyBaseInfo{}
    bp.SetStrenght(3)
    bp.SetVelocity(3)
    bp.SetMana(20)
    bp.SetDefence(2)
    bp.SetStamina(18)
    creature.SetBp(bp)
    prop :=&protobuf.PropertyInfo{}
    prop.SetMaxHP(106)
    prop.SetMaxMP(109)
    prop.SetATK(42)
    prop.SetArmor(38)
    prop.SetAgility(40)
    prop.SetSpirit(102)
    prop.SetRecovery(100)
    creature.SetProperty(prop)
    return  creature
}

func StartServices(cfg *config.SvrConfig, id *uint64) *FightServer {
    fightServer = &FightServer{}

    fightServer = &FightServer{
        rpcServer:server.NewServer(),
        battles:make(map[string]*Battle),
    }
    fightServer.rpcServer.Register(fightServer)

    fightServer.rpcServer.RegCallBackOnConn(
        func(conn server.RpcConn) {
            fightServer.onConn(conn)
        },
    )

    listener, err := net.Listen("tcp", cfg.FsHost[*id])
    if err != nil {
        logger.Fatal("net.Listen: %s", err.Error())
    }

    fightServer.id = *id

    go func() {
        for {
            //For Client/////////////////////////////
            time.Sleep(time.Millisecond * 5)
            conn, err := listener.Accept()
            if err != nil {
                logger.Error("fightserver StartServices %s", err.Error())
                break
            }
            go func() {
                rpcConn := server.NewTCPSocketConn(fightServer.rpcServer, conn, 1000, 0, 1)
                rpcConn.SetResultServer("Connector")

                defer func() {
                    if r := recover(); r != nil {
                        logger.Error("player rpc runtime error begin:", r)
                        debug.PrintStack()
                        rpcConn.Close()

                        logger.Error("player rpc runtime error end ")
                    }
                }()

                fightServer.rpcServer.ServeConn(rpcConn)
            }()
        }
    }()

    return fightServer
}

func WriteResult(conn server.RpcConn, value interface{}) bool {
    err := conn.WriteObj(value)
    if err != nil {
        logger.Info("WriteResult Error %s", err.Error())
        return false
    }
    return true
}

func (self *FightServer) onConn(conn server.RpcConn) {

    go func() {
        for {
            time.Sleep(time.Millisecond * 1000)
            req := &protobuf.Ping{}
            WriteResult(conn, req)
        }
    }()
}

//
func (self *FightServer) addBattle(b *Battle) {
    logger.Info("Connector:addBattle %v, %v", b.GetBid(), b)

    self.l.Lock()
    defer self.l.Unlock()

    //
    self.battles[b.GetBid()] = b
}

//
func (self *FightServer) delBattle(bid string) {
    logger.Info("Connector:delBattle %v", bid)

    _, exist := self.battles[bid]
    if exist {
        self.l.Lock()
        delete(self.battles, bid)
        self.l.Unlock()
    }
}

func (self *FightServer) StartBattleTest(conn server.RpcConn, test protobuf.BattleTest) error {

    logger.Debug("StartBattleWithMoster")

    id := common.GenUUID(fmt.Sprintf("%d", atomic.AddUint64(&self.id, 1)))
    base := &protobuf.BattleInfo{}
    base.SetBid(id)
    partners := make([]*protobuf.CreatureBaseInfo,0,10)
    mosters := make([]*protobuf.CreatureBaseInfo,0,10)
    partners = append(partners, ConvertPlayerToCreature(test.GetPlayer()))

    for _, enemy := range test.GetEnemy() {
        mosters = append(mosters, enemy)
    }

    if len(mosters) == 0 {
        mosters = append(mosters, RandomCreature("1"),RandomCreature("2"),RandomCreature("3"),RandomCreature("4"),RandomCreature("5"))
    }

    base.SetPartner(partners)
    base.SetEnemy(mosters)
    base.SetAttackunits(make([]*protobuf.AttackInfo,0,10))
    base.SetSpells(make([]*protobuf.SpellInfo,0,10))
    b := &Battle{BattleInfo: base}

    //WriteResult(conn, base)
    notify := &protobuf.NotifyBattleStart{}
    notify.SetBid(base.GetBid())
    notify.SetPartner(partners)
    notify.SetEnemy(mosters)
    WriteResult(conn, notify)

    self.addBattle(b)

    return nil
}

func (self *FightServer) StartBattle(conn server.RpcConn, player protobuf.PlayerBaseInfo) error {

    logger.Debug("StartBattle")

    id := common.GenUUID(fmt.Sprintf("%d", atomic.AddUint64(&self.id, 1)))
    base := &protobuf.BattleInfo{}
    base.SetBid(id)
    partners := make([]*protobuf.CreatureBaseInfo,0,10)
    mosters := make([]*protobuf.CreatureBaseInfo,0,10)
    partners = append(partners, ConvertPlayerToCreature(&player))
    mosters = append(mosters, RandomCreature("1"),RandomCreature("2"),RandomCreature("3"),RandomCreature("4"),RandomCreature("5"))
    base.SetPartner(partners)
    base.SetEnemy(mosters)
    base.SetAttackunits(make([]*protobuf.AttackInfo,0,10))
    base.SetSpells(make([]*protobuf.SpellInfo,0,10))
    b := &Battle{BattleInfo: base}

    //WriteResult(conn, base)
    notify := &protobuf.NotifyBattleStart{}
    notify.SetBid(base.GetBid())
    notify.SetPartner(partners)
    notify.SetEnemy(mosters)
    WriteResult(conn, notify)

    self.addBattle(b)

    return nil
}

func (self *FightServer) CalculateBattleResult(conn server.RpcConn, queue protobuf.BattleAttackQueue) error {

    logger.Debug("CalculateBattleResult")

    _, exist := self.battles[queue.GetBid()]

    if !exist {
        return nil
    }

    attackunits := self.battles[queue.GetBid()].GetAttackunits()

    for _, att := range queue.GetAttackunits() {
        attackunits = append(attackunits, att)
    }

    spells := self.battles[queue.GetBid()].GetSpells()

    for _, att := range queue.GetSpells() {
        spells = append(spells, att)
    }

    self.battles[queue.GetBid()].SetAttackunits(attackunits)
    self.battles[queue.GetBid()].SetSpells(spells)
    WriteResult(conn, self.battles[queue.GetBid()])

    end := true
    for _, p := range self.battles[queue.GetBid()].GetPartner() {
        stat := p.GetStat()
        if stat.GetHP() > 0 {
            end = false
        }
    }

    var exp uint32
    exp = 0
    if !end {
        exp = 100
        for _, e := range self.battles[queue.GetBid()].GetEnemy() {
            stat := e.GetStat()
            if stat.GetHP() > 0 {
                end = false
            }
        }
    }

    if end {
        for _, p := range self.battles[queue.GetBid()].GetPartner() {
            notify := &protobuf.NotifyBattleEnd{}
            notify.SetPlayerlid(p.GetUid())
            notify.SetExp(exp)
            WriteResult(conn, notify)
        }
    }

    return nil
}

func (self *FightServer) PingResult(conn server.RpcConn, login protobuf.PingResult) error {
    //keep the conect
    //logger.Debug("PingResult")
    return nil
}