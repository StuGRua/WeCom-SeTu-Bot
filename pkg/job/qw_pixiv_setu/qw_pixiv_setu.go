package qw_pixiv_setu

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"github.com/pterm/pterm"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"server/pkg/config"
	"server/pkg/domain/domain_service/static_entities"
	"server/pkg/domain/entity"
	"server/pkg/domain/repo"
	"server/pkg/infrastructure/repo_pixiv"
	"server/pkg/infrastructure/repo_setu"
	"time"
)

type CronTab struct {
	cron      *cron.Cron
	config    config.Config
	repoPixiv repo_pixiv.RepoPixiv
	repoSeTu  repo_setu.RepoSetu
	repoQW    repo.QWBotRepo
}

func NewCronTab(config config.Config, rp repo_pixiv.RepoPixiv, rs repo_setu.RepoSetu, rq repo.QWBotRepo) *CronTab {
	pterm.Info.Println("init cron job")
	l := cron.VerbosePrintfLogger(logrus.StandardLogger())
	c := CronTab{
		repoSeTu:  rs,
		repoPixiv: rp,
		repoQW:    rq,
		cron:      cron.New(cron.WithLocation(time.Local), cron.WithLogger(l), cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(l))),
		config:    config,
	}
	ctx := context.Background()
	// 每天0点执行
	addFunc, err := c.cron.AddFunc(c.config.SeTu.CronStr, c.SetuProcess)
	if err != nil {
		logrus.WithContext(ctx).Errorf("failed init cron SetuProcess: %v", err)
		return nil
	}
	logrus.WithContext(ctx).Println("success init cron: ", addFunc)
	// run
	c.cron.Start()
	pterm.Info.Println("init cron job success")
	return &c
}

func (m *CronTab) Run() error {
	m.cron.Start()
	return nil
}

type PackBotMsg map[int64][]*entity.BotMsgReq

func (p PackBotMsg) Json() string {
	marshal, _ := json.Marshal(p)
	return string(marshal)
}

func (m *CronTab) SetuProcess() {
	var err error
	ctx := context.Background()
	logrus.WithContext(ctx).Println("SeTuProcess start")
	// 获取图片描述信息
	queryResSlice, err := m.getSeTuDescSlice(ctx)
	if err != nil {
		return
	}
	// 下载图片
	archSlice, err := m.downloadSeTu(ctx, queryResSlice)
	if err != nil {
		return
	}
	// 发送机器人消息
	err = m.sendBotMessages(ctx, archSlice)
	if err != nil {
		return
	}
	logrus.WithContext(ctx).Println("SeTuProcess end")
	return
}

func (m *CronTab) sendBotMessages(ctx context.Context, archSlice []*entity.ArchiveWithData) error {
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
	for _, qw := range m.config.SeTu.QWAuth {
		logrus.WithContext(ctx).Println("ready to send qw: ", md5.Sum([]byte(qw)))
		for _, msg := range packBotMsg {
			err := m.repoQW.SendBotMessageBatch(ctx, qw, msg)
			if err != nil {
				logrus.WithContext(ctx).Errorln("SendBotMessageBatch failed: ", err)
				return err
			}
		}
	}
	return nil
}

func (m *CronTab) downloadSeTu(ctx context.Context, queryResSlice entity.QueryResult) ([]*entity.ArchiveWithData, error) {
	var archSlice []*entity.ArchiveWithData
	var err error
	// 从pixiv下载图片
	for _, archive := range queryResSlice.ArchiveSlice {
		var data []byte
		data, err = m.repoPixiv.FetchPixivPictureToMem(ctx, archive.Urls.Original)
		if err != nil {
			logrus.WithContext(ctx).Errorln("FetchPixivPictureToMem failed: ", err)
			return nil, err
		}
		archSlice = append(archSlice, &entity.ArchiveWithData{Info: archive, Data: data})
	}
	return archSlice, nil
}

func (m *CronTab) getSeTuDescSlice(ctx context.Context) (entity.QueryResult, error) {
	queryResSlice, err := m.repoSeTu.GetArchiveInfoSlice(ctx, &entity.Query{
		R18:   m.config.SeTu.R18,
		Num:   1,
		Tag:   m.config.SeTu.Tags,
		Size:  m.config.SeTu.PicSize,
		Proxy: m.config.SeTu.Proxy,
	})
	if err != nil {
		return entity.QueryResult{}, err
	}
	return queryResSlice, nil
}
