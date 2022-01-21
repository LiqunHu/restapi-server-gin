package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/LiqunHu/restapi-server-gin/models"
	"github.com/LiqunHu/restapi-server-gin/pkg/gredis"
	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
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
			// 获取API权限列表
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

			patha := strings.Split(c.Request.URL.Path, "/")
			if len(patha) < 3 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"errno": "-1",
					"msg":   "Auth Failed or session expired",
				})
				c.Abort()
				return
			}
			mdFuc := strings.ToUpper(patha[len(patha)-2])
			println(mdFuc)
			result := token2user(c, c.Request.Header["Authorization"][0])
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

func token2user(c *gin.Context, token string) int {
	tokenSplit := strings.Split(token, "_")
	if len(tokenSplit) != 5 {
		return -1
	}
	tokenType := tokenSplit[0]
	uid := tokenSplit[1]
	expires := tokenSplit[3]
	sha1 := tokenSplit[4]

	logger.Infof("tokenType=[%s] uid=[%s] expires=[%s] sha1=[%s]", tokenType, uid, expires, sha1)

	return 0
}
