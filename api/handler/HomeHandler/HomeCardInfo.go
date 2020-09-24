package HomeHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/home-service/client"
	proto "Lovers_srv/server/home-service/proto"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var home_client = client.NewHomeMicroClient()

func GetMainCard(c *gin.Context){

	mainCardList,err := home_client.Client_GetMainCard(c,nil)
	if err == nil && mainCardList != nil{
		total := len(mainCardList.MainCardInfo)
		Utils.CreateSuccessByList(c, total, mainCardList)
	}else{
		if mainCardList == nil{
			Utils.CreateErrorWithMsg(c, "GetMainCard failed error msg:"+err.Error() + "loginResp is nil",config.CODE_ERR_SERVER_INTERNAL)
		}else {
			code, err := strconv.Atoi(mainCardList.RespStatus.GetCardCode)
			Utils.CreateErrorWithMsg(c, "GetMainCard failed error msg:"+err.Error(), code)
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
	var cardInfo = &proto.PostCardInfoReq{}
	cardType, _ := strconv.Atoi(c.PostForm("CardType"))
	cardInfo.CardType = proto.CARDTYPE(cardType)

	adType, _ := strconv.Atoi(c.PostForm("AdType"))
	cardInfo.AdType = proto.ADTYPE(adType)

	infoType, _ := strconv.Atoi(c.PostForm("InfoType"))
	cardInfo.InfoType = proto.INFOTYPE(infoType)

	cardInfo.ImgUrl = c.PostForm("ImgUrl")
	cardInfo.Title = c.PostForm("Title")
	cardInfo.Content = c.PostForm("Content")
	cardInfo.TypeDesc = c.PostForm("TypeDesc")

	home_client.Client_PostCardInfo(c, cardInfo)
}
