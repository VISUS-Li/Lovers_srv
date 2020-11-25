package RunPush

import (
	"Lovers_srv/MsgPush/Channel"
	"Lovers_srv/MsgPush/SportPush/SportConst"
	"Lovers_srv/MsgPush/WSProtocol"
	"Lovers_srv/MsgPush/WSProtocol/SportProtocol"
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache/MsgPushCache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"encoding/binary"
	"github.com/gorilla/websocket"
	"time"
)


const(
	_RunPackOffset = 0
	_RunCurDisOffset = 4
	_RunCurPaceOffset = 8
	_RunCurTimeOffset = 12
	_RunCurCaloOffset = 16
	_RunPackLen = 16

	_MaxErrorCount = 100
)
/******
	实时接收的跑步数据
******/
type RunProto struct{
	CurDistance	int32 //当前运动长度
	CurPace		int32 //当前配速
	CurTime		int32 //当前耗时
	CurCalorie	int32 //当前消耗卡路里

	errorCount	int //解析数据错误次数
}

func (rp *RunProto)ParseProto(_body []byte) error{
	body := _body

	packLen := len(body)
	if packLen < _RunPackLen{
		return Utils.ErrorOutputf("[RunProto] invalid pack len:%d",packLen)
	}
	rp.CurDistance = int32(binary.BigEndian.Uint32(body[_RunPackOffset : _RunCurDisOffset]))
	rp.CurPace = int32(binary.BigEndian.Uint32(body[_RunCurDisOffset : _RunCurPaceOffset]))
	rp.CurTime = int32(binary.BigEndian.Uint32(body[_RunCurPaceOffset : _RunCurTimeOffset]))
	rp.CurCalorie = int32(binary.BigEndian.Uint32(body[_RunCurTimeOffset : _RunCurCaloOffset]))
	if rp.CurDistance <= 0 && rp.CurPace <= 0 && rp.CurTime <= 0 && rp.CurCalorie <= 0{
		rp.errorCount++
	}
	if rp.errorCount > _MaxErrorCount{
		return Utils.ErrorOutputf("[RunProto] error count more than _MaxErrorCount")
	}
	return nil
}

func (rp *RunProto)PrepareRunWriteData() []byte{
	buflen := _RunPackLen
	buf := make([]byte, buflen)
	binary.BigEndian.PutUint32(buf[_RunPackOffset:_RunCurDisOffset], uint32(rp.CurDistance))
	binary.BigEndian.PutUint32(buf[_RunCurDisOffset:_RunCurPaceOffset], uint32(rp.CurPace))
	binary.BigEndian.PutUint32(buf[_RunCurPaceOffset:_RunCurTimeOffset], uint32(rp.CurTime))
	binary.BigEndian.PutUint32(buf[_RunCurTimeOffset:_RunCurCaloOffset], uint32(rp.CurCalorie))
	return buf
}

/******
	跑步数据
******/
type RunData struct{
	CurProto		*RunProto		//本次消息传来的数据
	PreProto		*RunProto 		//上一个消息传来的数据，与CurProto一起计算TotalStatistics
	CurSportIndex	int32			//当前跑步序号
	CurSportId		string			//当前跑步的ID
	CurItemData		*DB.RunItemData
	TotalStatistics *DB.RunStatistics
	Ch				*Channel.Channel
	ReceiveCount	int64			//收到跑步消息的次数
	StartTime		int64			//本次跑步开始时间
	PackSeq			int64			//发送包的序号
}

func NewRunData(ch *Channel.Channel) *RunData{
	r  := new(RunData)
	cp := new(RunProto)
	pp := new(RunProto)
	rs := new(DB.RunStatistics)
	ri := new(DB.RunItemData)
	r.CurProto = cp
	r.PreProto = pp
	r.TotalStatistics = rs
	r.TotalStatistics.UserId = ch.UserId
	r.CurItemData = ri
	r.Ch			= ch
	r.StartTime = time.Now().Unix() //初始化RunData时，就表示开始本次跑步
	return r
}

