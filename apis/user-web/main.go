package main

import (
	"apis/user-web/global"
	"apis/user-web/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化 logger
	initialize.InitLogger()

	// 初始化配置
	initialize.InitConfig()

	// 初始化 srv 的连接
	initialize.InitSrvConn()

	// 初始化 routers
	Router := initialize.Routers(global.ServerConfig.Port)

	zap.S().Info("start server, port:", global.ServerConfig.Port)
	if err := Router.Run(); err != nil {
		zap.S().Panic("start error:", err.Error())
	}

	Router.Spin()
}
