package handler

import (
	"Lovers_srv/api/handler/JWTHandler"
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

func (user* UserHandler)RegisterUser(ctx context.Context, in *lovers_srv_user.RegisterReq, out *lovers_srv_user.RegisterResp) error{
	out.RegisteredInfo = &lovers_srv_user.LoginResp{}
	if in == nil{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_PARAM_nil
		return  errors.New("in param is nil")
	}
	if in.UserInfo == nil{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_PARAM_nil
		return  errors.New("registeredInfo param is nil")
	}
	if len(in.UserName) <= 0 || len(in.PassWord) <= 0{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_IN_EMPTY
		return  errors.New("UserName or PassWord is empty")
	}
	isphone := Utils.VerifyPhoneFormat(in.UserInfo.Phone)
	if !isphone{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_PHONE_ERR
		return errors.New("invalid phone number")
	}

	var dupliPhone []DB.UserBaseInfo
	user.DB.Where("Phone = ?",in.UserInfo.Phone).Find(&dupliPhone)
	if len(dupliPhone) > 0{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_EXIST
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
		Phone:  in.UserInfo.Phone,
	}

	err := user.DB.Create(&regLoginInf).Error

	if err != nil{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_REG_ERR
		err := errors.New("insert login info to db failed,err:" + err.Error())
		logrus.Error(err.Error())
		return err
	}
	err = user.DB.Create(&regBaseInfo).Error
	if err != nil{
		out.RegisteredInfo.LoginRes = config.MSG_DB_REG_REG_ERR
		err := errors.New("insert UserBaseInfo to db failed,err:" + err.Error())
		logrus.Error(err.Error())
		return err
	}

	out.RegisteredInfo.UserInfo = user.DBBaseInfoToRespBaseInfo(regBaseInfo)
	out.RegisteredInfo.LoginRes = config.MSG_DB_REG_OK
	out.RegisteredInfo.LoginTime = strconv.FormatInt(time.Now().Unix(),10)
	token,err := JWTHandler.GenerateToken(regBaseInfo.UserId,"") // 需要密码
	out.RegisteredInfo.Token = token

	return nil
}