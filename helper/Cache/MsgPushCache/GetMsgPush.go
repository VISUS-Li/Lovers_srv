package MsgPushCache

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	"encoding/json"
	"fmt"
)

const(
	_preRunStatisticsKey = "run_statistics_%s"
	_preRunItemKey = "run_item_%s_%s"
)
const(
	//设置redis hash结构的field key
	_fieldKeyPhone = "phone"
	_fieldKeyUserId = "user_id"
)
func statisticsKey(key string) string {
	return fmt.Sprintf(_preRunStatisticsKey, key)
}
func runItemKey(key1 string, Key2 string) string {
	return fmt.Sprintf(_preRunItemKey, key1, Key2)
}

/******
GetRunStatistics:
	根据用户ID获取该用户跑步统计
******/
func GetRunStatistics(userId string)(*DB.RunStatistics, int, error){
	var cacheFind = true //在缓存中是否找到，未找到则从数据库中查找
	cacheKey := statisticsKey(userId)
	jsonBytes, err := Cache.GetHashByJson(cacheKey,_fieldKeyUserId)
	if err != nil{
		Utils.ErrorOutputf("[GetRunStatistics] get has cache failed:%s",err.Error())
		cacheFind = false
	}
	if jsonBytes == nil{
		cacheFind = false
	}

	//如果从redis中找到了，则解析
	if cacheFind {
		runStati := new(DB.RunStatistics)
		err = json.Unmarshal(jsonBytes, runStati)
		if err != nil {
			Utils.ErrorOutputf("[GetRunStatistics] UnMarshal cache failed:%s", err.Error())
			cacheFind = false
		}else{
			return runStati, config.ENUM_ERR_OK, nil
		}
	}

	//如果cache中没有找到，则从数据库中查找，并且添加到cache中
	if cacheFind == false{
		runStati := new(DB.RunStatistics)
		err := Cache.DB.Where("user_id = ?", userId).Find(runStati).Error
		if err != nil{
			code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_QUERY_FAILED, err, "[GetRunStatistics] query table run_statistics by userId:%s failed:%s",userId, err.Error())
			return nil, code, err
		}

		//数据库中找到了，添加到cache中
		code, err := SetRunStatistics(userId, *runStati, false)
		return runStati, code, err
	}

	return nil, config.ENUM_ERR_OK, nil
}

func GetRunItemData(userId string, runId string)(*DB.RunItemData, int, error){
	var cacheFind = true //在缓存中是否找到，未找到则从数据库中查找
	cacheKey := runItemKey(userId, runId)
	jsonBytes, err := Cache.GetHashByJson(cacheKey,_fieldKeyUserId)
	if err != nil{
		Utils.ErrorOutputf("[GetRunItemData] get has cache failed:%s",err.Error())
		cacheFind = false
	}
	if jsonBytes == nil{
		cacheFind = false
	}

	//如果从redis中找到了，则解析
	if cacheFind {
		runItem := new(DB.RunItemData)
		err = json.Unmarshal(jsonBytes, runItem)
		if err != nil {
			Utils.ErrorOutputf("[GetRunItemData] UnMarshal cache failed:%s", err.Error())
			cacheFind = false
		}else{
			return runItem, config.ENUM_ERR_OK, nil
		}
	}

	//如果cache中没有找到，则从数据库中查找，并且添加到cache中
	if cacheFind == false{
		runItem := new(DB.RunItemData)
		err := Cache.DB.Where("user_id = ? and run_id = ?", userId, runId).Find(runItem).Error
		if err != nil{
			code, err := Utils.ErrorOutputMysqlf(config.ENUM_ERR_DB_QUERY_FAILED, err, "[GetRunItemData] query table run_item_data by userId:%s failed:%s",userId, err.Error())
			return nil, code, err
		}

		//数据库中找到了，添加到cache中
		code, err := SetRunItemData(userId, runId, *runItem, false)
		return runItem, code, err
	}

	return nil, config.ENUM_ERR_OK, nil
}