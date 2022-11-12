package main

import (
	"apis/goods-web/utils/register/consul"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"apis/goods-web/global"
	"apis/goods-web/initialize"
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

	// 初始化sentinel
	initialize.InitSentinel()

	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("sign out failed")
	} else {
		zap.S().Info("sign out success")
	}

	Router.Spin()
}
