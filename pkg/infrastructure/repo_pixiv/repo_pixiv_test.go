package repo_pixiv

import (
	"context"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/require"
	"server/pkg/config"
	"testing"
)

func Test_Download(t *testing.T) {
	t.Run("cannot download", func(t *testing.T) {
		ctx := context.Background()
		m := NewRepoPixiv(&config.Config{
			SeTu: &config.SeTuConfig{},
		})
		res, err := m.FetchPixivPictureToMem(ctx, "https://i.pixiv.re/img-original/img/2021/07/30/21/30/02/91605043_p1.jpg")
		require.Error(t, err)
		pterm.Info.Println(string(res))

	})
	t.Run("can download", func(t *testing.T) {
		ctx := context.Background()
		m := NewRepoPixiv(&config.Config{
			SeTu: &config.SeTuConfig{},
		})
		res, err := m.FetchPixivPictureToMem(ctx, "https://i.pixiv.re/img-original/img/2022/02/13/00/04/45/96196337_p0.jpg")
		require.NoError(t, err)
		pterm.Info.Println(len(res))
	})
}
