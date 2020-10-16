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
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)


type UserHandler struct {
	DB *gorm.DB
}


func (user* UserHandler) Login(ctx context.Context, in *lovers_srv_user.LoginReq, out *lovers_srv_user.LoginResp) error {

	if in.Type == config.ENUM_LOGIN_VERCODE {
		err := user.PhoneAndPwdLogin(in.Phone, in.PassWord, out)
		if err != nil {
			return err
		}
	}else{
		//默认采用用户名密码登录
		err := user.NameAndPwdLogin(in.UserName, in.PassWord, out)
		if err != nil {
			if Utils.VerifyPhoneFormat(in.UserName){
				//如果用户名为电话号码，通过电话号码登录
				err = user.PhoneAndPwdLogin(in.UserName, in.PassWord, out)
				return err
			}else if Utils.VerifyPhoneFormat(in.Phone){
					err = user.PhoneAndPwdLogin(in.Phone, in.PassWord, out)
					if err != nil {
						return err
					}
			}else{
				return user.loginFailResp(out, config.MSG_DB_REG_PHONE_ERR,config.CODE_ERR_REG_PHONE_ERR)
			}
		}
	}

return nil

}

/******
Type:1
通过用户名密码登录
******/
func (user* UserHandler)NameAndPwdLogin(userName string, password string, out *lovers_srv_user.LoginResp)(error){
	if(len(userName) <= 0 || len(password) <= 0){
		return user.loginFailResp(out, config.MSG_DB_LOGIN_IN_EMPTY,config.CODE_ERR_PARAM_EMPTY)
	}

	var logins []DB.LoginInfo
	err := user.DB.Where("user_name = ?",userName).Find(&logins).Error
	if err != nil{
		logrus.Error("query table user_name failed: " + err.Error())
		return user.loginFailResp(out,config.MSG_DB_LOGIN_QUERY_ERR,config.CODE_ERR_LOGIN_QUERY)
	}
	if len(logins) > 1{
		//该用户名不唯一，逻辑有问题
	}

	if len(logins) <= 0{
		return user.loginFailResp(out,config.MSG_DB_LOGIN_NO_USER,config.CODE_ERR_LOGIN_NO_USER)
	}

	if logins[0].PassWord != password{
		return user.loginFailResp(out,config.MSG_DB_LOGIN_PWD_ERROR,config.CODE_ERR_LOGIN_PWD_ERROR)
	}

	return user.loginSuccessResp(out,logins[0].UserId,userName,password)
}

/******
Type:1
通过电话号码密码登录
******/
func (user* UserHandler)PhoneAndPwdLogin(phone string, password string, out *lovers_srv_user.LoginResp)(error){
	if(len(phone) <= 0 || len(password) <= 0){
		return user.loginFailResp(out, config.MSG_DB_LOGIN_IN_EMPTY,config.CODE_ERR_PARAM_EMPTY)
	}

	var logins []DB.LoginInfo
	err := user.DB.Where("Phone = ?", phone).Find(&logins).Error
	if err != nil{
		logrus.Error("query table Phone failed: " + err.Error())
		return user.loginFailResp(out,config.MSG_DB_LOGIN_QUERY_ERR,config.CODE_ERR_LOGIN_QUERY)
	}
	if len(logins) > 1{
		//该用户名不唯一，逻辑有问题
	}

	if len(logins) <= 0{
		return user.loginFailResp(out,config.MSG_DB_LOGIN_NO_USER,config.CODE_ERR_LOGIN_NO_USER)
	}

	if logins[0].PassWord != password{
		return user.loginFailResp(out,config.MSG_DB_LOGIN_PWD_ERROR,config.CODE_ERR_LOGIN_PWD_ERROR)
	}

	return user.loginSuccessResp(out,logins[0].UserId,phone,password)
}


func (user* UserHandler)loginFailResp(out *lovers_srv_user.LoginResp, res string, code int)(error){
	out.Token = ""
	out.LoginTime = strconv.FormatInt(time.Now().Unix(),10)
	out.UserInfo = nil
	out.LoginRes = res
	out.LoginCode = strconv.Itoa(code)
	return errors.New(res)
}

//创建登录成功信息，在其中查询用户信息
func (user* UserHandler)loginSuccessResp(out *lovers_srv_user.LoginResp,userId string, username string, password string)(error){
	//通过用户ID，查询用户信息
	baseInfo := []DB.UserBaseInfo{}
	user.DB.Where("user_id = ?",userId).Find(&baseInfo)
	if len(baseInfo) <= 0{
		return user.loginFailResp(out, config.MSG_DB_LOGIN_NO_USER,config.CODE_ERR_LOGIN_NO_USER)
	}

	token,err := JWTHandler.GenerateToken(username, password)
	if err != nil{
		return user.loginFailResp(out, config.MSG_DB_LOGIN_TOKEN_ERROR, config.CODE_ERR_LOGIN_TOKEN_ERROR)
	}
	out.Token = token.Token
	out.TokenExpireTime = strconv.FormatInt(token.ExpireTime,10)
	out.LoginRes = config.MSG_DB_LOGIN_OK
	out.LoginCode = strconv.Itoa(config.CODE_ERR_SUCCESS)
	out.LoginTime = strconv.FormatInt(time.Now().Unix(),10)
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


func (user* UserHandler)BindLover(ctx context.Context, in *lovers_srv_user.BindLoverReq, out *lovers_srv_user.BindLoverResp) error{
	return nil
}
func (user* UserHandler)UnBindLover(ctx context.Context, in *lovers_srv_user.UnBindLoverReq, out *lovers_srv_user.UnBindLoverResp) error{
	return nil
}
func (user* UserHandler)GetLoverInfo(ctx context.Context, in *lovers_srv_user.GetLoverInfoReq, out *lovers_srv_user.GetLoverInfoResp) error{
	return nil
}

