package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	proto "Lovers_srv/server/home-service/proto"
	userClient "Lovers_srv/server/user-service/client"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type HomeHandler struct {
	DB *gorm.DB
}

var(
	user_clent = userClient.NewUserClient()
	totalCardCount = 0 // 整个数据库有多少张卡片
)

func (home* HomeHandler) GetMainCard(ctx context.Context, in *proto.GetMainCardReq, out *proto.GetMainCardResp) error{
	out.RespStatus = &proto.CardRespStatus{}
	weekStart, weekEnd := Utils.GetThisWeekStartEnd()
	weekStartUnix := weekStart.Unix()
	weekEndUnix := weekEnd.Unix()
	//先按ShowIndex升序 查询这一周相关的MainCard列表
	var thisWeekMainCard []DB.HomeCardInfo
	var lastWeekMainCard []DB.HomeCardInfo
	err := home.DB.Where("create_time >= ? and create_time < ?", strconv.FormatInt(weekStartUnix,10), strconv.FormatInt(weekEndUnix,10)).Order("show_index desc").Find(&thisWeekMainCard).Error

	//卡片数量不满足一周的需求，重新查询上一周
	if len(thisWeekMainCard) < 7 || err != nil{
		d, _ := time.ParseDuration("-168h") // 倒推一个星期7天
		weekStart = weekStart.Add(d)
		weekEnd = weekEnd.Add(d)
		logrus.Debug("weekStart:"+ weekStart.String())
		logrus.Debug("weekEnd:"+ weekEnd.String())
		weekStartUnix = weekStart.Unix()
		weekEndUnix = weekEnd.Unix()
		err = home.DB.Where("create_time >= ? AND create_time < ?", weekStartUnix, weekEndUnix).Order("show_index desc").Find(&lastWeekMainCard).Error
	}
	if err != nil{
		logrus.Error("query table CardInfo failed: " + err.Error())
		return GetMainCardFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
	}

	logrus.Debug("This Week MainCardList Len:"+ strconv.Itoa(len(thisWeekMainCard)))
	logrus.Debug("Last Week MainCardList Len:"+ strconv.Itoa(len(lastWeekMainCard)))
	if len(thisWeekMainCard) >= 7 {
		//查询到足够的卡片后，只取前7个card返回
		thisWeekMainCard = thisWeekMainCard[:7]
	}else if len(lastWeekMainCard) >= 7{
		//查询到足够的卡片后，只取前7个card返回
		lastWeekMainCard = lastWeekMainCard[:7]
	}
	GetMainCardSuccessResp(thisWeekMainCard,out)
	return nil
}


//上传Card
func (home* HomeHandler)PostCardInfo(ctx context.Context, in *proto.PostCardInfoReq, out *proto.PostCardInfoResp) error {
	out.RespStatus = &proto.CardRespStatus{}

	//先查看是否存在该用户
	upLoadUserId := in.PostCardInfo.UpLoadUserId
	queryExistReq := &lovers_srv_user.QueryUserIsExistByIdReq{UserId:upLoadUserId}
	queryExistResp,err := user_clent.Client_QueryUserExistById(ctx,queryExistReq)
	if err != nil{
		return err
	}else{
		if !queryExistResp.IsExist{
			out.RespStatus.OpCardCode = queryExistResp.QueryCode
			out.RespStatus.OpCardRes = queryExistResp.QueryRes
			return Utils.MicroErr(queryExistResp.QueryRes,int(queryExistResp.QueryCode))
		}
	}

	dbCardInfo := ReqCardToDBCard(*in.PostCardInfo)
	err = home.DB.Create(&dbCardInfo).Error
	if err != nil{
		out.RespStatus.OpCardCode = config.CODE_ERR_INSERT_DB_FAIL
		out.RespStatus.OpCardRes = config.MSG_ERR_INSERT_DB_FAIL
		logrus.Error("\"PostCardInfo\" insert table DB.HomeCardInfo fail: "+ err.Error())
		return Utils.MicroErr(config.MSG_ERR_INSERT_DB_FAIL,config.CODE_ERR_INSERT_DB_FAIL)
	}
	out.RespStatus.OpCardCode = config.CODE_ERR_SUCCESS
	out.RespStatus.OpCardRes = config.MSG_REQUEST_SUCCESS
	return nil
}


