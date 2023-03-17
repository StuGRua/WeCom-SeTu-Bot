package main

import (
	"server/pkg/config"
	"server/pkg/infrastructure/repo_pixiv"
	"server/pkg/infrastructure/repo_qw_bot"
	"server/pkg/infrastructure/repo_setu"
	"server/pkg/log"
	"server/server"
)

func main() {
	srv := server.NewServer() //创建一个服务器
	srv.Provide(
		log.GetLogger,            //依赖注入Logger
		config.NewConfig,         //依赖注入配置文件
		repo_setu.NewRepoSetu,    // setu mod
		repo_pixiv.NewRepoPixiv,  // pixiv mod
		repo_qw_bot.NewRepoQWBot, // qw bot mod
	)
	//jobCron := job.NewJob() //创建一个job
	//jobCron.Provide(
	//	config.NewConfig,        //依赖注入配置文件
	//	repo_setu.NewRepoSetu,   // setu mod
	//	repo_pixiv.NewRepoQWBot, // pixiv mod
	//)
	//jobCron.Run() //运行job
	srv.Run() //运行服务
}
