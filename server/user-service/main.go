package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/user-service/handler"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

func main(){
	config.Init()
	//初始化日志
	myLog := LogHelper.LoversLog{}
	var dbName string
	if (config.GlobalConfig.Srv_name == "") {
		dbName = Utils.GetDBNameFromSrvName(config.USER_SRV_NAME)
		myLog.SetOutPut(config.USER_SRV_NAME)
	}else{
		dbName = Utils.GetDBNameFromSrvName(config.USER_SRV_NAME)
		myLog.SetOutPut(config.GlobalConfig.Srv_name)
	}
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

	userHandler := handler.UserHandler{dbUtil.DB}

	//新建serivce
	service := micro.NewService(
			micro.Name(config.USER_SRV_NAME),
		)

	service.Init()

	err = lovers_srv_user.RegisterUserHandler(service.Server(), &userHandler)

	if err = service.Run(); err != nil{
		logrus.Error("service Run error,msg:"+ err.Error())
	}

}