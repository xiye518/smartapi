package apis

import "smartapi/internal/service/models"

// 公共的http请求返回值
type Header struct {
	Code int    `json:"code"` // 请求码
	Msg  string `json:"msg"`  // 请求码对应的说明
	Uuid string `json:"uuid"` // 请求的服务端响应唯一标识
}

type GetUserResp struct {
	Header
	Data models.User `json:"data"`
}
