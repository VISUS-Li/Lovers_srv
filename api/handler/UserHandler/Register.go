package UserHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	info := &lovers_srv_user.BaseInfo{}
	var regParam = &lovers_srv_user.RegisterReq{UserInfo:info}
	regParam.UserName = c.PostForm("UserName")
	regParam.PassWord = c.PostForm("PassWord")
	regParam.UserInfo.RealName = c.PostForm("RealName")
	regParam.UserInfo.Sex = c.PostForm("Sex")
	regParam.UserInfo.Phone = c.PostForm("Phone")
	regParam.UserInfo.HomeTown = c.PostForm("HomeTown")
	regParam.UserInfo.Sculpture = c.PostForm("Sculpture")
	regParam.UserInfo.Birth = c.PostForm("Birth")
	regParam.BindId = c.PostForm("BindId")
	regParam.RecommendID = c.PostForm("RecommendID")
	if len(regParam.UserName) <= 0 || len(regParam.PassWord) <= 0 {
		Utils.CreateErrorWithMsg(c, "UserName or PassWord is empty!")
	} else if !Utils.VerifyPhoneFormat(regParam.UserInfo.Phone){
		Utils.CreateErrorWithMsg(c, "Phone Format is invalid!")
	}else{
		regResp,err := user_clent.Client_Register(c,regParam)
		if regResp == nil || regResp.RegisteredInfo == nil|| regResp.RegisteredInfo.LoginRes != config.MSG_DB_REG_OK {
			if regResp != nil{
				if regResp.RegisteredInfo != nil{
					Utils.CreateErrorWithMsg(c,"register failed, error msg:" + regResp.RegisteredInfo.LoginRes)
					return
				}
			}
			Utils.CreateErrorWithMsg(c,"register failed,server internal error, error msg:"+ err.Error())
		}else{
			Utils.CreateSuccess(c,regResp)
		}
	}
}