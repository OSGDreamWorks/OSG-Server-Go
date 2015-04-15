package auth

import (
	"common/config"
	"common/logger"
	"component/db"
	"component/rpc"
	"net"
	"common/protobuf"
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

	info := &protobuf.Login{}
	result, err :=db.Query("playerinfo", uid, info)

	if err != nil {
		ret.SetResult(protobuf.LoginResult_SERVERERROR)
		ret.SetServerTime(uint32(time.Now().Unix()))
		return nil
	}

	if result == false {//用户注册
		info.SetUid(uid)
		info.SetAccount(req.GetAccount())
		info.SetPassword(common.GenPassword(req.GetAccount(), req.GetPassword()))
		info.SetLanguage(req.GetLanguage())
		info.SetOption(req.GetOption())
		info.SetSessionKey(common.GenSessionKey())
		info.SetUdid(req.GetUdid())
		info.SetCreateTime(uint32(time.Now().Unix()))
		db.Write("playerinfo", uid, info)
		logger.Info("playerinfo create")
	}else {//用户登陆
		if !common.CheckPassword(info.GetPassword(), req.GetAccount(), req.GetPassword()) {
			ret.SetResult(protobuf.LoginResult_AUTH_FAILED)
			ret.SetServerTime(uint32(time.Now().Unix()))
			return nil
		}
		info.SetSessionKey(common.GenSessionKey())
		db.Write("playerinfo", uid, info)
		logger.Info("playerinfo find")
	}

	ret.SetResult(protobuf.LoginResult_OK)
	ret.SetServerTime(uint32(time.Now().Unix()))
	ret.SetSessionKey(info.GetSessionKey())
	ret.SetUid(info.GetUid())

	logger.Info("ComeInto AuthServer.Login %v, %v", req, ret)

	return nil
}
