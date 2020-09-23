package HomeHandler

import (
	"Lovers_srv/helper/Utils"
	"Lovers_srv/server/home-service/client"
	"github.com/gin-gonic/gin"
)

var home_client = client.NewHomeMicroClient()

func GetMainCard(c *gin.Context){

	loginResp,err := home_client.Client_GetMainCard(c,nil)
	if err == nil && loginResp != nil{
		Utils.CreateSuccess(c, loginResp)
	}else{
		//if loginResp == nil{
		//	Utils.CreateErrorWithMsg(c, "login failed error msg:"+err.Error() + "loginResp is nil")
		//}else {
		//	Utils.CreateErrorWithMsg(c, "login failed error msg:"+err.Error())
		//}s
	}
}

func GetCardInfoByCount(c *gin.Context){

}
func GetCardInfoByIndx(c *gin.Context){

}

func GetCardInfoByType(c *gin.Context){

}
