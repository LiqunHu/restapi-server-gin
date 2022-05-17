package auth

import (
	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
	"github.com/LiqunHu/restapi-server-gin/pkg/util"
	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {

	var doc SigninIN
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(util.Fail(err))
		return
	}

	logger.Debug(doc)

	c.JSON(util.Success(nil))
}
