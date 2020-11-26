package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Cache/UserCache"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"time"
)


/******
	绑定Lover，给定双方的UserId进行绑定
*******/
func (user* UserHandler)BindLover(ctx context.Context, in *lovers_srv_user.BindLoverReq, out *lovers_srv_user.BindLoverResp) error{
	//处理参数有效性
	userId := in.UserId
	loverId := in.LoverId
	if !Utils.VerifyUUIDFormat(userId) || !Utils.VerifyUUIDFormat(loverId){
		out.BindLoverInfo = nil
		out.BindCode = config.INVALID_PARAMS
		out.BindRes = config.MSG_ERR_PARAM_WRONG
		return Utils.MicroErr(config.MSG_ERR_PARAM_WRONG, config.INVALID_PARAMS)
	}

	// step 1 检查用户是否存在
	queryExistReq := lovers_srv_user.QueryUserIsExistByIdReq{
		UserId:userId,
	}
	queryExistResp := new(lovers_srv_user.QueryUserIsExistByIdResp)

	user.QueryUserIsExistById(ctx, &queryExistReq, queryExistResp)
	if !queryExistResp.IsExist{
		out.BindLoverInfo = nil
		out.BindCode = config.CODE_ERR_USER_NOT_EXIST
		out.BindRes = config.MSG_USER_NOT_EXIST
		msg := fmt.Sprintf("user [%s] not exist", userId)
		return Utils.MicroErr(msg, config.CODE_ERR_USER_NOT_EXIST)
	}else if queryExistResp.QueryCode == config.CODE_ERR_SELECT_DB_FAIL{
		//查询失败
		out.BindLoverInfo = nil
		out.BindCode = config.CODE_ERR_SELECT_DB_FAIL
		out.BindRes = config.MSG_ERR_SELECT_DB_FAIL
		msg := fmt.Sprintf("user [%s] query db faild", userId)
		return Utils.MicroErr(msg, config.CODE_ERR_SELECT_DB_FAIL)
	}

	queryExistReq.UserId = loverId
	user.QueryUserIsExistById(ctx, &queryExistReq, queryExistResp)
	if !queryExistResp.IsExist{
		out.BindLoverInfo = nil
		out.BindCode = config.CODE_ERR_USER_NOT_EXIST
		out.BindRes = config.MSG_USER_NOT_EXIST
		msg := fmt.Sprintf("user [%s] not exist", loverId)
		return Utils.MicroErr(msg, config.CODE_ERR_USER_NOT_EXIST)
	}else if queryExistResp.QueryCode == config.CODE_ERR_SELECT_DB_FAIL{
		//查询失败
		out.BindLoverInfo = nil
		out.BindCode = config.CODE_ERR_SELECT_DB_FAIL
		out.BindRes = config.MSG_ERR_SELECT_DB_FAIL
		msg := fmt.Sprintf("user [%s] query db faild", loverId)
		return Utils.MicroErr(msg, config.CODE_ERR_SELECT_DB_FAIL)
	}

	//step 2 查询双方是否已经绑定过了
	userBaseInfo, _, _ := UserCache.GetUserBaseInfoByUserId(userId)
	//前面判断用户存在时已经获取了用户缓存，这里不用再判断

	//通过用户的情侣ID是否存在,判定用户是否绑定另一半
	if Utils.VerifyUUIDFormat(userBaseInfo.CoupleId){
		//已经绑定了另一半
		out.BindLoverInfo = nil
		out.BindCode = config.CODE_ERR_USER_ALREADY_BOUND_ANOTHER
		out.BindRes = config.MSG_USER_ALREADY_BOUND_ANOTHER
		msg := fmt.Sprintf("user [%s] already bound another user", userId)
		return Utils.MicroErr(msg, config.CODE_ERR_USER_ALREADY_BOUND_ANOTHER)
	}

	loverBaseInfo, _, _ := UserCache.GetUserBaseInfoByUserId(loverId)
	//前面判断用户存在时已经获取了用户缓存，这里不用再判断

	//通过用户的情侣ID是否存在,判定用户是否绑定另一半
	if Utils.VerifyUUIDFormat(loverBaseInfo.CoupleId){
		//已经绑定了另一半
		out.BindLoverInfo = nil
		out.BindCode = config.CODE_ERR_USER_ALREADY_BOUND_ANOTHER
		out.BindRes = config.MSG_USER_ALREADY_BOUND_ANOTHER
		msg := fmt.Sprintf("user [%s] already bound another user", loverId)
		return Utils.MicroErr(msg, config.CODE_ERR_USER_ALREADY_BOUND_ANOTHER)
	}

	//step 3 未绑定其他用户，开始绑定

	//setp 3.1 设置另一半信息
	coupleId := uuid.NewV1()
	boundUserInfo := SetBindInfo(userBaseInfo, loverBaseInfo,coupleId.String())
	SetBindInfo(loverBaseInfo, userBaseInfo,coupleId.String())

	//step 3.2 存储
	code ,err := UserCache.SetUserBaseInfoByUserId(userId, *userBaseInfo, true)
	if code != config.ENUM_ERR_OK{
		//存储失败
		out.BindLoverInfo = nil
		out.BindCode = int32(code)
		out.BindRes = err.Error()
		msg := fmt.Sprintf("user [%s] save cache and db failed", userId)
		return Utils.MicroErr(msg, code)
	}

	code ,err = UserCache.SetUserBaseInfoByUserId(loverId, *loverBaseInfo, true)
	if code != config.ENUM_ERR_OK{
		//存储失败
		out.BindLoverInfo = nil
		out.BindCode = int32(code)
		out.BindRes = err.Error()
		msg := fmt.Sprintf("user [%s] save cache and db failed", loverId)
		return Utils.MicroErr(msg, code)
	}

	out.BindLoverInfo = boundUserInfo
	out.BindCode = config.CODE_ERR_SUCCESS
	out.BindRes = config.MSG_REQUEST_SUCCESS
	return nil
}

