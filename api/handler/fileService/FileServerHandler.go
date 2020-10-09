package FilerServerHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	fileServerClient "Lovers_srv/server/fileServer/client"
	lovers_srv_file "Lovers_srv/server/fileServer/proto"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

var (
	fileserver_client = fileServerClient.NewFileServerClient()
	)

type fileDownGinParam_t struct{
	UserID 			string `form:"UserID"`
	DownLoadFileUrl	string `form:"DownLoadFileUrl"`
}

type fileUpGinParam_t struct {
	UserID 			string `form:"UserID"`
	FileInModule	string `form:"FileInModule"`
	FileType 		string `form:"FileType"`
	FileName		string `form:"FileName"`
}

type fileDelGinPatam_t struct {
	UserID 			string `form:"UserID"`
	FileUrl 		string `form:"FileUrl"`
}

func DownLoadFile(c *gin.Context) error {
	//use to proto
	fileServerParam := &lovers_srv_file.DownLoadFileReq{}
	//use to post
	fileDownGinParam := fileDownGinParam_t{}
	/*
	fileServerParam.DownLoadFileUrl = c.PostForm("DoWnLoadFileUrl")
	fileServerParam.UserID 			= c.PostForm("UserID")
	*/
	err := c.ShouldBind(&fileDownGinParam)
	if err != nil {
		logrus.Error("gin should bind failed:"+ err.Error())
		return err
	}

	if len(fileDownGinParam.UserID) <=0 ||
		len(fileDownGinParam.DownLoadFileUrl) <=0 {
		Utils.CreateErrorWithMsg(c,"Invalid arguments", config.INVALID_PARAMS)
	} else {
		fileServerParam.UserID = fileDownGinParam.UserID
		fileServerParam.DownLoadFileUrl = fileDownGinParam.DownLoadFileUrl
		fileDownresp, err := fileserver_client.Client_DownFile(c, fileServerParam)
		if err == nil && fileDownresp != nil { //no error & get response
			if len(fileDownresp.UserID) >= 0 &&
				len(fileDownresp.DownLoadFileRet) >= 0 { //error response param
					Utils.CreateErrorWithMsg(c, "error file download response", config.CODE_ERR_BAD_DOWN_RESPONSE)
			} else {	//file micro server return normal
			}
		} else {
			if	fileDownresp == nil {
				Utils.CreateErrorWithMsg(c, "error file download response", config.CODE_ERR_BAD_DOWN_RESPONSE)
			} else {
				Utils.CreateErrorWithMsg(c, err.Error(),config.CODE_ERR_BAD_DOWN_RESPONSE)
			}
		}
	}

	return nil
}

func UpLoadFile(c *gin.Context) error {
	//use to micro
	fileProtoParam := &lovers_srv_file.UpLoadFileReq{}
	//use to proto
	fileUpGinPatam := fileUpGinParam_t{}

	err := c.ShouldBind(&fileUpGinPatam)
	if err != nil {
		logrus.Error("gin should bind failed:" + err.Error())
		Utils.CreateErrorWithMsg(c, err.Error(), config.CODE_ERR_GET_PARAM_)
		return err
	}
	file, header, err := c.Request.FormFile("file")
	if err == nil {
		logrus.Error("failed to get file data, error:" + err.Error())
		Utils.CreateErrorWithMsg(c, err.Error(), config.CODE_ERR_GET_PARAM_)
		return err
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		logrus.Error("failed to get file data, err:" + err.Error())
		Utils.CreateErrorWithMsg(c, err.Error(), config.CODE_ERR_FAILED_UPFILE)
		return err
	}

	fileProtoParam.UserID 		= fileUpGinPatam.UserID
	fileProtoParam.FileInModule = fileUpGinPatam.FileInModule
	fileProtoParam.FileType 	= fileUpGinPatam.FileType
	fileProtoParam.FileName 	= header.Filename
	fileProtoParam.FileSize 	= header.Size
	fileProtoParam.FileBinData  = buf.Bytes()

	if (len(fileProtoParam.UserID) <= 0) ||
		(len(fileProtoParam.FileInModule) <= 0) ||
		(len(fileProtoParam.FileType) <= 0) ||
		(len(fileProtoParam.FileName) <= 0) {
		Utils.CreateErrorWithMsg(c, "Invalid argument", config.INVALID_PARAMS)
	} else {
		var FileUpResp = &lovers_srv_file.UpLoadFileResp{}
		FileUpResp, err = fileserver_client.Client_UpFile(c, fileProtoParam)
		if err == nil {
			Utils.CreateSuccess(c, FileUpResp)
		} else {
			Utils.CreateErrorWithMsg(c, "File upload failed, error msg: " + err.Error(),config.CODE_ERR_UNKNOW)
		}
	}
	return nil
}

func DelFile(c *gin.Context) error {
	var fileProtoParam = &lovers_srv_file.DelFileReq{}
	var fileDelGinPatam = fileDelGinPatam_t{}

	err := c.ShouldBind(fileDelGinPatam)
	if err != nil {
		logrus.Error("gin should bind failed: " + err.Error())
		Utils.CreateErrorWithMsg(c,err.Error(), config.CODE_ERR_GET_PARAM_)
		return err
	}

	if (len(fileDelGinPatam.UserID) <= 0) ||
		(len(fileDelGinPatam.FileUrl) <= 0) {
		logrus.Error("Invalid argument", config.INVALID_PARAMS)
		Utils.CreateErrorWithMsg(c, "Invalid argument", config.INVALID_PARAMS)
	} else {
		fileProtoParam.FileUrl = fileDelGinPatam.FileUrl
		fileProtoParam.UserID = fileDelGinPatam.UserID
		var FileDelResp = &lovers_srv_file.DelFileResp{}
		FileDelResp, err = fileserver_client.Clien_DelFile(c, fileProtoParam)
		if err == nil {
			Utils.CreateSuccess(c, FileDelResp)
		} else {
			Utils.CreateErrorWithMsg(c, "failed to delete file", config.CODE_ERR_UNKNOW)
			return err
		}
	}
	return nil
}


