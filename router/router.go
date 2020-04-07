package router

import (
	"net/http"
	"um_sys/middleware/jwt"
	"um_sys/pkg/upload"
	"um_sys/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/register", api.Register)
	r.POST("/api/login", api.Login)
	r.POST("/api/platform_login", api.PlatformLogin)

	chat := r.Group("/api/chat")
	chat.Use(jwt.Jwt())
	{

		chat.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
		chat.POST("/upload", api.UploadImage)
	}

	return r
}
