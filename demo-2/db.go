package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
}
