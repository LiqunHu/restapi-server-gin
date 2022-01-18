package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("111111")
		c.Set("user", "1111111111")
		c.Next()
	}
}
