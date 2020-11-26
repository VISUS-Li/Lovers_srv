package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache/UserCache"
	"Lovers_srv/helper/Utils"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"github.com/google/uuid"
)

func (user* UserHandler)QueryUserIsExistById(ctx context.Context, in *lovers_srv_user.QueryUserIsExistByIdReq, out *lovers_srv_user.QueryUserIsExistByIdResp) error{
	//验证参数有效性
	_, err := uuid.Parse(in.UserId)
	if err != nil{
		out.IsExist = false
		out.QueryCode = config.INVALID_PARAMS
		out.QueryRes = config.MSG_ERR_PARAM_WRONG
		return Utils.MicroErr(config.MSG_ERR_PARAM_WRONG, config.INVALID_PARAMS)
	}

	//step 1 根据userId 查询
	baseInfo, code, err :=UserCache.GetUserBaseInfoByUserId(in.UserId)
	if code != config.ENUM_ERR_OK{
		if code == config.ENUM_ERR_DB_QUERY_NOT_FOUND{
			//未查询到
			out.IsExist = false
			out.QueryCode = config.CODE_ERR_DB_RECORD_NOT_FOUND
			out.QueryRes  = config.MSG_ERR_DB_RECORD_NOT_FOUND
			return Utils.MicroErr(config.MSG_USER_NOT_EXIST, config.CODE_ERR_USER_NOT_EXIST)
		}else{
			//查询失败
			out.QueryCode = config.CODE_ERR_SELECT_DB_FAIL
			out.QueryRes  = config.MSG_ERR_SELECT_DB_FAIL
			return Utils.MicroErr(config.MSG_ERR_SELECT_DB_FAIL, config.CODE_ERR_SELECT_DB_FAIL)
		}
	}
	// step 2 验证查询到的结果
	if baseInfo == nil || !Utils.VerifyUUIDFormat(baseInfo.UserId){
		out.IsExist = false
		out.QueryCode = config.CODE_ERR_USER_NOT_EXIST
		out.QueryRes = config.MSG_USER_NOT_EXIST
		return Utils.MicroErr(config.MSG_USER_NOT_EXIST, config.CODE_ERR_USER_NOT_EXIST)
	}

	out.IsExist = true
	out.QueryCode = config.CODE_ERR_SUCCESS
	out.QueryRes = config.MSG_REQUEST_SUCCESS
	return  nil
}


func (uer *UserHandler)QueryLoverIdById(ctx context.Context, in *lovers_srv_user.QueryLoverIdByIdReq, out *lovers_srv_user.QueryLoverIdByIdResp) error{
	_, err := uuid.Parse(in.UserId)
	if err != nil{
		out.LoverId = ""
		out.QueryCode = config.INVALID_PARAMS
		out.QueryRes = config.MSG_ERR_PARAM_WRONG
		return Utils.MicroErr(config.MSG_ERR_PARAM_WRONG, config.INVALID_PARAMS)
	}
	baseInfo, code, err :=UserCache.GetUserBaseInfoByUserId(in.UserId)
	if code != config.ENUM_ERR_OK{
		out.LoverId = ""
		out.QueryCode = int32(code)
		out.QueryRes = err.Error()
		return Utils.MicroErr(err.Error(), code)
	}
	out.LoverId = baseInfo.LoverId
	out.QueryCode = config.ENUM_ERR_OK
	out.QueryRes = config.MSG_REQUEST_SUCCESS
	return nil
}