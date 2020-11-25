package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

var WSConfig WebSocketConfig

const(
	TCP_KEEP_ALIVE = false
	MEM_READ_POOL_COUNT = 1
	WS_LISTEN_ADDR		= ":8081"
)
type WebSocketConfig struct {
	WSListenAddr string
}

func GetWSJsonConfig()(WebSocketConfig, error){
	configPath := GetExeDstFileName("ws_config.json")
	file, err := os.Open(configPath)
	if err != nil {
		logrus.Error("open ws_config File fail:"+err.Error())
		return WebSocketConfig{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := WebSocketConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		logrus.Error("read ws_config File fail:"+err.Error())
		return WebSocketConfig{}, err
	}
	return conf, nil;
}

func GetWSDefaultConfig(){
	if len(WSConfig.WSListenAddr) <= 0 || WSConfig.WSListenAddr == ""{
		WSConfig.WSListenAddr = WS_LISTEN_ADDR
	}
}