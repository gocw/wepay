package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var db *gorm.DB

func initDB() {
	switch config.DB {
	case "mongodb":
		session, err := mgo.Dial(config.DBURL)
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
	case "mysql":
		var err error
		path := config.DBURL               //从env获取数据库连接地址
		logrus.Info("path:", string(path)) //打印数据库连接地址
		for {
			db, err = gorm.Open("mysql", string(path)) //使用gorm连接数据库
			if err != nil {
				logrus.Error(err, "Retry in 2 seconds!")
				time.Sleep(time.Second * 2)
				continue
			}
			logrus.Info("DB connect successful!")
			break
		}

	}

}
