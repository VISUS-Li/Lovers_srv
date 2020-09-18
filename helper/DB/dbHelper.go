package DB

import (
	"Lovers_srv/config"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"os"
)

type DBUtil struct{
	DB *gorm.DB
}

func (unit *DBUtil)CreateTable(tableModel interface{})(error){
	unit.DB = unit.DB.AutoMigrate(tableModel)
	return unit.DB.Error
}

func (unit *DBUtil)CreateConnect(dbName string)(error){

	//先用环境变量的设置，没有再使用默认设置
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	if len(host) <= 0 || len(user) <= 0 || len(pwd) <= 0{
		host = config.GlobalConfig.DB_host
		user = config.GlobalConfig.DB_user
		pwd  = config.GlobalConfig.DB_password
		logrus.Info("Env host/user/pwd is nil,use Config")
	}
	if(len(dbName) <= 0){
		var err = "Create DB Fail,dbName is empty"
		logrus.Error(err)
		return errors.New(err)
	}
	return unit.GetDBByHost(user, pwd, host, dbName)
}

func (unit *DBUtil) CloseConnect()(error){
	return unit.DB.Close()
}

func (unit *DBUtil)GetDBByHost(user string, password string, host string, dbName string) (error){
	var mysql = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True",
		user, password, host,dbName)
	var err error
	unit.DB,err = gorm.Open("mysql",mysql)
	if err != nil{
		logrus.Error("connect DB error:"+err.Error())
		logrus.Error("DB connect string:"+mysql)
		logrus.Error("DB Host Info: username:"+user + ",password:"+password + ",hostname:"+host + "\n")
		return err
	}
	logrus.Info("DB connect success")
	logrus.Info("DB connect string:"+mysql)
	logrus.Info("DB Host Info: username:"+user + ",password:"+password + ",hostname:"+host + "\n")

	unit.DB.DB().SetMaxIdleConns(10)
	unit.DB.DB().SetMaxOpenConns(100)
	unit.DB.AutoMigrate(&UserBaseInfo{})
	unit.DB.AutoMigrate(&LoginInfo{})
	return nil
}