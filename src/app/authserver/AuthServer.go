package auth

import (
	"common/config"
	"common/logger"
	"component/db"
	"component/rpc"
	"net"
	"protobuf"
	"time"
	"common"
)

type AuthServer struct {
	exit        chan bool
}

var pAuthServices *AuthServer

func CreateServices(authcfg config.AuthConfig)  *AuthServer {

	pAuthServices := NewAuthServer(authcfg)

	go func () {
		tsock, err := net.Listen("tcp", authcfg.AuthHost)
		if err != nil {
			logger.Fatal("net.Listen: %s", err.Error())
		}

		StartServices(pAuthServices, tsock)

		WaitForExit(pAuthServices)

		tsock.Close()
	}()

	return pAuthServices
}

func StartServices(self *AuthServer, listener net.Listener) {

	rpcServer := rpc.NewServer()
	rpcServer.Register(self)

	for {
		logger.Info("listener Connect")
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("StartServices %s", err.Error())
			break
		}
		go func() {
			rpcServer.ServeConn(conn)
			conn.Close()
		}()
	}
}

func WaitForExit(self *AuthServer) {
	<-self.exit
	close(self.exit)
}

func NewAuthServer(cfg config.AuthConfig) (server *AuthServer) {

	db.Init()

	server = &AuthServer{
		exit:        make(chan bool),
	}

	// http.Handle("/debug/state", debugHTTP{server})

	return server
}

func (self *AuthServer) LA_CheckAccount(req *protobuf.LA_CheckAccount, ret *protobuf.AL_CheckAccountResult) error {

	uid := common.GenUUID(req.GetAccount())

	if len(req.GetUid()) > 0{
		if req.GetUid() != uid {//客户端伪造uid
			ret.SetResult(protobuf.AL_CheckAccountResult_AUTH_FAILED)
			ret.SetServerTime(uint32(time.Now().Unix()))
			return nil
		}
	}

	account := &protobuf.AccountInfo{}
	result, err :=db.Query("AccountInfo", uid, account)

	if err != nil {
		ret.SetResult(protobuf.AL_CheckAccountResult_SERVERERROR)
		ret.SetServerTime(uint32(time.Now().Unix()))
		return nil
	}

	if result == false {//用户注册

		account.SetUid(uid)
		account.SetAccount(req.GetAccount())
		account.SetPassword(common.GenPassword(req.GetAccount(), req.GetPassword()))
		account.SetLanguage(req.GetLanguage())
		account.SetOption(req.GetOption())
		account.SetSessionKey(common.GenSessionKey())
		account.SetUdid(req.GetUdid())
		account.SetCreateTime(uint32(time.Now().Unix()))

		db.Write("AccountInfo", uid, account)
		logger.Info("Auth AccountInfo create")

	}else {//用户登陆
		if !common.CheckPassword(account.GetPassword(), req.GetAccount(), req.GetPassword()) {
			ret.SetResult(protobuf.AL_CheckAccountResult_AUTH_FAILED)
			ret.SetServerTime(uint32(time.Now().Unix()))
			return nil
		}
		account.SetSessionKey(common.GenSessionKey())//保存进缓存
		db.Write("AccountInfo", uid, account)
		logger.Info("Auth Account find")
	}

	ret.SetResult(protobuf.AL_CheckAccountResult_OK)
	ret.SetServerTime(uint32(time.Now().Unix()))
	ret.SetSessionKey(account.GetSessionKey())
	ret.SetUid(account.GetUid())

	logger.Info("ComeInto AuthServer.Login %v, %v", req, ret)

	return nil
}
