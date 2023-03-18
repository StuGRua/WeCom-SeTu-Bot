package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	HttpConfig struct {
		Port int `yaml:"port"`
	} `yaml:"http"`
	LogConfig struct {
		Output string `yaml:"output"`
	} `yaml:"log"`
	SeTu SeTuConfig `yaml:"se_tu"`
}

type SeTuConfig struct {
	SetuApiUrl string   `yaml:"setu_api_url"`
	QWAuth     []string `yaml:"qw_auth"`
	CronStr    string   `yaml:"cron_str"`
	R18        int64    `yaml:"r_18"`
	Tags       []string `yaml:"tags"`
	PicSize    []string `yaml:"pic_size"`
	Proxy      string   `yaml:"proxy"`
}

func NewConfig() *Config {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	return c
}
