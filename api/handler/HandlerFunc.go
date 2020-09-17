package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	noteListClient "Lovers_srv/server/note-list/client"
	userClient "Lovers_srv/server/user-service/client"
	lovers_srv_notelist "Lovers_srv/server/note-list/proto"
	"strconv"

	//handler "Lovers_srv/server/user-service/handler"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"github.com/gin-gonic/gin"
)

var (
	user_clent = userClient.NewUserClient()
	//user_handler = handler.UserHandler{}
	notelist_client = noteListClient.NewNoteListClient()
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
		CreateSuccess(c, loginResp)
	}else{
		if loginResp == nil{
			CreateErrorWithMsg(c, "login failed error msg:"+err.Error() + "loginResp is nil")
		}else {
			CreateErrorWithMsg(c, "login failed error msg:"+err.Error())
		}
	}
}

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
		CreateErrorWithMsg(c, "UserName or PassWord is empty!")
	} else if !Utils.VerifyPhoneFormat(regParam.UserInfo.Phone){
		CreateErrorWithMsg(c, "Phone Format is invalid!")
	}else{
		regResp,err := user_clent.Client_Register(c,regParam)
		if regResp == nil || regResp.RegisteredInfo == nil|| regResp.RegisteredInfo.LoginRes != config.DB_REG_OK{
			if regResp != nil{
				if regResp.RegisteredInfo != nil{
					CreateErrorWithMsg(c,"register failed, error msg:" + regResp.RegisteredInfo.LoginRes)
					return
				}
			}
			CreateErrorWithMsg(c,"register failed,server internal error, error msg:"+ err.Error())
		}else{
			CreateSuccess(c,regResp)
		}
	}
}

func NoteListUp (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListUpReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.NoteListLevel = c.PostForm("NoteListLevel")
	noteListParam.Timestamp = c.PostForm("Timestamp")
	noteListParam.NoeListData = c.PostForm("NoteListData")
	noteListParam.NoteListOpt, _ = strconv.ParseBool(c.PostForm("NoteListOpt"))
	noteListParam.ModTime = c.PostForm("ModTime")
	noteListParam.NoteListShare, _ = strconv.ParseBool(c.PostForm("NoteListShare"))
	noteListParam.NoteListStatus,_ = strconv.ParseBool(c.PostForm("NoteListStatus"))
	noteListParam.NoteListTitle = c.PostForm("NoteListTitle")
	if (len(noteListParam.UserID) <= 0) || (len(noteListParam.Timestamp) <= 0){
		CreateErrorWithMsg(c, "Invalid arguments")
	} else {
		var noteListUpResp = &lovers_srv_notelist.NoteListUpResp{}
		noteListUpResp, err := notelist_client.NoteList_Up(c, noteListParam)
		if err != nil {
			CreateSuccess(c, noteListUpResp)
		} else {
			CreateErrorWithMsg(c, "NoteListUp failed error msg:" + err.Error())
		}
	}
}

func NoteListDown (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListDownReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.StartIndex, _ = strconv.ParseInt(c.PostForm("StartIndex"), 10, 64)
	noteListParam.NoteListCnt,_ = strconv.ParseInt(c.PostForm("NoteListCnt"), 10, 64)
	noteListParam.NoteListStatus, _= strconv.ParseBool(c.PostForm("NoteListStatus"))
	noteListParam.NoteListShare,_ = strconv.ParseBool(c.PostForm("NoteListShare"))
	if len(noteListParam.UserID) <= 0 {
		CreateErrorWithMsg(c, "UserID is empty")
	} else {
		noteListDownResp, err := notelist_client.NoteList_Down(c, noteListParam)
		if err != nil {
			CreateErrorWithMsg(c, "NoteListDown failed error msg:" + err.Error())
		} else {
			CreateSuccess(c, noteListDownResp)
		}
	}
}

func NoteListDel (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListDelReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.Timestamp = c.PostForm("Timestamp")
	if (len(noteListParam.UserID) <=0) || (len(noteListParam.Timestamp)<=0) {
		CreateErrorWithMsg(c, "Invalid arguments")
		return
	}

	noteListDelResp, err := notelist_client.NoteList_Del(c,noteListParam)
	if err != nil  {
		CreateErrorWithMsg(c, "NoteListDel failed error msg:" + err.Error())
	} else {
		CreateSuccess(c, noteListDelResp)
	}

}