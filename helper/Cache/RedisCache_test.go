package Cache

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"testing"
)


func TestNewRedisPool(t *testing.T) {
	config.Init()
	NewRedisPool();
	data := &DB.LoginInfo{
		UserId:"123",
		UserName:"VISUS",
		PassWord:"123456",
		Phone:"15002326233",
	}
	err := SetHashByJson("login","123",data);
	if err != nil{
		Utils.ErrorOutputf("set hash cache failed:%s",err.Error())
	}else{
		logrus.Info("添加hash成功!")
	}

	jsonBytes, err := GetHashByJson("login","123")
	if err != nil{
		Utils.ErrorOutputf("get has cache failed:%s",err.Error())
	}else{
		logrus.Info("获取hash成功!")
	}
	if(jsonBytes != nil) {
		login := new(DB.LoginInfo)
		err = json.Unmarshal(jsonBytes, login)
		if err != nil {
			Utils.ErrorOutputf("UnMarshal cache failed:%s", err.Error())
		}else{
			logrus.Infof("UnMarshal json成功:%s, %s, %s, %s",login.UserId,login.UserName,login.PassWord,login.Phone)
		}
	}
}

