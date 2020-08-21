package handler

import (
	"Lovers_Micro_Test/helper/DB"
	lovers_srv_user "Lovers_Micro_Test/server/user-service/proto"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

const(
	DB_RES_OK = 0
	DB_RES_IN_EMPTY = 1 //用户名或密码为空
	DB_RES_NO_UNIQUE_IN_DB = 2 //内部错误，用户在数据库中不唯一
	DB_RES_NO_USER = 3 //用户未注册
	DB_RES_PWD_ERROR = 4 //密码错误
)

type UserHandler struct {
	DB *gorm.DB
}


func (user* UserHandler) Login(ctx context.Context, in *lovers_srv_user.LoginReq, out *lovers_srv_user.LoginResp) error{

	if(len(in.UserName) <= 0 || len(in.PassWord) <= 0){
		user.loginFailResp(out,DB_RES_IN_EMPTY)
		return errors.New("用户名或密码为空")
	}
	var logins []DB.LoginInfo
	user.DB.Where("user_name = ?",in.UserName).Find(&logins)
	if len(logins) > 1{
		user.loginFailResp(out,DB_RES_NO_UNIQUE_IN_DB)
		err := errors.New("user: " + in.UserName + "is not unique in database")
		logrus.Error(err.Error())
		return err
	}
	if len(logins) <= 0{
		//尝试查电话号码
		user.DB.Where("Phone = ?",in.UserName).Find(&logins)
		if len(logins) > 1{
			user.loginFailResp(out,DB_RES_NO_UNIQUE_IN_DB)
			err := errors.New("user: " + in.UserName + "is not unique in database")
			logrus.Error(err.Error())
			return err
		}
		if len(logins) <= 0 {
			user.loginFailResp(out, DB_RES_NO_USER)
			return errors.New("该用户未注册")
		}
	}

	if logins[0].PassWord != in.PassWord{
		user.loginFailResp(out,DB_RES_PWD_ERROR)
		return errors.New("密码错误")
	}
	return user.loginSuccessResp(out,logins[0].UserId)
}

func (user* UserHandler)loginDBQueryRes_Handle(res interface{}) (){}

func (user* UserHandler)loginFailResp(out *lovers_srv_user.LoginResp, res int32){
	out.Token = ""
	out.LoginTime = string(time.Now().Unix())
	out.UserInfo = nil
	out.LoginRes = res
}

func (user* UserHandler)loginSuccessResp(out *lovers_srv_user.LoginResp,userId string)(error){
	baseInfo := []DB.UserBaseInfo{}
	user.DB.Where("user_id = ?",userId).Find(&baseInfo)
	if len(baseInfo) <= 0{
		user.loginFailResp(out, DB_RES_NO_USER)
		return errors.New("该用户未注册")
	}
	out.LoginRes = DB_RES_OK
	out.LoginTime = string(time.Now().Unix())
	out.Token = ""
	out.UserInfo = user.DBBaseInfoToRespBaseInfo(baseInfo[0])
	return nil
}

func (user* UserHandler)DBBaseInfoToRespBaseInfo(dbBaseInfo DB.UserBaseInfo)( *lovers_srv_user.BaseInfo){
	anotherInfo := &lovers_srv_user.LoverInfo{}
	respBaseInfo := &lovers_srv_user.BaseInfo{AnotherInfo:anotherInfo}
	respBaseInfo.UserId = dbBaseInfo.UserId
	respBaseInfo.Phone = int32(dbBaseInfo.Phone)
	respBaseInfo.Birth = dbBaseInfo.Birth
	respBaseInfo.HomeTown = dbBaseInfo.HomeTown
	respBaseInfo.RealName = dbBaseInfo.RealName
	respBaseInfo.Sculpture = dbBaseInfo.Sculpture
	respBaseInfo.Sex = int32(dbBaseInfo.Sex)
	respBaseInfo.AnotherInfo.LoverId = dbBaseInfo.LoverId
	respBaseInfo.AnotherInfo.LoveDuration = dbBaseInfo.LoveDuration
	respBaseInfo.AnotherInfo.LoverNickName = dbBaseInfo.LoverNickName
	respBaseInfo.AnotherInfo.LoverPhone = int32(dbBaseInfo.LoverPhone)
	return respBaseInfo
}
func (user* UserHandler)Logout(ctx context.Context, in *lovers_srv_user.LogoutReq, out *lovers_srv_user.LogoutResp) error{
	return nil
}
func (user* UserHandler)RegisterUser(ctx context.Context, in *lovers_srv_user.RegisterReq, out *lovers_srv_user.RegisterResp) error{

	return nil
}
func (user* UserHandler)BindLover(ctx context.Context, in *lovers_srv_user.BindLoverReq, out *lovers_srv_user.BindLoverResp) error{
	return nil
}
func (user* UserHandler)UnBindLover(ctx context.Context, in *lovers_srv_user.UnBindLoverReq, out *lovers_srv_user.UnBindLoverResp) error{
	return nil
}
func (user* UserHandler)GetLoverInfo(ctx context.Context, in *lovers_srv_user.GetLoverInfoReq, out *lovers_srv_user.GetLoverInfoResp) error{
	return nil
}

