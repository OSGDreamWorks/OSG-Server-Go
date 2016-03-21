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
    "github.com/golang/protobuf/proto"
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

    pLoginServices.rpcServer.rpcServer.ApplyProtocol(protobuf.SL_Protocol_value)
    pLoginServices.rpcServer.rpcServer.Register(pLoginServices.rpcServer)

    pLoginServices.loginServer.loginServer.ApplyProtocol(protobuf.CL_Protocol_value)
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

    pLoginServices.m[*uConnId] = serverInfo{info.PlayerCount, info.TcpServerIp, info.HttpServerIp}
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
    result.Result = protobuf.LS_UpdatePlayerCountResult_OK
    result.ServerTime = uint32(time.Now().Unix())

    buf,err := proto.Marshal(&result)
    *ret = buf

    pLoginServices.l.Unlock()

    logger.Debug("recv cns msg [%v]: server %v , player count %v, player ip = %v | %v \n", *uConnId, info.ServerId, info.PlayerCount, info.TcpServerIp, info.HttpServerIp)
    return err
}

func (self *LoginServer)CL_CheckAccount(conn server.RpcConn, checkAccount protobuf.CL_CheckAccount) error {

    req := protobuf.LA_CheckAccount{}
    req.Uid = checkAccount.Uid
    req.Account = checkAccount.Account
    req.Password = checkAccount.Password
    req.Option = checkAccount.Option
    req.Language = checkAccount.Language
    req.Udid = checkAccount.Udid

    ret := protobuf.AL_CheckAccountResult{}

    err := self.authServer.Call(protobuf.Network_Protocol(protobuf.LA_Protocol_eLA_CheckAccount), &req, &ret)

    result := protobuf.LC_CheckAccountResult{}
    if err != nil {
        logger.Error("CL_CheckAccount Error : %v", err)
        result.Result = protobuf.LC_CheckAccountResult_SERVERERROR
        result.ServerTime = uint32(time.Now().Unix())
        result.SessionKey = ""
        result.Uid = ""
        result.GameServerIp = ""
    }else {
        result.Result = protobuf.LC_CheckAccountResult_OK
        result.ServerTime = uint32(time.Now().Unix())
        result.SessionKey = ret.SessionKey
        result.Uid = ret.Uid
        result.GameServerIp = pLoginServices.stableTcpServer
        if conn.IsWebConn() {
            result.GameServerIp = pLoginServices.stableHttpServer
        }

        logger.Debug("CL_CheckAccount : %v, %v, %v", ret, result.SessionKey, result.GameServerIp)
    }
    conn.WriteObj(&result)

    return err
}