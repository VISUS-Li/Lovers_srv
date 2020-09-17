package config

import (
	"Lovers_srv/helper/Utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

var GlobalConfig Config
const USER_SRV_NAME = "lovers.srv.user"
const NOTELIST_SRV_NAME = "lovers.srv.notelist"
const API_NAME = "lovers.api"

const (
	DB_HOST = "127.0.0.1"
	DB_USER = "root"
	DB_PASSWORD = "123456"
)

const(
	JWT_IDKEY = "lovers"
)

type Config struct{
	//服务名
	User_srv_name     string;
	Notelist_srv_name string;
	Api_name          string;

	//数据库
	DB_host     string;
	DB_user     string;
	DB_password string;

	//jwt
	JwtIDKey string
}

func Init(){
	GlobalConfig,_= getJsonConfig();
	getDefaultConfig();
}

func getDefaultConfig(){

	//服务名
	if(GlobalConfig.User_srv_name == ""){
		GlobalConfig.User_srv_name = USER_SRV_NAME;
	}
	if(GlobalConfig.Notelist_srv_name == ""){
		GlobalConfig.Notelist_srv_name = NOTELIST_SRV_NAME;
	}
	if(GlobalConfig.Api_name == ""){
		GlobalConfig.Api_name = API_NAME;
	}

	//数据库
	if(GlobalConfig.DB_host == ""){
		GlobalConfig.DB_host = DB_HOST;
	}
	if(GlobalConfig.DB_user == ""){
		GlobalConfig.DB_user = DB_USER;
	}
	if(GlobalConfig.DB_password == ""){
		GlobalConfig.DB_password = DB_PASSWORD;
	}

	//JWT
	if(GlobalConfig.JwtIDKey == ""){
		GlobalConfig.JwtIDKey = JWT_IDKEY;
	}
}

func getJsonConfig()(Config, error){
	configPath := Utils.GetExeDstFileName("config.json")
	file, _ := os.Open(configPath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		logrus.Error("read Config File fail:"+err.Error())
		return Config{}, err
	}
	return conf, nil;
}

