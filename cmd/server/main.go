package main

import (
	"flag"
	"github.com/fatih/structs"
	"github.com/liao0001/cloud_disk/config"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io"
	stdLog "log"
	"os"
)

var configPath string

func init() {
	// 配置文件
	configPath = os.Getenv("configPath")
	var cp string
	flag.StringVar(&cp, "c", "./config.yml", "配置文件")
	flag.Parse()

	if cp != "" {
		configPath = cp
	}
}

func main() {
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatal("load main config fail:", err.Error())
		return
	}
	structs.DefaultTagName = "json"
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)
	var logOut io.Writer
	if len(conf.LogDir) > 0 {
		lo := &lumberjack.Logger{
			Filename:  conf.LogDir,
			MaxSize:   50, // megabytes
			LocalTime: true,
		}
		defer lo.Rotate()
		logOut = lo
	} else {
		logOut = os.Stderr
	}
	logrus.SetOutput(logOut)
	stdLog.SetOutput(logOut)

	//

}
