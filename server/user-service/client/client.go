package client

import (
	"Lovers_srv/config"
	proto "Lovers_srv/server/user-service/proto"
	"context"
	"github.com/micro/go-micro/client"
)

const USER_SRV_NAME = "lovers.srv.user"
type UserClient struct{
	client proto.UserService
	serviceName string
}

func NewUserClient() *UserClient{
	serverName := config.GlobalConfig.Srv_name
	if serverName == ""{
		serverName = USER_SRV_NAME
	}
	c := proto.NewUserService(serverName, client.DefaultClient)
	return &UserClient{
		client: c,
		serviceName:serverName,
	}
}

func (user *UserClient) Client_Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginResp,error){
	resp, err := user.client.Login(ctx,req)
	if err != nil{
		return resp, err
	}

	return resp, nil
}

func (user *UserClient)Client_Register(ctx context.Context, req *proto.RegisterReq)(*proto.RegisterResp,error){
	resp ,err := user.client.RegisterUser(ctx, req)
	if err != nil{
		return resp,err
	}
	return resp,nil
}

func (user *UserClient)Client_QueryUserExistById(ctx context.Context, req *proto.QueryUserIsExistByIdReq)(*proto.QueryUserIsExistByIdResp,error){
	resp, err := user.client.QueryUserIsExistById(ctx,req)
	return resp, err
}

func (user *UserClient)Client_QueryLoverIdById(ctx context.Context, req *proto.QueryLoverIdByIdReq) (*proto.QueryLoverIdByIdResp, error){
	resp, err := user.client.QueryLoverIdById(ctx,req)
	return resp, err
}

func (user *UserClient)Client_BindLover(ctx context.Context, req *proto.BindLoverReq) (*proto.BindLoverResp, error){
	resp, err := user.client.BindLover(ctx, req)
	return resp, err
}