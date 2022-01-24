package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/LiqunHu/restapi-server-gin/models"
	"github.com/LiqunHu/restapi-server-gin/pkg/gredis"
	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
	"github.com/LiqunHu/restapi-server-gin/pkg/setting"
	"github.com/LiqunHu/restapi-server-gin/routers"
)

func init() {
	setting.Setup()
	logger.Setup()
	gredis.Setup()
	models.Setup()
}

// @title           Rest API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   Liqun Hu
// @contact.email  huliquns@126.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    10 * setting.ServerSetting.ReadTimeout,
		WriteTimeout:   10 * setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logger.Infof("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown: %s\n", err)
	}
	logger.Info("Server exiting")
	// s.ListenAndServe()
}
