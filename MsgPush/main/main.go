package main

import (
	"Lovers_srv/MsgPush"
	"Lovers_srv/MsgPush/Ticker"
	"Lovers_srv/MsgPush/handler"
	lovers_srv_msg_push "Lovers_srv/MsgPush/proto"
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/LogHelper"
	"Lovers_srv/helper/Utils"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)
var MSG_PUSH_SRV_NAME = "lovers.srv.msg_push"
var MSG_PUSH_DB_NAME  = "msg_push"
func main(){
	config.Init(MSG_PUSH_SRV_NAME)

	//初始化日志
	myLog := LogHelper.LoversLog{}

	//配置服务名和日志输出文件
	var serverName string
	if (config.GlobalConfig.Srv_name == "") {
		serverName = MSG_PUSH_SRV_NAME
		myLog.SetOutPut(serverName)
		config.GlobalConfig.Srv_name = MSG_PUSH_SRV_NAME
	}else{
		serverName = config.GlobalConfig.Srv_name
		myLog.SetOutPut(serverName)
	}

	//初始化数据库
	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect(MSG_PUSH_DB_NAME)
	if err != nil{
		panic(Utils.ErrorOutputf("create DB: %S error:", MSG_PUSH_DB_NAME, err.Error()))
	}
	defer dbUtil.CloseConnect()
	CreateAllTables(&dbUtil);
	//初始化Redis缓存
	Cache.NewRedisPool(dbUtil.DB)
	defer Cache.CloseRedis()
	mpHandler := handler.MsgPushHandler{dbUtil.DB}

	//初始化定时器
	Ticker.NewRunTicker()
	defer Ticker.StopTicker()

	//启动Web Socket 服务
	go MsgPush.StartWSServer(config.WSConfig.WSListenAddr)

	//注册中心为consul
	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = config.GlobalConfig.RegisterHosts
	//})

	//新建serivce
	service := micro.NewService(
		micro.Name(serverName),
		//micro.Registry(reg),
	)

	service.Init()

	err = lovers_srv_msg_push.RegisterMSG_PUSHHandler(service.Server(), &mpHandler)

	if err = service.Run(); err != nil{
		logrus.Error("service Run error,msg:"+ err.Error())
	}
}

func CreateAllTables(dbUtil *DB.DBUtil){

	err := dbUtil.CreateTable(DB.RunItemData{})
	if err != nil{
		Utils.ErrorOutputf("[CreateAllTables] create table RunItemData error:%",err.Error())
	}

	err = dbUtil.CreateTable(DB.RunStatistics{})
	if err != nil{
		Utils.ErrorOutputf("[CreateAllTables] create table RunItemData error:%",err.Error())
	}
}