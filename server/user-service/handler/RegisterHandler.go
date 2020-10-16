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
	//out.RegisteredInfo = &lovers_srv_user.LoginResp{}
	if in == nil{
		return user.RegisterFailResp(out,config.MSG_DB_REG_PARAM_nil,config.CODE_ERR_PARAM_EMPTY)
	}
	if len(in.Phone) <= 0 || len(in.PassWord) <= 0{
		return user.RegisterFailResp(out,config.MSG_DB_REG_IN_EMPTY,config.CODE_ERR_PARAM_EMPTY)
	}
	isphone := Utils.VerifyPhoneFormat(in.Phone)
	if !isphone{
		return user.RegisterFailResp(out,config.MSG_DB_REG_PHONE_ERR,config.CODE_ERR_REG_PHONE_ERR)
	}

	var dupliPhone []DB.UserBaseInfo
	user.DB.Where("Phone = ?",in.Phone).Find(&dupliPhone)
	if len(dupliPhone) > 0{
		return user.RegisterFailResp(out,config.MSG_DB_REG_EXIST,config.CODE_ERR_REG_PHONE_EXIST)
	}
	//创建UUID
	userUUID := uuid.NewV1()


	var regBaseInfo DB.UserBaseInfo

	regBaseInfo = DB.UserBaseInfo{
		Model:         gorm.Model{},
		UserId:        userUUID.String(),
		Phone:         in.Phone,
		Sex:           int(in.Gender),
	}
	regLoginInf := DB.LoginInfo{
		Model:    gorm.Model{},
		UserId:   userUUID.String(),
		PassWord: in.PassWord,
		Phone:  in.Phone,
	}

	err := user.DB.Create(&regLoginInf).Error

	if err != nil{
		logrus.Errorf("插入Login数据库失败,err:%s",err.Error())
		return user.RegisterFailResp(out,config.MSG_DB_REG_REG_ERR,config.CODE_ERR_SERVER_INTERNAL)

	}
	err = user.DB.Create(&regBaseInfo).Error
	if err != nil{
		logrus.Errorf("插入Register数据库失败,err:%s",err.Error())
		return user.RegisterFailResp(out,config.MSG_DB_REG_REG_ERR,config.CODE_ERR_SERVER_INTERNAL)
	}

	return user.RegisterSuccessResp(out,regBaseInfo)
}


func (user* UserHandler)RegisterFailResp(out *lovers_srv_user.RegisterResp, msg string, code int) error{
	out.RegisteredInfo = new(lovers_srv_user.LoginResp)
	out.RegisteredInfo.LoginRes = msg
	out.RegisteredInfo.LoginCode = strconv.Itoa(code)
	return errors.New(msg+"_"+strconv.Itoa(code))
}

func (user* UserHandler)RegisterSuccessResp(out *lovers_srv_user.RegisterResp, regBaseInfo DB.UserBaseInfo) error{
	token,err := JWTHandler.GenerateToken(regBaseInfo.UserId,"") // 需要密码
	if err != nil{
		logrus.Debugf("生成token失败,UserId:%s,Password:%s",regBaseInfo.UserId,"")
		return user.RegisterFailResp(out, config.MSG_SERVER_INTERNAL, config.CODE_ERR_SERVER_INTERNAL)
	}
	out.RegisteredInfo = new(lovers_srv_user.LoginResp)
	out.RegisteredInfo.UserInfo = user.DBBaseInfoToRespBaseInfo(regBaseInfo)
	out.RegisteredInfo.LoginRes = config.MSG_DB_REG_OK
	out.RegisteredInfo.LoginTime = strconv.FormatInt(time.Now().Unix(),10)
	out.RegisteredInfo.Token = token.Token
	out.RegisteredInfo.TokenExpireTime =  strconv.FormatInt(token.ExpireTime,10)
	return nil
}