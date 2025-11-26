package config

import (
	"log"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化GORM数据库连接
func InitDB() error {
	// 数据库连接配置
	dsn := "root:1qaz@WSX@tcp(localhost:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("数据库连接成功")
	return nil
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate() error {
	// 自动创建或更新表结构
	err := DB.AutoMigrate(
		&UserSettingdemo{},
	)
	if err != nil {
		return err
	}

	log.Println("数据库表迁移成功")
	return nil
}

// UserSettings 用户设置表结构
type UserSettingdemo struct {
	ID     uint           `gorm:"primaryKey;autoIncrement"`
	UserID uint           `gorm:"uniqueIndex;not null"` // 每个用户一份配置
	Data   datatypes.JSON `gorm:"type:json;not null"`   // 保存 data JSON
}
