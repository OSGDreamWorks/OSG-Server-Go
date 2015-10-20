package connector

import (
	"common/config"
	"common/logger"
	"protobuf"
	"component/db"
	"component/rpc"
	"component/server"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"time"
	"runtime/debug"
	"sync"
	"fmt"
	"code.google.com/p/goprotobuf/proto"
	"github.com/garyburd/redigo/redis"
)

type serverInfo struct {
	PlayerCount uint16
	ServerIp    string
}

type Connector struct {
	m            map[uint32]serverInfo
	stableServer string
	maincache    *db.CachePool
	authserver  *rpc.Client
	loginserver   *rpc.Client
	FsMgr        FServerConnMgr
	rpcServer    *server.Server
	players         map[uint64]*Player
	playersbyid     map[string]*Player
	l               sync.RWMutex
	id              uint32
	listenTcpIp        string
	listenHttpIp        string
}

var pConnector *Connector

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsServeConnHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "rpc: error")
		return
	}

	logger.Debug("wsServeConnHandler : %v", r.FormValue("method"))

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Info("Upgrade:", err.Error())
		conn, bufrw, err := w.(http.Hijacker).Hijack()
		if err != nil {
			logger.Debug("rpc hijacking %v : %v", r.RemoteAddr, err.Error())
			return
		}else {

			httpConn := server.NewTCPSocketConn(pConnector.rpcServer, conn, 1, 1, 1)
			logger.Debug("rpc hijacking %v : %v",  r.RemoteAddr, r.FormValue("method"))

			fmt.Fprintln(w, "rpc: hello")

			buf := make([]byte, 10)
			for i := 0; i < len(buf); i++ {
				buf[i] = byte(i)
			}

			_, err = bufrw.Write(buf)
			if err == nil {
				err = bufrw.Flush()
			}
			if err != nil {
				fmt.Printf("ResponseError: %s\\n", err)
			} else {
				fmt.Println("Bye, Jack!")
			}

			logger.Debug("httpConn WriteObj %v", httpConn.GetRemoteIp())
		}
		return
	}

	rpcConn := server.NewWebSocketConn(pConnector.rpcServer, *conn, 128, 45, 2)
	defer func() {
		rpcConn.Close() // 客户端退出减去计数
	}()

	pConnector.rpcServer.ServeConn(rpcConn)
}

func CreateConnectorServerForClient(cfg config.SvrConfig) *Connector {

	db.Init()

	var authCfg config.AuthConfig
	if err := config.ReadConfig("etc/authserver.json", &authCfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
	}
	authConn, err := net.Dial("tcp", authCfg.AuthHost)
	if err != nil {
		logger.Fatal("connect logserver failed %s", err.Error())
	}

	var gsCfg config.LoginConfig
	if err = config.ReadConfig("etc/loginserver.json", &gsCfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
	}
	gsConn, err := net.Dial("tcp", gsCfg.LoginHost)
	if err != nil {
		logger.Fatal("%s", err.Error())
	}

	pConnector = &Connector{
		m:           make(map[uint32]serverInfo),
		authserver: rpc.NewClient(authConn),
		loginserver: rpc.NewClient(gsConn),
		rpcServer:   server.NewServer(),
		players:       make(map[uint64]*Player),
		playersbyid:   make(map[string]*Player),
	}

	//初始化cache
	logger.Info("Init Cache %v", authCfg.MainCacheProfile)
	pConnector.maincache = db.NewCachePool(authCfg.MainCacheProfile)

	pConnector.rpcServer.ApplyProtocol(protobuf.CS_Protocol_value)
	pConnector.rpcServer.Register(pConnector)

	pConnector.rpcServer.RegCallBackOnConn(
	func(conn server.RpcConn) {
		pConnector.onConn(conn)
	},
	)

	pConnector.rpcServer.RegCallBackOnDisConn(
	func(conn server.RpcConn) {
		pConnector.onDisConn(conn)
	},
	)

	pConnector.rpcServer.RegCallBackOnCallBefore(
	func(conn server.RpcConn) {
		conn.Lock()
	},
	)

	pConnector.rpcServer.RegCallBackOnCallAfter(
	func(conn server.RpcConn) {
		conn.Unlock()
	},
	)

	//开始对fightserver的RPC服务
	pConnector.FsMgr.Init(pConnector.rpcServer, cfg)

	listener, err := net.Listen("tcp", cfg.TcpHost)
	if err != nil {
		logger.Fatal("net.Listen: %s", err.Error())
	}

	pConnector.id = cfg.ServerID
	pConnector.listenTcpIp = cfg.TcpHost
	pConnector.listenHttpIp = cfg.HttpHost

	pConnector.sendPlayerCountToGateServer()

	go func() {
		for {
			//For Client/////////////////////////////
			time.Sleep(time.Millisecond * 5)
			conn, err := listener.Accept()

			if err != nil {
				logger.Error("cns StartServices %s", err.Error())
				break
			}

			go func() {
				rpcConn := server.NewTCPSocketConn(pConnector.rpcServer, conn, 128, 45, 1)
				defer func() {
					if r := recover(); r != nil {
						logger.Error("player rpc runtime error begin:", r)
						debug.PrintStack()
						rpcConn.Close()

						logger.Error("player rpc runtime error end ")
					}
				}()

				pConnector.rpcServer.ServeConn(rpcConn)
			}()
		}
	}()

	http.HandleFunc("/", wsServeConnHandler)
	http.ListenAndServe(cfg.HttpHost, nil)

	return pConnector
}

