package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	proto "Lovers_srv/server/home-service/proto"
	userClient "Lovers_srv/server/user-service/client"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"errors"
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
)

func (home* HomeHandler) GetMainCard(ctx context.Context, in *proto.GetMainCardReq, out *proto.GetMainCardResp) error{
	out.RespStatus = &proto.CardRespStatus{}
	weekStart, weekEnd := Utils.GetThisWeekStartEnd()
	weekStartUnix := weekStart.Unix()
	weekEndUnix := weekEnd.Unix()
	//先按ShowIndex升序 查询这一周相关的MainCard列表
	var thisWeekMainCard []DB.HomeCardInfo
	err := home.DB.Where("create_time >= ? and create_time < ?", strconv.FormatInt(weekStartUnix,10), strconv.FormatInt(weekEndUnix,10)).Order("show_index desc").Find(&thisWeekMainCard).Error

	//卡片数量不满足一周的需求，重新查询上一周
	if len(thisWeekMainCard) < 7 || err != nil{

		//清空切片
		var newCardList []DB.HomeCardInfo
		thisWeekMainCard = newCardList

		d, _ := time.ParseDuration("-168h") // 倒推一个星期7天
		weekStart = weekStart.Add(d)
		weekEnd = weekEnd.Add(d)
		logrus.Info("weekStart:"+ weekStart.String())
		logrus.Info("weekEnd:"+ weekEnd.String())
		weekStartUnix = weekStart.Unix()
		weekEndUnix = weekEnd.Unix()
		err = home.DB.Where("create_time >= ? AND create_time < ?", weekStartUnix, weekEndUnix).Order("show_index desc").Find(&thisWeekMainCard).Error
	}
	if err != nil{
		logrus.Error("query table CardInfo failed: " + err.Error())
		return GetHomeCardFailResp(out, config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL)
	}else if len(thisWeekMainCard) < 7 {
		logrus.Error("not enough CardInfo")
		return GetHomeCardFailResp(out, config.MSG_HOME_NOT_ENOUGH_CARD,config.CODE_ERR_HOME_NOT_ENOUGH_CARD)
	}

	//查询到足够的卡片后，只取前7个card返回
	thisWeekMainCard = thisWeekMainCard[:7]
	GetHomeCardSuccessResp(thisWeekMainCard,out)
	return nil
}

func GetHomeCardSuccessResp(cardList []DB.HomeCardInfo, out *proto.GetMainCardResp){
	out.MainCardInfo = DBHomeCardToRespHomeCard(cardList)
	out.RespStatus.OpCardCode = config.CODE_ERR_SUCCESS
	out.RespStatus.OpCardRes = config.MSG_REQUEST_SUCCESS
}

func GetHomeCardFailResp(out *proto.GetMainCardResp, msg string, code int) error{
	out.MainCardInfo = nil
	out.RespStatus.OpCardRes = msg
	out.RespStatus.OpCardCode = int32(code)
	return errors.New(msg)
}



//上传Card
func (home* HomeHandler)PostCardInfo(ctx context.Context, in *proto.PostCardInfoReq, out *proto.PostCardInfoResp) error {
	out.RespStatus = &proto.CardRespStatus{}

	//先查看是否存在该用户
	upLoadUserId := in.PostCardInfo.UpLoadUserId
	queryExistReq := &lovers_srv_user.QueryUserIsExistByIdReq{UserId:upLoadUserId}
	queryExistResp,err := user_clent.Client_QueryUserExistById(ctx,queryExistReq)
	if err != nil{
		if queryExistResp == nil{
			out.RespStatus.OpCardCode = config.INVALID_PARAMS
			out.RespStatus.OpCardRes = config.MSG_ERR_PARAM_WRONG
			return errors.New(config.MSG_ERR_PARAM_WRONG)
		}
		out.RespStatus.OpCardCode = queryExistResp.QueryCode
		out.RespStatus.OpCardRes = queryExistResp.QueryRes
		return errors.New(queryExistResp.QueryRes)
	}else{
		if !queryExistResp.IsExist{
			out.RespStatus.OpCardCode = queryExistResp.QueryCode
			out.RespStatus.OpCardRes = queryExistResp.QueryRes
			return errors.New(queryExistResp.QueryRes)
		}
	}

	dbCardInfo := ReqCardToDBCard(*in.PostCardInfo)
	err = home.DB.Create(&dbCardInfo).Error
	if err != nil{
		out.RespStatus.OpCardCode = config.CODE_ERR_INSERT_DB_FAIL
		out.RespStatus.OpCardRes = config.MSG_ERR_INSERT_DB_FAIL
		logerr := errors.New("\"PostCardInfo\" insert table DB.HomeCardInfo fail: "+ err.Error())
		logrus.Error(logerr)
		return logerr
	}
	out.RespStatus.OpCardCode = config.CODE_ERR_SUCCESS
	out.RespStatus.OpCardRes = config.MSG_REQUEST_SUCCESS
	return nil
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
		dbCardList = append(dbCardList, tmpDBCardList)
	}
	return dbCardList
}