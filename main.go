package main

func main() {
	r := setupRouter() // 调用 router.go 中的 SetupRouter
	r.Run(":8080")     // 监听端口
}
