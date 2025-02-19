package config

import (
	"encoding/json"
	"flag"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	SeTuCronStr string `yaml:"se_tu_cron_str"`
	//HttpConfig struct {
	//	Port int `yaml:"port"`
	//} `yaml:"http"`
	//LogConfig struct {
	//	Output string `yaml:"output"`
	//} `yaml:"log"`
	SeTu *SeTuConfig `yaml:"se_tu"`
}

type SeTuConfig struct {
	SetuApiUrl  string   `yaml:"setu_api_url"`
	QWAuth      []string `yaml:"qw_auth"`
	R18         int64    `yaml:"r_18"`
	Tags        []string `yaml:"tags"`
	PicSize     []string `yaml:"pic_size"`
	Proxy       string   `yaml:"proxy"`
	DirectProxy string   `yaml:"direct_proxy"`
}

func NewConfigFromFile() *Config {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	GlobalConfig = c
	return c
}

func NewConfigFromFlags() *Config {
	ac := &Config{
		SeTu: &SeTuConfig{
			QWAuth:      []string{},
			R18:         0,
			Tags:        []string{},
			PicSize:     []string{"original"},
			Proxy:       "i.pixiv.re",
			DirectProxy: "https://pixiv.re",
		},
	}
	cfg := ac.SeTu
	// 解析flag
	qwAuth := flag.String("qw_auth", "", "企微api地址")
	r18 := flag.Int64("r18", 0, "是否启用r18")
	tags := flag.String("tags", "[]", "想要的tag")
	picSize := flag.String("pic_size", `["original"]`, "图片尺寸")
	proxy := flag.String("proxy", "i.pixiv.re", "proxy")
	directProxy := flag.String("direct_proxy", "https://pixiv.re", "direct proxy")
	flag.Parse()
	if *qwAuth == "" {
		panic("qw_auth is required")
	}
	cfg.QWAuth = append(cfg.QWAuth, *qwAuth)
	if *r18 != 0 {
		cfg.R18 = *r18
	}
	if *tags != "" {
		err := json.Unmarshal([]byte(*tags), &cfg.Tags)
		if err != nil {
			logrus.Errorf("tags format err: %v", err)
			panic("tags format err")
		}
	}
	if *picSize != "" {
		err := json.Unmarshal([]byte(*picSize), &cfg.PicSize)
		if err != nil {
			logrus.Errorf("pic_size format err: %v", err)
			panic("pic_size format err")
		}
	}
	if *proxy != "" {
		cfg.Proxy = *proxy
	}
	if *directProxy != "" {
		cfg.DirectProxy = *directProxy
	}
	GlobalConfig = ac
	return ac
}

var GlobalConfig = &Config{}
