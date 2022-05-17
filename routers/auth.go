package routers

import (
	"github.com/LiqunHu/restapi-server-gin/service/v1/auth"
	"github.com/gin-gonic/gin"
)

var AuthHandleMap map[string]gin.HandlerFunc

func init() {
	AuthHandleMap = make(map[string]gin.HandlerFunc)
	AuthHandleMap["/signin"] = auth.Signin
}
