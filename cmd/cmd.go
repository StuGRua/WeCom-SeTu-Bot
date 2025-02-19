package main

import (
	"github.com/sirupsen/logrus"
	"server/pkg/biz/se_tu"
	"server/pkg/config"
	"server/pkg/infrastructure/repo_pixiv"
	"server/pkg/infrastructure/repo_qw_bot"
	"server/pkg/infrastructure/repo_setu"
)

func main() {
	cfg := config.NewConfigFromFlags()
	// init repo
	repoQW := repo_qw_bot.NewRepoQWBot(cfg)
	repoPixiv := repo_pixiv.NewRepoPixiv(cfg)
	repoSeTu := repo_setu.NewRepoSetu(cfg)
	// init biz
	biz := se_tu.New(cfg, repoQW, repoPixiv, repoSeTu)
	err := biz.SeTuProcess()
	if err != nil {
		logrus.Errorf("SeTuProcess err: %v", err)
	}
}
