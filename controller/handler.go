package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/liao0001/cloud_disk"
	"github.com/liao0001/cloud_disk/config"
)

type Handler struct {
	conf        *config.Conf
	app         *iris.Application
	res         func(ctx iris.Context) *simpleRes
	resWithUser func(ctx iris.Context) *simpleRes
	needCaptcha bool // 是否需要验证码
	manager     *cloud_disk.Manager
	// service     *service.Service
}

func NewHandler(conf *config.Conf, app *iris.Application, manager *cloud_disk.Manager) {
	a := &Handler{
		conf:    conf,
		app:     app,
		res:     NewSimpleRes,
		manager: manager,
	}

	// 访问静态页面
	if manager.Config.Storage.DefaultStorage.Driver == "local" {
		sconf := manager.Config.Storage.DefaultStorage
		app.HandleDir(sconf.Endpoint, sconf.Bucket, router.DirOptions{
			IndexName: sconf.Endpoint,
			ShowList:  false,
		})
	}

	// 本地的upload  直接复制文件
	// app.Post("/upload", a.Upload)

	app.Get("/test1", a.Test)
	app.Get("/test2", a.Test2)
	app.Get("/test3", a.Test3)

	// 权限验证
	auth := app.Party("/api", authCheck)
	auth.Get("/test", func(ctx iris.Context) {
		_ = ctx.JSON("ok")
	})

	// 此处为脚本更新占位//
}

func (a *Handler) Test(ctx iris.Context) {
	var err error
	msg := a.res(ctx)
	defer msg.Flush(err)

	msg.Success(ctx.Request().Host)
}

func (a *Handler) Test2(ctx iris.Context) {
	var err error
	msg := a.res(ctx)
	defer msg.Flush(err)

	msg.Error(fmt.Errorf("失败的请求"))
}

// 每一步都需要返回
func (a *Handler) Test3(ctx iris.Context) {
	var err error
	msg := a.res(ctx)
	defer msg.Flush(err)
	return
	msg.Error(fmt.Errorf("失败的请求"))
}
