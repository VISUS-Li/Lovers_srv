package main

import (
	"Lovers_srv/api/handler"
	"Lovers_srv/config"
	"Lovers_srv/helper/LogHelper"
	"github.com/micro/go-micro/web"
	"github.com/sirupsen/logrus"
)
func main(){
	//从配置文件获取配置信息
	config.Init()

	//初始化日志
	myLog := LogHelper.LoversLog{}
	var serverName string
	if (config.GlobalConfig.Srv_name == "") {
		serverName = config.API_NAME;
	}else{
		serverName = config.GlobalConfig.Srv_name;
	}
	myLog.SetOutPut(serverName)
	//service := micro.NewService()
	//新建Web服务
	webSrv := web.NewService(
		web.Name(serverName),
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
