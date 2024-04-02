package mqc

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/internal/svc"
)

type LogConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *LogConsumer {
	return &LogConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogConsumer) Consume(key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	return nil
}
