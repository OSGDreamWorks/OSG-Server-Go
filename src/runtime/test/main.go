package main

import (
    "common/logger"
    "flag"
    "common"
    "common/config"
    "net"
    "component/server"
    "reflect"
    "protobuf"
    "time"
)
var (
    gateConfigFile = flag.String("c", "etc/loginserver.json", "config file name for the game server")
    svrConfigFile = flag.String("g", "etc/gameserver.json", "config file name for the fight server")
)

func main() {
    logger.Info("start test server")
    testLogin()
    logger.Info("stop test server")
}

func testLogin() {

    var cfg config.LoginConfig
    if err := config.ReadConfig(*gateConfigFile, &cfg); err != nil {
        logger.Fatal("load config failed, error is: %v", err)
        return
    }

    conn, err := net.Dial("tcp", cfg.TcpHostForClient)
    if err != nil {
        logger.Fatal("%s", err.Error())
    }

    rpcConn := server.NewTCPSocketConn(nil, conn, 4, 0, 1)

    login := &protobuf.CL_CheckAccount{}
    login.SetAccount("account")
    login.SetPassword("password")
    rpcConn.Call("LoginServer.CL_CheckAccount", login)

    rst := new(server.RequestWrap)
    err = rpcConn.ReadRequest(&rst.Request)

    // argv guaranteed to be a pointer now.
    argv := reflect.New(reflect.TypeOf(protobuf.LC_CheckAccountResult{}))
    rpcConn.GetRequestBody(&rst.Request, argv.Interface())

    info := argv.Interface().(*protobuf.LC_CheckAccountResult)

    logger.Info("LoginServer Info : %v", argv.Interface())
    logger.Info("                 %v", &rst.Request)
    logger.Info("                 %v", info.GetSessionKey())

    rpcConn.Close()

    time.Sleep(time.Millisecond * 100)

    conn, err = net.Dial("tcp", info.GetGameServerIp())
    if err != nil {
        logger.Fatal("%s", err.Error())
    }

    rpcConn = server.NewTCPSocketConn(nil, conn, 4, 0, 1)
    defer rpcConn.Close()

    check := &protobuf.CS_CheckSession{}
    check.SetUid(info.GetUid())
    check.SetSessionKey(info.GetSessionKey())
    check.SetTimestamp(uint32(time.Now().Unix()))
    rpcConn.Call("GameServer.CS_CheckSession", check)

    rst = new(server.RequestWrap)
    err = rpcConn.ReadRequest(&rst.Request)

    // argv guaranteed to be a pointer now.
    argv = reflect.New(reflect.TypeOf(protobuf.SC_CheckSessionResult{}))
    rpcConn.GetRequestBody(&rst.Request, argv.Interface())
    logger.Info("GameServer.CS_CheckSession : %v", argv.Interface())
    logger.Info("                 %v", &rst.Request)

    go testPingForever(rpcConn)

    time.Sleep(time.Millisecond * 100)

    testEnterFight(rpcConn)
}

func testPingForever(rpcConn server.RpcConn) {
    for i := 0; i < 100; i++ {

        rpcConn.Lock()

        req := &protobuf.CS_Ping{}
        rpcConn.Call("GameServer.CS_Ping", req)

        rst := new(server.RequestWrap)
        err := rpcConn.ReadRequest(&rst.Request)

        if err != nil {
            logger.Error("Error : %v", err)
        }

        // argv guaranteed to be a pointer now.
        argv := reflect.New(reflect.TypeOf(protobuf.SC_PingResult{}))
        rpcConn.GetRequestBody(&rst.Request, argv.Interface())
        logger.Info("GameServer.CS_Ping : %v", argv.Interface())
        logger.Info("                 %v", &rst.Request)

        rpcConn.Unlock()

        time.Sleep(time.Millisecond * 1000)

    }
}

func testEnterFight(rpcConn server.RpcConn) {

    rpcConn.Lock()

    req := &protobuf.CS_EnterFight{}
    rpcConn.Call("GameServer.CS_EnterFight", req)

    rst := new(server.RequestWrap)
    err := rpcConn.ReadRequest(&rst.Request)

    if err != nil {
        logger.Error("Error : %v", err)
    }

    // argv guaranteed to be a pointer now.
    argv := reflect.New(reflect.TypeOf(protobuf.SC_EnterClientScene{}))
    rpcConn.GetRequestBody(&rst.Request, argv.Interface())
    logger.Info("GameServer.SC_EnterClientScene : %v", argv.Interface())
    logger.Info("                 %v", &rst.Request)

    rpcConn.Unlock()

}

func testCommon() {
    logger.Info("uuid: %v", common.GenUUID("123"))
    logger.Info("uuid: %v", common.GenUUID("account"))

    passwdhash := common.GenPassword("account", "passwd")
    logger.Info("passwdhash: %v", passwdhash)
    logger.Info("check: %v", common.CheckPassword(passwdhash,"account", "passwd"))

    sessionKey := common.GenSessionKey()
    logger.Info("session: %v", sessionKey)
    logger.Info("check: %v", common.CheckSessionKey(sessionKey))
}

func testFightServer() {
//
//    var cfg config.SvrConfig
//    if err := config.ReadConfig(*svrConfigFile, &cfg); err != nil {
//        logger.Fatal("load config failed, error is: %v", err)
//        return
//    }
//
//    conn, err := net.Dial("tcp", cfg.FsHost[0])
//    if err != nil {
//        logger.Fatal("%s", err.Error())
//    }
//
//    rpcConn := server.NewTCPSocketConn(nil, conn, 4, 0, 1)
//    player := &protobuf.PlayerBaseInfo{}
//    player.SetUid("test")
//    player.SetName("Account")
//    player.SetLevel(1)
//    player.SetExperience(0)
//    player.SetHP(106)
//    player.SetMP(109)
//    player.SetRage(109)
//    player.SetMaxHP(106)
//    player.SetMaxMP(109)
//    player.SetMaxRage(109)
//    rpcConn.Call("FightServer.StartBattle", player)
//
//    rst := new(server.RequestWrap)
//    err = rpcConn.ReadRequest(&rst.Request)
//
//    // argv guaranteed to be a pointer now.
//    argv := reflect.New(reflect.TypeOf(protobuf.LoginResult{}))
//    rpcConn.GetRequestBody(&rst.Request, argv.Interface())
//    logger.Info("FightServer.StartBattle : %v", argv.Interface())
//    logger.Info("                 %v", &rst.Request)
}