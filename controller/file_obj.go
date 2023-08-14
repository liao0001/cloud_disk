package controller

import "github.com/kataras/iris/v12"

// FilesDirCreate 添加目录
func (a *Handler) FilesDirCreate(ctx iris.Context) {
	var err error
	msg := a.res(ctx)
	defer func() {
		msg.Flush(err)
	}()

	var param struct {
		ParentID string `json:"parent_id"`
		DirName  string `json:"dir_name"`
	}
	if err = ctx.ReadJSON(&param); err != nil {
		msg.Error(err)
		return
	}

	err = a.service.FileObject().CreateDir(msg.userId, param.ParentID, param.DirName)
	return
}
