package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/user-service/handler"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)
var USER_SRV_NAME = "lovers.srv.user"

func main(){
	config.Init(USER_SRV_NAME)

	//初始化日志
	myLog := LogHelper.LoversLog{}

	//配置服务名和日志输出文件
	var dbName string
	var serverName string
	if (config.GlobalConfig.Srv_name == "") {
		serverName = USER_SRV_NAME
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	}else{
		serverName = config.GlobalConfig.Srv_name
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	}

	//初始化数据库
	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect(dbName)
	if err != nil{
		logrus.Error("create DB:" + dbName + "error:"+ err.Error())
		return
	}
	defer dbUtil.CloseConnect()

	err = dbUtil.CreateTable(DB.LoginInfo{})
	if err != nil{
		logrus.Error("create table LoginInfo error:"+err.Error())
	}
	err = dbUtil.CreateTable(DB.UserBaseInfo{})
	if err != nil{
		logrus.Error("create table UserBaseInfo error:"+err.Error())
	}

	//初始化Redis缓存
	Cache.NewRedisPool(dbUtil.DB)
	defer Cache.CloseRedis()

	userHandler := handler.UserHandler{dbUtil.DB}

	//注册中心为consul
	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = config.GlobalConfig.RegisterHosts
	//})

	//新建serivce
	service := micro.NewService(
			micro.Name(serverName),
			//micro.Registry(reg),
		)

	service.Init()

	err = lovers_srv_user.RegisterUserHandler(service.Server(), &userHandler)

	if err = service.Run(); err != nil{
		logrus.Error("service Run error,msg:"+ err.Error())
	}

}