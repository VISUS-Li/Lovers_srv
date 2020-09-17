package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/server/note-list/handler"
	lovers_srv_user "Lovers_srv/server/note-list/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

func main() {
	//create log
	myLog := LogHelper.LoversLog{}
	myLog.SetOutPut(config.NOTELIST_SRV_NAME)
	//create database
	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect()
	if err != nil {
		logrus.Error("NoteList database create failed, msg:" +err.Error())
		return
	}

	err = dbUtil.CreateTable(DB.NoteListDB{})
	if err != nil {
		logrus.Error("create table NoteListDB error:"+err.Error())
		return
	}
	defer dbUtil.CloseConnect()

	noteListHandler := handler.NoteListHandler{dbUtil.DB}

	service := micro.NewService(
		micro.Name(config.NOTELIST_SRV_NAME),
	)

	service.Init()

	err = lovers_srv_user.RegisterNoteListHandler(service.Server(), &noteListHandler)

	if err = service.Run(); err != nil {
		logrus.Error("NoteList service Run error, msg:" + err.Error())
	}
}

