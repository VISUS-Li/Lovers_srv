package UserHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	userClient "Lovers_srv/server/user-service/client"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/gin-gonic/gin"
	"strconv"
)
var(
	user_clent = userClient.NewUserClient()
)

func Login(c *gin.Context){
	var login = &lovers_srv_user.LoginReq{}
	login.UserName = c.PostForm("UserName")
	if len(login.UserName) <= 0{
		login.UserName = c.Query("UserName")
	}
	login.PassWord = c.PostForm("PassWord")
	if len(login.PassWord) <= 0{
		login.PassWord = c.Query("PassWord")
	}
	login.Phone = c.PostForm("Phone")
	if len(login.Phone) <= 0{
		login.Phone = c.Query("Phone")
	}

	login.VertifyCode = c.Query("VertifyCode")
	login.Type = c.Query("Type") // 登录类型,0.手机号登录,1.账号密码,2.手机验证码,3.微信,4.QQ
	var loginResp = &lovers_srv_user.LoginResp{}
	loginResp,err := user_clent.Client_Login(c,login)
	if err == nil && loginResp != nil{
		Utils.CreateSuccess(c, loginResp)
	}else{
		if loginResp == nil{
			Utils.CreateErrorWithMsg(c, err.Error() ,config.CODE_ERR_UNKNOW)
		}else {
			code,err := strconv.Atoi(loginResp.LoginCode)
			Utils.CreateErrorWithMsg(c, err.Error(),code)
		}
	}
}
