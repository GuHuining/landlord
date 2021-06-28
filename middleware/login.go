package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginCheckMiddleware 检查是否登录
func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user_id")
		if id == nil {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			c.Next()
		}
	}
}