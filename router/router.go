package router

import (
	"um_sys/middleware/jwt"
	"um_sys/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.POST("/api/register", api.Register)
	r.POST("/api/login", api.Login)
	r.POST("/api/platform_login", api.PlatformLogin)

	userGroup := r.Group("/api/user")
	userGroup.Use(jwt.Jwt())
	{

	}
}
