package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/lovers-tools/note-list/handler"
	lovers_srv_user "Lovers_srv/server/lovers-tools/note-list/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

var NOTELIST_SRV_NAME = "lovers.srv.notelist"

func main() {
	config.Init(NOTELIST_SRV_NAME)
	//create log
	myLog := LogHelper.LoversLog{}
	var dbName string
	var serverName string
	if (config.GlobalConfig.Srv_name == "") {
		serverName = NOTELIST_SRV_NAME
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	} else {
		serverName = config.GlobalConfig.Srv_name
		dbName = Utils.GetDBNameFromSrvName(serverName)
		myLog.SetOutPut(serverName)
	}
	//create database
	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect(dbName)

	if err != nil {
		logrus.Error("NoteList database create failed, msg:" +err.Error())
		return
	}
	defer dbUtil.CloseConnect()

	err = dbUtil.CreateTable(DB.NoteListDB{})
	if err != nil {
		logrus.Error("create table NoteListDB error:"+err.Error())
		return
	}

	Cache.NewRedisPool(dbUtil.DB)
	defer Cache.CloseRedis()

	noteListHandler := handler.NoteListHandler{dbUtil.DB}

	//注册中心为consul
	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = config.GlobalConfig.RegisterHosts
	//})

	service := micro.NewService(
		micro.Name(serverName),
	)

	service.Init()

	err = lovers_srv_user.RegisterNoteListHandler(service.Server(), &noteListHandler)

	if err = service.Run(); err != nil {
		logrus.Error("NoteList service Run error, msg:" + err.Error())
	}
}