func (home* HomeHandler) GetCardByCount(ctx context.Context, in *proto.GetCardByCountReq, out *proto.GetCardByCountResp) error {
	var startTime = in.StartTime
	var endTime = in.EndTime
	var searchCount = int(in.CardCount)
	var bStartTime = true //是否查询开始时间
	var bEndTime = true //是否查询结束时间
	if startTime <= 0{
		bStartTime = false
	}
	if endTime <= 0{
		bEndTime = false
	}
	if searchCount <= 0{
		searchCount = config.GlobalConfig.DefaultCardCount
	}
	var cardList []DB.HomeCardInfo
	if(bStartTime && bEndTime){
		//限定开始和结束时间查询
		err := home.DB.Where("create_time >= ? and create_time < ?", strconv.FormatInt(startTime,10), strconv.FormatInt(endTime,10)).Order("show_index desc").Limit(searchCount).Find(&cardList).Error
		if err != nil{
			logrus.Error("query table CardInfo failed: " + err.Error())
			return GetCardByCountFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
		}
		GetCardByCountSuccessResp(cardList, out)
		return nil
	}else if(bStartTime && !bEndTime){
		//只限定开始
		err := home.DB.Where("create_time >= ? ", strconv.FormatInt(startTime,10)).Order("show_index desc").Limit(searchCount).Find(&cardList).Error
		if err != nil{
			logrus.Error("query table CardInfo failed: " + err.Error())
			return GetCardByCountFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
		}
		GetCardByCountSuccessResp(cardList, out)
		return nil
	}else if(bEndTime && !bStartTime){
		//只限定结束时间
		err := home.DB.Where("create_time < ?",strconv.FormatInt(endTime,10)).Order("show_index desc").Limit(searchCount).Find(&cardList).Error
		if err != nil{
			logrus.Error("query table CardInfo failed: " + err.Error())
			return GetCardByCountFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
		}
		GetCardByCountSuccessResp(cardList, out)
		return nil
	}else{
		//不限定时间
		err := home.DB.Order("show_index desc").Limit(searchCount).Find(&cardList).Error
		if err != nil{
			logrus.Error("query table CardInfo failed: " + err.Error())
			return GetCardByCountFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
		}
		GetCardByCountSuccessResp(cardList, out)
		return nil
	}
}

func (home* HomeHandler) GetCardByIndex(ctx context.Context, in *proto.GetCardByIndexReq, out *proto.GetCardByIndexResp) error {
	var startIndex = in.StartIndex
	var endIndex = in.EndIndex
	var searchCount = endIndex - startIndex

	if searchCount <= 0{
		logrus.Debug("传入startIndex和endIndex错误")
		return GetCardByIndexFailResp(out, config.MSG_ERR_PARAM_WRONG,config.CODE_ERR_PARAM_WRONG)
	}

	var cardList []DB.HomeCardInfo
	err := home.DB.Order("show_index desc").Find(&cardList).Error
	if err != nil{
		logrus.Error("query table CardInfo failed: " + err.Error())
		return GetCardByIndexFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
	}
	totalCardCount = len(cardList)
	if totalCardCount <= 0{
		GetCardByIndexSuccessResp(cardList, out)
		return nil
	}

	err = home.DB.Where("ID >= ? and ID < ?", strconv.Itoa(int(startIndex)), strconv.Itoa(int(endIndex))).Order("show_index desc").Find(&cardList).Error
	if err != nil{
		logrus.Error("query table CardInfo failed: " + err.Error())
		return GetCardByIndexFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
	}
	GetCardByIndexSuccessResp(cardList, out)
	return nil
}

//获取MainCard成功的操作
func GetMainCardSuccessResp(cardList []DB.HomeCardInfo, out *proto.GetMainCardResp){
	out.MainCardInfo = DBHomeCardToRespHomeCard(cardList)
	out.RespStatus.OpCardCode = config.CODE_ERR_SUCCESS
	out.RespStatus.OpCardRes = config.MSG_REQUEST_SUCCESS
}

//获取MainCard失败的操作
func GetMainCardFailResp(out *proto.GetMainCardResp, msg string, code int) error{
	out.MainCardInfo = nil
	out.RespStatus.OpCardRes = msg
	out.RespStatus.OpCardCode = int32(code)
	return Utils.MicroErr(msg, code)
}

//获取CardByCount成功的操作
func GetCardByCountSuccessResp(cardList []DB.HomeCardInfo, out *proto.GetCardByCountResp){
	out.CardList = DBHomeCardToRespHomeCard(cardList)
}

//获取CardByCount失败的操作
func GetCardByCountFailResp(out *proto.GetCardByCountResp, msg string, code int) error{
	out.CardList = nil
	return Utils.MicroErr(msg,code)
}

//获取CardByIndex成功的操作
func GetCardByIndexSuccessResp(cardList []DB.HomeCardInfo, out *proto.GetCardByIndexResp){
	out.CardList = DBHomeCardToRespHomeCard(cardList)
	out.GetCardCount = int32(len(cardList));
	out.TotalCardCount = int32(totalCardCount);
}

