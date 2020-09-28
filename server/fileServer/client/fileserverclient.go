package client

import (
	"Lovers_srv/config"
	proto "Lovers_srv/server/fileserver/proto"
	"context"
	"github.com/micro/go-micro/client"
	)

const FILE_SRV_NAME = "lovers.srv.files"

type FileServerClient struct {
	client proto.FileServerService
	serviceName string
}

func NewFileServerClient() *FileServerClient {
	serverName := config.GlobalConfig.Srv_name
	if serverName == "" {
		serverName = FILE_SRV_NAME
	}

	c := proto.NewFileServerService(serverName, client.DefaultClient)
	return &FileServerClient{
		client:      c,
		serviceName: serverName,
	}

}

func (home *FileServerClient) Client_DownFile(ctx context.Context, req *proto.DownLoadFileReq) (*proto.DownLoadFileResp, error) {
	resp, err := home.client.DownLoadFile(ctx, req)
	return resp,err
}

func (home *FileServerClient) Client_UpFile(ctx context.Context, req *proto.UpLoadFileReq) (*proto.UpLoadFileResp, error) {
	resp, err := home.client.UpLoadFile(ctx,req)
	return resp,err
}

func (home *FileServerClient) Clien_DelFile(ctx context.Context, req *proto.DelFileReq) (*proto.DelFileResp, error) {
	resp, err := home.client.DelFile(ctx, req)
	return resp, err
}


