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
    login.Account = "account"
    login.Password = "password"
    rpcConn.Call(protobuf.Network_Protocol(protobuf.CL_Protocol_eCL_CheckAccount), login)

    rst := new(server.RequestWrap)
    err = rpcConn.ReadRequest(&rst.Packet)

    // argv guaranteed to be a pointer now.
    argv := reflect.New(reflect.TypeOf(protobuf.LC_CheckAccountResult{}))
    rpcConn.GetRequestBody(&rst.Packet, argv.Interface())

    info := argv.Interface().(*protobuf.LC_CheckAccountResult)

    logger.Info("LoginServer Info : %v", argv.Interface())
    logger.Info("                 %v", &rst.Packet)
    logger.Info("                 %v", info.SessionKey)

    rpcConn.Close()

    time.Sleep(time.Millisecond * 1000)

    conn, err = net.Dial("tcp", info.GameServerIp)
    if err != nil {
        logger.Fatal("%s", err.Error())
    }

    rpcConn = server.NewTCPSocketConn(nil, conn, 4, 0, 1)

    check := &protobuf.CS_CheckSession{}
    check.Uid = info.Uid
    check.SessionKey = info.SessionKey
    check.Timestamp = uint32(time.Now().Unix())
    rpcConn.Call(protobuf.Network_Protocol(protobuf.CS_Protocol_eCS_CheckSession), check)

    rst = new(server.RequestWrap)
    err = rpcConn.ReadRequest(&rst.Packet)

    // argv guaranteed to be a pointer now.
    argv = reflect.New(reflect.TypeOf(protobuf.SC_CheckSessionResult{}))
    rpcConn.GetRequestBody(&rst.Packet, argv.Interface())
    logger.Info("GameServer.CS_CheckSession : %v", argv.Interface())
    logger.Info("                 %v", &rst.Packet)

    for i := 0; i < 100; i++ {
        time.Sleep(time.Millisecond * 1000)
        req := &protobuf.CS_Ping{}
        rpcConn.Call(protobuf.Network_Protocol(protobuf.CS_Protocol_eCS_Ping), req)

        rst = new(server.RequestWrap)
        err = rpcConn.ReadRequest(&rst.Packet)

        // argv guaranteed to be a pointer now.
        argv = reflect.New(reflect.TypeOf(protobuf.SC_PingResult{}))
        rpcConn.GetRequestBody(&rst.Packet, argv.Interface())
        logger.Info("GameServer.CS_Ping : %v", argv.Interface())
        logger.Info("                 %v", &rst.Packet)
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
}