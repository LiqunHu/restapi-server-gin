package util

import (
	"github.com/LiqunHu/restapi-server-gin/pkg/e"
	"github.com/gin-gonic/gin"
)

func Success(obj interface{}) (int, gin.H) {
	return 200, gin.H{
		"errno": "0",
		"msg":   "success",
		"info":  obj,
	}
}

func Error(code string) (int, gin.H) {
	return 700, gin.H{
		"errno": code,
		"msg":   e.GetMsg(code),
	}
}

func Fail(err error) (int, gin.H) {
	return 500, gin.H{
		"errno": "-1",
		"msg":   err.Error(),
	}
}
