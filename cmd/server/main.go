package main

import (
	"flag"
	"fmt"
	"github.com/fatih/structs"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/liao0001/cloud_disk"
	"github.com/liao0001/cloud_disk/config"
	"github.com/liao0001/cloud_disk/controller"
	"github.com/liao0001/cloud_disk/db"
	"github.com/liao0001/cloud_disk/models"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io"
	stdLog "log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/liao0001/cloud_disk/storage/local"
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

	globalManager := cloud_disk.NewManager(
		func(manager *cloud_disk.Manager) {
			// 初始化配置
			manager.Config = conf

			// 初始化数据库
			engine, err := db.NewEngine(&conf.DB)
			if err != nil {
				logrus.Fatal("初始化数据库失败:%s", err.Error())
				return
			}
			manager.DB = engine

			// 自动创建表
			_ = initTables(manager)
		},
	)

	var cerr = make(chan error)
	go startHttp(cerr, conf, globalManager)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		cerr <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("exit:%s", <-cerr)
}

func startHttp(cerr chan error, conf *config.Conf, manager *cloud_disk.Manager) {
	app := iris.New()

	// 跨域设置
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	app.WrapRouter(crs.ServeHTTP)

	app.Use(recover2.New())
	app.Use(logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		MessageContextKeys: []string{"user_name"},
	}))

	// 初始化handler
	controller.NewHandler(conf, app, manager)

	err := app.Run(iris.Addr(fmt.Sprintf("0.0.0.0:%d", conf.Http.Port)))
	if err != nil {
		cerr <- err
		return
	}
}

func initTables(manager *cloud_disk.Manager) error {
	return manager.DB.DB.AutoMigrate(
		&models.FileObject{},
	)
}