func (self *Connector) sendPlayerCountToGateServer() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Info("sendPlayerCountToGateServer runtime error:", r)

				debug.PrintStack()
			}
		}()

		for {

			time.Sleep(5 * time.Second)

			self.l.RLock()
			playerCount := uint32(len(self.players))
			self.l.RUnlock()

			var ret []byte
			req := protobuf.SL_UpdatePlayerCount{}
			req.SetServerId(self.id)
			req.SetPlayerCount(playerCount)
			req.SetTcpServerIp(self.listenTcpIp)
			req.SetHttpServerIp(self.listenHttpIp)

			//logger.Debug("playerCount %v", playerCount)

			buf,err := proto.Marshal(&req)
			if err != nil {
				logger.Error("Error On Connector.sendPlayerCountToGateServer : %s", err.Error())
				return

			}
			err = self.loginserver.Call("LoginRpcServer.SL_UpdatePlayerCount", &buf, &ret)

			if err != nil {
				logger.Error("Error On LoginRpcServer.SL_UpdatePlayerCount : %s", err.Error())
				return
			}

		}

	}()
}

func (self *Connector) onConn(conn server.RpcConn) {
}

func (self *Connector) onDisConn(conn server.RpcConn) {
	logger.Info("Connector:onDisConn  %v", conn.GetId())
	self.delPlayer(conn.GetId())
}

//添加玩家到全局表中
func (self *Connector) addPlayer(connId uint64, p *Player) {
	logger.Info("Connector:addPlayer %v, %v", connId, p.GetUid())

	self.l.Lock()
	defer self.l.Unlock()

	//进入服务器全局表
	self.players[connId] = p
	self.playersbyid[p.GetUid()] = p
}

//销毁玩家
func (self *Connector) delPlayer(connId uint64) {
	logger.Info("Connector:delPlayer %v", connId)

	p, exist := self.players[connId]
	if exist {
		p.OnQuit()

		self.l.Lock()
		delete(self.players, connId)
		delete(self.playersbyid, p.GetUid())
		self.l.Unlock()
	}
}

func WriteResult(conn server.RpcConn, value interface{}) bool {
	err := conn.WriteObj(value)
	if err != nil {
		logger.Info("WriteResult Error %s", err.Error())
		return false
	}
	return true
}

func (self *Connector) CS_CheckSession(conn server.RpcConn, login protobuf.CS_CheckSession) (err error) {

	rep := protobuf.SC_CheckSessionResult{}
	uid := login.GetUid()
	var rst []byte

	rst, err = redis.Bytes(self.maincache.Do("GET", "SessionKey_" + uid))
	rep.SetResult(protobuf.SC_CheckSessionResult_AUTH_FAILED)
	rep.SetServerTime(uint32(time.Now().Unix()))
	if rst != nil || err == nil{
		if login.GetSessionKey() == string(rst) {
			rep.SetResult(protobuf.SC_CheckSessionResult_OK)
		}
	}

	logger.Debug("SC_CheckSessionResult %v", rep)

	rep.SetResult(rep.GetResult())

	if rep.GetResult() == protobuf.SC_CheckSessionResult_OK {
		WriteResult(conn, &rep)
		if p, ok := self.playersbyid[login.GetUid()]; ok {
			if err := p.conn.Close(); err == nil {
				logger.Info("kick the online player")
			}
		}

		var base protobuf.PlayerBaseInfo
		logger.Info("query db : %v", login.GetUid())
		result, err :=db.Query("playerbase", login.GetUid(), &base)
		if result == false {
			base = protobuf.PlayerBaseInfo{}
			base.SetUid(login.GetUid())

			stat := &protobuf.StatusInfo{}
			stat.SetName("test_" + uid)
			stat.SetLevel(1)

			base.SetStat(stat)
			db.Write("playerbase", login.GetUid(), &base)
			logger.Info("playerbase create %v", login.GetUid())
		}else {
			if err != nil {
				logger.Info("err query db : %v", err)
				return err
			}
			logger.Info("playerbase find")
		}

		p := &Player{PlayerBaseInfo: &base, conn: conn}

		p.SetUid(uid)

		//进入服务器全局表

		self.addPlayer(conn.GetId(), p)

	}else {
		WriteResult(conn, &rep)

		go func() {
			time.Sleep(time.Millisecond * 1000)
			defer func() {
				conn.Close()
			}()
		}()
	}

	return nil
}

func (self *Connector) CS_Ping(conn server.RpcConn, login protobuf.CS_Ping) error {

	rep := protobuf.SC_PingResult{}
	rep.SetServerTime(uint32(time.Now().Unix()))

	WriteResult(conn, &rep)
	return nil
}
