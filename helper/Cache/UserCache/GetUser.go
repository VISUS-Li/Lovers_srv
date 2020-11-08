package UserCache
//用户缓存包
//查找时，先从缓存中找，找不到再去数据库找
//插入时，先插入到缓存中，再根据需要持久化到数据库中

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"encoding/json"
	"fmt"
)

const(
	_preLoginKey = "login_%s" //手机号或userId从redis获取login信息的前缀: login_15002326233
	_preBaseInfoKey = "base_info_%s" //通過user_id从redis获取baseInfo的前缀:base_info_a91559d8-0c99-11eb-9650-00163e2ed191
)

const(
	//设置redis hash结构的field key
	_fieldKeyPhone = "phone"
	_fieldKeyUserId = "user_id"
)

func loginKey(key string) string{
	return fmt.Sprintf(_preLoginKey, key)
}

func baseInfoKey(key string) string{
	return fmt.Sprintf(_preBaseInfoKey,key)
}
/******
 通过电话号码获取用户登录信息,先从缓存中找，没找到再从数据库找
******/
func GetUserLoginByPhone(phone string)(*DB.LoginInfo, int, error){
	var cacheFind = true //在缓存中是否找到，未找到则从数据库中查找
	cacheKey := loginKey(phone)
	jsonBytes, err := Cache.GetHashByJson(cacheKey,_fieldKeyPhone)
	if err != nil{
		Utils.ErrorOutputf("get has cache failed:%s",err.Error())
		cacheFind = false
	}
	if jsonBytes == nil{
		cacheFind = false
	}

	//如果从redis中找到了，则解析
	if cacheFind {
		login := new(DB.LoginInfo)
		err = json.Unmarshal(jsonBytes, login)
		if err != nil {
			Utils.ErrorOutputf("UnMarshal cache failed:%s", err.Error())
			cacheFind = false
		}else{
			return login,config.ENUM_ERR_OK, nil
		}
	}

	//如果cache中没有找到，则从数据库中查找，并且添加到cache中
	if cacheFind == false{
		login := new(DB.LoginInfo)
		err := Cache.DB.Where("phone = ?", phone).Find(login).Error
		if err != nil{
			code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_QUERY_FAILED, err, "query table login_info by phone:%s failed:%s",phone, err.Error())
			return nil, code, err
		}

		//数据库中找到了，添加到cache中
		code, err := SetUserLoginInfobyPhone(phone, *login, false)
		return login, code, err
	}

	return nil, config.ENUM_ERR_OK, nil
}


func GetUserLoginInfoByUserId(userId string)(*DB.LoginInfo, int, error){
	var cacheFind = true //在缓存中是否找到，未找到则从数据库中查找
	cacheKey := loginKey(userId)
	jsonBytes, err := Cache.GetHashByJson(cacheKey,_fieldKeyUserId)
	if err != nil{
		Utils.ErrorOutputf("get has cache failed:%s",err.Error())
		cacheFind = false
	}
	if jsonBytes == nil{
		cacheFind = false
	}

	//如果从redis中找到了，则解析
	if cacheFind {
		loginInfo := new(DB.LoginInfo)
		err = json.Unmarshal(jsonBytes, loginInfo)
		if err != nil {
			Utils.ErrorOutputf("UnMarshal cache failed:%s", err.Error())
			cacheFind = false
		}else{
			return loginInfo, config.ENUM_ERR_OK, nil
		}
	}

	//如果cache中没有找到，则从数据库中查找
	if cacheFind == false{
		loginInfo := new(DB.LoginInfo)
		err := Cache.DB.Where("user_id = ?", userId).Find(loginInfo).Error
		if err != nil{
			code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_QUERY_FAILED, err, "query table login_info by userId:%s failed:%s",userId, err.Error())
			return nil, code, err
		}


		//数据库中找到了，添加到cache中
		code, err := SetUserLoginInfobyUserId(userId, *loginInfo, false)
		return loginInfo, code, err
	}
	return nil, config.ENUM_ERR_OK, nil
}


func GetUserBaseInfoByPhone(phone string)(*DB.UserBaseInfo, int, error){
	var cacheFind = true //在缓存中是否找到，未找到则从数据库中查找
	cacheKey := baseInfoKey(phone)
	jsonBytes, err := Cache.GetHashByJson(cacheKey,_fieldKeyPhone)
	if err != nil{
		Utils.ErrorOutputf("get has cache failed:%s",err.Error())
		cacheFind = false
	}
	if jsonBytes == nil{
		cacheFind = false
	}

	//如果从redis中找到了，则解析
	if cacheFind {
		baseInfo := new(DB.UserBaseInfo)
		err = json.Unmarshal(jsonBytes, baseInfo)
		if err != nil {
			Utils.ErrorOutputf("UnMarshal cache failed:%s", err.Error())
			cacheFind = false
		}else{
			return baseInfo,config.ENUM_ERR_OK, nil
		}
	}

	//如果cache中没有找到，则从数据库中查找，并且添加到cache中
	if cacheFind == false{
		baseInfo := new(DB.UserBaseInfo)
		err := Cache.DB.Where("phone = ?", phone).Find(baseInfo).Error
		if err != nil{
			code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_QUERY_FAILED, err, "query table user_base_info by phone:%s failed:%s",phone, err.Error())
			return nil, code, err
		}

		//数据库中找到了，添加到cache中
		code, err := SetUserBaseInfoByPhone(phone, *baseInfo, false)
		return baseInfo, code, err
	}

	return nil, config.ENUM_ERR_OK, nil
}



/******
 通过UserId获取用户基本信息,先从缓存中找，没找到再从数据库找
数据库中找到了，会更新到cache中
******/
func GetUserBaseInfoByUserId(userId string)(*DB.UserBaseInfo, int, error){
	var cacheFind = true //在缓存中是否找到，未找到则从数据库中查找
	cacheKey := baseInfoKey(userId)
	jsonBytes, err := Cache.GetHashByJson(cacheKey,_fieldKeyUserId)
	if err != nil{
		Utils.ErrorOutputf("get has cache failed:%s",err.Error())
		cacheFind = false
	}
	if jsonBytes == nil{
		cacheFind = false
	}

	//如果从redis中找到了，则解析
	if cacheFind {
		baseInfo := new(DB.UserBaseInfo)
		err = json.Unmarshal(jsonBytes, baseInfo)
		if err != nil {
			Utils.ErrorOutputf("UnMarshal cache failed:%s", err.Error())
			cacheFind = false
		}else{
			return baseInfo, config.ENUM_ERR_OK, nil
		}
	}

	//如果cache中没有找到，则从数据库中查找
	if cacheFind == false{
		baseInfo := new(DB.UserBaseInfo)
		err := Cache.DB.Where("user_id = ?", userId).Find(baseInfo).Error
		if err != nil{
			code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_QUERY_FAILED, err, "query table UserBaseInfo by userId:%s failed:%s",userId, err.Error())
			return nil, code, err
		}


		//数据库中找到了，添加到cache中
		code, err := SetUserBaseInfoByUserId(userId, *baseInfo, false)
		return baseInfo, code, err
	}
	return nil, config.ENUM_ERR_OK, nil
}