package db

import (
	gp "code.google.com/p/goprotobuf/proto"
	"common/config"
	"common/logger"
	"component/rpc"
)

var pPollBase *rpc.ClientPool

func Init() {
	//base
	var dbCfg config.DBConfig
	if err := config.ReadConfig("etc/dbBase.json", &dbCfg); err != nil {
		logger.Fatal("%v", err)
	}

	aHosts := make([]string, 0)
	aHosts = append(aHosts, dbCfg.DBHost)
	pPollBase = rpc.CreateClientPool(aHosts)
	if pPollBase == nil {
		logger.Fatal("create failed")
	}
}

//基础信息库
func Query(table, uid string, value gp.Message) (exist bool, err error) {
	err, conn := pPollBase.RandomGetConn()
	if err != nil {
		return
	}

	return KVQuery(conn, table, uid, value)
}

func Write(table, uid string, value gp.Message) (result bool, err error) {
	err, conn := pPollBase.RandomGetConn()
	if err != nil {
		return
	}

	return KVWrite(conn, table, uid, value)
}

func Delete(table, uid string) (exist bool, err error) {
	err, conn := pPollBase.RandomGetConn()
	if err != nil {
		return
	}

	return KVDelete(conn, table, uid)
}
