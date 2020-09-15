package handler

import (
	lovers_srv_notelist "Lovers_srv/server/note-list/proto"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

//PATH: NOTELISTPATHBASE + USERID + "Note"
const NOTELISTPATHBASE = "/opt/noteList/"

//定义互斥锁
var noteFileLock sync.Locker
var noteReadLock sync.Locker
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
}

//下载事件清单
func (notelist* NoteListHandler) NoteListDown(ctx context.Context, in *lovers_srv_notelist.NoteListDownReq, out *lovers_srv_notelist.NoteListDownResp) error {
}

//删除事件清单
func (notelist* NoteListHandler) NoteListDel(ctx context.Context, in *lovers_srv_notelist.NoteListDelReq, out *lovers_srv_notelist.NoteListDelResp) error {

}



func (notelist* NoteListHandler) FileIsExist(path string) bool {
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

func (notelist* NoteListHandler) NoteListDir(path string) error {
	_,err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			return err
		}
	}
	return err
}

//打开文件，文件路径/opt/UserId/Note
func (notelist* NoteListHandler) NoteListOpen(path string) (*os.File, error) {
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
func (notelist* NoteListHandler) NoteListClose(fpNote *os.File) error{
	err := fpNote.Close()
	return err
}

func (notelist* NoteListHandler) NoteListWrite(fpNote *os.File, data []byte, dataLen int) (int, error) {
	if fpNote == nil {
		return -1, errors.New("empty fpNote")
	}
	var writeLen = 0
	for writeLen < dataLen {
		ret, err := fpNote.Write(data)
		if err != nil{
			return -1, errors.New("writ file failed")
		}
		writeLen += ret
	}

	return writeLen, nil
}

func (notelist* NoteListHandler) NoteListRead(fpNote *os.File, data []byte,  dateLen int) (int , error) {
	if fpNote == nil {
		return -1, errors.New("empty fpNote")
	} else if data == nil {
		return -1, errors.New("empty data")
	}

	readLen, err := fpNote.Read(data)
	if err != nil {
		logrus.Error("read noteList failed")
		return -1, err
	}

	return readLen,nil

}


