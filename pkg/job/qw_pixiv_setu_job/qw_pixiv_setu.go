package qw_pixiv_setu_job

import (
	"context"
	"github.com/pterm/pterm"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"server/pkg/biz/se_tu"
	"server/pkg/config"
	"time"
)

type CronTab struct {
	cron    *cron.Cron
	config  config.Config
	seTuBiz se_tu.Biz
}

func NewCronTab(config config.Config, seTuBiz se_tu.Biz) *CronTab {
	pterm.Info.Println("init cron job")
	l := cron.VerbosePrintfLogger(logrus.StandardLogger())
	c := CronTab{
		cron:    cron.New(cron.WithLocation(time.Local), cron.WithLogger(l), cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(l))),
		config:  config,
		seTuBiz: seTuBiz,
	}
	ctx := context.Background()
	// 每天0点执行
	addFunc, err := c.cron.AddFunc(c.config.SeTuCronStr, func() {
		execErr := c.seTuBiz.SeTuProcess()
		if execErr != nil {
			logrus.WithContext(ctx).Error(execErr)
		}
	})
	if err != nil {
		logrus.WithContext(ctx).Errorf("failed init cron SeTuProcess: %v", err)
		return nil
	}
	logrus.WithContext(ctx).Println("success init cron: ", addFunc)
	// run
	c.cron.Start()
	pterm.Info.Println("init cron job success")
	return &c
}

func (m *CronTab) Run() error {
	m.cron.Start()
	return nil
}
