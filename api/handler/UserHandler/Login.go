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
	login.PassWord = c.PostForm("PassWord")
	login.Phone = c.Query("Phone")
	login.VertifyCode = c.Query("VertifyCode")
	login.Type = c.Query("Type") // 登录类型,0.手机号登录,1.账号密码,2.手机验证码,3.微信,4.QQ
	var loginResp = &lovers_srv_user.LoginResp{}
	loginResp,err := user_clent.Client_Login(c,login)
	if err == nil && loginResp != nil{
		Utils.CreateSuccess(c, loginResp)
	}else{
		if loginResp == nil{
			Utils.CreateErrorWithMsg(c, "login failed error msg:"+err.Error() + "loginResp is nil",config.CODE_ERR_UNKNOW)
		}else {
			code,err := strconv.Atoi(loginResp.LoginCode)
			Utils.CreateErrorWithMsg(c, "login failed error msg:"+err.Error(),code)
		}
	}
}
