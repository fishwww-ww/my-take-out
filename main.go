package main

import (
	"my-take-out/initialize"
)

func main() {
	// 初始化配置
	router := initialize.GlobalInit()

	router.Run(":8080")
}
