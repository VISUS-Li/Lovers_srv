package main

import (
	"Lovers_Micro_Test/api/handler"
	"github.com/micro/go-micro/web"
)

const SRV_NAME = "shop.srv.apigateway"
func main(){
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

	}

	//开始运行Web服务
	if err := webSrv.Run(); err != nil{

	}

}
