package UserHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/gin-gonic/gin"
)

func BindLover(c *gin.Context){
	var bindReq = &lovers_srv_user.BindLoverReq{}
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