//获取CardByCount失败的操作
func GetCardByIndexFailResp(out *proto.GetCardByIndexResp, msg string, code int) error{
	out.CardList = nil
	out.GetCardCount = 0;
	out.TotalCardCount = 0;
	return Utils.MicroErr(msg,code)
}

func DBHomeCardToRespHomeCard(dbCardList []DB.HomeCardInfo) []*proto.HomeCardInfo{
	var respCardList [] *proto.HomeCardInfo
	for _, v := range dbCardList{
		tmpRespCard := &proto.HomeCardInfo{}
		tmpRespCard.CardType = proto.CARDTYPE(v.CardType)
		tmpRespCard.AdType = proto.ADTYPE(v.AdType)
		tmpRespCard.InfoType = proto.INFOTYPE(v.InfoType)
		tmpRespCard.Title = v.Title
		tmpRespCard.Content = v.Content
		tmpRespCard.TypeDesc = v.TypeDesc
		tmpRespCard.CreateTime = v.CreateTime
		tmpRespCard.ShowIndex = int32(v.ShowIndex)
		tmpRespCard.IsMainCard = v.IsMainCard
		tmpRespCard.UpLoadUserId = v.UpLoadUserId
		tmpRespCard.CardId = v.CardId
		tmpRespCard.HomeHtmlUrl = v.HomeHtmlUrl
		tmpRespCard.HomeImgUrl = v.HomeImgUrl
		tmpRespCard.CardMediaType = proto.MEDIATYPE(v.CardMediaType)
		tmpRespCard.AudioFileUrl = v.AudioFileUrl
		tmpRespCard.AudioLength = v.AudioLength
		tmpRespCard.ImgMaskType = int32(v.ImgMaskType)
		respCardList = append(respCardList, tmpRespCard)
	}
	return respCardList
}
func ReqCardToDBCard(reqCardList proto.HomeCardInfo) *DB.HomeCardInfo{
	tmpDBCardList := &DB.HomeCardInfo{}
	tmpDBCardList.CardType = int(reqCardList.CardType)
	tmpDBCardList.AdType = int(reqCardList.AdType)
	tmpDBCardList.InfoType = int(reqCardList.InfoType)
	tmpDBCardList.Title = reqCardList.Title
	tmpDBCardList.Content = reqCardList.Content
	tmpDBCardList.TypeDesc = reqCardList.TypeDesc
	tmpDBCardList.ShowIndex = int(reqCardList.ShowIndex)
	tmpDBCardList.IsMainCard = reqCardList.IsMainCard
	tmpDBCardList.UpLoadUserId = reqCardList.UpLoadUserId
	tmpDBCardList.CardId = reqCardList.CardId
	tmpDBCardList.CreateTime = reqCardList.CreateTime
	tmpDBCardList.HomeImgUrl = reqCardList.HomeImgUrl
	tmpDBCardList.HomeHtmlUrl = reqCardList.HomeHtmlUrl
	tmpDBCardList.CardMediaType = int(reqCardList.CardMediaType)
	tmpDBCardList.AudioFileUrl = reqCardList.AudioFileUrl
	tmpDBCardList.AudioLength = reqCardList.AudioLength
	tmpDBCardList.ImgMaskType = int(reqCardList.ImgMaskType)
	if reqCardList.CreateTime <= 0{
		tmpDBCardList.CreateTime = time.Now().Unix()
	}
	return tmpDBCardList
}
func ReqCardListToDBCardList(reqCardList []proto.HomeCardInfo) []*DB.HomeCardInfo{
	var dbCardList [] *DB.HomeCardInfo
	for _, v := range reqCardList{
		tmpDBCardList := &DB.HomeCardInfo{}
		tmpDBCardList.CardType = int(v.CardType)
		tmpDBCardList.AdType = int(v.AdType)
		tmpDBCardList.InfoType = int(v.InfoType)
		tmpDBCardList.Title = v.Title
		tmpDBCardList.Content = v.Content
		tmpDBCardList.TypeDesc = v.TypeDesc
		tmpDBCardList.CreateTime = v.CreateTime
		tmpDBCardList.ShowIndex = int(v.ShowIndex)
		tmpDBCardList.IsMainCard = v.IsMainCard
		tmpDBCardList.UpLoadUserId = v.UpLoadUserId
		tmpDBCardList.CardId = v.CardId
		tmpDBCardList.HomeHtmlUrl = v.HomeHtmlUrl
		tmpDBCardList.HomeImgUrl = v.HomeImgUrl
		tmpDBCardList.CardMediaType = int(v.CardMediaType)
		tmpDBCardList.AudioFileUrl = v.AudioFileUrl
		tmpDBCardList.AudioLength = v.AudioLength
		dbCardList = append(dbCardList, tmpDBCardList)
	}
	return dbCardList
}



