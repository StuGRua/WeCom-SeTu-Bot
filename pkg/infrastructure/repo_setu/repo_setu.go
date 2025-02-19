package repo_setu

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/pkg/config"
	"server/pkg/model/entity"
	"server/pkg/model/repo_interface"
	"time"
)

type repoSeTu struct {
	config     *config.Config
	httpClient *http.Client
}

func NewRepoSetu(config *config.Config) repo_interface.SeTuRepo {
	r := repoSeTu{
		config:     config,
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
	return &r
}

// GetArchiveInfoSlice 从api获取setu信息
func (r *repoSeTu) GetArchiveInfoSlice(ctx context.Context, query *entity.Query) (result entity.QueryResult, err error) {
	jsonStr, err := json.Marshal(query)
	if err != nil {
		logrus.WithContext(ctx).Println("[GetArchiveInfoSlice] Marshal json failed.", err)
		return
	}
	req, err := http.NewRequest("POST", "https://api.lolicon.app/setu/v2", bytes.NewBuffer(jsonStr))
	if err != nil {
		logrus.WithContext(ctx).Println("[GetArchiveInfoSlice] Http request failed.", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := r.httpClient.Do(req)
	if err != nil {
		logrus.WithContext(ctx).Println("[GetArchiveInfoSlice] Http Do failed.", err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		logrus.WithContext(ctx).Println("[GetArchiveInfoSlice] Http Get status is", resp.StatusCode, "not 200")
		return
	}
	bodyStr, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithContext(ctx).Println("[GetArchiveInfoSlice] Read Http Get body failed.", err)
		return
	}
	err = json.Unmarshal(bodyStr, &result)
	if err != nil {
		logrus.WithContext(ctx).Println("[GetArchiveInfoSlice] Json unmarshal failed.", err)
		return
	}
	logrus.WithContext(ctx).Printf("%+v\n", result)
	return
}
