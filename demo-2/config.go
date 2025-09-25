package main

import "os"

type Config struct {
	DSN string
}

func NewConfig() *Config {
	// 从环境变量读取，默认 root:123456
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "gin:1qaz@WSX@tcp(oneformadevops.mysql.database.azure.com:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return &Config{DSN: dsn}
}
