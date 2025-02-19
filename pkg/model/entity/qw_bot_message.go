package entity

type BotMsgType string

const (
	BotMsgText  BotMsgType = "text"
	BotMsgNews  BotMsgType = "news"
	BotMsgImage BotMsgType = "image"
)

type BotMsgReq struct {
	MsgType BotMsgType `json:"msgtype"`
	News    *BotNews   `json:"news,omitempty"`
	Image   *BotImage  `json:"image,omitempty"`
	Text    *BotText   `json:"text,omitempty"`
}

type BotNews struct {
	Articles []BotArticle `json:"articles"`
}

type BotArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type BotImage struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

type BotText struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list"`
}
