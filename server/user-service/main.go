package main

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/server/user-service/handler"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

func main(){
	myLog := LogHelper.LoversLog{}
	myLog.SetOutPut(config.USER_SRV_NAME)
	db_util := DB.DBUtil{}
	err := db_util.CreateConnect()
	defer db_util.CloseConnect()
	if err != nil{
	}
	userHandler := handler.UserHandler{db_util.DB}

	//新建serivce
	service := micro.NewService(
			micro.Name(config.USER_SRV_NAME),
		)

	service.Init()

	err = lovers_srv_user.RegisterUserHandler(service.Server(), &userHandler)

	logrus.Info("run Service ...!")
	if err = service.Run(); err != nil{
		logrus.Error("service Run error,msg:"+ err.Error())
	}

}