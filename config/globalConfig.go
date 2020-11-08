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
	"time"
)

var GlobalConfig Config
const USER_SRV_NAME = "lovers.srv.user"
const NOTELIST_SRV_NAME = "lovers.srv.notelist"
const FILE_SRV_NAME		= "lovers.srv.file"
const API_NAME = "lovers.api"
const REGISTER_HOST = "127.0.0.1"

//运行模式，开发、生产
const (
	RUNMODE_DEF = "dev"
	RUNMODE_DEV = "dev" //开发模式
	RUNMODE_PRO = "pro" //生产模式
)

const (
	DB_HOST = "127.0.0.1"
	DB_USER = "root"
	DB_PASSWORD = "123456"
)

const (
	REDIS_NETWORK = "tcp"
	REDIS_ADDR = "127.0.0.1:6379"
	REDIS_PWD = ""
	REDIS_EXPIRETIME = "30m"
	REDIS_MAXIDLE = 1024
	REDIS_MAXACTIVE = 60000
	REDIS_IDLETIMEOUT = "120s"
	REDIS_DIALTIMEOUT = "200ms"
	REDIS_READTIMEOUT = "5s"
	REDIS_WRITETIMEOUT = "5s"
)

const(
	JWT_IDKEY = "lovers"
	JWT_EXPIRETIME  = 3
	JWT_SECRET 		= "liningtao"
)

//主页
const(
	DEFAULT_CARD_COUNT = 10
)
type Config struct{
	//运行模式
	RunMode string;
	//服务名
	Srv_name     string;

	//服务发现注册地址
	RegisterHosts []string

	//数据库
	DB_host     string;
	DB_user     string;
	DB_password string;

	//jwt
	JwtIDKey      string;
	JwtExpireTime int; // token过期时间，单位小时
	JwtSecret     string;

	//主页
	DefaultCardCount int; //默认获取卡片数量

	//Redis
	Redis_NetWork string;
	Redis_Addr string;
	Redis_Pwd string
	Redis_MaxIdle int;
	Redis_MaxActive int;
	redis_IdleTimeoutStr string;
	Redis_IdleTimeout time.Duration;
	redis_DialTimeoutStr string;
	Redis_DialTimeout time.Duration;
	redis_ReadTimeoutStr string;
	Redis_ReadTimeout time.Duration;
	redis_WriteTimeout string;
	Redis_WriteTimeout time.Duration;
	redis_ExpireTimeStr string;
	Redis_ExpireTime time.Duration;
}

func Init(){
	GlobalConfig,_= getJsonConfig();
	getDefaultConfig();
}

func getDefaultConfig(){
	//运行模式
	if(GlobalConfig.RunMode == ""){
		GlobalConfig.RunMode = RUNMODE_DEF;
	}

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
	if(GlobalConfig.JwtExpireTime <= 0){
		GlobalConfig.JwtExpireTime = JWT_EXPIRETIME;
	}
	if(GlobalConfig.JwtSecret == ""){
		GlobalConfig.JwtSecret = JWT_SECRET;
	}

	//主页
	if(GlobalConfig.DefaultCardCount <= 0){
		GlobalConfig.DefaultCardCount = DEFAULT_CARD_COUNT;
	}

	//Redis
	if(GlobalConfig.Redis_Addr == ""){
		GlobalConfig.Redis_Addr = REDIS_ADDR;
	}
	if(GlobalConfig.Redis_MaxIdle <= 0){
		GlobalConfig.Redis_MaxIdle = REDIS_MAXIDLE;
	}
	if(GlobalConfig.Redis_MaxActive <= 0){
		GlobalConfig.Redis_MaxActive = REDIS_MAXACTIVE;
	}
	if(GlobalConfig.Redis_NetWork == ""){
		GlobalConfig.Redis_NetWork = REDIS_NETWORK;
	}
	if(GlobalConfig.Redis_Pwd == ""){
		GlobalConfig.Redis_Pwd = REDIS_PWD
	}

	UnMarshDuration(&GlobalConfig.Redis_IdleTimeout, &GlobalConfig.redis_IdleTimeoutStr, REDIS_IDLETIMEOUT);

	UnMarshDuration(&GlobalConfig.Redis_DialTimeout, &GlobalConfig.redis_DialTimeoutStr, REDIS_DIALTIMEOUT);

	UnMarshDuration(&GlobalConfig.Redis_ReadTimeout, &GlobalConfig.redis_ReadTimeoutStr, REDIS_READTIMEOUT);

	UnMarshDuration(&GlobalConfig.Redis_WriteTimeout, &GlobalConfig.redis_WriteTimeout, REDIS_WRITETIMEOUT);

	UnMarshDuration(&GlobalConfig.Redis_ExpireTime, &GlobalConfig.redis_ExpireTimeStr, REDIS_EXPIRETIME);
}

func UnMarshDuration(target *time.Duration, confStr *string, defaultStr string){
	var tmpConfStr string
	if(confStr == nil){
		tmpConfStr = defaultStr;
	}else if *confStr == ""{
		tmpConfStr = defaultStr;
		*confStr = defaultStr;
	}else{
		tmpConfStr = *confStr;
	}

	d,err := time.ParseDuration(tmpConfStr)
	if err != nil{
		logrus.Error("ParseDuration Error:"+err.Error())
		*target,_ = time.ParseDuration(defaultStr)
	}else{
		*target = d
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



