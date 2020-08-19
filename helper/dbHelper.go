package helper

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type DBUtil struct{
	DB *gorm.DB
}

type UserBaseInfo struct{
	gorm.Model
	UserId 		string `gorm:"unique;not null"`
	RealName 	string
	Phone 		int
	Sex 		int /*1.男，2.女*/
	Birth 		string
	Sculpture	string /*头像地址*/
	HomeTown	string /*所在地*/
	LoverId			string
	LoverPhone		int
	LoverNickName 	string /*对另一半的昵称*/
	LoveDuration	int64 `gorm:type:BIGINT`
}

func (unit *DBUtil)CreateTable(tableModel interface{})(error){
	unit.DB = unit.DB.AutoMigrate(tableModel)
	return unit.DB.Error
}

func (unit *DBUtil)CreateConnect()(error){
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd  := os.Getenv("DB_PASSWORD")
	return unit.GetDBByHost(user, pwd, host)
}

func (unit *DBUtil) CloseConnect()(error){
	return unit.DB.Close()
}

func (unit *DBUtil)GetDBByHost(user string, password string, host string) (error){
	var mysql = fmt.Sprintf("%s:%s@tcp(%s:3306)/lovers?charset=utf8mb4&parseTime=True",
		user, password, host)
	var err error
	unit.DB,err = gorm.Open("mysql",mysql)
	if err != nil{
		return err
	}
	unit.DB.DB().SetMaxIdleConns(10)
	unit.DB.DB().SetMaxOpenConns(100)
	return nil
}