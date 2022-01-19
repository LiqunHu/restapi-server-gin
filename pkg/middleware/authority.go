package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LiqunHu/restapi-server-gin/models"
	"github.com/LiqunHu/restapi-server-gin/pkg/gredis"
	"github.com/LiqunHu/restapi-server-gin/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthAPI struct {
	ApiFunction string `json:"api_function"`
	AuthFlag    string `json:"auth_flag"`
}

func AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Request.Header["Authorization"]; ok {
			//存在
			var apis map[string]string
			exist := gredis.Exists("AUTHAPI")
			if exist {
				apibuf, err := gredis.Get("AUTHAPI")
				if err != nil {
					c.JSON(util.Fail(err))
					c.Abort()
					return
				}
				json.Unmarshal(apibuf, &apis)
			} else {
				rows, err := models.GDB.Raw("SELECT api_function, auth_flag FROM tbl_common_api where state = '1' and api_function != ''").Rows()
				if err != nil {
					c.JSON(util.Fail(err))
					c.Abort()
					return
				}
				var api, auth string
				apis = make(map[string]string)
				for rows.Next() {
					rows.Scan(&api, &auth)
					apis[api] = auth
				}
			}

			fmt.Println(c.Request.Header["Authorization"][0])
			c.Set("user", "1111111111")
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errno": "-1",
				"msg":   "Auth Failed or session expired",
			})
			c.Abort()
			return
		}
	}
}
