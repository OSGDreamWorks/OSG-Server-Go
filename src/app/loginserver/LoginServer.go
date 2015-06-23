package gateserver

import (
    "sync"
    "component/server"
    "component/rpc"
    "common/config"
    "protobuf"
    "common/logger"
    "net"
    "time"
    "code.google.com/p/goprotobuf/proto"
)

type serverInfo struct {
    PlayerCount     uint32
    TcpServerIp    string
    HttpServerIp    string
}

type LoginRpcServer struct {
    rpcServer       *rpc.Server
    ConnId          uint64
}

type LoginServer struct {
    loginServer     *server.Server
    authServer      *rpc.Client
}

type LoginMasterServer struct {
    rpcServer       *LoginRpcServer
    loginServer     *LoginServer
    authHost        string
    l                sync.RWMutex
    m               map[uint64]serverInfo
    stableTcpServer  string
    stableHttpServer string
}

var pLoginServices *LoginMasterServer

func CreateServices(lgcfg config.LoginConfig, authcfg config.AuthConfig) *LoginMasterServer {

    pLoginServices = &LoginMasterServer{
        rpcServer : &LoginRpcServer{rpcServer:rpc.NewServer()},
        loginServer : &LoginServer{loginServer:server.NewServer()},
        authHost : authcfg.AuthHost,
        m: make(map[uint64]serverInfo),
    }

    pLoginServices.reConnect()

    pLoginServices.rpcServer.rpcServer.Register(pLoginServices.rpcServer)
    pLoginServices.loginServer.loginServer.Register(pLoginServices.loginServer)

    pLoginServices.rpcServer.rpcServer.RegCallBackOnConn(
    func(uConnId interface{}) {
        uConnIdPtr := uConnId.(*uint64)
        *uConnIdPtr++
        pLoginServices.m[*uConnIdPtr] = serverInfo{0, "", ""}
    },
    )
    pLoginServices.rpcServer.rpcServer.RegCallBackOnConn(
    func(uConnId interface{}) {
        uConnIdPtr := uConnId.(*uint64)
        delete(pLoginServices.m, *uConnIdPtr)
    },
    )

    pLoginServices.rpcServer.rpcServer.ListenAndServe(lgcfg.LoginHost, &pLoginServices.rpcServer.ConnId)

    pLoginServices.loginServer.loginServer.ListenAndServe(lgcfg.TcpHostForClient, lgcfg.HttpHostForClient)

    return pLoginServices
}

func (self *LoginMasterServer)reConnect() {

    authConn, err := net.Dial("tcp", self.authHost)
    for {
        if err == nil {
            break
        }
        logger.Error("AuthServer Connect Error : %v", err.Error())
        time.Sleep(time.Second * 3)
        authConn, err = net.Dial("tcp", self.authHost)
    }

    self.loginServer.authServer = rpc.NewClient(authConn)

    self.loginServer.authServer.AddDisCallback(func(err error) {
        logger.Info("disconnected error:", err)
        self.reConnect()
    })
}

func (self *LoginRpcServer)SL_UpdatePlayerCount(uConnId *uint64, req *[]byte, ret *[]byte) error {

    pLoginServices.l.Lock()

    info := protobuf.SL_UpdatePlayerCount{}
    err := proto.Unmarshal(*req, &info)

    pLoginServices.m[*uConnId] = serverInfo{info.GetPlayerCount(), info.GetTcpServerIp(), info.GetHttpServerIp()}
    playerCountMax := uint32(0xffffffff) // max count
    pLoginServices.stableTcpServer = ""
    pLoginServices.stableHttpServer = ""
    for _, v := range pLoginServices.m {
        if (len(v.TcpServerIp) > 0 && len(v.HttpServerIp) > 0) && v.PlayerCount < playerCountMax {
            playerCountMax = v.PlayerCount
            pLoginServices.stableTcpServer = v.TcpServerIp
            pLoginServices.stableHttpServer = v.HttpServerIp
        }
    }

    result := protobuf.LS_UpdatePlayerCountResult{}
    result.SetResult(protobuf.LS_UpdatePlayerCountResult_OK)
    result.SetServerTime(uint32(time.Now().Unix()))

    buf,err := proto.Marshal(&result)
    *ret = buf

    pLoginServices.l.Unlock()

    logger.Debug("recv cns msg [%v]: server %v , player count %v, player ip = %v | %v \n", *uConnId, info.GetServerId(), info.GetPlayerCount(), info.GetTcpServerIp(), info.GetHttpServerIp())
    return err
}

func (self *LoginServer)CL_CheckAccount(conn server.RpcConn, checkAccount protobuf.CL_CheckAccount) error {

    req := protobuf.LA_CheckAccount{}
    req.SetUid(checkAccount.GetUid())
    req.SetAccount(checkAccount.GetAccount())
    req.SetPassword(checkAccount.GetPassword())
    req.SetOption(checkAccount.GetOption())
    req.SetLanguage(checkAccount.GetLanguage())
    req.SetUdid(checkAccount.GetUdid())

    ret := protobuf.AL_CheckAccountResult{}

    err := self.authServer.Call("AuthServer.LA_CheckAccount", &req, &ret)

    result := protobuf.LC_CheckAccountResult{}
    if err != nil {
        logger.Error("CL_CheckAccount Error : %v", err)
        result.SetResult(protobuf.LC_CheckAccountResult_SERVERERROR)
        result.SetServerTime(uint32(time.Now().Unix()))
        result.SetSessionKey("")
        result.SetUid("")
        result.SetGameServerIp("")
    }else {
        result.SetResult(protobuf.LC_CheckAccountResult_OK)
        result.SetServerTime(uint32(time.Now().Unix()))
        result.SetSessionKey(ret.GetSessionKey())
        result.SetUid(ret.GetUid())
        result.SetGameServerIp(pLoginServices.stableTcpServer)
        if conn.IsWebConn() {
            result.SetGameServerIp(pLoginServices.stableHttpServer)
        }

        logger.Debug("CL_CheckAccount : %v, %v, %v", ret, result.GetSessionKey(), result.GetGameServerIp())
    }
    conn.WriteObj(&result)

    return err
}