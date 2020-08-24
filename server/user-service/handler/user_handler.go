package handler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/DB"
	"Lovers_srv/helper/Utils"
	lovers_srv_user "Lovers_srv/server/user-service/proto"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)


type UserHandler struct {
	DB *gorm.DB
}


func (user* UserHandler) Login(ctx context.Context, in *lovers_srv_user.LoginReq, out *lovers_srv_user.LoginResp) error{

	if(len(in.UserName) <= 0 || len(in.PassWord) <= 0){
		user.loginFailResp(out, config.DB_LOGIN_IN_EMPTY)
		return errors.New("用户名或密码为空")
	}
	var logins []DB.LoginInfo
	user.DB.Where("user_name = ?",in.UserName).Find(&logins)
	if len(logins) > 1{
		user.loginFailResp(out, config.DB_LOGIN_NO_UNIQUE_IN_DB)
		err := errors.New("user: " + in.UserName + "is not unique in database")
		logrus.Error(err.Error())
		return err
	}
	if len(logins) <= 0{
		//尝试查电话号码
		user.DB.Where("Phone = ?",in.UserName).Find(&logins)
		if len(logins) > 1{
			user.loginFailResp(out, config.DB_LOGIN_NO_UNIQUE_IN_DB)
			err := errors.New("user: " + in.UserName + "is not unique in database")
			logrus.Error(err.Error())
			return err
		}
		if len(logins) <= 0 {
			user.loginFailResp(out, config.DB_LOGIN_NO_USER)
			return errors.New("该用户未注册")
		}
	}

	if logins[0].PassWord != in.PassWord{
		user.loginFailResp(out, config.DB_LOGIN_PWD_ERROR)
		return errors.New("密码错误")
	}
	return user.loginSuccessResp(out,logins[0].UserId)
}

func (user* UserHandler)loginDBQueryRes_Handle(res interface{}) (){}

func (user* UserHandler)loginFailResp(out *lovers_srv_user.LoginResp, res string){
	out.Token = ""
	out.LoginTime = string(time.Now().Unix())
	out.UserInfo = nil
	out.LoginRes = res
}

//创建登录成功信息，在其中查询用户信息
func (user* UserHandler)loginSuccessResp(out *lovers_srv_user.LoginResp,userId string)(error){
	baseInfo := []DB.UserBaseInfo{}
	user.DB.Where("user_id = ?",userId).Find(&baseInfo)
	if len(baseInfo) <= 0{
		user.loginFailResp(out, config.DB_LOGIN_NO_USER)
		return errors.New("该用户未注册")
	}
	out.LoginRes = config.DB_LOGIN_OK
	out.LoginTime = string(time.Now().Unix())
	out.Token = ""
	out.UserInfo = user.DBBaseInfoToRespBaseInfo(baseInfo[0])
	return nil
}

