package common

import (
	"time"

	"github.com/LiqunHu/restapi-server-gin/models"
	"gorm.io/gorm"
)

type CommonUser struct {
	models.Model

	UserId            string    `gorm:"primaryKey;comment:用户主键"`
	UserUsername      string    `gorm:"comment:用户名"`
	UserType          string    `gorm:"comment:类型"`
	UserEmail         string    `gorm:"comment:Email"`
	UserCountryCode   string    `gorm:"comment:国家代码"`
	UserPhone         string    `gorm:"comment:手机"`
	UserPassword      string    `gorm:"comment:密码请从各种查询中删除"`
	UserPasswordError int       `gorm:"comment:'密码错误次数 -1未设置密码 0正常'"`
	UserLoginTime     time.Time `gorm:"comment:末次登陆时间"`
	UserName          string    `gorm:"comment:姓名"`
	UserGender        string    `gorm:"comment:性别"`
	UserAvatar        string    `gorm:"comment:用户头像"`
	UserProvince      string    `gorm:"comment:省"`
	UserCity          string    `gorm:"comment:'市/县'"`
	UserDistrict      string    `gorm:"comment:区"`
	UserAddress       string    `gorm:"comment:'地址'"`
	UserZipcode       string    `gorm:"comment:'邮编'"`
	UserCompany       string    `gorm:"comment:'公司'"`
	UserRemark        string    `gorm:"comment:'备注'"`
}

// func (User) TableName() string {
// 	return "tbl_common_user"
// }

func GetUserByPhone(phone string) (*CommonUser, error) {
	var user CommonUser
	err := models.GDB.Where("user_phone = ? ", phone).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
