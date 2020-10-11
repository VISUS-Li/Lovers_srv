package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/fileserver/handler"
	"github.com/micro/go-micro"

	lovers_srv_file "Lovers_srv/server/fileserver/proto"
	"github.com/sirupsen/logrus"

)

var FILE_SRV_NAME = "lovers.srv.file"

func main() {
	config.Init()
	//初始化日志
	myLog := LogHelper.LoversLog{}

	var dbName string
	var serverName string
	if (config.GlobalConfig.Srv_name == "") {
		serverName = FILE_SRV_NAME
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	} else {
		serverName = config.GlobalConfig.Srv_name
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	}

	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect(dbName)
	if err != nil {
		logrus.Error("create DB:" + dbName + "error:" + err.Error())
		return;
	}
	defer dbUtil.CloseConnect()

	err = dbUtil.CreateTable(DB.FileServerInfo{})
	if err != nil {
		logrus.Error("create table FileServerInfo error" + err.Error())
	}

	var fileHandler = handler.FileHandler{dbUtil.DB}
	service := micro.NewService(
		micro.Name(serverName),
	)

	service.Init()

	err = lovers_srv_file.RegisterFileServerHandler(service.Server(), &fileHandler)

	if err = service.Run(); err != nil {
		logrus.Error("service Run error, msg:" + err.Error())
	}

}
