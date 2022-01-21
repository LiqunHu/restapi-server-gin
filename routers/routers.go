package routers

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
	"github.com/LiqunHu/restapi-server-gin/pkg/middleware"
	"github.com/LiqunHu/restapi-server-gin/service"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(middleware.Ginzap(logger.Logger().Desugar(), time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(middleware.RecoveryWithZap(logger.Logger().Desugar(), true))

	r.POST("/api/auth", service.Echo)

	var serversV1 = [...]map[string]gin.HandlerFunc{UserHandleMap, TestHandleMap}

	apiv1 := r.Group("/api/v1")

	apiv1.Use(middleware.AUTH())
	{
		for idx, _ := range serversV1 {
			for k, v := range serversV1[idx] {
				apiv1.POST(k, v)
			}
		}
	}
	return r
}
