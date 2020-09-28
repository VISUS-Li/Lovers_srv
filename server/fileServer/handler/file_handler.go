package handler

import (
	"context"
	"github.com/jinzhu/gorm"

	proto "Lovers_srv/server/fileServer/proto"
)

type FileHandler struct {
	DB *gorm.DB
}

func (file *FileHandler) DownLoadFile(ctx context.Context, in *proto.DownLoadFileReq, out *proto.DownLoadFileResp) error {
	return nil
}

func (file *FileHandler) UpLoadFile(ctx context.Context, in *proto.UpLoadFileReq, out *proto.UpLoadFileResp) error {
	return nil
}

func (file *FileHandler) DelFile(ctx context.Context, in *proto.DelFileReq, out *proto.DelFileResp) error {
	return nil
}