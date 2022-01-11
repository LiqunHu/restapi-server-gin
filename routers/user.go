package routers

import (
	"github.com/LiqunHu/restapi-server-gin/service/v1/common"
	"github.com/gin-gonic/gin"
)

var UserHandleMap map[string]gin.HandlerFunc

func init() {
	UserHandleMap = make(map[string]gin.HandlerFunc)
	UserHandleMap["/User/GetUser"] = common.GetUser
}
