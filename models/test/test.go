package test

import (
	"github.com/LiqunHu/restapi-server-gin/models"
)

type Test struct {
	models.Model

	TestId uint   `gorm:"primaryKey;autoIncrement"`
	A      string `gorm:"comment:A"`
	B      string `gorm:"comment:B"`
	C      string `gorm:"comment:V"`
}

// func (User) TableName() string {
// 	return "tbl_common_user"
// }

func CreatTest(test *Test) error {
	err := models.GDB.Create(test).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTestByID(id int) (*Test, error) {
	var test Test
	err := models.GDB.Where("test_id = ? ", id).First(&test).Error
	if err != nil {
		return nil, err
	}

	return &test, nil
}
