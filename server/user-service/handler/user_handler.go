package handler

import (
	lovers_srv_user "Lovers_Micro_Test/server/user-service/proto"
	"context"
	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}


func (user* UserHandler) Login(ctx context.Context, in *lovers_srv_user.LoginReq, out *lovers_srv_user.LoginResp) error{

	//if(in.UserName == "VISUS" && in.PassWord == "123"){
	//	out.Token = "aaa"
	//	out.UserInfo = nil
	//	out.LoginTime = "202008161430"
	//	return nil
	//}
	//out.Token = "null"
	return nil
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

