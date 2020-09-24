package client

import (
	"Lovers_srv/config"
	proto "Lovers_srv/server/home-service/proto"
	"context"
	"github.com/micro/go-micro/client"
)

const HOME_SRV_NAME = "lovers.srv.home"
type HomeMicroClient struct{
	client proto.HomeService
	serviceName string
}

func NewHomeMicroClient() *HomeMicroClient{
	serverName := config.GlobalConfig.Srv_name
	if serverName == ""{
		serverName = HOME_SRV_NAME
	}
	c := proto.NewHomeService(serverName, client.DefaultClient)
	return &HomeMicroClient{
		client: c,
		serviceName:serverName,
	}
}

func (home *HomeMicroClient)Client_GetMainCard(ctx context.Context, req *proto.GetMainCardReq)(*proto.GetMainCardResp,error){
	resp, err := home.client.GetMainCard(ctx,req)
	return resp,err
}

func (home *HomeMicroClient)Client_PostCardInfo(ctx context.Context, req *proto.PostCardInfoReq)(* proto.PostCardInfoResp,error){
	resp, err := home.client.PostCardInfo(ctx,req)
	return resp, err
}