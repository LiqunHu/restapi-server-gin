package test

import (
	"fmt"

	"github.com/LiqunHu/restapi-server-gin/models"
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
		c.JSON(util.Fail(err))
		return
	}
	test, err := test.GetTestByID(doc.Id)
	if err != nil {
		c.JSON(util.Fail(err))
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

func UpdateTestById(c *gin.Context) {
	var doc UpdateTestIN
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(util.Fail(err))
		return
	}
	var tdata test.Test
	models.GDB.Model(&tdata).Where("test_id = ?", doc.Id).Updates(test.Test{A: doc.A, B: doc.B, C: doc.C})
	c.JSON(util.Success(nil))
}

func DeleteTestById(c *gin.Context) {
	var doc DeleteTestIN
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(util.Fail(err))
		return
	}
	err := test.DeleteTestByID(doc.Id)
	if err != nil {
		c.JSON(util.Fail(err))
		return
	}
	c.JSON(util.Success(nil))
}

func GetTests(c *gin.Context) {
	var tests []TestResult
	models.GDB.Raw("SELECT a, b, c FROM tbl_test Where test_id > ?", 1).Scan(&tests)

	rows, err := models.GDB.Raw("SELECT a, b, c FROM tbl_test Where test_id > ?", 1).Rows()
	if err != nil {
		c.JSON(util.Fail(err))
		return
	}
	defer rows.Close()
	var tdata TestResult
	for rows.Next() {
		models.GDB.ScanRows(rows, &tdata)
		fmt.Println(tdata)
	}

	c.JSON(util.Success(tests))
}
