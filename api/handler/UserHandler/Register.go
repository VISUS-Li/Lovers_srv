package UserHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var registerReq = &lovers_srv_user.RegisterReq{}
	err := c.ShouldBind(registerReq)
	if err != nil{
		Utils.CreateErrorWithMsg(c,err.Error(),config.INVALID_PARAMS)
		return
	}


	if len(registerReq.Phone) <= 0|| len(registerReq.PassWord) <= 0 {
		Utils.CreateErrorWithMsg(c, config.MSG_DB_LOGIN_IN_EMPTY,config.CODE_ERR_PARAM_EMPTY)
	} else if !Utils.VerifyPhoneFormat(registerReq.Phone){
		Utils.CreateErrorWithMsg(c, config.MSG_DB_REG_PHONE_ERR,config.CODE_ERR_REG_PHONE_ERR)
	}else{
		regResp,err := user_clent.Client_Register(c,registerReq)
		if err != nil{
			msg,code := Utils.SplitMicroErr(err)
			Utils.CreateErrorWithMsg(c,msg,code)
			return
		}
		Utils.CreateSuccess(c,regResp.RegisteredInfo.UserInfo)
		return
	}
}