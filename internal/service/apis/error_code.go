package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	Success = 200 // 200 请求成功

	// 请求失败的错误码
	AuthFailed       = iota + 1000 // 1000
	UserNotFound                   // 1001
	UserAlreadyExist               // 1002
	UuidNotFound                   // 1003
	ParamInvalid                   // 1004
)

var ErrCode = map[int]string{
	Success: "success", // 200 请求成功

	AuthFailed:       "auth failed",        // 1000 认证未通过
	UserNotFound:     "user not found",     // 1001 用户未找到
	UserAlreadyExist: "user already exist", // 1002 用户已存在
	UuidNotFound:     "uuid not found",     // 1003 uuid未找到，请检查授权中间件
	ParamInvalid:     "Param Invalid",      //1004 参数非法，解析请求参数失败
}

func getErrCodeMsg(code int) string {
	msg, ok := ErrCode[code]
	if !ok {
		msg = fmt.Sprintf("错误码对应msg未找到，请检查定义,code: %d", code)
	}
	return msg
}

func ErrorMsgWithExtraInfo(c *gin.Context, code int, r Header) {
	r.Code = code
	r.Msg = getErrCodeMsg(code)
	c.JSON(400, r)
}
