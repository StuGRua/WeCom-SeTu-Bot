package repo_pixiv

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/pkg/config"
	"server/pkg/domain/repo"
)

type RepoPixiv interface {
	repo.PicRepo
}

type repoPixiv struct {
	config     *config.Config
	httpClient *http.Client
}

func NewRepoPixiv(config *config.Config) RepoPixiv {
	r := repoPixiv{
		config:     config,
		httpClient: &http.Client{},
	}
	return &r
}

func (r repoPixiv) FetchPixivPictureToMem(ctx context.Context, url string) ([]byte, error) {
	var dlReq *http.Request
	var dlResp *http.Response
	var err error
	dlReq, err = http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	if err != nil {
		logrus.WithContext(ctx).Println(err)
		return nil, err
	}
	cli := http.Client{}
	dlResp, err = cli.Do(dlReq)
	if err != nil {
		logrus.WithContext(ctx).Println("[FetchPixivPictureToMem] Download picture failed.", err)
		return nil, err
	}
	picData, err := io.ReadAll(dlResp.Body)
	if err != nil {
		logrus.WithContext(ctx).Println("[FetchPixivPictureToMem] io.ReadAll(dlResp.Body) failed.", err)
		return nil, err
	}
	logrus.WithContext(ctx).Println(url, dlResp.Status, " || ", len(picData))
	_ = dlResp.Body.Close()
	return picData, nil
}
