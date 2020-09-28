package FilerServerHandler

import (
	fileServerClient "Lovers_srv/server/fileServer/client"
	lovers_srv_file "Lovers_srv/server/fileServer/proto"
	"github.com/gin-gonic/gin"
)

var (
	fileserver_client = fileServerClient.NewFileServerClient()
	)

func DownLoadFile(c *gin.Context) {
	fileServerParam := &lovers_srv_file.DownLoadFileReq{}
	/*
	fileServerParam.DownLoadFileUrl = c.PostForm("DoWnLoadFileUrl")
	fileServerParam.UserID 			= c.PostForm("UserID")
	*/
	c.ShouldBind(&fileServerParam)


}

func UpLoadFile(c *gin.Context) {

}

func DelFile(c *gin.Context) {

}


