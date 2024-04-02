package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common"
	"strconv"
	"time"
)

func MergeTask() {

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	reportCron := cron.New(cron.WithLocation(timezone))

	hours := make([]string, 0, 24)
	for i := 0; i < 24; i++ {
		hours = append(hours, strconv.Itoa(i))
	}

	_, err := reportCron.AddFunc("* * * * *", common.MergeLogic)
	if err != nil {
		logx.Errorf("MergeTask，添加合并定时任务失败，ERR: [%v]", err)
	}

	logx.Info("MergeTask，添加合并定时任务成功")
	reportCron.Start()
}
