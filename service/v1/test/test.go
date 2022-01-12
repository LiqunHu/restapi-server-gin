package test

import (
	"github.com/LiqunHu/restapi-server-gin/models/test"
	"github.com/LiqunHu/restapi-server-gin/pkg/util"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetTestById(c *gin.Context) {
	var doc GetTestByIdIN
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(700, gin.H{
			"errno": "-1",
			"msg":   err.Error(),
		})
		return
	}
	test, err := test.GetTestByID(doc.Id)
	if err != nil {
		c.JSON(700, gin.H{
			"errno": "-1",
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(util.Success(test))
}

func CreateTest(c *gin.Context) {
	tobj := test.Test{A: "1111", B: "2222", C: "3333"}
	test.CreatTest(&tobj)
	c.JSON(200, gin.H{
		"TestId": tobj.TestId,
	})
}
