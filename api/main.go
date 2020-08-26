package main

import (
	"Lovers_srv/api/handler"
	"Lovers_srv/config"
	"Lovers_srv/helper/LogHelper"
	"github.com/micro/go-micro/web"
	"github.com/sirupsen/logrus"
)

const SRV_NAME = "shop.srv.apigateway"
func main(){
	myLog := LogHelper.LoversLog{}
	myLog.SetOutPut(config.API_NAME)
	//service := micro.NewService()
	//新建Web服务
	webSrv := web.NewService(
		web.Name(SRV_NAME),
		web.Address(":20050"),
		)

	//构造Gin的Engine
	router := handler.ClientEngine()
	//注册Web服务的处理Engine
	webSrv.Handle("/",router)

	//初始化Web服务
	if err := webSrv.Init(); err != nil{
		logrus.Error("go-micro web server init error:"+err.Error())
	}

	//开始运行Web服务
	if err := webSrv.Run(); err != nil{
		logrus.Error("go-micro web server run error:"+err.Error())
	}

}
