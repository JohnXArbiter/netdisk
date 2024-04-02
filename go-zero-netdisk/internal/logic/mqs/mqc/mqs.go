package mqc

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"lc/netdisk/internal/config"
	"lc/netdisk/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		kq.MustNewQueue(c.KqConsumerConf, NewLogConsumer(ctx, svcContext)),
	}
}
