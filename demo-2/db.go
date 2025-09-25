package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DBConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
