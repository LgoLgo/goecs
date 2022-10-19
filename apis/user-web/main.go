package main

import (
	"E-commerce-system/apis/user-web/global"
	"E-commerce-system/apis/user-web/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化 logger
	initialize.InitLogger()

	initialize.InitConfig()

	// 初始化 routers
	Router := initialize.Routers(global.ServerConfig.Port)

	zap.S().Info("start server, port:", global.ServerConfig.Port)
	if err := Router.Run(); err != nil {
		zap.S().Panic("start error:", err.Error())
	}

	Router.Spin()
}