func (r *RunData)RunDataHandle(status int32, body []byte, curChannel *Channel.Channel) error {
	//运动数据处理，包括：
	//1.实时运动数据推送
	//2.总运动量统计
	//3.运动数据存入缓存和数据库(完成、暂停、异常终止时写入)

	//stpe 1 解析跑步数据
	err := r.CurProto.ParseProto(body)
	if err != nil{
		return Utils.ErrorOutputf("[RunData] ParseProto failed:%s", err.Error())
	}

	//step 2 根据不同状态作处理
	switch status {
	case SportConst.SPROT_STATUS_ING:
		err = r.RunningHandle()
		break
	case SportConst.SPORT_STATUS_PAUSE:
		err = r.PauseHandle()
		break

	case SportConst.SPORT_STATUS_STOP:
	case SportConst.SPORT_STATUS_ABNORMALEXIT:
	case SportConst.SPORT_STATUS_UNKNOWN:
		err = r.StopHandle(status)
		break
	}
	return err
}

/******
	发送跑步数据
******/
func (r *RunData)WriteRunData(status int32, ws *websocket.Conn) error{
	//准备跑步数据
	runData := r.CurProto.PrepareRunWriteData()

	sportProto := new(SportProtocol.SportProto)
	sportProto.SportType = SportConst.SPORT_TYPE_RUN
	sportProto.SportStatus = status
	sportProto.Body = runData
	sportData := sportProto.PrepareSportProto()

	r.PackSeq += 1
	wsProto := new(WSProtocol.Proto)
	wsProto.Op = WSProtocol.PACK_MESSAGE
	wsProto.Seq = int32(r.PackSeq)
	wsProto.Body = sportData

	//发送跑步数据
	err := wsProto.WriteWebsocket(ws)
	if err != nil{
		return Utils.ErrorOutputf("write web socket message failed:%s", err.Error())
	}
	return nil
}

/******
	正在跑步
******/
func (r *RunData)RunningHandle() error{
	//step 1 更新数据
	r.AddStatistics()
	r.AddRunItemData()
	//更新完数据就记录当前数据为上一次的数据了
	r.PreProto = r.CurProto

	//step 2 推送给另一半
	//获取另一半的Channel
	loverCh := Channel.SportChMng.Get(r.Ch.LoverId)
	if loverCh == nil{
		return Utils.ErrorOutputf("[RunningHandle] Can not find userId:%s 's lover channel", r.Ch.UserId)
	}
	//发送消息
	err := r.WriteRunData(SportConst.SPROT_STATUS_ING, loverCh.WSConn)
	if err != nil{
		return Utils.ErrorOutputf("[RunningHandle] :%s", err.Error())
	}

	return nil
}

func (r *RunData)PauseHandle() error {
	//step 1 更新数据
	r.AddPauseTime()

	//step 2 推送给另一半
	//获取另一半的Channel
	loverCh := Channel.SportChMng.Get(r.Ch.LoverId)
	if loverCh == nil{
		return Utils.ErrorOutputf("[PauseHandle] Can not find userId:%s 's lover channel", r.Ch.UserId)
	}

	//发送消息
	err := r.WriteRunData(SportConst.SPORT_STATUS_PAUSE, loverCh.WSConn)
	if err != nil{
		return Utils.ErrorOutputf("[PauseHandle] :%s", err.Error())
	}
	return nil
}

func (r *RunData)StopHandle(status int32) error{
	// step 1 更新数据
	r.AddEndRunItem()

	//step 2 推送给另一半
	//获取另一半的Channel
	loverCh := Channel.SportChMng.Get(r.Ch.LoverId)
	if loverCh == nil{
		return Utils.ErrorOutputf("[StopHandle] Can not find userId:%s 's lover channel", r.Ch.UserId)
	}
	//发送消息
	err :=r.WriteRunData(status, loverCh.WSConn)
	if err != nil{
		return Utils.ErrorOutputf("[StopHandle] :%s", err.Error())
	}
	return nil
}

func (r *RunData) SaveStatisticsToCache(saveDB bool) (int, error) {
	code, err := MsgPushCache.SetRunStatistics(r.Ch.UserId, *r.TotalStatistics,saveDB)
	return code, err
}

