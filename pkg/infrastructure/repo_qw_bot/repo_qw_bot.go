package repo_qw_bot

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/pkg/config"
	"server/pkg/domain/entity"
	"server/pkg/domain/repo"
)

type RepoQWBot interface {
	repo.QWBotRepo
}

type qwBotRepo struct {
	config     *config.Config
	httpClient *http.Client
}

func NewRepoQWBot(config *config.Config) RepoQWBot {
	r := qwBotRepo{
		config:     config,
		httpClient: &http.Client{},
	}
	return &r
}

func (m *qwBotRepo) SendBotMessageBatch(ctx context.Context, url string, reqSlice []*entity.BotMsgReq) (err error) {
	for _, item := range reqSlice {
		err = m.SendBotMessage(ctx, url, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *qwBotRepo) SendBotMessage(ctx context.Context, url string, req *entity.BotMsgReq) error {
	postStr, err := json.Marshal(req)
	if err != nil {
		logrus.WithContext(ctx).Errorln("Json marshal post failed.", err)
		return err
	}
	respPost, err := http.Post(url, "application/json", bytes.NewBuffer(postStr))
	if err != nil {
		logrus.WithContext(ctx).Errorln("Post to wechat failed", err)
		return err
	}
	msg, err := io.ReadAll(respPost.Body)
	if err != nil {
		logrus.WithContext(ctx).Errorln(err)
		return err
	}
	logrus.WithContext(ctx).Println(string(msg))
	_ = respPost.Body.Close()
	return nil
}
