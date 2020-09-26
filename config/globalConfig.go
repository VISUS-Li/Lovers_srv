package config

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var GlobalConfig Config
const USER_SRV_NAME = "lovers.srv.user"
const NOTELIST_SRV_NAME = "lovers.srv.notelist"
const API_NAME = "lovers.api"
const REGISTER_HOST = "127.0.0.1"

const (
	DB_HOST = "127.0.0.1"
	DB_USER = "root"
	DB_PASSWORD = "123456"
)

const(
	JWT_IDKEY = "lovers"
	JWT_EXPIRETIME  = 3
	JWT_SECRET 		= "liningtao"
)

type Config struct{
	//服务名
	Srv_name     string;

	//服务发现注册地址
	RegisterHosts []string

	//数据库
	DB_host     string;
	DB_user     string;
	DB_password string;

	//jwt
	JwtIDKey string;
	ExpireTime int; // token过期时间，单位小时
	JwtSecret string;
}

func Init(){
	GlobalConfig,_= getJsonConfig();
	getDefaultConfig();
}

func getDefaultConfig(){
	//服务发现注册
	if(len(GlobalConfig.RegisterHosts) <= 0){
		GlobalConfig.RegisterHosts = append(GlobalConfig.RegisterHosts, REGISTER_HOST)
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
	if(GlobalConfig.ExpireTime <= 0){
		GlobalConfig.ExpireTime = JWT_EXPIRETIME;
	}
	if(GlobalConfig.JwtSecret == ""){
		GlobalConfig.JwtSecret = JWT_SECRET;
	}
}

func getJsonConfig()(Config, error){
	configPath := GetExeDstFileName("config.json")
	file, err := os.Open(configPath)
	if err != nil {
		logrus.Error("open Config File fail:"+err.Error())
		return Config{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		logrus.Error("read Config File fail:"+err.Error())
		return Config{}, err
	}
	return conf, nil;
}

func GetCurrentExecPath()(string,error){
	file,err := exec.LookPath(os.Args[0])
	if err != nil{
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil{
		return "", err
	}
	if runtime.GOOS !="windows"{
		path = strings.Replace(path, "\\", "/", -1)
	}

	i := strings.LastIndex(path, "/")
	if i < 0{
		return "", errors.New(`Can't find "/" or  "\".`)
	}

	return string(path[0:i+1]), nil
}

/******
传入文件名，得到exe目录下该文件名的全路径
******/
func GetExeDstFileName(dstName string)(string){
	exePath,_ := GetCurrentExecPath()
	path := exePath + "\\" + dstName
	if runtime.GOOS !="windows"{
		path = strings.Replace(path, "\\", "/", -1)
	}
	return path
}



