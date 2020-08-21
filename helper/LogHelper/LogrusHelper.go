package LogHelper

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type LoversLog struct{

}


func (l *LoversLog)SetOutPut(srvName string) (error){
	execPath,err := GetCurrentExecPath()
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
	logrus.SetOutput(file)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)
	return err
}


func GetCurrentExecPath()(string,error){
	file,err := exec.LookPath(os.Args[0])
	if err != nil{
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil{
		return "", err
	}

	if runtime.GOOS =="windows"{
		path = strings.Replace(path, "\\", "/", -1)
	}

	i := strings.LastIndex(path, "/")
	if i < 0{
		return "", errors.New(`Can't find "/" or  "\".`)
	}

	return string(path[0:i+1]), nil
}