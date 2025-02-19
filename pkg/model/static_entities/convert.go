package static_entities

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"server/pkg/model/entity"
	"server/pkg/utils"
)

// ConvertArchiveWithDataToBotTextMsg 将画作稿件信息转换为企微机器人文本消息
func ConvertArchiveWithDataToBotTextMsg(data *entity.ArchiveWithData) *entity.BotMsgReq {
	var MentionedList []string
	proxyUrl := data.Info.Urls.Original
	rawPixivUrl := fmt.Sprintf("https://www.pixiv.net/artworks/%d", data.Info.Pid)
	txt := &entity.BotText{
		Content:       fmt.Sprintf("proxy图源：%s\npixiv图源：%s", proxyUrl, rawPixivUrl),
		MentionedList: MentionedList,
	}
	postText := &entity.BotMsgReq{
		MsgType: entity.BotMsgText,
		Text:    txt,
	}
	return postText
}

func ConvertArchiveWithDataToBotNewsMsg(data *entity.ArchiveWithData) *entity.BotMsgReq {
	desc := fmt.Sprintf("Author: %s, Tags: ", data.Info.Author)
	for _, tag := range data.Info.Tags {
		desc += tag + " | "
	}
	article := entity.BotArticle{Title: data.Info.Title,
		Description: desc,
		Url:         data.Info.Urls.Original,
		Picurl:      data.Info.Urls.Original}
	var articles []entity.BotArticle
	articles = append(articles, article)
	news := &entity.BotNews{Articles: articles}
	postNews := &entity.BotMsgReq{MsgType: entity.BotMsgNews, News: news}
	return postNews
}

func ConvertArchiveWithDataToPictureMsg(data *entity.ArchiveWithData) *entity.BotMsgReq {
	picData := data.Data
	//picDataSize := len(picData)
	var err error
	picDataJpg, err := utils.TransferPicDataToJpg(picData)
	if err != nil {
		log.Println("[ConvertArchiveWithDataToPictureMsg] TransferPicDataToJpg failed: ", err)
		return nil
	}
	picComp, err := utils.CompressPictureUntilSize(picDataJpg, 2*1024*1024)
	if err != nil {
		log.Println("[ConvertArchiveWithDataToPictureMsg] CompressPictureUntilSize failed: ", err)
		return nil
	}
	picCompSize := len(picComp)
	log.Println("[ConvertArchiveWithDataToPictureMsg] after CompressPictureUntilSize process, pic raw size is: ", picCompSize)
	picBase64 := base64.StdEncoding.EncodeToString(picComp)
	md5Hash := md5.New()
	md5Hash.Write(picComp)
	md5Str := hex.EncodeToString(md5Hash.Sum(nil))
	img := &entity.BotImage{Base64: picBase64, Md5: md5Str}
	postPic := &entity.BotMsgReq{MsgType: entity.BotMsgImage, Image: img}
	return postPic
}
