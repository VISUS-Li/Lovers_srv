package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"

	proto "Lovers_srv/server/fileserver/proto"
)

const fileBasePath = "/user/LoversFile"

type FileHandler struct {
	DB *gorm.DB
}


func (file *FileHandler) DownLoadFile(ctx context.Context, in *proto.DownLoadFileReq, out *proto.DownLoadFileResp) error {
	if (len(in.UserID) <= 0) ||
		(len(in.DownLoadFileUrl) <= 0)  {
		out.UserID 					= in.UserID
		out.FileBinData 			= nil
		out.DownLoadFileRet 		= "failed"
		out.FileSha 				= ""
		out.RespStatus.FileRespMsg  = config.MSG_ERR_PARAM
		out.RespStatus.FileRespCode = strconv.Itoa(config.CODE_ERR_GET_PARAM_)

		logrus.Error("error param")
		err := errors.New("error Download File param")
		return err
	}

	openFile, err := os.Open(in.DownLoadFileUrl)
	if err != nil {
		logrus.Error("cannot open file: " + err.Error())
		out.UserID					= in.UserID
		out.FileBinData				= nil
		out.FileSha					= ""
		out.DownLoadFileRet			= "failed"
		out.RespStatus.FileRespCode	= strconv.Itoa(config.CODE_ERR_SERVER_INTERNAL)
		out.RespStatus.FileRespMsg  = config.MSG_ERR_OPENFILE

		return err
	}

	out.FileBinData, err = ioutil.ReadAll(openFile)
	if err != nil {
		out.UserID					= in.UserID
		out.FileBinData				= nil
		out.FileSha					= ""
		out.DownLoadFileRet			= "failed"
		out.RespStatus.FileRespCode	= strconv.Itoa(config.CODE_ERR_SERVER_INTERNAL)
		out.RespStatus.FileRespMsg  = config.MSG_ERR_OPENFILE

		return err
	}

	out.UserID 						= in.UserID
	out.DownLoadFileRet 			= "success"
	out.FileSha						= Utils.Sha1(out.FileBinData)
	out.RespStatus.FileRespMsg		= config.MSG_FILE_DOWN_OK
	out.RespStatus.FileRespCode		= ""

	return nil
}

func (file *FileHandler) UpLoadFile(ctx context.Context, in *proto.UpLoadFileReq, out *proto.UpLoadFileResp) error {
	if (len(in.UserID) <= 0) ||
		(len(in.FileInModule) <= 0) ||
		(len(in.FileType) <= 0) ||
		(len(in.FileName) <= 0) {
		out.UserID = in.UserID
		out.FileSha1 = ""
		out.UpLoadFileRet = "failed"
		out.RespStatus.FileRespMsg = config.MSG_SERVER_INTERNAL
		out.RespStatus.FileRespCode = strconv.Itoa(config.CODE_ERR_BAD_UP_RESPONSE)
		err := errors.New("error param")
		logrus.Error(err.Error())
		return err
	}

	var fileDir = fileBasePath + "/" + in.FileInModule + "/" + in.UserID + "/" + in.FileType + "/"
	var filePath = fileDir + in.FileName
	exist, _ := Utils.PathExists(filePath)

	if !exist {
		err := os.Mkdir(fileDir, os.ModePerm)
		if err != nil {
			logrus.Error("fileDir create failed!")
			out.UserID 				= in.UserID
			out.FileSha1 			= ""
			out.UpLoadFileRet 		= "failed"
			out.UpLoadFileUrl 		= ""

			return err
		}
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		logrus.Error("Failed to create a new file:" + err.Error())
		out.UserID 					= in.UserID
		out.FileSha1 				= ""
		out.UpLoadFileUrl 			= ""
		out.UpLoadFileRet 			= "failed"
		out.RespStatus.FileRespMsg  = config.MSG_SERVER_INTERNAL
		out.RespStatus.FileRespCode = strconv.Itoa(config.CODE_ERR_UNKNOW)

		return err
	}

	nByte, err := newFile.Write(in.FileBinData)
	if 	int64(nByte) != in.FileSize || err != nil {
		logrus.Error("Failed save data to file, write size:%d, but need: %d. err:" + err.Error(), nByte, in.FileSize)
		out.UserID 					= in.UserID
		out.FileSha1 				= in.FileSha1
		out.UpLoadFileRet 			= "failed"
		out.UpLoadFileUrl 			= ""
		out.RespStatus.FileRespCode = strconv.Itoa(config.CODE_ERR_UNKNOW)
		out.RespStatus.FileRespMsg  = config.MSG_SERVER_INTERNAL

		return err
	}

	newFile.Seek(0,0)

	out.UserID 					= in.UserID
	out.FileSha1				= Utils.FileSha1(newFile)
	out.UpLoadFileUrl			= filePath
	out.UpLoadFileRet			= "success"
	out.RespStatus.FileRespMsg  = config.MSG_FILE_UP_OK
	out.RespStatus.FileRespCode = ""


	//save the file information to database
	fileUpDB := DB.FileServerInfo{
		UserID:       out.UserID,
		FileUrl:      out.UpLoadFileUrl,
		FileInModule: in.FileInModule,
		FileType:     in.FileType,
		FileName:     in.FileName,
		FileSha1:     out.FileSha1,
	}
	file.DB.Create(fileUpDB)

	return nil
}

func (file *FileHandler) DelFile(ctx context.Context, in *proto.DelFileReq, out *proto.DelFileResp) error {
	if (len(in.UserID) <= 0) ||
		(len(in.FileUrl) <= 0) {
		out.UserID 						= in.UserID
		out.DelFileRet					= "failed"
		out.RespStatus.FileRespCode		= strconv.Itoa(config.CODE_ERR_GET_PARAM_)
		out.RespStatus.FileRespMsg		= config.MSG_ERR_PARAM

		err := errors.New("Invalid Param")
		logrus.Error("error delete file param")
		return err
	}

	err := os.Remove(in.FileUrl)
	if err != nil {
		logrus.Error("delete file failed, error: " + err.Error())
		out.UserID 						= in.UserID
		out.DelFileRet					= "failed"
		out.RespStatus.FileRespMsg		= config.MSG_SERVER_INTERNAL
		out.RespStatus.FileRespCode		= strconv.Itoa(config.CODE_ERR_FAILED_DELFILE)

		return err
	}

	err = file.DB.Delete("UserID = ? and FileUrl = ?",
		in.UserID,	in.FileUrl).Error
	if err != nil {
		out.UserID						= in.UserID
		out.DelFileRet					= "failed"
		out.RespStatus.FileRespCode		= strconv.Itoa(config.CODE_DB_DELETE_FAILED)
		out.RespStatus.FileRespMsg		= config.MSG_ERR_DELETE_DB

		logrus.Error("file delete database failed:" + err.Error())

		return err
	}

	out.UserID 							= in.UserID
	out.DelFileRet						= "success"
	out.RespStatus.FileRespMsg			= config.MSG_FILE_DEL_OK
	out.RespStatus.FileRespCode			= ""

	return nil
}