package HomeHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/home-service/client"
	proto "Lovers_srv/server/home-service/proto"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var home_client = client.NewHomeMicroClient()

func GetMainCard(c *gin.Context){
	getMainCardReq := &proto.GetMainCardReq{}
	mainCardList,err := home_client.Client_GetMainCard(c,getMainCardReq)
	if err == nil && mainCardList != nil{
		total := len(mainCardList.MainCardInfo)
		Utils.CreateSuccessByList(c, total, mainCardList)
	}else{
		if mainCardList == nil{
			Utils.CreateErrorWithMsg(c, err.Error(),config.CODE_ERR_SERVER_INTERNAL)
		}else {
			Utils.CreateErrorWithMsg(c, err.Error(), int(mainCardList.RespStatus.OpCardCode))
		}
	}
}

func GetCardByCount(c *gin.Context){
	getCardByCountReq := &proto.GetCardByCountReq{}
	err := c.ShouldBind(getCardByCountReq)
	if err != nil{
		Utils.CreateErrorWithMsg(c,err.Error(),config.INVALID_PARAMS)
		return
	}
	HomeCardList,err := home_client.Client_GetCardByCount(c,getCardByCountReq)
	if err == nil && HomeCardList != nil{
		total := len(HomeCardList.CardList)
		Utils.CreateSuccessByList(c, total, HomeCardList)
	}else{
			msg,code := Utils.SplitMicroErr(err)
			Utils.CreateErrorWithMsg(c, msg,code)
	}
}
func GetCardByIndex(c *gin.Context){
	getCardByIndexReq := &proto.GetCardByIndexReq{}
	err := c.ShouldBind(getCardByIndexReq)
	if err != nil{
		Utils.CreateErrorWithMsg(c,err.Error(),config.INVALID_PARAMS)
		return
	}
	HomeCardList,err := home_client.Client_GetCardByIndex(c, getCardByIndexReq)
	if err == nil && HomeCardList != nil{
		total := len(HomeCardList.CardList)
		Utils.CreateSuccessByList(c, total, HomeCardList)
	}else{
		msg,code := Utils.SplitMicroErr(err)
		Utils.CreateErrorWithMsg(c, msg,code)
	}
}

func GetCardByType(c *gin.Context){

}


func PostCardInfo(c* gin.Context){
	var NewCardInfo = &proto.HomeCardInfo{}
	var cardInfo = &proto.PostCardInfoReq{PostCardInfo:NewCardInfo}
	err := c.ShouldBind(cardInfo)
	if err != nil{
		Utils.CreateErrorWithMsg(c,err.Error(),config.INVALID_PARAMS)
		return
	}

	cardInfo.PostCardInfo.CardId = uuid.NewV1().String()

	_, err = home_client.Client_PostCardInfo(c, cardInfo)
	if err != nil{
		Utils.CreateErrorWithMicroErr(c, err)
	}else{
		Utils.CreateSuccess(c,nil)
	}
}
