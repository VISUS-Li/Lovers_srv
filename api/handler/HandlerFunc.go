package handler

import (
	userClient "Lovers_Micro_Test/server/user-service/client"
	//handler "Lovers_Micro_Test/server/user-service/handler"
	lovers_srv_user "Lovers_Micro_Test/server/user-service/proto"
	"github.com/gin-gonic/gin"
)

var (
	user_clent = userClient.NewUserClient()
	//user_handler = handler.UserHandler{}
)

func Login(c *gin.Context){
	var login = &lovers_srv_user.LoginReq{}
	login.UserName = c.PostForm("UserName")
	login.PassWord = c.PostForm("PassWord")
	login.Phone = c.Query("Phone")
	login.VertifyCode = c.Query("VertifyCode")
	var loginResp = &lovers_srv_user.LoginResp{}
	loginResp,_ = user_clent.Client_Login(c,login)
	if loginResp.Token != "null" {
		CreateSuccess(c, loginResp)

	}

}
