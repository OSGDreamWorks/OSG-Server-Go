package config

import (
	"common/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type MySQLConfig struct {
	Host        string  `json:"host"`
	Port        uint16  `json:"port"`
	Uname       string  `json:"uname"`
	Pass        string  `json:"pass"`
	Vnode       uint8   `json:"vnode"`
	NodeName    string  `json:"nodename"`
	Dbname      string  `json:"dbname"`
	Charset     string  `json:"charset"`
	PoolSize    uint16  `json:"pool"`
	IdleTimeOut float64 `json:"idle"`
	MaxRetry    uint8   `json:"retry"`
}

type CacheConfig struct {
	Host        string  `json:"host"`
	Port        uint16  `json:"port"`
	Index       uint8   `json:"index"`
	Vnode       uint8   `json:"vnode"`
	NodeName    string  `json:"nodename"`
	PoolSize    uint16  `json:"pool"`
	IdleTimeOut float64 `json:"idle"`
	MaxRetry    uint8   `json:"retry"`
}

type TableConfig struct {
	DBProfile    string `json:"db-profile"`
	CacheProfile string `json:"cache-profile"`
	DeleteExpiry uint64 `json:"expiry"`
}

type DBConfig struct {
	DBHost        string
	DebugHost     string
	GcTime        uint8
	DBProfiles    map[string][]MySQLConfig `json:"database"`
	CacheProfiles map[string][]CacheConfig `json:"cache"`
	Tables        map[string]TableConfig   `json:"tables"`
}

type AuthConfig struct {
	AuthHost        string
	DebugHost     string
	GcTime        uint8
	MainCacheProfile CacheConfig `json:"maincache"`
}

type SvrConfig struct {
	ServerID	uint32
	TcpHost        string
	HttpHost        string
	DebugHost     string
	GcTime        uint8
	FsHost        []string
}

type GateConfig struct {
	GateHost        string
	TcpHostForClient     string
	HttpHostForClient    string
	DebugHost     string
	GcTime        uint8
}

//读取配置表
func ReadConfig(file string, cfg interface{}) error {
	cfgpath, _ := os.Getwd()

	if err := ReadJson(path.Join(cfgpath, file), cfg); err != nil {
		logger.Fatal("read config failed, %v", err)
		return err
	}

	return nil
}

func ToJson(val interface{}) string {
	data, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func WriteJson(filename string, val interface{}) error {
	data, err := json.MarshalIndent(val, "  ", "  ")
	if err != nil {
		return fmt.Errorf("WriteJson failed: %v %v", filename, err)
	}
	return ioutil.WriteFile(filename, data, 0660)
}

func ReadJson(filename string, val interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("ReadJson failed: %T %v", val, err)
	}
	if err = json.Unmarshal(data, val); err != nil {
		return fmt.Errorf("ReadJson failed: %T %v %v", val, filename, err)
	}
	return nil
}

func ReadJsonByDataStr(jDataStr string, val interface{}) error {
	data := []byte(jDataStr)
	if err := json.Unmarshal(data, val); err != nil {
		return fmt.Errorf("ReadJson failed: %T %v %v", val, jDataStr, err)
	}
	return nil
}
