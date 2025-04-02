package repo_setu

import (
	"context"
	"server/pkg/config"
	"server/pkg/model/entity"
	"testing"
)

// RepoSetu.GetArchiveInfoSlice 单元测试
func Test_GetArchiveInfoSlice(t *testing.T) {
	r := NewRepoSetu(&config.Config{
		SeTu: &config.SeTuConfig{},
	})
	rp, err := r.GetArchiveInfoSlice(context.Background(), &entity.Query{
		Num:  1,
		Size: []string{"original"},
		R18:  0,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(rp)
	return
}