func SetBindInfo(userInfo *DB.UserBaseInfo, loverInfo *DB.UserBaseInfo, coupleId string) *lovers_srv_user.LoverInfo{
	retInfo := new(lovers_srv_user.LoverInfo)
	userInfo.CoupleId = coupleId
	retInfo.CoupleId = coupleId

	userInfo.LoverId = loverInfo.UserId
	retInfo.LoverId = loverInfo.UserId

	userInfo.LoverPhone = loverInfo.Phone
	retInfo.LoverPhone = loverInfo.Phone

	userInfo.LoveDuration = 0
	retInfo.LoveDuration = 0

	userInfo.LoverNickName = loverInfo.Phone
	retInfo.LoverNickName =loverInfo.Phone
	return retInfo
}

func (user *UserHandler) UnBindLover(context.Context, *lovers_srv_user.UnBindLoverReq, *lovers_srv_user.UnBindLoverResp) error {
	return nil
}

func (user *UserHandler)GetBindWaitCode(ctx context.Context, in *lovers_srv_user.GetBindWaitCodeReq, out *lovers_srv_user.GetBindWaitCodeResp) error{
	//先获取用户基本信息
	userBaseInfo, code, err := UserCache.GetUserBaseInfoByUserId(in.UserId)
	if code != config.ENUM_ERR_OK{
		if code ==config.ENUM_ERR_DB_QUERY_NOT_FOUND{
			out.UserId = ""
			out.WaitCode = 0
			out.WaitCodeCode = config.CODE_ERR_DB_RECORD_NOT_FOUND
			out.WaitCodeRes = config.MSG_ERR_DB_RECORD_NOT_FOUND
			msg := fmt.Sprintf("user [%s] not found ", in.UserId)
			return Utils.MicroErr(msg, config.CODE_ERR_DB_RECORD_NOT_FOUND)
		}else{
			out.UserId = ""
			out.WaitCode = 0
			out.WaitCodeCode = config.CODE_ERR_SELECT_DB_FAIL
			out.WaitCodeRes = config.MSG_ERR_SELECT_DB_FAIL
			msg := fmt.Sprintf("user [%s] find failed err:(%s)", in.UserId,err.Error())
			return Utils.MicroErr(msg, config.CODE_ERR_SELECT_DB_FAIL)
		}
	}

	//step 1
	//通过用户的情侣ID是否存在,判定用户是否绑定另一半
	if Utils.VerifyUUIDFormat(userBaseInfo.CoupleId){
		//已经绑定了另一半
		out.UserId = in.UserId
		out.WaitCode = 0
		out.WaitCodeCode = config.CODE_ERR_USER_ALREADY_BOUND_ANOTHER
		out.WaitCodeRes = config.MSG_USER_ALREADY_BOUND_ANOTHER
		msg := fmt.Sprintf("user [%s] already bound another user", in.UserId)
		return Utils.MicroErr(msg, config.CODE_ERR_USER_ALREADY_BOUND_ANOTHER)
	}

	//step 2
	//生成WaitCode，判断该waitCode是否已经存在
	var waitCode int
	for {
		rand.Seed(time.Now().Unix())
		waitCode = rand.Intn(100000)
		exist, _ :=UserCache.IsExistWaitCode(strconv.Itoa(waitCode))
		if exist {
			continue
		}else{
			break
		}
	}

	//加入缓存
	_, err = UserCache.SetWaitUser(strconv.Itoa(waitCode), in.UserId)
	if code != config.ENUM_ERR_OK{
		out.UserId = in.UserId
		out.WaitCode = 0
		out.WaitCodeCode = config.CODE_ERR_SERVER_INTERNAL
		out.WaitCodeRes = config.MSG_SERVER_INTERNAL
		msg := fmt.Sprintf("[GetBindWaitCode] user [%s] set cache failed:%s", in.UserId,err.Error())
		return Utils.MicroErr(msg, config.CODE_ERR_SERVER_INTERNAL)
	}

	out.UserId = in.UserId
	out.WaitCode = int32(waitCode)
	out.WaitCodeCode = config.CODE_ERR_SUCCESS
	out.WaitCodeRes = config.MSG_REQUEST_SUCCESS
	return nil
}

func (user *UserHandler)GetWaitingUser(ctx context.Context, in *lovers_srv_user.GetWaitingUserReq, out *lovers_srv_user.GetWaitingUserResp) error{
	return nil
}