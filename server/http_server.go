package server

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/macaron.v1"
	"net/http"
	"server/pkg/config"
	"server/pkg/infrastructure/repo_pixiv"
	"server/pkg/infrastructure/repo_qw_bot"
	"server/pkg/infrastructure/repo_setu"
	"server/pkg/job/qw_pixiv_setu"
	"server/pkg/router"
)

type HttpServer struct {
	cfg    *config.Config
	logger *zap.Logger
	mar    *macaron.Macaron
	tab    *qw_pixiv_setu.CronTab
}

func NewHttpServer(cfg *config.Config, logger *zap.Logger, repoPixiv repo_pixiv.RepoPixiv, repoSetu repo_setu.RepoSetu, repoQW repo_qw_bot.RepoQWBot) *HttpServer {
	return &HttpServer{
		cfg:    cfg,
		logger: logger.Named("http_server"),
		mar:    macaron.Classic(),
		tab:    qw_pixiv_setu.NewCronTab(*cfg, repoPixiv, repoSetu, repoQW),
	}
}
func (srv *HttpServer) Run() error {
	router.Register(srv.mar.Router)
	addr := fmt.Sprintf("0.0.0.0:%v", srv.cfg.HttpConfig.Port)
	srv.logger.Info("http run ", zap.String("addr", addr))
	return http.ListenAndServe(addr, srv.mar)
}
