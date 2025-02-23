package main

import (
	"os"
	"server/pkg/biz/se_tu"
	"server/pkg/config"
	"server/pkg/infrastructure/repo_pixiv"
	"server/pkg/infrastructure/repo_qw_bot"
	"server/pkg/infrastructure/repo_setu"
	"server/pkg/job/qw_pixiv_setu_job"
	"syscall"
)

func main() {
	// init config
	cfg := config.NewConfigFromFile()
	// init repo
	repoQW := repo_qw_bot.NewRepoQWBot(cfg)
	repoPixiv := repo_pixiv.NewRepoPixiv(cfg)
	repoSeTu := repo_setu.NewRepoSetu(cfg)
	// init biz
	biz := se_tu.New(cfg, repoQW, repoPixiv, repoSeTu)
	cronTab := qw_pixiv_setu_job.NewCronTab(*cfg, biz)
	err := cronTab.Run()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
