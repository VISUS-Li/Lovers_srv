package handler

import (
	"Lovers_srv/helper/Utils"
	proto "Lovers_srv/server/home-service/proto"
	"context"
	"github.com/jinzhu/gorm"
)

type HomeHandler struct {
	DB *gorm.DB
}
func (home* HomeHandler) GetMainCard(ctx context.Context, in *proto.GetMainCardReq, out *proto.GetMainCardResp) error{
	Utils.GetThisWeekStartEnd()
	return nil
}

