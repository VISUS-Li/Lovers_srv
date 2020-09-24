package HomeHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/home-service/client"
	"github.com/gin-gonic/gin"
	"strconv"
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
