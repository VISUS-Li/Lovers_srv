package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"errors"
	"github.com/google/uuid"
)

func (user* UserHandler)QueryUserIsExistById(ctx context.Context, in *lovers_srv_user.QueryUserIsExistByIdReq, out *lovers_srv_user.QueryUserIsExistByIdResp) error{
	_, err := uuid.Parse(in.UserId)
	if err != nil{
		out.IsExist = false
		out.QueryCode = config.INVALID_PARAMS
		out.QueryRes = config.MSG_ERR_PARAM_WRONG
		return errors.New(config.MSG_ERR_PARAM_WRONG)
	}
	baseInfo := &DB.UserBaseInfo{}
	err = user.DB.Where("user_id = ?",in.UserId).Find(&baseInfo).Error
	if err != nil{
		out.IsExist = false
		out.QueryCode = config.CODE_ERR_SELECT_DB_FAIL
		out.QueryRes  = config.MSG_ERR_SELECT_DB_FAIL
		return errors.New(config.MSG_ERR_SELECT_DB_FAIL)
	}
	_, err = uuid.Parse(baseInfo.UserId)
	if err != nil {
		out.IsExist = false
		out.QueryCode = config.CODE_ERR_USER_NOT_EXIST
		out.QueryRes = config.MSG_USER_NOT_EXIST
		return errors.New(config.MSG_USER_NOT_EXIST)
	}
	out.IsExist = true
	out.QueryCode = config.CODE_ERR_SUCCESS
	out.QueryRes = config.MSG_REQUEST_SUCCESS
	return  nil
}
