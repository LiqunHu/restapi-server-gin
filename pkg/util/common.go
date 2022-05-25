package util

import (
	"reflect"

	"github.com/LiqunHu/restapi-server-gin/pkg/e"
	"github.com/gin-gonic/gin"
)

func Success(obj interface{}) (int, gin.H) {
	if IsNull(obj) {
		return 200, gin.H{
			"errno": "0",
			"msg":   "success",
			"info":  gin.H{},
		}
	} else {
		return 200, gin.H{
			"errno": "0",
			"msg":   "success",
			"info":  obj,
		}
	}
}

func Error(code string) (int, gin.H) {
	return 700, gin.H{
		"errno": code,
		"msg":   e.GetMsg(code),
	}
}

func Fail(err error) (int, gin.H) {
	// typeError := errors.New("github.com/go-playground/validator/v10.ValidationErrors")

	// if errors.As(err, &typeError) {
	// 	return 700, gin.H{
	// 		"errno": "INPUT",
	// 		"msg":   err.Error(),
	// 	}
	// }

	return 500, gin.H{
		"errno": "-1",
		"msg":   err.Error(),
	}
}

func IsNull(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
