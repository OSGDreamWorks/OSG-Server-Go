package connector

import (
	"fmt"
	"net"
	"runtime/debug"
	"sync"
	"component/server"
	"common/logger"
	"common/config"
	"protobuf"
)

type FServerConnMgr struct {
	poollock  sync.RWMutex
	connpool  []server.RpcConn
	poolsize  uint8
	workindex int8
	poolid    uint8
	fightjob  chan interface{}
	quit      chan bool
}

func (self *FServerConnMgr) GetConn() server.RpcConn {
	return self.connpool[self.workindex]
}

func (self *FServerConnMgr) Open(poolsize uint8) {
	self.poolid = 0
	self.workindex = -1
	self.poolsize = poolsize
	self.connpool = make([]server.RpcConn, poolsize)
	self.quit = make(chan bool)

}

func (self *FServerConnMgr) GetNewConnId() uint8 {
	self.poollock.Lock()
	defer self.poollock.Unlock()

	self.poolid++
	return self.poolid - 1
}

func (self *FServerConnMgr) NewConn(conn server.RpcConn, uConnId uint8) {
	self.poollock.Lock()
	defer self.poollock.Unlock()

	self.connpool[uConnId] = conn
}

func (self *FServerConnMgr) GetWorkConn() server.RpcConn {
	self.poollock.Lock()
	defer self.poollock.Unlock()

	self.workindex++
	if uint8(self.workindex) >= self.poolsize {
		self.workindex = 0
	}
	logger.Info("FServerConnMgr GetWorkConn -----> %v", self.workindex)
	return self.connpool[self.workindex]
}

func (self *FServerConnMgr) Call(serviceMethod protobuf.Network_Protocol, arg interface{}) error {
	logger.Info("Call FServerConnMgr -----> (%v) %v", serviceMethod, arg)
	return self.GetWorkConn().Call(serviceMethod, arg)
}

func (self *FServerConnMgr) Quit() {
	self.poollock.Lock()
	defer self.poollock.Unlock()
	for i, v := range self.connpool {

		logger.Info("ShutDown FServerConnMgr -----> %d", i)
		v.Close()
		self.quit <- true

	}
}

func (self *FServerConnMgr) Init(connnector *server.Server, cfg config.SvrConfig) {

	fsCount := len(cfg.FsHost)
	self.Open(uint8(fsCount))

	for i := 0; i < fsCount; i++ {
		go func() {

			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("FServerConnMgr runtime error:", r)

					debug.PrintStack()
				}
			}()

			connId := self.GetNewConnId()
			host := cfg.FsHost[connId]

			for {
				select {
				case <-self.quit:
					{
						logger.Info("FServerConnMgr Goroutine Quit ----->")
						return
					}
				default:
					{
						var err error
						var fsConn net.Conn

						for {
							fsConn, err = net.Dial("tcp", host)
							if err != nil {
								//logger.Fatal("Connect FightServer Error :%s", err.Error())
							} else {
								break
							}
						}

						logger.Info("Connect to FightServer : %s ok!!!!", host)

						fsRpcConn := server.NewTCPSocketConn(connnector, fsConn, 1000, 0, 1)
						fsRpcConn.SetResultServer("FightServer")
						self.NewConn(fsRpcConn, connId)

						connnector.ServeConn(fsRpcConn)
					}
				}
			}
		}()

	}
}
