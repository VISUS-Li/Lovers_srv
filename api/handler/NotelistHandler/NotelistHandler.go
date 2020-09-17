package NotelistHandler

import (
	"Lovers_srv/helper/Utils"
	noteListClient "Lovers_srv/server/note-list/client"
	lovers_srv_notelist "Lovers_srv/server/note-list/proto"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (

	//user_handler = handler.UserHandler{}
	notelist_client = noteListClient.NewNoteListClient()
)

func NoteListUp (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListUpReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.NoteListLevel = c.PostForm("NoteListLevel")
	noteListParam.Timestamp = c.PostForm("Timestamp")
	noteListParam.NoeListData = c.PostForm("NoteListData")
	if (len(noteListParam.UserID) <= 0) || (len(noteListParam.Timestamp) <= 0){
		Utils.CreateErrorWithMsg(c, "Invalid arguments")
	} else {
		var noteListUpResp = &lovers_srv_notelist.NoteListUpResp{}
		noteListUpResp, err := notelist_client.NoteList_Up(c, noteListParam)
		if err != nil {
			Utils.CreateSuccess(c, noteListUpResp)
		} else {
			Utils.CreateErrorWithMsg(c, "NoteListUp failed error msg:" + err.Error())
		}
	}
}

func NoteListDown (c *gin.Context) {
	noteListParam := &lovers_srv_notelist.NoteListDownReq{}
	noteListParam.UserID = c.PostForm("UserID")
	noteListParam.StartIndex, _ = strconv.ParseInt(c.PostForm("StartIndex"), 10, 64)
	noteListParam.NoteListCnt,_ = strconv.ParseInt(c.PostForm("NoteListCnt"), 10, 64)
	if len(noteListParam.UserID) <= 0 {
		Utils.CreateErrorWithMsg(c, "UserID is empty")
	} else {
		noteListDownResp, err := notelist_client.NoteList_Down(c, noteListParam)
		if err != nil {
			Utils.CreateErrorWithMsg(c, "NoteListDown failed error msg:" + err.Error())
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
		Utils.CreateErrorWithMsg(c, "Invalid arguments")
		return
	}

	noteListDelResp, err := notelist_client.NoteList_Del(c,noteListParam)
	if err != nil  {
		Utils.CreateErrorWithMsg(c, "NoteListDel failed error msg:" + err.Error())
	} else {
		Utils.CreateSuccess(c, noteListDelResp)
	}

}
