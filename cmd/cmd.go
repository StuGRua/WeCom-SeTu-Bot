package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"server/pkg/biz/se_tu"
	"server/pkg/config"
	"server/pkg/infrastructure/repo_pixiv"
	"server/pkg/infrastructure/repo_qw_bot"
	"server/pkg/infrastructure/repo_setu"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfigFromFlags()
	logrus.WithContext(ctx).Warnf("config: %+v", cfg.String())
	// init repo
	repoQW := repo_qw_bot.NewRepoQWBot(cfg)
	repoPixiv := repo_pixiv.NewRepoPixiv(cfg)
	repoSeTu := repo_setu.NewRepoSetu(cfg)
	// init biz
	biz := se_tu.New(cfg, repoQW, repoPixiv, repoSeTu)
	err := biz.SeTuProcess(ctx)
	if err != nil {
		logrus.Errorf("SeTuProcess err: %v", err)
	}
}
