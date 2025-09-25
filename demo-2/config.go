package main

type Config struct {
	DBConn string
	Addr   string
}

func NewConfig() *Config {
	return &Config{
		DBConn: "gin:@tcp(1qaz@WSX:3306)/gin",
		Addr:   ":8080",
	}
}
