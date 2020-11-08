package client

import (
	"Lovers_srv/config"
	proto "Lovers_srv/server/lovers-tools/note-list/proto"
	"context"
	"github.com/micro/go-micro/client"
	//文件IO

)

const NOTELIST_SRV_NAME = "lovers.srv.notelist"

type NoteListClient struct {
	client      proto.NoteListService
	serviceName string
}

func NewNoteListClient() *NoteListClient{
	//新建一个服务
	serverName := config.GlobalConfig.Srv_name
	if serverName == "" {
		serverName = NOTELIST_SRV_NAME
	}

	c := proto.NewNoteListService(serverName, client.DefaultClient)
	return &NoteListClient{
		client:      c,
		serviceName: serverName,
	}
}

//提供给gin路由的API接口
func (notelist *NoteListClient)NoteList_Up(ctx context.Context, req *proto.NoteListUpReq) (*proto.NoteListUpResp, error) {
	//调用微服务接口
	resp, err := notelist.client.NoteListUp(ctx, req)
	if err != nil{
		return resp, err
	}

	return resp, nil
}

func (notelist *NoteListClient)NoteList_Down(ctx context.Context, req *proto.NoteListDownReq) (*proto.NoteListDownResp, error) {
	//调用微服务接口
	resp, err := notelist.client.NoteListDown(ctx, req)
	if err != nil {
		return resp,err
	}
	return resp, nil

}

func (notelis *NoteListClient)NoteList_Del(ctx context.Context, req *proto.NoteListDelReq) (*proto.NoteListDelResp, error) {
	//调用微服务接口
	resp, err := notelis.client.NoteListDel(ctx, req)

	return resp, err
}

