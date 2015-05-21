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

func (self *AuthServer) Login(req *protobuf.Login, ret *protobuf.LoginResult) error {

	uid := common.GenUUID(req.GetAccount())

	if len(req.GetUid()) > 0{
		if req.GetUid() != uid {//客户端伪造uid
			ret.SetResult(protobuf.LoginResult_AUTH_FAILED)
			ret.SetServerTime(uint32(time.Now().Unix()))
			return nil
		}
	}

	info := &protobuf.PlayerInfo{}
	result, err :=db.Query("playerinfo", uid, info)

	if err != nil {
		ret.SetResult(protobuf.LoginResult_SERVERERROR)
		ret.SetServerTime(uint32(time.Now().Unix()))
		return nil
	}

	if result == false {//用户注册

		account := &protobuf.Login{}
		account.SetAccount(req.GetAccount())
		account.SetPassword(common.GenPassword(req.GetAccount(), req.GetPassword()))
		account.SetLanguage(req.GetLanguage())
		account.SetOption(req.GetOption())
		account.SetSessionKey(common.GenSessionKey())
		account.SetUdid(req.GetUdid())
		account.SetCreateTime(uint32(time.Now().Unix()))

		info.SetUid(uid)
		info.SetAccount(account)

		db.Write("playerinfo", uid, info)
		logger.Info("playerinfo create")
	}else {//用户登陆

		account := info.GetAccount()

		if !common.CheckPassword(account.GetPassword(), req.GetAccount(), req.GetPassword()) {
			ret.SetResult(protobuf.LoginResult_AUTH_FAILED)
			ret.SetServerTime(uint32(time.Now().Unix()))
			return nil
		}
		account.SetSessionKey(common.GenSessionKey())

		info.SetAccount(account)

		db.Write("playerinfo", uid, info)
		logger.Info("playerinfo find")
	}

	ret.SetResult(protobuf.LoginResult_OK)
	ret.SetServerTime(uint32(time.Now().Unix()))
	account := info.GetAccount()
	if account != nil {
		ret.SetSessionKey(account.GetSessionKey())
	}
	ret.SetUid(info.GetUid())

	logger.Info("ComeInto AuthServer.Login %v, %v", req, ret)

	return nil
}
