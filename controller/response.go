package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/url"
	"time"
)

type HttpCode int

const (
	HttpCodeOK              HttpCode = 0
	HttpCodeUnauthorized    HttpCode = 401 // 没有访问权限
	HttpCodePaymentRequired HttpCode = 402
	HttpCodeNotImplemented  HttpCode = 501
	HttpCodeGatewayTimeout  HttpCode = 504

	HttpCodeSystemError    HttpCode = 500
	HttpCodeResultNotFound HttpCode = 1001
	HttpCodeDuplicateKey   HttpCode = 1002
	HttpCodeParamInvalid   HttpCode = 1003
)

type HttpRes struct {
	Code    HttpCode    `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Type    string      `json:"type"`
}

type simpleRes struct {
	ctx       iris.Context
	hasReturn bool
	userId    string // 用户id
	user      any    // 用户信息
}

func NewSimpleRes(ctx iris.Context) *simpleRes {
	// uid, _ := ctx.Value("_user_id").(string)
	uid := "test_001"
	return &simpleRes{
		ctx:    ctx,
		userId: uid,
	}
}

func (s *simpleRes) Write(p []byte) (n int, err error) {
	if s.hasReturn {
		return
	}
	s.hasReturn = true
	return s.ctx.Write(p)
}

func (s *simpleRes) Success(data interface{}) {
	if s.hasReturn {
		return
	}
	s.hasReturn = true
	if data == nil {
		data = map[string]any{}
	}
	_ = s.ctx.JSON(HttpRes{
		Code:    HttpCodeOK,
		Result:  data,
		Type:    "success",
		Message: "ok",
	})
}

func (s *simpleRes) SuccessDefault() {
	if s.hasReturn {
		return
	}
	s.hasReturn = true
	_ = s.ctx.JSON(HttpRes{
		Code:    HttpCodeOK,
		Result:  map[string]any{},
		Type:    "success",
		Message: "ok",
	})
}

func (s *simpleRes) Error(err error) {
	if s.hasReturn {
		return
	}
	s.hasReturn = true
	_ = s.ctx.JSON(HttpRes{
		Code:    HttpCodeSystemError,
		Result:  map[string]any{},
		Type:    "error",
		Message: parseError(err),
	})
}

func (s *simpleRes) Errorf(format string, args ...any) {
	if s.hasReturn {
		return
	}
	s.hasReturn = true
	_ = s.ctx.JSON(HttpRes{
		Code:    HttpCodeSystemError,
		Result:  map[string]any{},
		Type:    "error",
		Message: fmt.Sprintf(format, args...),
	})
}

func (s *simpleRes) Flush(err error) {
	if err != nil {
		s.Error(err)
		return
	}
	s.Success("未设置返回值")
}

func (s *simpleRes) ErrorWithCode(err error, code HttpCode) {
	if s.hasReturn {
		return
	}
	s.hasReturn = true
	_ = s.ctx.JSON(HttpRes{
		Code:    code,
		Result:  map[string]any{},
		Type:    "error",
		Message: parseError(err),
	})
}

func (s *simpleRes) SetExcelHeader(fileName string) *simpleRes {
	s.ctx.Header(context.ContentDispositionHeaderKey, "attachment;filename="+url.PathEscape(fileName))
	s.ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=UTF-8")
	s.ctx.Header("Transfer-Encoding", "chunked")
	s.ctx.Header("Last-Modified", time.Now().Format(time.RFC1123))
	return s
}

// ------------------------ 读取数据 --------------------

// 这里读取query和url参数 读取的是 string,前端需要转成对应的
func (s *simpleRes) readQuery(ptr any) error {
	params := s.ctx.URLParams()
	bs, err := json.Marshal(&params)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}

func (s *simpleRes) readQueryDefault(ptr any) {
	params := s.ctx.URLParams()
	bs, _ := json.Marshal(&params)
	_ = json.Unmarshal(bs, ptr)
}
