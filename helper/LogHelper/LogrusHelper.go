package LogHelper

import (
	"Lovers_srv/helper/Utils"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type LoversLog struct{

}


func (l *LoversLog)SetOutPut(srvName string) (error){
	execPath,err := Utils.GetCurrentExecPath()
	if err != nil{
		return err
	}
	if srvName == ""{
		srvName = "global.log"
	}else {
		i := strings.LastIndex(srvName, "/")
		j := strings.LastIndex(srvName, "\\")
		if i >= 0 || j >= 0{
			return errors.New("Loverslog SetOutPut param only can be server name")
		}
		srvName = strings.Replace(srvName,".","_",-1)
		srvName += ".log"
	}
	logPath := execPath + "Logs/"
	_, err = os.Stat(logPath)
	if err != nil{
		os.MkdirAll(logPath, os.ModePerm)
	}
	logFilePath := logPath +  srvName

	file, err := os.OpenFile(logFilePath, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		logrus.SetOutput(fileAndStdoutWriter)
	} else {
		logrus.SetOutput(file)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	return err
}




func OutPutErrorInfo(str string) {
	errors.New(str)
	logrus.Error(str)
}