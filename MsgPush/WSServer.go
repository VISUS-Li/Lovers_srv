package MsgPush

import (
	"Lovers_srv/MsgPush/Channel"
	"Lovers_srv/MsgPush/SportPush"
	"Lovers_srv/MsgPush/SportPush/SportConst"
	"Lovers_srv/MsgPush/Ticker"
	"Lovers_srv/MsgPush/WSProtocol"
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache/MsgPushCache"
	"Lovers_srv/helper/Utils"
	userClient "Lovers_srv/server/user-service/client"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

var (
	user_clent = userClient.NewUserClient()
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 65536,
	CheckOrigin:func(r *http.Request) bool{
		return true
	},
}

/******
 	Web Socket的其他模块处理
******/
func wsHandler(w http.ResponseWriter, r *http.Request){


	_, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		Utils.ErrorOutputf("Upgrade http to websocket failed:%s",err.Error())
		return
	}

}

/******
 	Web Socket的运动模块处理

	目前的处理是：每一次调用wsSportHandler,都是一个新的跑步连接
******/
func wsSportHandler(w http.ResponseWriter, r *http.Request){
	//获取到用户ID，用以判定是哪个用户的数据
	reqParams := r.URL.Query()
	userId :=reqParams.Get("userId")
	keepSport := reqParams.Get("continue")
	querySportId := reqParams.Get("sportId")
	querySportType := reqParams.Get("sportType")

	if len(userId) <= 0 || userId == "" {
		err := Utils.ErrorOutputf("[SportWSHandle] param userId is nil")
		Utils.HttpCreateErrorBadReq(w, err.Error(),config.CODE_ERR_PARAM_EMPTY)
		return
	}

	//通过userId获取到loverId
	ctx, _ := context.WithCancel(context.Background())
	queryReq := &lovers_srv_user.QueryLoverIdByIdReq{UserId:userId}
	queryResp,err := user_clent.Client_QueryLoverIdById(ctx,queryReq)

	var loverId string
	if queryResp == nil || err != nil || len(queryResp.LoverId) <= 0{
		var errStr string
		if err != nil{
			errStr = err.Error()
		}
		Utils.ErrorOutputf("[SportWSHandle] Query lover Id failed:%s", errStr)
	}else {
		loverId = queryResp.LoverId
	}

	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		err = Utils.ErrorOutputf("[SportWSHandle] Upgrade http to websocket failed:%s",err.Error())
		Utils.HttpCreateErrorInterErr(w, err.Error(),config.CODE_ERR_SERVER_INTERNAL)
		return
	}



	//升级协议完成后，验证用户token
	//_, err = JWTHandler.SportJWTAuth(r)
	sportId := uuid.NewV1()
	wsMsg := new (WSProtocol.Proto)
	if err != nil{
		//验证token失败，退出本次连接
		wsMsg.WriteWSStatusWithOp(conn,WSProtocol.PACK_AUTH_FAIL,"")
		conn.Close()
		return
	}
	////验证成功后，先判断是否是继续运动，再返回Auth结果
	var sportIdStr string
	//判断是否是断线重连运动
	if keepSport == "1"{
		//断线重连，判断参数传的运动guid是否有效
		if querySportId != ""{
			sType,err := strconv.Atoi(querySportType)
			if err != nil{
				err = Utils.ErrorOutputf("param continue invalid")
				wsMsg.WriteWSStatusWithOp(conn,WSProtocol.PACK_CONTINUE_ERR,err.Error())
				return
			}
			switch sType {
			case SportConst.SPORT_TYPE_RUN:
				runItem, code, err :=MsgPushCache.GetRunItemData(userId, querySportId)
				if code != config.ENUM_ERR_OK || runItem == nil{
					wsMsg.WriteWSStatusWithOp(conn,WSProtocol.PACK_CONTINUE_ERR,err.Error())
					return
				}
				break
			}
			sportIdStr = querySportId;
		}
	}else{
		sportIdStr = sportId.String()
	}

	//验证成功，返回本次运动的GUID给客户端
	wsMsg.Op = WSProtocol.PACK_AUTH_SUCC
	wsMsg.Seq = 1
	wsMsg.Body = []byte(sportIdStr)
	wsMsg.WriteWebsocket(conn)

	Utils.InfoOutputf("sportId bytes:%d",sportId.Bytes())
	Utils.InfoOutputf("sportId string:%s",sportId.String())



	//运动模块的处理对象，包含建立的通道
	sportHandle := SportPush.NewSportHandler(conn, userId, loverId, sportIdStr)

	//定时存储统计数据
	go func() {
		for{
			select {
			case <-Ticker.RunTicker.C:
				sportHandle.SaveSportDataToDB(true)
			}
		}
	}()
	//将通道添加到通道管理map中
	Channel.SportChMng.Set(userId, sportHandle.Ch)

	go sportHandle.SportWSHandle()
}

/******
	启动Web Socket服务监听
******/
func StartWSServer(addr string) error {

	//初始化运动模块Channel管理map
	Channel.SportChMng = Channel.NewSportChMng()

	//监听sport模块
	http.HandleFunc("/ws/sport",wsSportHandler)

	//监听其他模块
	http.HandleFunc("/ws", wsHandler)

	Utils.InfoOutputf("Web Socket Server Start")
	err := http.ListenAndServe(addr, nil)
	return err
}