# 用户
```$xslt



```

## 1.查询用户信息接口
* get    
* /v1/api/user?id=1

### 接口说明
* 用户id和token必传

### 请求参数
|参数|类型|必传|说明|
---|---|---|---
|id|int|是|用户id|
|token|string|是|用户授权码，用于接口鉴权|

### 请求示例
```$xslt
curl -i -X POST -H "Content-Type: application/json" -d '{
  "id": 1,
  "token": "token value",
}'  http://127.0.0.1/v1/api/user

```

### 响应
```
{	
	"code":200,	
	"msg":"success",
	"uuid":"d6627b03-75fa-412e-07c6-0dd6ac22aa19",
	"data": {
	  "id": 1,
	  
	}
}

```