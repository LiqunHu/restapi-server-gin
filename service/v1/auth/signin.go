package auth

import (
	"errors"

	"github.com/LiqunHu/restapi-server-gin/models"
	"github.com/LiqunHu/restapi-server-gin/models/common"
	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
	"github.com/LiqunHu/restapi-server-gin/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Signin
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param object query SigninIN true "请求参数"
// @Router /api/auth/signin [post]
func Signin(c *gin.Context) {

	var doc SigninIN
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(util.Fail(err))
		return
	}

	var user common.CommonUser
	err := models.GDB.Where("(user_phone = ? OR user_username = ?) AND state = '1'", doc.Username, doc.Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(util.Error("auth_01"))
		return
	} else if err != nil {
		c.JSON(util.Fail(err))
		return
	}

	decrypted, err := util.AesECBDecrypt("Fuh/MikqJyX+/Ca0P887WA==", []byte("e10adc3949ba59abbe56e057f20f883e"))
	if err != nil {
		c.JSON(util.Fail(err))
		return
	}
	logger.Debug(decrypted)

	c.JSON(util.Success(user))
}
