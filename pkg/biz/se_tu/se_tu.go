package se_tu

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"server/pkg/config"
	"server/pkg/model/entity"
	"server/pkg/model/repo_interface"
	"server/pkg/model/static_entities"
	"strconv"
)

type Biz interface {
	SeTuProcess(ctx context.Context) (err error)
}

type biz struct {
	config    config.Config
	repoQW    repo_interface.QWBotRepo
	repoPixiv repo_interface.PicRepo
	repoSeTu  repo_interface.SeTuRepo
}

func New(cfg *config.Config, repoQW repo_interface.QWBotRepo, repoPixiv repo_interface.PicRepo, repoSeTu repo_interface.SeTuRepo) Biz {
	return &biz{
		config:    *cfg,
		repoQW:    repoQW,
		repoPixiv: repoPixiv,
		repoSeTu:  repoSeTu,
	}
}

// SeTuProcess 定时发送涩图
func (b *biz) SeTuProcess(ctx context.Context) (err error) {
	// 重试3次
	for i := 0; i < 3; i++ {
		var err error
		logrus.WithContext(ctx).Println("SeTuProcess start")
		// 获取图片描述信息
		queryResSlice, err := b.getSeTuDescSlice(ctx, b.config.SeTu)
		if err != nil {
			logrus.WithContext(ctx).Errorf("getSeTuDescSlice err: %v", err)
			continue
		}
		// 下载图片
		archSlice, err := b.downloadSeTu(ctx, queryResSlice)
		if err != nil {
			logrus.WithContext(ctx).Errorf("downloadSeTu err: %v", err)
			continue
		}
		// 发送机器人消息
		err = b.sendBotMessages(ctx, archSlice)
		if err != nil {
			logrus.WithContext(ctx).Errorf("sendBotMessages err: %v", err)
			continue
		}
		logrus.WithContext(ctx).Println("SeTuProcess end")
		break
	}
	return
}

func (b *biz) sendBotMessages(ctx context.Context, archSlice []*entity.ArchiveWithData) error {
	// 打包为机器人消息
	packBotMsg := PackBotMsg{}
	for _, data := range archSlice {
		var singleP []*entity.BotMsgReq
		singleP = append(singleP, static_entities.ConvertArchiveWithDataToBotTextMsg(data))
		singleP = append(singleP, static_entities.ConvertArchiveWithDataToBotNewsMsg(data))
		singleP = append(singleP, static_entities.ConvertArchiveWithDataToPictureMsg(data))
		packBotMsg[data.Info.Pid] = singleP
	}
	logrus.WithContext(ctx).Println("packBotMsg ", len(packBotMsg.Json()))
	for _, qw := range b.config.SeTu.QWAuth {
		logrus.WithContext(ctx).Println("ready to send qw: ", md5.Sum([]byte(qw)))
		for _, msg := range packBotMsg {
			err := b.repoQW.SendBotMessageBatch(ctx, qw, msg)
			if err != nil {
				logrus.WithContext(ctx).Errorln("SendBotMessageBatch failed: ", err)
				return err
			}
		}
	}
	return nil
}

func (b *biz) downloadSeTu(ctx context.Context, queryResSlice entity.QueryResult) ([]*entity.ArchiveWithData, error) {
	var archSlice []*entity.ArchiveWithData
	var err error
	// 从pixiv下载图片
	for _, archive := range queryResSlice.ArchiveSlice {
		var data []byte
		//data, err = m.repoPixiv.FetchPixivPictureToMem(ctx, archive.Urls.Original)
		//if err != nil {
		//	logrus.WithContext(ctx).Errorln("[FetchPixivPictureToMem] via rev proxy failed: ", err)
		//	return nil, err
		//}
		directUrl := config.GlobalConfig.SeTu.DirectProxy + "/" + strconv.FormatInt(archive.Pid, 10) + "." + archive.Ext
		for _, s := range b.config.SeTu.PicSize {
			switch s {
			case "original":
				directUrl = archive.Urls.Original
			case "regular":
				directUrl = archive.Urls.Regular
			case "small":
				directUrl = archive.Urls.Small
			case "thumb":
				directUrl = archive.Urls.Thumb
			case "mini":
				directUrl = archive.Urls.Mini
			}
		}
		data, err = b.repoPixiv.FetchPixivPictureToMem(ctx, directUrl)
		if err != nil {
			logrus.WithContext(ctx).Errorln("[FetchPixivPictureToMem] via direct proxy failed: ", directUrl, " | ", err)
			return nil, err
		}
		archSlice = append(archSlice, &entity.ArchiveWithData{Info: archive, Data: data})
	}
	return archSlice, nil
}

func (b *biz) getSeTuDescSlice(ctx context.Context, seTuConfig *config.SeTuConfig) (entity.QueryResult, error) {
	queryResSlice, err := b.repoSeTu.GetArchiveInfoSlice(ctx, &entity.Query{
		R18:   seTuConfig.R18,
		Num:   1,
		Tag:   seTuConfig.Tags,
		Size:  seTuConfig.PicSize,
		Proxy: seTuConfig.Proxy,
	})
	if err != nil {
		return entity.QueryResult{}, err
	}
	return queryResSlice, nil
}

type PackBotMsg map[int64][]*entity.BotMsgReq

func (p PackBotMsg) Json() string {
	marshal, _ := json.Marshal(p)
	return string(marshal)
}
