package HomeHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/home-service/client"
	proto "Lovers_srv/server/home-service/proto"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strconv"
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

	stCardType, exist := c.GetPostForm("CardType")
	if exist{
		cardType, _ := strconv.Atoi(stCardType)
		cardInfo.PostCardInfo.CardType = proto.CARDTYPE(cardType)
	}else{
		cardInfo.PostCardInfo.CardType = proto.CARDTYPE_CARDTYPE_UNKNOWN
	}

	stAdType,exist := c.GetPostForm("AdType")
	if exist{
		adType, _ := strconv.Atoi(stAdType)
		cardInfo.PostCardInfo.AdType = proto.ADTYPE(adType)
	}else{
		cardInfo.PostCardInfo.AdType = proto.ADTYPE_ADTYPE_UNKNOWN
	}


	stInfoType,exist := c.GetPostForm("InfoType")
	if exist{
		infoType, _ := strconv.Atoi(stInfoType)
		cardInfo.PostCardInfo.InfoType = proto.INFOTYPE(infoType)
	}else{
		cardInfo.PostCardInfo.InfoType = proto.INFOTYPE_INFOTYPE_UNKNOWN
	}


	cardInfo.PostCardInfo.Title = c.PostForm("Title")
	cardInfo.PostCardInfo.Content = c.PostForm("Content")
	cardInfo.PostCardInfo.TypeDesc = c.PostForm("TypeDesc")

	showIndex, _ := strconv.Atoi(c.PostForm("ShowIndex"))
	cardInfo.PostCardInfo.ShowIndex = int32(showIndex)
	isMainCard,_ := strconv.Atoi(c.PostForm("IsMainCard"))
	if isMainCard != 0{
		cardInfo.PostCardInfo.IsMainCard = true
	}else{
		cardInfo.PostCardInfo.IsMainCard = false
	}

	cardInfo.PostCardInfo.UpLoadUserId = c.PostForm("UpLoadUserId")
	if len(cardInfo.PostCardInfo.UpLoadUserId) <= 0{
		Utils.CreateErrorWithMsg(c, "PostUpLoadUserId is null", config.INVALID_PARAMS)
		return
	}

	//cardInfo.PostCardInfo.HomeImgUrl = c.PostForm("HomeImgUrl")
	//if len(cardInfo.PostCardInfo.HomeImgUrl) <= 0{
	//	Utils.CreateErrorWithMsg(c, "HomeImgUrl is null", config.INVALID_PARAMS)
	//	return
	//}
	//
	//cardInfo.PostCardInfo.HomeHtmlUrl = c.PostForm("HomeHtmlUrl")
	//if len(cardInfo.PostCardInfo.HomeImgUrl) <= 0{
	//	Utils.CreateErrorWithMsg(c, "HomeHtmlUrl is null", config.INVALID_PARAMS)
	//	return
	//}

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
