package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	lovers_srv_notelist "Lovers_srv/server/note-list/proto"
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
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
	DB *gorm.DB
}

//上传事件清单
func (notelist* NoteListHandler) NoteListUp(ctx context.Context, in *lovers_srv_notelist.NoteListUpReq, out *lovers_srv_notelist.NoteListUpResp) error {
	if (len(in.UserID) <= 0) ||
		(len(in.Timestamp) <= 0) ||
		(len(in.NoteListTitle) <= 0) ||
		(len(in.NoteListData) <= 0) {
		out.UserID = in.UserID
		notelist.uPFailResp(out, config.NOTELISTUP_INVALID_PARAM)
		return errors.New("invalid param")
	}
	out.UserID = in.UserID
	//创建数据库结构体
	var noteListInfo DB.NoteListDB
	noteListInfo.UserID = in.UserID
	noteListInfo.Timestamp = in.Timestamp
	noteListInfo.NoteListStatus = in.NoteListStatus
	noteListInfo.NoteListShare = in.NoteListShare
	noteListInfo.NoteListTitle = in.NoteListTitle
	noteListInfo.NoteListData = in.NoteListData
	noteListInfo.ModTime = in.ModTime
	noteListInfo.BackImage = in.BackImage

	if in.NoteListOpt { //新建事件清单
		if err := notelist.NoteListNew(&noteListInfo); err != nil {
			logrus.Error(err.Error())
			notelist.uPFailResp(out, err.Error())
			return err
		}

	} else { //修改事件清单
		if err := notelist.NoteListMod(&noteListInfo); err != nil {
			logrus.Error(err.Error())
			notelist.uPFailResp(out, err.Error())
			return err
		}
	}
	return nil
}

//下载事件清单
func (notelist* NoteListHandler) NoteListDown(ctx context.Context, in *lovers_srv_notelist.NoteListDownReq, out *lovers_srv_notelist.NoteListDownResp) (error) {
	if len(in.UserID) <= 0 {
		err := errors.New("UserID不能为空")
		logrus.Error(err.Error())
		return  err
	}

	out.UserID = in.UserID;

	var getAllList []DB.NoteListDB
	notelist.DB.Where("UserID = ?", in.UserID).Find(&getAllList)
	if len(getAllList) > 0 {
		//未来查找到事件，获取失败
		err := errors.New("没有查找到事件,请求失败")
		return err
	}

	for _, temp := range getAllList {
		var list lovers_srv_notelist.AllNoteList
		list.UserID = temp.UserID
		list.Timestamp = temp.Timestamp
		list.NoteListShare = temp.NoteListShare
		list.ModTime = temp.ModTime
		list.NoteListTitle = temp.NoteListTitle
		list.NoteListLevel = temp.NoteListLevel
		list.NoteListStatus = temp.NoteListStatus
		list.BackImage = temp.BackImage
		list.NoeListData = temp.NoteListData

		out.NoteList = append(out.NoteList, &list)
	}
	return  nil
}

//删除事件清单
func (notelist* NoteListHandler) NoteListDel(ctx context.Context, in *lovers_srv_notelist.NoteListDelReq, out *lovers_srv_notelist.NoteListDelResp) error {
	if len(in.UserID) <= 0 {
		out.NoteListDelRet = "failed"
		out.Err = "UserId不能为空"
	}
	if len(in.Timestamp) <= 0{
		out.NoteListDelRet = "failed"
		out.Err = "Timestamp不能为空"
	}

	err := notelist.DB.Delete("UserId = ? and Timestamp = ?",
		in.UserID, in.Timestamp).Error
	if err != nil {
		out.Err = err.Error()
	} else {
		out.Err = ""
	}
	out.NoteListDelRet = "success"
	return nil
}

//新建事件
func (notelist* NoteListHandler) NoteListNew(dbinfo* DB.NoteListDB) error {
	//查询该UserID和时间戳是否存在
	//接收查询结果
	var haveNoteList []DB.NoteListDB
	notelist.DB.Where("UserID = ? and Timestamp = ?",
		dbinfo.UserID,dbinfo.Timestamp).Find(&haveNoteList)
	if len(haveNoteList) > 0 {
		//存在该事件，结果与请求不符，失败
		err :=errors.New("操作请求参数出错，该事件已经存在，停止新建")
		return err
	}

	if err := notelist.DB.Create(dbinfo); err != nil {
		err := errors.New("insert notelistDB to db failed")
		return err
	}

	return nil
}

func (notelist* NoteListHandler) NoteListMod(dbinfo* DB.NoteListDB) error {
	var haveNoteList []DB.NoteListDB
	//查找该事件是否存在
	notelist.DB.Where("UserId = ? and Timestamp = ?",
		dbinfo.UserID, dbinfo.Timestamp).Find(&haveNoteList)

	fmt.Println("NoteListMod ---> request quantity of notelist: ",
		len(haveNoteList))

	//不存在该事件，更新失败
	if len(haveNoteList) <= 0 {
		err := errors.New("未查询到该事件，修改失败")
		return err
	}

	err := notelist.DB.Where("UserId = ? and Timestamp = ?",
		dbinfo.UserID, dbinfo.Timestamp).Update(DB.NoteListDB{
		NoteListStatus: dbinfo.NoteListStatus,
		NoteListLevel:  dbinfo.NoteListLevel,
		NoteListTitle:  dbinfo.NoteListTitle,
		ModTime:        dbinfo.ModTime,
		NoteListShare:  dbinfo.NoteListShare,
		NoteListData:   dbinfo.NoteListData,
	}).Error
	if err != nil {
		return nil
	}


	return nil
}



func (notelist *NoteListHandler) uPFailResp(out *lovers_srv_notelist.NoteListUpResp, res string) {
	if res != "success" {
		out.NoteListUpResult = "failed"
	}

	out.BackImage = ""
}



/********************************************************************/
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


