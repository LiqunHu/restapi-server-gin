package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/LiqunHu/restapi-server-gin/pkg/middleware"
	"github.com/LiqunHu/restapi-server-gin/service"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
