package main

import (
	"fmt"
	"net/http"
	"um_sys/module"
	"um_sys/pkg/setting"
	"um_sys/pkg/util"
	"um_sys/router"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	module.Setup()
	util.Setup()
}

func main() {

	fmt.Println("========")
	gin.SetMode(setting.ServerSetting.RunMode)

	engine := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := http.Server{
		Addr:           endPoint,
		Handler:        engine,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()

}
