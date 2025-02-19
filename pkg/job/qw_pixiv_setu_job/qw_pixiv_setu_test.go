package qw_pixiv_setu_job

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"server/pkg/config"
	"server/pkg/infrastructure/repo_setu"
	"server/pkg/model/entity"
	"testing"
)

// mock_repoSetu 是通过mockgen自动生成的，用于模拟repoSetu接口的mock文件

func TestCronTab_getSeTuDescSlice(t *testing.T) {
	testCases := []struct {
		name                   string
		seTuConfig             config.SeTuConfig
		getArchiveInfoSliceRes entity.QueryResult
		getArchiveInfoSliceErr error
		wantErr                bool
	}{
		{
			name: "success",
			seTuConfig: config.SeTuConfig{
				R18:     0,
				Tags:    []string{"tag1", "tag2"},
				PicSize: []string{"large"},
				Proxy:   "",
			},
			getArchiveInfoSliceRes: entity.QueryResult{
				ArchiveSlice: []entity.Archive{
					{
						Title: "pic1",
					},
				},
			},
			getArchiveInfoSliceErr: nil,
			wantErr:                false,
		},
		{
			name: "error",
			seTuConfig: config.SeTuConfig{
				R18:     1,
				Tags:    []string{"tag3"},
				PicSize: []string{"small", "medium"},
				Proxy:   "",
			},
			getArchiveInfoSliceRes: entity.QueryResult{},
			getArchiveInfoSliceErr: errors.New("get archive info slice error"),
			wantErr:                true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			// 创建mock对象
			mockRepoSeTu := repo_setu.NewMockRepoSetu(ctrl)

			m := &CronTab{repoSeTu: mockRepoSeTu}

			// 设置mock对象返回值
			mockRepoSeTu.EXPECT().GetArchiveInfoSlice(ctx, &entity.Query{
				R18:   tc.seTuConfig.R18,
				Num:   1,
				Tag:   tc.seTuConfig.Tags,
				Size:  tc.seTuConfig.PicSize,
				Proxy: tc.seTuConfig.Proxy,
			}).Return(tc.getArchiveInfoSliceRes, tc.getArchiveInfoSliceErr)

			res, err := m.getSeTuDescSlice(ctx, tc.seTuConfig)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.getArchiveInfoSliceRes, res)
			}
		})
	}
}
