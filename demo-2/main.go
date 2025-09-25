package main

import (
	"log"
)

func main() {
	engine, err := InitApp() // 只接收 2 个返回值
	if err != nil {
		log.Fatal(err)
	}
	engine.Run(":8080")
}
