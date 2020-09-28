package Utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
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
从请求的Form中获取参数
******/
//func GetParamFromForm(c *gin.Context, stu interface{})(interface{},error){
//	fields, err := GetFieldName(&stu)
//	if err != nil{
//		return nil, err
//	}
//	if len(fields) <= 0{
//		return nil, errors.New("parse struct field fail, fields are empty")
//	}
//
//	for i, v := range fields{
//		value := reflect.ValueOf(stu)
//		postValue := c.PostForm(v)
//		convPostV, err := strconv.Atoi(postValue)
//		if err != nil{
//			reflect.ValueOf(postValue).Type()
//		}
//
//
//		fieldType := value.Type().String()
//		if fieldType == "int"{
//
//		}
//
//		value.FieldByName(v).Set(postValue)
//	}
//
//}