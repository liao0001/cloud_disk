package controller

import (
	"github.com/kataras/iris/v12"
)

//
var authCheck = iris.Handler(func(ctx iris.Context) {

	ctx.Next()
})
