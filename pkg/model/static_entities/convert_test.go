package static_entities

import (
	"github.com/stretchr/testify/require"
	"server/pkg/model/entity"
	"testing"
)

func Test_ConvertArchiveWithDataToBotTextMsg(t *testing.T) {
	testCases := []struct {
		name     string
		data     *entity.ArchiveWithData
		expected *entity.BotMsgReq
	}{
		{
			name: "regular data",
			data: &entity.ArchiveWithData{
				Info: entity.Archive{
					Pid: 123456,
					Urls: entity.PicUrl{
						Original: "https://example.com/image.jpg",
					},
				},
				Data: []byte{},
			},
			expected: &entity.BotMsgReq{
				MsgType: entity.BotMsgText,
				Text: &entity.BotText{
					Content: "proxy图源：https://example.com/image.jpg\npixiv图源：https://www.pixiv.net/artworks/123456",
				},
			},
		},
		{
			name: "data with special characters",
			data: &entity.ArchiveWithData{
				Info: entity.Archive{
					Pid: 789012,
					Urls: entity.PicUrl{
						Original: "https://example.com/image(with)special&characters.jpg",
					},
				},
				Data: []byte{},
			},
			expected: &entity.BotMsgReq{
				MsgType: entity.BotMsgText,
				Text: &entity.BotText{
					Content: "proxy图源：https://example.com/image(with)special&characters.jpg\npixiv图源：https://www.pixiv.net/artworks/789012",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ConvertArchiveWithDataToBotTextMsg(tc.data)
			require.Equal(t, tc.expected, result)
		})
	}
}
