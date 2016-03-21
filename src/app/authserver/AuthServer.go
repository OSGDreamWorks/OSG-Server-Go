package auth

import (
	"common/config"
	"common/logger"
	"component/db"
	"component/rpc"
	"protobuf"
	"common"
	"time"
)

type AuthServer struct {
	maincache    *db.CachePool
	rpcServer	 *rpc.Server
	exit        chan bool
}

var pAuthServices *AuthServer

func CreateServices(authcfg config.AuthConfig)  *AuthServer {

	//初始化db
	logger.Info("Init DB")
	db.Init()

	//初始化cache
	var cacheCfg config.CacheConfig
	if err := config.ReadConfig("etc/maincache.json", &cacheCfg); err != nil {
		logger.Fatal("load config failed, error is: %v", err)
	}
	logger.Info("Init Cache %v", cacheCfg)

	pAuthServices = &AuthServer{
		maincache : db.NewCachePool(cacheCfg),
		rpcServer : rpc.NewServer(),
		exit:        make(chan bool),
	}

	pAuthServices.rpcServer.ApplyProtocol(protobuf.LA_Protocol_value)
	pAuthServices.rpcServer.Register(pAuthServices)

	pAuthServices.rpcServer.ListenAndServe(authcfg.AuthHost, nil)

	return pAuthServices
}

func WaitForExit(self *AuthServer) {
	<-self.exit
	close(self.exit)
}

func (self *AuthServer) LA_CheckAccount(req *protobuf.LA_CheckAccount, ret *protobuf.AL_CheckAccountResult) error {

	uid := common.GenUUID(req.Account)

	if len(req.Uid) > 0{
		if req.Uid != uid {//客户端伪造uid
			ret.Result = protobuf.AL_CheckAccountResult_AUTH_FAILED
			return nil
		}
	}

	account := &protobuf.AccountInfo{}
	result, err :=db.Query("AccountInfo", uid, account)

	if err != nil {
		ret.Result = protobuf.AL_CheckAccountResult_SERVERERROR
		return nil
	}

	if result == false {//用户注册

		account.Uid = uid
		account.Account = req.Account
		account.Password = common.GenPassword(req.Account, req.Password)
		account.Language = req.Language
		account.Option = req.Option
		account.SessionKey = common.GenSessionKey()
		account.Udid = req.Udid
		account.CreateTime = uint32(time.Now().Unix())

		db.Write("AccountInfo", uid, account)
		logger.Info("Auth AccountInfo create")

	}else {//用户登陆
		if !common.CheckPassword(account.Password, req.Account, req.Password) {
			ret.Result = protobuf.AL_CheckAccountResult_AUTH_FAILED
			return nil
		}
		account.SessionKey = common.GenSessionKey()//保存进缓存
		db.Write("AccountInfo", uid, account)
		logger.Info("Auth Account find")
	}

	self.maincache.Do("SET", "SessionKey_" + uid, []byte(account.SessionKey))

	ret.Result = protobuf.AL_CheckAccountResult_OK
	ret.SessionKey = account.SessionKey
	ret.Uid = account.Uid

	logger.Info("ComeInto AuthServer.Login %v, %v", req, ret)

	return nil
}
