package NotelistHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	noteListClient "Lovers_srv/server/lovers-tools/note-list/client"
	lovers_srv_notelist "Lovers_srv/server/lovers-tools/note-list/proto"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	notelist_client = noteListClient.NewNoteListClient()
)

func NoteListUp (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListUpReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.NoteListStatus, _ = strconv.ParseBool(c.PostForm("NoteListStatus"))
	noteListParam.NoteListLevel = c.PostForm("NoteListLevel")
	noteListParam.NoteListOpt, _ = strconv.ParseBool(c.PostForm("NoteListOpt"))
	noteListParam.NoteListTitle = c.PostForm("NoteListTitle")
	noteListParam.Timestamp = c.PostForm("Timestamp")
	noteListParam.ModTime = c.PostForm("ModTime")
	noteListParam.NoteListShare, _ = strconv.ParseBool(c.PostForm("NoteListShare"))
	noteListParam.NoteListData = c.PostForm("NoteListData")

	if (len(noteListParam.UserID) <= 0) ||
		(len(noteListParam.Timestamp) <= 0) ||
		(len(noteListParam.NoteListTitle) <= 0) ||
		(len(noteListParam.Timestamp) <= 0) ||
		(len(noteListParam.NoteListData) < 0){
			Utils.CreateErrorWithMsg(c, "Invalid arguments",config.INVALID_PARAMS)
		} else {
			var noteListUpResp = &lovers_srv_notelist.NoteListUpResp{}
			noteListUpResp, err := notelist_client.NoteList_Up(c, noteListParam)
		if err != nil {
			Utils.CreateErrorWithMsg(c, "NoteListUp failed error msg:" + err.Error(),config.CODE_ERR_UNKNOW)
		} else {
			Utils.CreateSuccess(c, noteListUpResp)
		}
	}
}

func NoteListDown (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListDownReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.NoteListStatus, _ = strconv.ParseBool(c.PostForm("NoteListStatus"))
	noteListParam.NoteListShare, _ = strconv.ParseBool(c.PostForm("NoteListShare"))
	noteListParam.StartIndex, _ = strconv.ParseInt(c.PostForm("StartIndex"), 10, 64)
	noteListParam.NoteListCnt,_ = strconv.ParseInt(c.PostForm("NoteListCnt"), 10, 64)
	if len(noteListParam.UserID) <= 0 {
		Utils.CreateErrorWithMsg(c, "UserID is empty",config.CODE_ERR_PARAM_EMPTY)
	} else {
		noteListDownResp, err := notelist_client.NoteList_Down(c, noteListParam)
		if err != nil {
			Utils.CreateErrorWithMsg(c, "NoteListDown failed error msg:" + err.Error(),config.CODE_ERR_UNKNOW)
		} else {
			Utils.CreateSuccess(c, noteListDownResp)
		}
	}
}

func NoteListDel (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListDelReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.Timestamp = c.PostForm("Timestamp")
	if (len(noteListParam.UserID) <=0) || (len(noteListParam.Timestamp)<=0) {
		Utils.CreateErrorWithMsg(c, "Invalid arguments",config.INVALID_PARAMS)
		return
	}

	noteListDelResp, err := notelist_client.NoteList_Del(c,noteListParam)
	if err != nil  {
		Utils.CreateErrorWithMsg(c, "NoteListDel failed error msg:" + err.Error(),config.CODE_ERR_UNKNOW)
	} else {
		Utils.CreateSuccess(c, noteListDelResp)
	}

}
