package client

import (
	"Lovers_srv/config"
	proto "Lovers_srv/server/note-list/proto"
	"context"
	"github.com/micro/go-micro/client"
	//文件IO

)

type NoteListClient struct {
	client proto.NoteListService
	serviceName  string
}

func NewNoteListClient() *NoteListClient{
	//新建一个服务
	c := proto.NewNoteListService(config.NOTELIST_SRV_NAME, client.DefaultClient)
	return &NoteListClient{
		client:      c,
		serviceName: config.NOTELIST_SRV_NAME,
	}
}

func (notelist *NoteListClient)NoteList_Request(ctx context.Context, req *proto.BaseInfo) error {

}


