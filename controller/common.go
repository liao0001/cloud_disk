package controller

import (
	"github.com/kataras/iris/v12"
)

// UploadFile
func (a *Handler) UploadFile(ctx iris.Context) {
	var err error
	msg := a.res(ctx)

	file, header, err := ctx.FormFile("file")
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()
	// 添加文件

}
