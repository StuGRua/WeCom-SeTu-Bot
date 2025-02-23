package repo_pixiv

import (
	"bytes"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/pkg/config"
	"server/pkg/model/repo_interface"
)

type repoPixiv struct {
	config     *config.Config
	httpClient *http.Client
}

func NewRepoPixiv(config *config.Config) repo_interface.PicRepo {
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
	if dlResp.StatusCode != 200 {
		err = errors.New(dlResp.Status)
		return nil, err
	}
	_ = dlResp.Body.Close()
	return picData, nil
}
