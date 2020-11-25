package handler

import (
	mpProto "Lovers_srv/MsgPush/proto"
	"context"
	"github.com/jinzhu/gorm"
)

type MsgPushHandler struct {
	DB *gorm.DB
}

func(mp *MsgPushHandler)GetRunStatisticsByUserId(ctx context.Context, in *mpProto.RunStatisByUserIdReq, out *mpProto.RunStatisByUserIdResp) (error) {
	return nil
}
