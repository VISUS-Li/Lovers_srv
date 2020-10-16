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

func GetCardInfoByCount(c *gin.Context){

}
func GetCardInfoByIndx(c *gin.Context){

}

func GetCardInfoByType(c *gin.Context){

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

	PostCardResp,err := home_client.Client_PostCardInfo(c, cardInfo)
	if err != nil{
		if PostCardResp == nil{
			Utils.CreateErrorWithMsg(c,err.Error(),config.CODE_ERR_SERVER_INTERNAL)
		}else{
			Utils.CreateErrorWithMsg(c, PostCardResp.RespStatus.OpCardRes, int(PostCardResp.RespStatus.OpCardCode))
		}
	}else{
		Utils.CreateSuccess(c,nil)
	}
}
