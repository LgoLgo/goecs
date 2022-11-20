package main

import (
	"apis/oss-web/utils/register/consul"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"apis/oss-web/global"
	"apis/oss-web/initialize"
)

func main() {
	// 初始化 logger
	initialize.InitLogger()

	// 初始化配置
	initialize.InitConfig()

	// 初始化 routers
	Router := initialize.Routers(global.ServerConfig.Port)

	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("service registry failed:", err.Error())
	}
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	//  receive termination signal
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
