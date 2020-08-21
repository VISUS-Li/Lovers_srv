package DB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type DBUtil struct{
	DB *gorm.DB
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
	unit.DB.AutoMigrate(&UserBaseInfo{})
	unit.DB.AutoMigrate(&LoginInfo{})
	return nil
}