package UserHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	userClient "Lovers_srv/server/user-service/client"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/gin-gonic/gin"
)
var(
	user_clent = userClient.NewUserClient()
)

func Login(c *gin.Context){
	var login = &lovers_srv_user.LoginReq{}
	err := c.ShouldBind(login)
	if err != nil{
		Utils.CreateErrorWithMsg(c,err.Error(),config.INVALID_PARAMS)
		return
	}
	var loginResp = &lovers_srv_user.LoginResp{}
	loginResp,err = user_clent.Client_Login(c,login)
	if err != nil{
		Utils.CreateErrorWithMicroErr(c, err)
	}else {
		Utils.CreateSuccess(c, loginResp)
	}
}
