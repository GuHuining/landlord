package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"landlord/service"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	server := gin.Default()
	server.Use(sessions.Sessions("landlord", service.Store)) // 设置session
	api := server.Group("/api") // api统一以/api开头

	userGroupRegister(api)

	if err := server.Run(":8082"); err != nil {
		log.Fatalf("starting: %v", err)
	}
}

// userGroupRegister 注册与用户信息有关的api。 /api/user
func userGroupRegister(api *gin.RouterGroup) {
	group := api.Group("/user")
	//	/api/user/register
	group.POST("/register", service.Register)
	group.POST("/validate_code", service.SendValidateCode)
	group.POST("/login", service.Login)
	group.POST("/login_check", service.LoginCheck)
}
