package Utils

import (
	"Lovers_srv/config"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func VerifyEmailFormat(email string) bool{
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyPhoneFormat(phone string) bool{
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(pattern)
	return  reg.MatchString(phone)
}
func VerifyUUIDFormat(uuid string) bool{
	pattern := "^[A-Fa-f0-9]{8}(-[A-Fa-f0-9]{4}){3}-[A-Fa-f0-9]{12}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(uuid)
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
	if runtime.GOOS !="windows"{
		path = strings.Replace(path, "\\", "/", -1)
	}

	i := strings.LastIndex(path, "/")
	if i < 0{
		return "", errors.New(`Can't find "/" or  "\".`)
	}

	return string(path[0:i]), nil
}

/******
传入文件名，得到exe目录下该文件名的全路径
******/
func GetExeDstFileName(dstName string)(string){
	exePath,_ := GetCurrentExecPath()
	path := exePath + "\\" + dstName
	if runtime.GOOS !="windows"{
		path = strings.Replace(path, "\\", "/", -1)
	}
	return path
}

/******
通过服务名获取数据库名称
******/
func GetDBNameFromSrvName(srvName string)(string){
	splitRes := strings.Split(srvName,".")
	splitLen := len(splitRes)
	return splitRes[splitLen-1]
}



/******
获取结构体中字段的名称
******/
func GetFieldName(stu *interface{}) (map [int]string, error) {

	t := reflect.TypeOf(stu)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil,errors.New("Check type error not Struct")
	}
	fieldNum := t.NumField()
	fieldList := make(map[int]string)
	for i := 0; i < fieldNum; i++ {
		fieldList[i] = t.Field(i).Name
	}
	return fieldList,nil
}

/******
	从微服务返回的错误中分割出错误码和错误信息
******/
func SplitMicroErr(err error) (string, int){
	if err == nil{
		return "",0
	}
	errInfo := err.Error()
	errVec := strings.Split(errInfo,"_")
	var msg string
	var code int
	if len(errVec) > 0 {
		msg = errVec[0]
	}else{
		msg = config.MSG_SERVER_INTERNAL
		code = config.CODE_ERR_SERVER_INTERNAL
	}
	if len(errVec) > 1{
		codeErr, covErr := strconv.Atoi(errVec[1])
		if covErr != nil{
			logrus.Errorf("分割错误码失败，err:%s",covErr)
			return config.MSG_SERVER_INTERNAL,config.CODE_ERR_SERVER_INTERNAL
		}
		code = codeErr
	}else{
		code = config.CODE_ERR_SERVER_INTERNAL
	}


	return msg, code
}

/******
	构建微服务返回的错误格式
******/
func MicroErr(msg string, code int) error{
	return errors.New(msg + "_" + strconv.Itoa(code))
}

/******
ErrorOutputf，构建error的同时，将这个error的信息打印到日志中，支持格式化字符串
******/
func ErrorOutputf(format string, args ...interface{}) error{
	msg := fmt.Sprintf(format, args...)
	if(config.GlobalConfig.RunMode == config.RUNMODE_DEV){
		logrus.Error(msg)
	}
	return errors.New(msg)
}

/******
	处理mysql返回的错误，并将错误码转换成自定义错误码
******/
func ErrorOutputMysqlf(defCode int, err error,format string, args ...interface{}) (int, error){
	msg := fmt.Sprintf(format, args...)
	if(config.GlobalConfig.RunMode == config.RUNMODE_DEV){
		logrus.Error(msg)
	}
	t := reflect.TypeOf(err)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	var errCode int
	if t.Name() == "MySQLError"{
		mySqlErr := err.(*mysql.MySQLError)
		//将mysql错误码，转换成自定义错误码
		switch mySqlErr.Number {
		case 1062:
			errCode = config.ENUM_ERR_DB_INSERT_DUPLICATE //插入时，主键重复
			break;
		default:
			errCode = defCode
			if(config.GlobalConfig.RunMode == config.RUNMODE_DEV){
				logrus.Errorf("MYSQLError msg:%, code:%d",mySqlErr.Message, mySqlErr.Number)
			}
			break;
		}
		return errCode, mySqlErr
	}else{
		if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND_ENG{
			errCode = config.ENUM_ERR_DB_QUERY_NOT_FOUND
		}else{
			errCode = defCode
		}
		return errCode, err
	}
}

func ParseError(){

}
