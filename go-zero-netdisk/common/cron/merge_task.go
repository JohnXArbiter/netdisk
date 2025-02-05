package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common"
	"log"
	"time"
)

func MergeTask() {

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	mergeCron := cron.New(cron.WithLocation(timezone))

	_, err := mergeCron.AddFunc("*/10 * * * *", common.MergeLogic)
	if err != nil {
		log.Fatalf("MergeTask，添加合并定时任务失败，ERR: [%v]", err)
	}

	logx.Info("MergeTask，添加合并定时任务成功")
	mergeCron.Start()
}
