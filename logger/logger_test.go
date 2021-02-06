package logger

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type qlogItem struct {
	FieldOne string `json:"fieldOne"`
	FieldTwo int    `json:"fieldTwo"`
}

func TestLog(t *testing.T) {

	item := &qlogItem{
		FieldOne: "some",
		FieldTwo: 134,
	}
	Convey("Logger", t, func() {
		Convey("直接输出", func() {
			logger.Info("log item 1")
			logger.Info("log item 2")
		})
		Convey("输出到控制台", func() {
			consoleConfig := &Config{
				FileName:        "../log/log.log",
				Level:           "info",
				Output:          "file",
				TimestampFormat: "2006-01-02 15:04:05",
				MaxLogItemSize:  5000,
			}
			initLoggerWithConfig(consoleConfig)
			Convey("不带参数的log", func() {
				logger.Info("log item 1")
				logger.Info("log item 2")
			})
			Convey("带参数的log", func() {
				itemStr, _ := json.Marshal(item)
				logger.InfoF("log item 1 %s %d", itemStr, 156)
			})
		})
	})
}
