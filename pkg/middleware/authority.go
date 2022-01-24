package middleware

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LiqunHu/restapi-server-gin/models"
	"github.com/LiqunHu/restapi-server-gin/pkg/gredis"
	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
	"github.com/LiqunHu/restapi-server-gin/pkg/setting"
	"github.com/LiqunHu/restapi-server-gin/pkg/util"
	"github.com/gin-gonic/gin"
)

type AuthAPI struct {
	ApiFunction string `json:"api_function"`
	AuthFlag    string `json:"auth_flag"`
}

type UserStruct struct {
	UserId                  string `json:"user_id"`
	UserUsername            string `json:"user_username"`
	UserType                string `json:"user_type"`
	UserEmail               string `json:"user_email"`
	UserCountryCode         string `json:"user_country_code"`
	UserPhone               string `json:"user_phone"`
	UserPasswordError       int    `json:"user_password_error"`
	UserLoginTime           string `json:"user_login_time"`
	UserName                string `json:"user_name"`
	UserGender              string `json:"user_gender"`
	UserAvatar              string `json:"user_avatar"`
	UserProvince            string `json:"user_province"`
	UserCity                string `json:"user_city"`
	UserDistrict            string `json:"user_district"`
	UserAddress             string `json:"user_address"`
	UserZipcode             string `json:"user_zipcode"`
	UserCompany             string `json:"user_company"`
	UserRemark              string `json:"user_remark"`
	DefaultOrganization     int    `json:"default_organization"`
	DefaultOrganizationCode string `json:"default_organization_code"`
	CreatedAt               string `json:"created_at"`
}
type ApiStruct struct {
	ApiName     string `json:"api_name"`
	ApiFunction string `json:"api_function"`
}
type UserCache struct {
	SessionToken string      `json:"session_token"`
	User         UserStruct  `json:"user"`
	AuthApis     []ApiStruct `json:"authApis"`
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

			checkResult := token2user(c, c.Request.Header["Authorization"][0], mdFuc)

			exitFlag := false
			for k, _ := range apis {
				if mdFuc == k {
					exitFlag = true
				}
			}

			if exitFlag {
				if apis[mdFuc] == "1" {
					if checkResult != 0 {
						if checkResult == 2 {
							c.JSON(http.StatusUnauthorized, gin.H{
								"errno": "-2",
								"msg":   "Login from other place",
							})
							c.Abort()
							return
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
			} else {
				if mdFuc != "AUTH" {
					c.JSON(http.StatusUnauthorized, gin.H{
						"errno": "-1",
						"msg":   "Auth Failed or session expired",
					})
					c.Abort()
					return
				}
			}
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

func token2user(c *gin.Context, token string, mdFuc string) int {
	tokenSplit := strings.Split(token, "_")
	if len(tokenSplit) != 5 {
		return -1
	}
	tokenType := tokenSplit[0]
	uid := tokenSplit[1]
	expires := tokenSplit[3]
	sha1str := tokenSplit[4]

	logger.Infof("tokenType=[%s] uid=[%s] expires=[%s] sha1=[%s]", tokenType, uid, expires, sha1str)
	exp, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return -1
	}
	logger.Debug(time.Now().Unix())
	if exp < time.Now().Unix() {
		logger.Debug("expires")
		return -3
	}
	userKey := "AUTH_" + tokenType + "_" + uid
	kexist := gredis.Exists(userKey)
	if kexist {
		authBuf, err := gredis.Get(userKey)
		if err != nil {
			logger.Error(err.Error())
			return -1
		}
		var uCache UserCache
		err = json.Unmarshal(authBuf, &uCache)
		if err != nil {
			logger.Error(err.Error())
			return -1
		}

		if uCache.SessionToken != token {
			logger.Debug("login from other place")
			return -2
		}

		s := strings.Join([]string{tokenType, uid, uCache.User.CreatedAt, expires, setting.AppSetting.SecretKey}, "_")
		hash := sha1.New()
		hash.Write([]byte(s))
		ss := hash.Sum(nil)
		if fmt.Sprintf("%x", ss) != sha1str {
			logger.Error("invalid sha1")
			return -1
		}

		exist := false

		for _, item := range uCache.AuthApis {
			if item.ApiName == mdFuc {
				exist = true
				break
			}
		}

		if exist {
			c.Set("User", uCache.User)
			return 0
		}
	} else {
		return -1
	}

	return -1
}
