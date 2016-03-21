package db

import (
	gp "github.com/golang/protobuf/proto"
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

	logger.Debug("Init DBClient : %v ", dbCfg.DBHost)
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

//基础信息库二进制查询
func QueryBinary(table, uid string, value *[]byte) (exist bool, err error) {
	err, conn := pPollBase.RandomGetConn()
	if err != nil {
		return
	}
	exist, err = KVQuery(conn, table, uid, value)
	return exist, err
}

func WriteBinary(table, uid string, value []byte) (result bool, err error) {
	err, conn := pPollBase.RandomGetConn()
	if err != nil {
		return
	}

	return KVWrite(conn, table, uid, value)
}

func DeleteBinary(table, uid string) (exist bool, err error) {
	err, conn := pPollBase.RandomGetConn()
	if err != nil {
		return
	}

	return KVDelete(conn, table, uid)
}
