package main

import (
	"flag"
	"fmt"
	"lc/netdisk/common/cron"
	"lc/netdisk/internal/middleware"

	"lc/netdisk/internal/config"
	"lc/netdisk/internal/handler"
	"lc/netdisk/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./etc/netdisk-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(middleware.HandleCors)

	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svcCtx)

	//svcGroup := service.NewServiceGroup()
	//svcGroup.Add(server)
	//for _, mq := range mqc.Consumers(c, context.Background(), svcCtx) {
	//	svcGroup.Add(mq)
	//}

	cron.MergeTask()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//svcGroup.Start()
	server.Start()
}
