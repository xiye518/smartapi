package apis

import (
	"github.com/gin-gonic/gin"
	"smartapi/internal/log"
	"smartapi/internal/tools"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//检查访问的api是否放行
		uuid := tools.GetGuid()
		c.Set("uuid", uuid)
		log.Infof("url:%s ,uuid: %s", c.Request.URL.String(), uuid)
		// 检查api授权。
		if !isPass() {
			log.Errorf("check auth not pass,uuid: %s", uuid)
			c.Abort()
		}

		// 授权通过
		c.Set("account", "")
		c.Next()
	}
}

func isPass() bool {
	// 检查授权，一般是验证token

	return false
}
