package routers

import (
	"github.com/LiqunHu/restapi-server-gin/service/v1/test"
	"github.com/gin-gonic/gin"
)

var TestHandleMap map[string]gin.HandlerFunc

func init() {
	TestHandleMap = make(map[string]gin.HandlerFunc)
	TestHandleMap["/Test/Test"] = test.Test
	TestHandleMap["/Test/GetTestById"] = test.GetTestById
	TestHandleMap["/Test/CreateTest"] = test.CreateTest
}
