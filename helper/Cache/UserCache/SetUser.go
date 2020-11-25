package UserCache

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"github.com/jinzhu/gorm"
)

func SetUserLoginInfobyPhone(phone string, data DB.LoginInfo, saveDB bool) (int, error) {
	//先进行参数判断
	if len(phone) <=0 || !Utils.VerifyPhoneFormat(phone){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetUserLoginInfobyPhone] phone format error, param phone:%s",phone)
	}
	if len(data.UserId) <= 0 || !Utils.VerifyUUIDFormat(data.UserId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetUserLoginInfobyPhone] data userId format error, param data.UserId:%s",data.UserId)
	}

	//更新cache
	cacheKey := loginKey(phone)
	Cache.DelHash(cacheKey,_fieldKeyPhone)
	err := Cache.SetHashByJson(cacheKey,_fieldKeyPhone,data)
	if err != nil{
		//缓存必须设置成功才能进行下一步操作，否则会导致缓存和数据库不同步
		return config.ENUM_ERR_SETHASH_FAILED, Utils.ErrorOutputf("[SetUserLoginInfobyPhone] set hash cache failed:%s",err.Error())
	}

	//将该记录持久化到数据库中
	if saveDB{
		//先通过userid查询数据库中是否存在
		findLogin := new(DB.LoginInfo)
		err := Cache.DB.Where("user_id = ?",data.UserId).Find(&findLogin).Error
		if err != nil{
			if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND {
				findLogin = nil
			}else{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyPhone)
				return config.ENUM_ERR_DB_QUERY_FAILED, Utils.ErrorOutputf("[SetUserLoginInfobyPhone] query table login_info failed, user_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}

		//数据库中有该信息，更新数据库
		if findLogin != nil{
			updateLogin := DB.LoginInfo{}
			updateLogin = data
			err := Cache.DB.Model(DB.LoginInfo{}).Where("user_id = ?", data.UserId).Updates(&updateLogin).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyPhone)
				return config.ENUM_ERR_DB_UPDATE_FAILED, Utils.ErrorOutputf("[SetUserLoginInfobyPhone] update table login_info failed, uer_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}else{
			//数据库中没有该信息，新添加
			newLoginInf := DB.LoginInfo{}
			newLoginInf = data
			newLoginInf.Model = gorm.Model{}
			err := Cache.DB.Create(&newLoginInf).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyPhone)
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_INSERT_FAILED, err, "[SetUserLoginInfobyPhone] save to db failed, user_id:%s errinfo:%s",data.UserId, err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}

func SetUserLoginInfobyUserId(userId string, data DB.LoginInfo, saveDB bool) (int, error) {
	//先进行参数判断
	if len(userId) <=0 || !Utils.VerifyUUIDFormat(userId) || len(data.UserId) <= 0 || !Utils.VerifyUUIDFormat(data.UserId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetUserLoginInfobyUserId] userId format error, param UserId:%s data.UserId:%s",userId, data.UserId)
	}

	//更新cache
	cacheKey := loginKey(userId)
	Cache.DelHash(cacheKey,_fieldKeyUserId)
	err := Cache.SetHashByJson(cacheKey,_fieldKeyUserId, data)
	if err != nil{
		//缓存必须设置成功才能进行下一步操作，否则会导致缓存和数据库不同步
		return config.ENUM_ERR_SETHASH_FAILED, Utils.ErrorOutputf("[SetUserLoginInfobyUserId] set hash cache failed:%s",err.Error())
	}

	//将该记录持久化到数据库中
	if saveDB{
		//先通过userid查询数据库中是否存在
		findLogin := new(DB.LoginInfo)
		err := Cache.DB.Where("user_id = ?",data.UserId).Find(&findLogin).Error
		if err != nil{
			if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND {
				findLogin = nil
			}else{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				return config.ENUM_ERR_DB_QUERY_FAILED, Utils.ErrorOutputf("[SetUserLoginInfobyUserId] query table login_info failed, user_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}

		//数据库中有该信息，更新数据库
		if findLogin != nil{
			updateLogin := DB.LoginInfo{}
			updateLogin = data
			err := Cache.DB.Model(DB.LoginInfo{}).Where("user_id = ?", data.UserId).Updates(&updateLogin).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				return config.ENUM_ERR_DB_UPDATE_FAILED, Utils.ErrorOutputf("[SetUserLoginInfobyUserId] update table login_info failed, uer_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}else{
			//数据库中没有该信息，新添加
			newLoginInf := DB.LoginInfo{}
			newLoginInf = data
			newLoginInf.Model = gorm.Model{}
			err := Cache.DB.Create(&newLoginInf).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_INSERT_FAILED, err, "[SetUserLoginInfobyUserId] save to db failed, user_id:%s errinfo:%s",data.UserId, err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}


func SetUserBaseInfoByPhone(phone string, data DB.UserBaseInfo, saveDB bool) (int, error) {
	//先进行参数判断
	if len(phone) <=0 || !Utils.VerifyPhoneFormat(phone){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetUserBaseInfoByPhone] phone format error, param phone:%s",phone)
	}
	if len(data.UserId) <= 0 || !Utils.VerifyUUIDFormat(data.UserId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetUserBaseInfoByPhone] data userId format error, param data.UserId:%s",data.UserId)
	}

	//更新cache
	cacheKey := baseInfoKey(phone)
	Cache.DelHash(cacheKey,_fieldKeyPhone)
	err := Cache.SetHashByJson(cacheKey,_fieldKeyPhone, data)
	if err != nil{
		//缓存必须设置成功才能进行下一步操作，否则会导致缓存和数据库不同步
		return config.ENUM_ERR_SETHASH_FAILED, Utils.ErrorOutputf("[SetUserBaseInfoByPhone] set hash cache failed:%s",err.Error())
	}

	//将该记录持久化到数据库中
	if saveDB{
		//先通过userid查询数据库中是否存在
		findBase := new(DB.UserBaseInfo)
		err := Cache.DB.Where("user_id = ?",data.UserId).Find(&findBase).Error
		if err != nil{
			if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND {
				findBase = nil
			}else{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyPhone)
				return config.ENUM_ERR_DB_QUERY_FAILED, Utils.ErrorOutputf("[SetUserBaseInfoByPhone] query table user_base_info failed, user_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}

		//数据库中有该信息，更新数据库
		if findBase != nil{
			updateBase := DB.UserBaseInfo{}
			updateBase = data
			err := Cache.DB.Model(DB.UserBaseInfo{}).Where("user_id = ?", data.UserId).Updates(&updateBase).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyPhone)
				return config.ENUM_ERR_DB_UPDATE_FAILED, Utils.ErrorOutputf("[SetUserBaseInfoByPhone] update table login_info failed, uer_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}else{
			//数据库中没有该信息，新添加
			newBaseInfo := DB.UserBaseInfo{}
			newBaseInfo = data
			newBaseInfo.Model = gorm.Model{}

			err := Cache.DB.Create(&newBaseInfo).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyPhone)
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_INSERT_FAILED, err, "[SetUserBaseInfoByPhone] save to db failed, user_id:%s errinfo:%s",data.UserId, err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}

func SetUserBaseInfoByUserId(userId string, data DB.UserBaseInfo, saveDB bool) (int, error) {
	//先进行参数判断
	if len(userId) <=0 || !Utils.VerifyUUIDFormat(userId) || len(data.UserId) <= 0 || !Utils.VerifyUUIDFormat(data.UserId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetUserLoginInfobyUserId] userId format error, param UserId:%s data.UserId:%s",userId, data.UserId)
	}

	//更新cache
	cacheKey := baseInfoKey(userId)
	Cache.DelHash(cacheKey,_fieldKeyUserId)
	err := Cache.SetHashByJson(cacheKey,_fieldKeyUserId, data)
	if err != nil{
		//缓存必须设置成功才能进行下一步操作，否则会导致缓存和数据库不同步
		return config.ENUM_ERR_SETHASH_FAILED, Utils.ErrorOutputf("[SetUserBaseInfoByUserId] set hash cache failed:%s",err.Error())
	}

	//将该记录持久化到数据库中
	if saveDB{
		//先通过userid查询数据库中是否存在
		findBase := new(DB.UserBaseInfo)
		err := Cache.DB.Where("user_id = ?",data.UserId).Find(&findBase).Error
		if err != nil{
			if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND {
				findBase = nil
			}else{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				return config.ENUM_ERR_DB_QUERY_FAILED, Utils.ErrorOutputf("[SetUserBaseInfoByUserId] query table user_base_info failed, user_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}

		//数据库中有该信息，更新数据库
		if findBase != nil{
			updateBase := DB.UserBaseInfo{}
			updateBase = data
			err := Cache.DB.Model(DB.UserBaseInfo{}).Where("user_id = ?", data.UserId).Updates(&updateBase).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				return config.ENUM_ERR_DB_UPDATE_FAILED, Utils.ErrorOutputf("[SetUserBaseInfoByUserId] update table user_base_info failed, uer_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}else{
			//数据库中没有该信息，新添加
			newBaseInfo := DB.UserBaseInfo{}
			newBaseInfo = data
			newBaseInfo.Model = gorm.Model{}
			err := Cache.DB.Create(&newBaseInfo).Error
			if err != nil{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_INSERT_FAILED, err, "[SetUserBaseInfoByUserId] save to db failed, user_id:%s errinfo:%s",data.UserId, err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}

/******
	通过手机号删除登录信息
	phone:要删除的手机号
	delDB:是否删除数据库中记录
	phyDel:删除数据库时，采用软删除还是物理删除
******/
func DelUserLoginInfobyPhone(phone string, delDB bool, phyDel bool) (int, error) {
	if len(phone) <=0 || !Utils.VerifyPhoneFormat(phone){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[DelUserLoginInfobyPhone] phone format error, param phone:%s",phone)
	}
	cacheKey := loginKey(phone)
	err := Cache.DelHash(cacheKey,_fieldKeyPhone)
	if err != nil{
		return config.ENUM_ERR_DELHASH_FAILED, Utils.ErrorOutputf("[DelUserLoginInfobyPhone] del hash cache failed:%s",err.Error())
	}

	//通过phone删除数据库中的登录数据
	if delDB{
		if phyDel{
			err = Cache.DB.Unscoped().Where("phone = ?", phone).Delete(DB.LoginInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserLoginInfobyPhone] del from login_info with phone:%s failed:%s",phone,err.Error())
				return code, err
			}
		}else{
			err = Cache.DB.Where("phone = ?", phone).Delete(DB.LoginInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserLoginInfobyPhone] del from login_info with phone:%s failed:%s",phone,err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}

/******
	通过用户ID删除登录信息
	userId:要删除的用户ID
	delDB:是否删除数据库中记录
	phyDel:删除数据库时，采用软删除还是物理删除
******/
func DelUserLoginInfobyUserId(userId string, delDB bool, phyDel bool) (int, error) {
	if len(userId) <=0 || !Utils.VerifyUUIDFormat(userId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[DelUserLoginInfobyUserId] userId format error, param userId:%s",userId)
	}
	cacheKey := loginKey(userId)
	err := Cache.DelHash(cacheKey,_fieldKeyUserId)
	if err != nil{
		return config.ENUM_ERR_DELHASH_FAILED,Utils.ErrorOutputf("[DelUserLoginInfobyUserId] del hash cache failed:%s",err.Error())
	}

	//通过UserId删除数据库中的登录数据
	if delDB{
		if phyDel{
			err = Cache.DB.Unscoped().Where("user_id = ?", userId).Delete(DB.LoginInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserLoginInfobyUserId] del from login_info with userId:%s failed:%s",userId,err.Error())
				return code, err
			}
		}else{
			err = Cache.DB.Where("user_id = ?", userId).Delete(DB.LoginInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserLoginInfobyUserId] del from login_info with userId:%s failed:%s",userId,err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}

/******
	通过手机号删除用户基本信息
	phone:要删除的手机号
	delDB:是否删除数据库中记录
	phyDel:删除数据库时，采用软删除还是物理删除
******/
func DelUserBaseInfobyPhone(phone string, delDB bool, phyDel bool) (int, error) {
	if len(phone) <=0 || !Utils.VerifyPhoneFormat(phone){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[DelUserBaseInfobyPhone] phone format error, param phone:%s",phone)
	}
	cacheKey := baseInfoKey(phone)
	err := Cache.DelHash(cacheKey,_fieldKeyPhone)
	if err != nil{
		return config.ENUM_ERR_DELHASH_FAILED, Utils.ErrorOutputf("[DelUserBaseInfobyPhone] del hash cache failed:%s",err.Error())
	}

	//通过phone删除数据库中的用户基础信息
	if delDB{
		if phyDel{
			err = Cache.DB.Unscoped().Where("phone = ?", phone).Delete(DB.UserBaseInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserBaseInfobyPhone] del from user_base_info with phone:%s failed:%s",phone,err.Error())
				return code, err
			}
		}else{
			err = Cache.DB.Where("phone = ?", phone).Delete(DB.UserBaseInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserBaseInfobyPhone] del from user_base_info with phone:%s failed:%s",phone,err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}


/******
	通过用户ID删除用户基本信息
	userId:要删除的用户ID
	delDB:是否删除数据库中记录
	phyDel:删除数据库时，采用软删除还是物理删除
******/
func DelUserBaseInfobyUserId(userId string, delDB bool, phyDel bool) (int, error) {
	if len(userId) <=0 || !Utils.VerifyUUIDFormat(userId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[DelUserBaseInfobyUserId] userId format error, param userId:%s",userId)
	}
	cacheKey := baseInfoKey(userId)
	err := Cache.DelHash(cacheKey,_fieldKeyUserId)
	if err != nil{
		return config.ENUM_ERR_DELHASH_FAILED, Utils.ErrorOutputf("[DelUserBaseInfobyUserId] del hash cache failed:%s",err.Error())
	}

	//通过UserId删除数据库中的基本信息
	if delDB{
		if phyDel{
			err = Cache.DB.Unscoped().Where("user_id = ?", userId).Delete(DB.UserBaseInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserBaseInfobyUserId] del from user_base_info with userId:%s failed:%s",userId,err.Error())
				return code, err
			}
		}else{
			err = Cache.DB.Where("user_id = ?", userId).Delete(DB.UserBaseInfo{}).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_DELETE_FAILED, err, "[DelUserBaseInfobyUserId] del from user_base_info with userId:%s failed:%s",userId,err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}