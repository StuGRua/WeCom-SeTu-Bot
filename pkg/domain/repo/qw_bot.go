package repo

import (
	"context"
	"server/pkg/domain/entity"
)

type QWBotRepo interface {
	SendBotMessage(ctx context.Context, url string, req *entity.BotMsgReq) error
	SendBotMessageBatch(ctx context.Context, url string, reqSlice []*entity.BotMsgReq) error
}
