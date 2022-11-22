package main

import (
	"apis/order-web/utils/register/consul"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"apis/order-web/global"
	"apis/order-web/initialize"
)

func main() {
	// Init logger
	initialize.InitLogger()

	// Init config
	initialize.InitConfig()

	// Init connection to microservice
	initialize.InitSrvConn()

	// Init routers
	Router := initialize.Routers(global.ServerConfig.Port)

	// 初始化sentinel
	initialize.InitSentinel()

	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("service registry failed:", err.Error())
	}
	zap.S().Debugf("Start the server, port: %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(); err != nil {
			zap.S().Panic("start error:", err.Error())
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
