package main

import (
	"E-commerce-system/apis/user-web/initialize"
	"go.uber.org/zap"
)

func main() {
	port := 8021

	// 初始化 logger
	initialize.InitLogger()
	// 初始化 routers
	Router := initialize.Routers(port)

	zap.S().Info("start server, port:", port)
	if err := Router.Run(); err != nil {
		zap.S().Panic("start error:", err.Error())
	}

	Router.Spin()
}
