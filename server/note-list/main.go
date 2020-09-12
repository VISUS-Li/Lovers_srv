package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/server/note-list/handler"
	lovers_srv_user "Lovers_srv/server/note-list/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

func main() {
	myLog := LogHelper.LoversLog{}
	myLog.SetOutPut(config.NOTELIST_SRV_NAME)

	service := micro.NewService(
		micro.Name(config.NOTELIST_SRV_NAME),
	)

	notelistHandler := handler.NoteListHandler{}

	service.Init()

	err := lovers_srv_user.RegisterNoteListHandler(service.Server(), &notelistHandler)

	if err = service.Run(); err != nil {
		logrus.Error("notelist service Run error, msg:" + err.Error())
	}
}

