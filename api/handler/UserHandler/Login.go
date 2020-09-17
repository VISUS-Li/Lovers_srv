package UserHandler

import (
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
	login.UserName = c.PostForm("UserName")
	login.PassWord = c.PostForm("PassWord")
	login.Phone = c.Query("Phone")
	login.VertifyCode = c.Query("VertifyCode")
	var loginResp = &lovers_srv_user.LoginResp{}
	loginResp,err := user_clent.Client_Login(c,login)
	if err == nil && loginResp != nil{
		Utils.CreateSuccess(c, loginResp)
	}else{
		if loginResp == nil{
			Utils.CreateErrorWithMsg(c, "login failed error msg:"+err.Error() + "loginResp is nil")
		}else {
			Utils.CreateErrorWithMsg(c, "login failed error msg:"+err.Error())
		}
	}
}

