// @Title
// @Description
// @Author  Wangwengang  2023/12/22 08:48
// @Update  Wangwengang  2023/12/22 08:48
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// // CasbinHandler 拦截器
func BaseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if span, _, err := plugin.GenerateSpanWithContext(c, "BaseHandler"); err == nil {
		// 	defer span.Finish()
		// }
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		contentType := c.Request.Header.Get("Content-Type")
		var isJson, isProtobuf bool
		isJson = strings.Contains(contentType, "application/json")
		isProtobuf = strings.Contains(contentType, "application/x-protobuf")
		c.Set("isJson", isJson)
		c.Set("isProtobuf", isProtobuf)

		if isJson || isProtobuf {
			// 处理请求
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}

	}
}
