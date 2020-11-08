package UserCache

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"testing"
)

func initTest() error{
	config.Init()
	config.GlobalConfig.RunMode = config.RUNMODE_DEV
	dbUtil := DB.DBUtil{}
	err := dbUtil.CreateConnect("user")
	if err != nil{
		return Utils.ErrorOutputf("connect db failed:%s", err.Error())
	}
	Cache.NewRedisPool(dbUtil.DB)
	return nil
}

func TestSetUserLoginInfo(t *testing.T) {
	initTest()
	data := DB.LoginInfo{
		UserId:   "a91559d8-0c99-11eb-9650-00163e2ed191",
		PassWord: "abcdefg",
		Phone:    "15002323233",
	}
	_,err := SetUserLoginInfobyPhone(data.Phone,data,true)
	if err != nil{
		Utils.ErrorOutputf("通过手机号设置数据库中[已存在]用户登录信息失败，手机号为:%s, 用户ID为:%s, 错误信息为:%s", data.Phone, data.UserId, err.Error())
	}else{
		Utils.ErrorOutputf("通过手机号设置数据库中[已存在]用户登录信息成功!")
	}

	data = DB.LoginInfo{
		UserId:   "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
		PassWord: "111111",
		Phone:    "18823232334",
	}
	_, err = SetUserLoginInfobyPhone(data.Phone,data,true)
	if err != nil{
		Utils.ErrorOutputf("通过手机号设置数据库中[不存在]用户登录信息失败，手机号为:%s, 用户ID为:%s, 错误信息为:%s", data.Phone, data.UserId, err.Error())
	}else{
		Utils.ErrorOutputf("通过手机号设置数据库中[不存在]用户登录信息成功!")
	}



	data = DB.LoginInfo{
		UserId:   "a91559d8-0c99-11eb-9650-00163e2ed191",
		PassWord: "cccc",
		Phone:    "13002323255",
	}
	_, err = SetUserLoginInfobyUserId(data.UserId,data,true)
	if err != nil{
		Utils.ErrorOutputf("通过用户ID设置数据库中[已存在]用户登录信息失败，手机号为:%s, 用户ID为:%s, 错误信息为:%s", data.Phone, data.UserId, err.Error())
	}else{
		Utils.ErrorOutputf("通过用户ID设置数据库中[已存在]用户登录信息成功!")
	}

	data = DB.LoginInfo{
		UserId:   "aabb5678-bbbb-cccc-dddd-eeeeeeeeeeee",
		PassWord: "111111",
		Phone:    "18823232335",
	}
	code, err := SetUserLoginInfobyUserId(data.UserId,data,true)
	if err != nil{
		if code == config.ENUM_ERR_DB_INSERT_DUPLICATE{
			Utils.ErrorOutputf("通过用户ID设置数据库中[不存在]用户登录信息失败，手机号为:%s, 用户ID为:%s, 错误信息为:手机号重复", data.Phone, data.UserId)
		}else{
			Utils.ErrorOutputf("通过用户ID设置数据库中[不存在]用户登录信息失败，手机号为:%s, 用户ID为:%s, 错误信息为:%s", data.Phone, data.UserId, err.Error())
		}
	}else{
		Utils.ErrorOutputf("通过用户ID设置数据库中[不存在]用户登录信息成功!")
	}

}

func TestGetUserLoginInfo(t *testing.T) {
	initTest()
	data := DB.LoginInfo{
		UserId:   "a91559d8-0c99-11eb-9650-00163e2ed191",
		PassWord: "abcdefg",
		Phone:    "18823232334",
	}
	login, code, err := GetUserLoginByPhone(data.Phone)
	if code != config.ENUM_ERR_OK && err != nil{
		Utils.ErrorOutputf("通过电话号码获取用户登录信息失败:%s", err.Error())
	}else{
		Utils.ErrorOutputf("通过电话号码获取用户登录信息成功! %v",login)
	}
}

func TestDelUserLoginInfo(t *testing.T) {
	initTest()
	var Phone = "15002326237"
	//数据库存在的电话
	code, err := DelUserLoginInfobyPhone(Phone, true, false)
	if code != config.ENUM_ERR_OK{
		Utils.ErrorOutputf("通过手机号删除[存在的]登录信息失败,code:%d msg:%s",code, err.Error())
	}else{
		Utils.ErrorOutputf("通过手机号删除[存在的]登录信息成功!")
	}

	var noPhone = "15002326239"
	//数据库不存在的电话
	code, err = DelUserLoginInfobyPhone(noPhone, true, false)
	if code != config.ENUM_ERR_OK{
		Utils.ErrorOutputf("通过手机号删除[不存在的]登录信息失败,code:%d msg:%s",code, err.Error())
	}else{
		Utils.ErrorOutputf("通过手机号删除[不存在的]登录信息成功!")
	}
}
