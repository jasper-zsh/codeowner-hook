package app

import (
	"encoding/json"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	QyWeixinBot string `json:"qyweixin_bot"`
	GithubToken string `json:"github_token"`
}

var Config AppConfig

func InitConfig() {
	file, err := os.Open("conf/config.json")
	if err != nil {
		logrus.Fatalf("Failed to open config file. %+v", err)
		panic(err)
	}
	raw, err := io.ReadAll(file)
	if err != nil {
		logrus.Fatalf("Failed to read config file. %+v", err)
		panic(err)
	}
	err = json.Unmarshal(raw, &Config)
	if err != nil {
		logrus.Fatalf("Failed to parse config file. %+v", err)
		panic(err)
	}
}
