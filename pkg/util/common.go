package util

import (
	"encoding/json"
	"errors"
	"fmt"
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
	var typeError *json.UnmarshalTypeError

	if errors.As(err, &typeError) {
		JsonErr := err.(*json.UnmarshalTypeError)
		return 700, gin.H{
			"errno": "INPUT",
			"msg":   fmt.Sprintf(("'%s' 字段错误"), JsonErr.Field),
		}
	}
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
