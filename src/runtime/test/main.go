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
    gateConfigFile = flag.String("c", "etc/gateserver.json", "config file name for the game server")
    svrConfigFile = flag.String("g", "etc/gameserver.json", "config file name for the fight server")
)

func main() {
    logger.Info("start test server")
    testLogin()
    logger.Info("stop test server")
}

func testLogin() {

    var cfg config.GateConfig
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

    logger.Info("Connector Info : %v", argv.Interface())
    logger.Info("                 %v", &rst.Request)
    logger.Info("                 %v", info.GetSessionKey())

    rpcConn.Close()

    time.Sleep(time.Millisecond * 1000)

    conn, err = net.Dial("tcp", info.GetGameServerIp())
    if err != nil {
        logger.Fatal("%s", err.Error())
    }

    rpcConn = server.NewTCPSocketConn(nil, conn, 4, 0, 1)

    check := &protobuf.CS_CheckSession{}
    check.SetUid(info.GetUid())
    check.SetSessionKey(info.GetSessionKey())
    check.SetTimestamp(uint32(time.Now().Unix()))
    rpcConn.Call("GameServer.Login", check)

    rst = new(server.RequestWrap)
    err = rpcConn.ReadRequest(&rst.Request)

    // argv guaranteed to be a pointer now.
    argv = reflect.New(reflect.TypeOf(protobuf.SC_CheckSessionResult{}))
    rpcConn.GetRequestBody(&rst.Request, argv.Interface())
    logger.Info("Connector.Login : %v", argv.Interface())
    logger.Info("                 %v", &rst.Request)

    for i := 0; i < 100; i++ {
        time.Sleep(time.Millisecond * 1000)
        req := &protobuf.CS_Ping{}
        rpcConn.Call("GameServer.Ping", req)

        rst = new(server.RequestWrap)
        err = rpcConn.ReadRequest(&rst.Request)

        // argv guaranteed to be a pointer now.
        argv = reflect.New(reflect.TypeOf(protobuf.SC_PingResult{}))
        rpcConn.GetRequestBody(&rst.Request, argv.Interface())
        logger.Info("Connector.Ping : %v", argv.Interface())
        logger.Info("                 %v", &rst.Request)
    }

    rpcConn.Close()
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