package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	homeHandler "Lovers_srv/server/home-service/handler"
	lovers_srv_home "Lovers_srv/server/home-service/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)
var HOME_SRV_NAME = "lovers.srv.home"

func main(){
	config.Init(HOME_SRV_NAME)
	//初始化日志
	myLog := LogHelper.LoversLog{}
	var dbName string
	var serverName string
	if (config.GlobalConfig.Srv_name == "") {
		serverName = HOME_SRV_NAME
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	}else{
		serverName = config.GlobalConfig.Srv_name
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	}
	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect(dbName)
	if err != nil{
		logrus.Error("create DB:" + dbName + "error:"+ err.Error())
		return
	}
	defer dbUtil.CloseConnect()

	err = dbUtil.CreateTable(DB.HomeCardInfo{})
	if err != nil{
		logrus.Error("create table HomeCardInfo error:"+err.Error())
	}

	homeHandler := homeHandler.HomeHandler{dbUtil.DB}

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

	err = lovers_srv_home.RegisterHomeHandler(service.Server(), &homeHandler)

	if err = service.Run(); err != nil{
		logrus.Error("service Run error,msg:"+ err.Error())
	}

}