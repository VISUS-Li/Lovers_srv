package handler

import (
	lovers_srv_notelist "Lovers_srv/server/note-list/proto"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"sync"
	"unsafe"
)

//PATH: NOTELISTPATHBASE + USERID + "Note"
const NOTELISTPATHBASE = "/opt/noteList/"

//定义互斥锁
var noteFileLock sync.Locker
//NoteListOpen() (*os.File, error)
//NoteListWrite() int
//NoteListRead() int
//NoteList索引: 内容编号， 内容长度
type NoteListIndexT struct {
	NoteNum int
	NoteLen int
}

//NoteList: 事件是否完成， 内容实体
type NoteListDateT struct {
	NoteOpt	bool
	NoteData []byte
}



type NoteListHandler struct {

}

//上传事件清单
func (notelist* NoteListHandler) NoteListUp(ctx context.Context, in *lovers_srv_notelist.NoteListUpReq, out *lovers_srv_notelist.NoteListUpResp) error {
	notelistindex := *(**NoteListIndexT)(unsafe.Pointer(&in.NoteListIndex))
	notePath := NOTELISTPATHBASE+in.UserId+"/"+ strconv.Itoa(notelistindex.NoteNum)
	fpNote, err := notelist.NoteListOpen(notePath)
	if err != nil {
		res := "打开事件清单失败"
		notelist.noteUpResp(out, res)
		logrus.Error("NoteList open file failed" + err.Error())
		return err
	}
	/*
	fpIndex, err := notelist.NoteListOpen(indexPath)
	if err != nil {
		res := "打开事件清单索引失败"
		notelist.NoteListClose(fpNote)
		notelist.noteUpResp(out, res)
		logrus.Error("Notelist open index failed" + err.Error())
		return err
	}*/

	//
	//调用写文件接口
	//
	notelist.NoteListWrite(fpNote, in.NoteListIndex, len(in.NoteListIndex))
	notelist.NoteListWrite(fpNote, in.NoeListData, len(in.NoeListData))


}

func (notelist* NoteListHandler) NoteListDown(ctx context.Context, in *lovers_srv_notelist.NoteListDownReq, out *lovers_srv_notelist.NoteListDownResp) error {
	notelistindex := *(**NoteListIndexT)(unsafe.Pointer(&in.NoteListIndex))
	notePath := NOTELISTPATHBASE+in.UserId+"/"+ strconv.Itoa(notelistindex.NoteNum)
	//indexPath := NOTELISTPATHBASE + in.UserId+ "/"+"Index"
	fpNote, err := notelist.NoteListOpen(notePath)
	if err != nil {
		res := "打开事件清单失败"
		notelist.noteDownFailResp(out, res)
		logrus.Error("NoteList open file failed" + err.Error())
		return err
	}
	/*
	fpIndex, err := notelist.NoteListOpen(indexPath)
	if err != nil {
		res := "打开事件清单索引失败"
		notelist.NoteListClose(fpNote)
		notelist.noteDownFailResp(out, res)
		logrus.Error("Notelist open index failed" + err.Error())
		return err
	}
	*/
	notelist.NoteListRead(fpNote, out.NoteListIndex, len(out.NoteListIndex))
	notelist.NoteListRead(fpNote, out.NoteListData, notelistindex.NoteLen)

}



func (notelist* NoteListHandler) FileIsExist(path string) (bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

//打开文件，文件路径/opt/UserId/Note
func (notelist* NoteListHandler) NoteListOpen(path string) (*os.File, error) {
	//fpNote, err := os.OpenFile(NOTELISTPATHBASE+req.UserId+"Note", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	logrus.Error("open NoteList file failed" + err.Error())
	//	return nil, err
	//}
	//path := NOTELISTPATHBASE+req.UserId+"Note"
	err := notelist.FileIsExist(path)
	if err { //file exist
		fpNote, err := os.OpenFile(path, os.O_RDWR | os.O_APPEND, 0666)
		if err != nil {
			logrus.Error("open NoteList file failed" + err.Error())
			return nil,err
		}
		return fpNote,err
	} else { //file note exist
		fpNote, err := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			logrus.Error("file not exist,create failed!" + err.Error())
			return nil, err
		}
		return fpNote, err
	}

}

//关闭文件
func (notelist* NoteListHandler) NoteListClose(fpNote *os.File) error {
	defer fpNote.Close()
	return nil
}

func (notelist* NoteListHandler) NoteListWrite(fpNote *os.File, data []byte, dataLen int) (int, error) {
	if fpNote == nil {
		return -1, errors.New("empty fpNote")
	}
	var writeLen = 0
	for writeLen < dataLen {
		ret, err := fpNote.Write(data)
		if err != nil {
			return -1, errors.New("writ file failed")
		}
		writeLen += ret
	}


}

func (notelist* NoteListHandler) NoteListRead(fpNote *os.File, data []byte, NoteNum int) (int , error) {
	if fpNote == nil {
		return -1, errors.New("empty fpNote")
	} else if data == nil {
		return -1, errors.New("empty data")
	}
	var stNoteIndex NoteListIndexT
	headLen := int(unsafe.Sizeof(stNoteIndex))
	var offset = NoteNum
	bufIndex := make([]byte, headLen)
	for readLen := 0 ; readLen < headLen{
		readLen, _ = fpNote.ReadAt(bufIndex, int64(offset))
	}

}


/**************
@brief NoteList upload response
@param NoteListUpResp: create in micro
		res: types of errors
**********/
func (notelist* NoteListHandler) noteUpResp(out *lovers_srv_notelist.NoteListUpResp, res string) {
	if res != "success" {
		out.NoteListUpResult = "failed"
		out.Err = res
		return
	}

	out.NoteListUpResult = "success"
	out.Err = "NULL"
}

/*
@brief NoteList download response
@param NoteListDownResp: create in micro
	    res: types of errors
@return error
 */
func (notelist* NoteListHandler)noteDownFailResp(out* lovers_srv_notelist.NoteListDownResp, res string) {
	out.NoteListDownResult = "failed"
	out.Err = res
	return
}

func (notelist *NoteListHandler)noteDownSuccessResp(out* lovers_srv_notelist.NoteListDownResp) {
	out.NoteListDownResult = "success"
	out.Err = "NULL"
}





