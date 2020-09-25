package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	proto "Lovers_srv/server/home-service/proto"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
)

type HomeHandler struct {
	DB *gorm.DB
}
func (home* HomeHandler) GetMainCard(ctx context.Context, in *proto.GetMainCardReq, out *proto.GetMainCardResp) error{
	weekStart, weekEnd := Utils.GetThisWeekStartEnd()
	weekStartUnix := weekStart.Unix()
	weekEndUnix := weekEnd.Unix()
	//先按ShowIndex升序 查询这一周相关的MainCard列表
	var thisWeekMainCard []DB.HomeCardInfo
	err := home.DB.Where("CreateTime >= ? AND CreateTime < ?", weekStartUnix, weekEndUnix).Order("ShowIndex desc").Find(&thisWeekMainCard).Error

	//卡片数量不满足一周的需求，重新查询上一周
	if len(thisWeekMainCard) < 7 || err != nil{

		//清空切片
		var newCardList []DB.HomeCardInfo
		thisWeekMainCard = newCardList

		weekStart.Add(-7)
		weekEnd.Add(-7)
		weekStartUnix = weekStart.Unix()
		weekEndUnix = weekEnd.Unix()
		err = home.DB.Where("CreateTime >= ? AND CreateTime < ?", weekStartUnix, weekEndUnix).Find(&thisWeekMainCard).Error
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
	out.RespStatus.GetCardCode = strconv.Itoa(config.CODE_ERR_SUCCESS)
	out.RespStatus.GetCardRes = config.MSG_REQUEST_SUCCESS
}

func GetHomeCardFailResp(out *proto.GetMainCardResp, msg string, code int) error{
	out.MainCardInfo = nil
	out.RespStatus.GetCardRes = msg
	out.RespStatus.GetCardCode = strconv.Itoa(code)
	return errors.New(msg)
}



//上传Card
func (home* HomeHandler)PostCardInfo(ctx context.Context, in *proto.PostCardInfoReq, out *proto.PostCardInfoResp) error {

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
		respCardList = append(respCardList, tmpRespCard)
	}
	return respCardList
}

func ReqCardInfoToDBHomeCard(reqCardList []proto.HomeCardInfo) []*DB.HomeCardInfo{
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
		dbCardList = append(dbCardList, tmpDBCardList)
	}
	return dbCardList
}