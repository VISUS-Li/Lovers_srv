package SportPush

import (
	"Lovers_srv/MsgPush/Channel"
	"Lovers_srv/MsgPush/SportPush/RunPush"
	"Lovers_srv/MsgPush/SportPush/SportConst"
	"Lovers_srv/MsgPush/WSProtocol"
	"Lovers_srv/MsgPush/WSProtocol/SportProtocol"
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache/MsgPushCache"
	"Lovers_srv/helper/Utils"
	"github.com/gorilla/websocket"
	"time"
)

type SportHandler struct{
	Ch 				*Channel.Channel
	WSProtoMsg		*WSProtocol.Proto
	SportProtoMsg	*SportProtocol.SportProto
	RunData			*RunPush.RunData
	CurSportId		string
}

func NewSportHandler(conn *websocket.Conn, userId, loverId string, sportId string) *SportHandler{
	sp := new(SportHandler)
	sp.Ch = Channel.NewChannel(conn, userId, loverId)
	sp.WSProtoMsg = new(WSProtocol.Proto)
	sp.SportProtoMsg = new(SportProtocol.SportProto)
	sp.RunData	= RunPush.NewRunData(sp.Ch)
	sp.CurSportId = sportId
	return sp
}

/******
	断开Sport Web Socket TCP通信
******/
func (sport *SportHandler)Release(){
	Utils.InfoOutputf("[SportWSHandle] Release Web Socket, userId:%s",sport.Ch.UserId)
	sport.SaveSportDataToDB(false) //关闭通信前，要存储数据
	err := sport.Ch.WSConn.Close()
	if err != nil{
		Utils.ErrorOutputf("[SportWSHandle]Sport Release failed:%s",err.Error())
	}
}
/******
	Sport Web Socket模块处理
******/
func (sport *SportHandler) SportWSHandle(){
	for{
		//设置web socket 读数据超时时间
		err := sport.Ch.WSConn.SetReadDeadline(time.Now().Add(12 * time.Minute))
		err = sport.WSProtoMsg.ReadWebsocket(sport.Ch.WSConn)
		if err != nil{
			Utils.ErrorOutputf("[SportWSHandle] Read Web Socket failed:%s", err.Error())
			sport.Release()
			return
		}

		switch sport.WSProtoMsg.Op {
		case WSProtocol.PACK_HEARTBEAT:
			sport.HearbetHandle()
			break
		case WSProtocol.PACK_MESSAGE:
			sport.MsgHandle()
			break
		}

	}
}

/******
	接收到心跳数据，回复
******/
func (sport *SportHandler)HearbetHandle(){
	sport.WSProtoMsg.WriteWebsocketHeart(sport.Ch.WSConn)
}

/******
	接收到消息数据，处理
******/
func (sport *SportHandler)MsgHandle(){
	body := sport.WSProtoMsg.Body
	sport.SportProtoMsg.ReadSportProto(body)
	switch sport.SportProtoMsg.SportType {
	//跑步运动
	case SportConst.SPORT_TYPE_RUN:
		sport.RunData.ReceiveCount++ //收到跑步消息，更新跑步消息收到数量
		if sport.RunData.TotalStatistics.LastRunId == "" || sport.RunData.TotalStatistics.LastRunIndex <= 0{
			//查询数据库，获取总跑步次数，和上次跑步的ID号
			runStatis, errCode, err := MsgPushCache.GetRunStatistics(sport.Ch.UserId)
			if errCode != config.ENUM_ERR_OK && errCode != config.ENUM_ERR_DB_QUERY_NOT_FOUND{
				err = Utils.ErrorOutputf("[MsgHandle] Query run statistics failed:%s", err.Error())
				return
			}
			if runStatis == nil{ //没有查到
				//赋值
				sport.RunData.CurSportIndex = 1

			}else{ //查到了
				sport.RunData.TotalStatistics = runStatis
				sport.RunData.CurSportIndex = runStatis.LastRunIndex + 1
			}
			sport.RunData.CurSportId = sport.CurSportId
			sport.RunData.TotalStatistics.LastRunIndex = sport.RunData.CurSportIndex
			sport.RunData.TotalStatistics.LastRunId = sport.CurSportId
		}

		sport.RunData.RunDataHandle(sport.SportProtoMsg.SportStatus, sport.SportProtoMsg.Body,sport.Ch)
		break

	}
}

/******
	存储运动数据
*****/
func (sp *SportHandler)SaveSportDataToDB(bInterval bool){

	if bInterval {
		//达到接收最大消息数时，自动存储统计数据
		if sp.RunData.ReceiveCount > SportConst.MAX_RECIEVE_RUN_SAVE_COUNT {
			//存储统计数据
			sp.RunData.SaveStatisticsToCache(true)

			//存储当前运动数据
			sp.RunData.SaveCurDataToCache(true)

			//存储之后，消息数量重置
			sp.RunData.ReceiveCount = 0;
		}
	}else{ 		//主动存储数据

		//存储统计数据
		sp.RunData.SaveStatisticsToCache(true)

		//存储当前运动数据
		sp.RunData.SaveCurDataToCache(true)
	}

}