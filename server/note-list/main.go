package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/note-list/handler"
	lovers_srv_user "Lovers_srv/server/note-list/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

func main() {
	//create log
	myLog := LogHelper.LoversLog{}
	var dbName string
	if (config.GlobalConfig.Srv_name == "") {
		dbName = Utils.GetDBNameFromSrvName(config.NOTELIST_SRV_NAME)
		myLog.SetOutPut(config.NOTELIST_SRV_NAME)
	} else {
		dbName = Utils.GetDBNameFromSrvName(config.NOTELIST_SRV_NAME)
		myLog.SetOutPut(config.NOTELIST_SRV_NAME)
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

	noteListHandler := handler.NoteListHandler{dbUtil.DB}

	//注册中心为consul
	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = config.GlobalConfig.RegisterHosts
	//})

	service := micro.NewService(
		micro.Name(config.NOTELIST_SRV_NAME),
		//micro.Registry(reg),
	)

	service.Init()

	err = lovers_srv_user.RegisterNoteListHandler(service.Server(), &noteListHandler)

	if err = service.Run(); err != nil {
		logrus.Error("NoteList service Run error, msg:" + err.Error())
	}
}