func (user* UserHandler)DBBaseInfoToRespBaseInfo(dbBaseInfo DB.UserBaseInfo)( *lovers_srv_user.BaseInfo){
	anotherInfo := &lovers_srv_user.LoverInfo{}
	respBaseInfo := &lovers_srv_user.BaseInfo{AnotherInfo:anotherInfo}
	respBaseInfo.UserId = dbBaseInfo.UserId
	respBaseInfo.Phone = dbBaseInfo.Phone
	respBaseInfo.Birth = dbBaseInfo.Birth
	respBaseInfo.HomeTown = dbBaseInfo.HomeTown
	respBaseInfo.RealName = dbBaseInfo.RealName
	respBaseInfo.Sculpture = dbBaseInfo.Sculpture
	respBaseInfo.Sex = strconv.Itoa(dbBaseInfo.Sex)
	respBaseInfo.AnotherInfo.LoverId = dbBaseInfo.LoverId
	respBaseInfo.AnotherInfo.LoveDuration = strconv.FormatInt(dbBaseInfo.LoveDuration,10)
	respBaseInfo.AnotherInfo.LoverNickName = dbBaseInfo.LoverNickName
	respBaseInfo.AnotherInfo.LoverPhone = dbBaseInfo.LoverPhone
	return respBaseInfo
}
func (user* UserHandler)Logout(ctx context.Context, in *lovers_srv_user.LogoutReq, out *lovers_srv_user.LogoutResp) error{
	return nil
}
func (user* UserHandler)RegisterUser(ctx context.Context, in *lovers_srv_user.RegisterReq, out *lovers_srv_user.RegisterResp) error{
	out.RegisteredInfo = &lovers_srv_user.LoginResp{}
	if in == nil{
		out.RegisteredInfo.LoginRes = config.DB_REG_PARAM_nil
		return  errors.New("in param is nil")
	}
	if in.UserInfo == nil{
		out.RegisteredInfo.LoginRes = config.DB_REG_PARAM_nil
		return  errors.New("registeredInfo param is nil")
	}
	if len(in.UserName) <= 0 || len(in.PassWord) <= 0{
		out.RegisteredInfo.LoginRes = config.DB_REG_IN_EMPTY
		return  errors.New("UserName or PassWord is empty")
	}
	isphone := Utils.VerifyPhoneFormat(in.UserInfo.Phone)
	if !isphone{
		out.RegisteredInfo.LoginRes = config.DB_REG_PHONE_ERR
		return errors.New("invalid phone number")
	}

	var dupliPhone []DB.UserBaseInfo
	user.DB.Where("Phone = ?",in.UserInfo.Phone).Find(&dupliPhone)
	if len(dupliPhone) > 0{
		out.RegisteredInfo.LoginRes = config.DB_REG_EXIST
		return errors.New("phone number is exists")
	}
	//创建UUID
	userUUID := uuid.NewV1()
	sex, _ := strconv.Atoi(in.UserInfo.Sex)

	var regBaseInfo DB.UserBaseInfo

	if in.UserInfo.AnotherInfo == nil {
		regBaseInfo = DB.UserBaseInfo{
			Model:         gorm.Model{},
			UserId:        userUUID.String(),
			RealName:      in.UserInfo.RealName,
			Phone:         in.UserInfo.Phone,
			Sex:           sex,
			Birth:         in.UserInfo.Birth,
			Sculpture:     in.UserInfo.Sculpture,
			HomeTown:      in.UserInfo.HomeTown,
		}
	}else{
		regBaseInfo = DB.UserBaseInfo{
			Model:         gorm.Model{},
			UserId:        userUUID.String(),
			RealName:      in.UserInfo.RealName,
			Phone:         in.UserInfo.Phone,
			Sex:           sex,
			Birth:         in.UserInfo.Birth,
			Sculpture:     in.UserInfo.Sculpture,
			HomeTown:      in.UserInfo.HomeTown,
			LoverId:       in.UserInfo.AnotherInfo.LoverId,
			LoverPhone:    in.UserInfo.AnotherInfo.LoverPhone,
			LoverNickName: in.UserInfo.AnotherInfo.LoverNickName,
			LoveDuration:  0,
		}
	}
	regLoginInf := DB.LoginInfo{
		Model:    gorm.Model{},
		UserId:   userUUID.String(),
		UserName: in.UserName,
		PassWord: in.PassWord,
		Phone:  in.UserInfo.Phone  ,
	}

	err := user.DB.Create(&regLoginInf).Error

	if err != nil{
		out.RegisteredInfo.LoginRes = config.DB_REG_REG_ERR
		return errors.New("insert login info to db failed,err:" + err.Error())
	}
	err = user.DB.Create(&regBaseInfo).Error
	if err != nil{
		out.RegisteredInfo.LoginRes = config.DB_REG_REG_ERR
		return errors.New("insert UserBaseInfo to db failed,err:" + err.Error())
	}

	out.RegisteredInfo.UserInfo = user.DBBaseInfoToRespBaseInfo(regBaseInfo)
	out.RegisteredInfo.LoginRes = config.DB_REG_OK
	out.RegisteredInfo.LoginTime = string(time.Now().Unix())
	out.RegisteredInfo.Token = ""

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

