package gateserver

import (
    "common/logger"
    "sync"
    "component/server"
    "component/rpc"
    "common/config"
    "net"
    "net/http"
    "github.com/gorilla/websocket"
    "common/protobuf"
    "fmt"
    "time"
    "common"
)

type serverInfo struct {
    PlayerCount     uint32
    TcpServerIp    string
    HttpServerIp    string
}

type GateServices struct {
    l            sync.RWMutex
    m            map[uint32]serverInfo
    stableTcpServer string
    stableHttpServer string
}

var pGateServices *GateServices

func CreateGateServicesForServer(cfg config.GateConfig) *GateServices {
    pGateServices = &GateServices{m: make(map[uint32]serverInfo)}
    rpcServer := rpc.NewServer()

    rpcServer.Register(pGateServices)

    listener, err := net.Listen("tcp", cfg.GateHost)
    if err != nil {
        logger.Fatal("net.Listen: %s", err.Error())
    }

    var uConnId uint32 = 0
    for {
        conn, err := listener.Accept()
        if err != nil {
            logger.Error("gateserver StartServices %s", err.Error())
            break
        }

        uConnId++
        go func(uConnId uint32) {

            pGateServices.l.Lock()
            pGateServices.m[uConnId] = serverInfo{0, "", ""}
            pGateServices.l.Unlock()

            rpcServer.ServeConnWithContext(conn, uConnId)

            pGateServices.l.Lock()
            delete(pGateServices.m, uConnId)
            pGateServices.l.Unlock()

        }(uConnId)
    }

    return pGateServices
}

func (self *GateServices) UpdateCnsPlayerCount(uConnId uint32, info *protobuf.ConnectorInfo, result *protobuf.ConnectorInfoResult) error {
    self.l.Lock()
    self.m[uConnId] = serverInfo{info.GetPlayerCount(), info.GetTcpServerIp(), info.GetHttpServerIp()}

    playerCountMax := uint32(0xffffffff) // max count
    self.stableTcpServer = ""
    self.stableHttpServer = ""
    for _, v := range self.m {
        if (len(v.TcpServerIp) > 0 && len(v.HttpServerIp) > 0) && v.PlayerCount < playerCountMax {
            playerCountMax = v.PlayerCount
            self.stableTcpServer = v.TcpServerIp
            self.stableHttpServer = v.HttpServerIp
        }
    }

    self.l.Unlock()

    logger.Debug("recv cns msg : server %v , player count %v, player ip = %v | %v \n", info.GetServerId(), info.GetPlayerCount(), info.GetTcpServerIp(), info.GetHttpServerIp())
    return nil
}

func (self *GateServices) getStableAddress(isHttp bool) (cnsIp string) {
    self.l.RLock()
    defer self.l.RUnlock()
    if isHttp {
        // websocket or http connector
        return self.stableHttpServer
    }
    return self.stableTcpServer
}

type GateServicesForClient struct {
    m string
    rpcServer    *server.Server
}

var gateServicesForClient *GateServicesForClient

var upgrader = websocket.Upgrader{
    ReadBufferSize:  4096,
    WriteBufferSize: 4096,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}
func wsServeConnHandler(w http.ResponseWriter, r *http.Request) {

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        logger.Info("Upgrade:", err)
        return
    }

    rpcConn := server.NewWebSocketConn(gateServicesForClient.rpcServer, *conn, 128, 45, 2)
    defer func() {
        rpcConn.Close()
    }()

    gateServicesForClient.rpcServer.ServeConn(rpcConn)
}

func CreateGateServicesForClient(cfg config.GateConfig) *GateServicesForClient {
    gateServicesForClient = &GateServicesForClient{}
    rpcServer := server.NewServer()
    rpcServer.Register(gateServicesForClient)

    rpcServer.RegCallBackOnConn(
    func(conn server.RpcConn) {
        gateServicesForClient.onConn(conn)
    },
    )

    listener, err := net.Listen("tcp", cfg.TcpHostForClient)
    if err != nil {
        logger.Fatal("net.Listen: %s", err.Error())
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            logger.Error("gateserver StartServices %s", err.Error())
            break
        }
        go func() {
            rpcConn := server.NewTCPSocketConn(rpcServer, conn, 4, 0, 1)
            rpcServer.ServeConn(rpcConn)
        }()
    }

    http.HandleFunc("/", wsServeConnHandler)
    http.ListenAndServe(cfg.HttpHostForClient, nil)

    return gateServicesForClient
}

func (self *GateServicesForClient) onConn(conn server.RpcConn) {
    rep := protobuf.LoginInfo{}

    serverIp := pGateServices.getStableAddress(conn.IsWebConn())
    rep.ServerIp = &serverIp
    gasinfo := fmt.Sprintf("%s;%d", conn.GetRemoteIp(), time.Now().Unix())
    logger.Info("Client(%s) -> Server(%s)", conn.GetRemoteIp(), serverIp)
    // encode
    encodeInfo := common.Base64Encode([]byte(gasinfo))

    gasinfo = fmt.Sprintf("%s;%s", gasinfo, encodeInfo)

    //fmt.Printf("%s \n", gasinfo)

    rep.GsInfo = &gasinfo

    conn.WriteObj(&rep)

    time.Sleep(10 * time.Second)
    conn.Close()
}

func (self *GateServicesForClient) Login(conn server.RpcConn, login protobuf.Login) error {
    return nil
}