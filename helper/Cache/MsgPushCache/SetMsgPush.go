package MsgPushCache

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"github.com/jinzhu/gorm"
)

func SetRunStatistics(userId string, data DB.RunStatistics, saveDB bool) (int, error) {
	//先进行参数判断
	if len(userId) <=0 || !Utils.VerifyUUIDFormat(userId) || len(data.UserId) <= 0 || !Utils.VerifyUUIDFormat(data.UserId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetRunStatistics] userId format error, param UserId:%s data.UserId:%s",userId, data.UserId)
	}

	//更新cache
	cacheKey := statisticsKey(userId)
	Cache.DelHash(cacheKey,_fieldKeyUserId)
	err := Cache.SetHashByJson(cacheKey,_fieldKeyUserId, data)
	if err != nil{
		//缓存必须设置成功才能进行下一步操作，否则会导致缓存和数据库不同步
		return config.ENUM_ERR_SETHASH_FAILED, Utils.ErrorOutputf("[SetRunStatistics] set hash cache failed:%s",err.Error())
	}

	//将该记录持久化到数据库中
	if saveDB{
		//先通过userid查询数据库中是否存在
		findRunStati := new(DB.RunStatistics)
		err := Cache.DB.Where("user_id = ?",data.UserId).Find(&findRunStati).Error
		if err != nil{
			if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND {
				findRunStati = nil
			}else{
				//对数据库操作失败，必须删除缓存，否则会导致缓存和数据库不同步
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				return config.ENUM_ERR_DB_QUERY_FAILED, Utils.ErrorOutputf("[SetRunStatistics] query table run_statistics failed, user_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}

		//数据库中有该信息，更新数据库
		if findRunStati != nil{
			updateRunStati := DB.RunStatistics{}
			updateRunStati = data
			err := Cache.DB.Model(DB.RunStatistics{}).Where("user_id = ?", data.UserId).Updates(&updateRunStati).Error
			if err != nil{
				//操作数据库失败了都要删除缓存
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				return config.ENUM_ERR_DB_UPDATE_FAILED, Utils.ErrorOutputf("[SetRunStatistics] update table run_statistics failed, uer_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}else{
			//数据库中没有该信息，新添加
			newRunStati := DB.RunStatistics{}
			newRunStati = data
			newRunStati.Model = gorm.Model{}
			err := Cache.DB.Create(&newRunStati).Error
			if err != nil{
				//操作数据库失败了都要删除缓存
				Cache.DelHash(cacheKey,_fieldKeyUserId)
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_INSERT_FAILED, err, "[SetRunStatistics] save to db failed, user_id:%s errinfo:%s",data.UserId, err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}


func SetRunItemData(userId string, runId string, data DB.RunItemData, saveDB bool) (int, error) {
	//先进行参数判断
	if len(userId) <=0 || !Utils.VerifyUUIDFormat(userId) || len(data.UserId) <= 0 || !Utils.VerifyUUIDFormat(data.UserId) || len(runId) <= 0 || !Utils.VerifyUUIDFormat(runId){
		return config.ENUM_ERR_INVALID_PARAM, Utils.ErrorOutputf("[SetRunItemData] userId format error, param UserId:%s data.UserId:%s",userId, data.UserId)
	}

	//更新cache
	cacheKey := runItemKey(userId, runId)
	Cache.DelHash(cacheKey,_fieldKeyUserId)
	err := Cache.SetHashByJson(cacheKey,_fieldKeyUserId, data)
	if err != nil{
		if saveDB {
			Utils.ErrorOutputf("[SetRunItemData] failed:%s",err.Error())
		}else{
			return config.ENUM_ERR_SETHASH_FAILED, Utils.ErrorOutputf("[SetRunItemData] set hash cache failed:%s",err.Error())
		}
	}

	//将该记录持久化到数据库中
	if saveDB{
		//先通过userid查询数据库中是否存在
		findRunItem := new(DB.RunItemData)
		err := Cache.DB.Where("user_id = ? and run_id = ?",data.UserId, runId).Find(&findRunItem).Error
		if err != nil{
			if err.Error() == config.MSG_ERR_DB_RECORD_NOT_FOUND {
				findRunItem = nil
			}else{
				return config.ENUM_ERR_DB_QUERY_FAILED, Utils.ErrorOutputf("[SetRunItemData] query table run_item_data failed, user_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}

		//数据库中有该信息，更新数据库
		if findRunItem != nil{
			updateRunItem := DB.RunItemData{}
			updateRunItem = data
			err := Cache.DB.Model(DB.RunItemData{}).Where("user_id = ? and run_id = ?", data.UserId, runId).Updates(&updateRunItem).Error
			if err != nil{
				return config.ENUM_ERR_DB_UPDATE_FAILED, Utils.ErrorOutputf("[SetRunItemData] update table run_statistics failed, uer_id:%s errinfo:%s",data.UserId, err.Error())
			}
		}else{
			//数据库中没有该信息，新添加
			newRunItem := DB.RunItemData{}
			newRunItem = data
			newRunItem.Model = gorm.Model{}
			err := Cache.DB.Create(&newRunItem).Error
			if err != nil{
				code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_INSERT_FAILED, err, "[SetRunItemData] save to db failed, user_id:%s errinfo:%s",data.UserId, err.Error())
				return code, err
			}
		}
	}
	return config.ENUM_ERR_OK, nil
}