func (r *RunData)SaveCurDataToCache(saveDB bool) (int, error){
	code, err := MsgPushCache.SetRunItemData(r.Ch.UserId, r.CurSportId, *r.CurItemData, saveDB)
	return code, err
}

func (r *RunData)AddStatistics(){
	disOffset := r.CurProto.CurDistance - r.PreProto.CurDistance
	if disOffset > 0 {
		r.TotalStatistics.TotalDis += disOffset
	}
	calOffset := r.CurProto.CurCalorie - r.PreProto.CurCalorie
	if calOffset > 0{
		r.TotalStatistics.TotalCalorie += calOffset
	}

	timeOffset := r.CurProto.CurTime - r.PreProto.CurTime
	if timeOffset > 0{
		r.TotalStatistics.TotalRunTime += timeOffset
	}


	if r.TotalStatistics.Farthest < r.CurProto.CurDistance{
		r.TotalStatistics.Farthest = r.CurProto.CurDistance
	}

	if r.TotalStatistics.FastestPace <= 0 || r.TotalStatistics.FastestPace > r.CurProto.CurPace{
		r.TotalStatistics.FastestPace = r.CurProto.CurPace
	}

	if r.TotalStatistics.LongestTime < r.CurProto.CurTime{
		r.TotalStatistics.LongestTime = r.CurProto.CurTime
	}

	//更新最近跑步日期
	startTime, _ := Utils.GetNowDayStartEnd()
	if startTime > r.TotalStatistics.LatestTime {
		r.TotalStatistics.TotalDays ++
		r.TotalStatistics.LatestTime = time.Now().Unix()
	}
}

/******
	结束跑步数据计算
******/
func (r *RunData)AddEndRunItem(){
	r.AddStatistics()
	r.AddRunItemData()
	r.ReceiveCount = 0
	r.CurProto.errorCount = 0
	//保存数据
	r.SaveStatisticsToCache(true)
	r.SaveCurDataToCache(true)
}

/******
	暂停跑步时数据计算
******/
func (r *RunData)AddPauseTime() {
	r.CurItemData.PauseTime += r.CurProto.CurTime //前端记录暂停了多久，时间是个区间，在暂停时上传
}

/*****
	跑步中数据计算
******/
func (r *RunData)AddRunItemData() error {
	//通过开始时间判断是否已经从数据库中获取了该次的跑步数据
	if r.CurItemData.StartTime <= 0{
		//没有取，先从数据库取数据，再更新
		runItem, code, err := MsgPushCache.GetRunItemData(r.Ch.UserId, r.CurSportId)
		if code != config.ENUM_ERR_OK && code != config.ENUM_ERR_DB_QUERY_NOT_FOUND{
			return err
		}
		if code == config.ENUM_ERR_DB_QUERY_NOT_FOUND{
			//数据库没有该数据
			r.CurItemData.UserId = r.Ch.UserId
			r.CurItemData.RunId = r.CurSportId
			r.CurItemData.StartTime = r.StartTime
			r.CurItemData.RunIndex = r.CurSportIndex
		}else{
			if runItem != nil{
				r.CurItemData = runItem
			}else{
				return err
			}
		}
	}
	r.CurItemData.EndTime = time.Now().Unix()
	if r.CurProto.CurDistance > r.CurItemData.RunDistance {
		r.CurItemData.RunDistance = r.CurProto.CurDistance
	}

	if r.CurProto.CurCalorie > r.CurItemData.RunCalorie {
		r.CurItemData.RunCalorie = r.CurProto.CurCalorie
	}

	if r.CurItemData.FastestPace > r.CurProto.CurPace || r.CurItemData.FastestPace <= 0{
		r.CurItemData.FastestPace = r.CurProto.CurPace
	}

	if r.CurItemData.SlowestPace < r.CurProto.CurPace || r.CurItemData.SlowestPace <= 0{
		r.CurItemData.SlowestPace = r.CurProto.CurPace
	}

	return nil
}
