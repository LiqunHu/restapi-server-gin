package main

import (
	"fmt"
	"net/http"

	"github.com/LiqunHu/restapi-server-gin/models"
	"github.com/LiqunHu/restapi-server-gin/pkg/gredis"
	"github.com/LiqunHu/restapi-server-gin/pkg/setting"
	"github.com/LiqunHu/restapi-server-gin/routers"
)

func init() {
	setting.Setup()
	gredis.Setup()
	models.Setup()
}

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    10 * setting.ServerSetting.ReadTimeout,
		WriteTimeout:   10 * setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
