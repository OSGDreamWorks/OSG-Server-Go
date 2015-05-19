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
    creature.SetName(p.GetName())
    creature.SetLevel(p.GetLevel())
    creature.SetExperience(p.GetExperience())
    creature.SetHP(p.GetHP())
    creature.SetMP(p.GetMP())
    creature.SetMaxHP(p.GetMaxHP())
    creature.SetMaxMP(p.GetMaxMP())
    creature.SetGender(p.GetGender())
    creature.SetModelid(p.GetModelid())
    creature.SetTransform(p.GetTransform())
    creature.SetStrenght(p.GetStrenght())
    creature.SetVelocity(p.GetVelocity())
    creature.SetMana(p.GetMana())
    creature.SetDefence(p.GetDefence())
    creature.SetStamina(p.GetStamina())
    creature.SetATK(p.GetATK())
    creature.SetArmor(p.GetArmor())
    creature.SetAgility(p.GetAgility())
    creature.SetSpirit(p.GetSpirit())
    creature.SetRecovery(p.GetRecovery())
    return  creature
}

func RandomCreature(uid string) *protobuf.CreatureBaseInfo {
    creature  := &protobuf.CreatureBaseInfo{}
    creature.SetUid(uid)
    creature.SetName("daocaoren")
    creature.SetLevel(1)
    creature.SetExperience(0)
    creature.SetHP(106)
    creature.SetMP(109)
    creature.SetMaxHP(106)
    creature.SetMaxMP(109)
    creature.SetGender(0)
    creature.SetModelid(0)
    creature.SetStrenght(3)
    creature.SetVelocity(3)
    creature.SetMana(20)
    creature.SetDefence(2)
    creature.SetStamina(18)
    creature.SetATK(42)
    creature.SetArmor(38)
    creature.SetAgility(40)
    creature.SetSpirit(102)
    creature.SetRecovery(100)
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

    logger.Debug("StartBattleTest")

    id := common.GenUUID(fmt.Sprintf("%d", atomic.AddUint64(&self.id, 1)))
    base := &protobuf.BattleInfo{}
    base.SetBid(id)
    partners := make([]*protobuf.CreatureBaseInfo,0,10)
    mosters := make([]*protobuf.CreatureBaseInfo,0,10)
    partners = append(partners, ConvertPlayerToCreature(test.GetPlayer()))

    for _, enemy := range test.GetMoster() {
        mosters = append(mosters, enemy)
    }

    base.SetPartner(partners)
    base.SetMoster(mosters)
    base.SetAttackunits(make([]*protobuf.AttackInfo,0,10))
    base.SetSpells(make([]*protobuf.SpellInfo,0,10))
    b := &Battle{BattleInfo: base}

    //WriteResult(conn, base)
    notify := &protobuf.NotifyBattleStart{}
    notify.SetBid(base.GetBid())
    notify.SetPartner(partners)
    notify.SetMoster(mosters)
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
    base.SetMoster(mosters)
    base.SetAttackunits(make([]*protobuf.AttackInfo,0,10))
    base.SetSpells(make([]*protobuf.SpellInfo,0,10))
    b := &Battle{BattleInfo: base}

    //WriteResult(conn, base)
    notify := &protobuf.NotifyBattleStart{}
    notify.SetBid(base.GetBid())
    notify.SetPartner(partners)
    notify.SetMoster(mosters)
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
        if p.GetHP() > 0 {
            end = false
        }
    }

    var exp uint32
    exp = 0
    if !end {
        exp = 100
        for _, m := range self.battles[queue.GetBid()].GetMoster() {
            if m.GetHP() > 0 {
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