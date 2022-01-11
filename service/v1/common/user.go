package common

import (
	"github.com/LiqunHu/restapi-server-gin/models/common"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	user, err := common.GetUserByPhone("18698729476")
	if err != nil {
		c.JSON(700, gin.H{
			"message": "error",
		})
	}

	c.JSON(200, gin.H{
		"message": user.UserName,
	})
}
