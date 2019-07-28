# 公共说明
* 所有接口都会返回如下三个参数，即公共返回参数

|参数|类型|说明|
---|---|---
|code|int|请求码|
|msg|string|请求码对应的说明|
|uuid|string|请求的服务端响应唯一标识|

## 错误码对应msg
```
var ErrCode = map[int]string{
	Success: "success", // 200 请求成功

	AuthFailed:       "auth failed",        // 1000 认证未通过
	UserNotFound:     "user not found",     // 1001 用户未找到
	UserAlreadyExist: "user already exist", // 1002 用户已存在
	UuidNotFound:     "uuid not found",     // 1003 uuid未找到，请检查授权中间件
	ParamInvalid:     "Param Invalid",      //1004 参数非法，解析请求参数失败
}


```
