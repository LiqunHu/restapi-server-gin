package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"

	"github.com/LiqunHu/restapi-server-gin/pkg/logger"
	"github.com/LiqunHu/restapi-server-gin/pkg/setting"
)

var GDB *gorm.DB

type Model struct {
	State     string
	Version   int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	log := zapgorm2.New(logger.Logger().Desugar())
	log.SetAsDefault()
	GDB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Datebase)), &gorm.Config{
		Logger: log,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_", // 表名前缀
			SingularTable: true,   // 使用单数表名
		},
	})

	if err != nil {
		logger.Fatalf("models.Setup err: %v", err)
	}

	GDB.Callback().Update().Replace("gorm:before_update", updateVersionForUpdateCallback)

	sqlDB, err := GDB.DB()
	if err != nil {
		logger.Fatalf("models.Setup err: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateVersionForUpdateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		fmt.Println(db.Statement.Schema)
		field := db.Statement.Schema.LookUpField("Version")
		if field != nil {
			val, _ := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
			db.Statement.SetColumn("Version", val.(int)+1)
		}
	}
}
