package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/koding/multiconfig"
	"github.com/joho/godotenv"
)

// Config 配置结构体
type Config struct {
	DB    string `default:"mysql"`
	DBURL string `default:"localhost"`
	Debug bool `default:"true"`
}

var config Config

// initConfig 初始化config
func initConfig() {
	//获取.env文件
	err := godotenv.Overload()
	if err != nil {
		logrus.Infof("Error loading .env file: %+v", err)
	}
	//读取config default值
	m := multiconfig.MultiLoader(
		&multiconfig.TagLoader{},
		&multiconfig.EnvironmentLoader{
			CamelCase: true,
		},
	)
	//载入全局变量config
	m.Load(&config)
	logrus.Info(config)
	//设置日志等级
	if config.Debug {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "06-01-02 15:04:05.00",
		})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "06-01-02 15:04:05.00",
		})
		logrus.SetLevel(logrus.InfoLevel)
	}
}
