package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/LiqunHu/restapi-server-gin/pkg/setting"
)

var GDB *gorm.DB

type Model struct {
	state     string
	version   int
	CreatedAt time.Time `gorm:autoCreateTime"`
	UpdatedAt time.Time `gorm:autoUpdateTime"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	GDB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Datebase)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_", // 表名前缀
			SingularTable: true,   // 使用单数表名
		},
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	GDB.Callback().Update().Replace("gorm:autoUpdateTime", updateTimeStampForUpdateCallback)

	sqlDB, err := GDB.DB()
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		fmt.Println(db.Statement.Schema)
	}
	// if _, ok := scope.Get("gorm:update_column"); !ok {
	// 	scope.SetColumn("ModifiedOn", time.Now().Unix())
	// }
}
