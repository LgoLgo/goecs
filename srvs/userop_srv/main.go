package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"srvs/user_srv/utils"
	"srvs/userop_srv/global"
	"srvs/userop_srv/handler"
	"srvs/userop_srv/initialize"
	proto "srvs/userop_srv/proto/gen"
	"syscall"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "address")
	Port := flag.Int("port", 0, "post")
	// Init
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()

		zap.S().Info("port: ", *Port)

		server := grpc.NewServer()
		proto.RegisterAddressServer(server, &handler.UserOpServer{})
		proto.RegisterMessageServer(server, &handler.UserOpServer{})
		proto.RegisterUserFavServer(server, &handler.UserOpServer{})
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
		if err != nil {
			panic("failed to listen:" + err.Error())
		}
		// Registration Service Health Check
		grpc_health_v1.RegisterHealthServer(server, health.NewServer())

		// service registry
		cfg := api.DefaultConfig()
		cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port)

		client, err := api.NewClient(cfg)
		if err != nil {
			panic(err)
		}
		// Generate corresponding inspection objects
		check := &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.Host, *Port),
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "15s",
		}

		// Generate registration object
		registration := new(api.AgentServiceRegistration)
		registration.Name = global.ServerConfig.Name
		serviceID := fmt.Sprintf("%s", uuid.NewV4())
		registration.ID = serviceID
		registration.Port = *Port
		registration.Tags = []string{"L2ncE", "userop", "srv"}
		registration.Address = global.ServerConfig.Host
		registration.Check = check

		err = client.Agent().ServiceRegister(registration)
		if err != nil {
			panic(err)
		}

		go func() {
			err = server.Serve(lis)
			if err != nil {
				panic("failed to start grpc:" + err.Error())
			}
		}()

		// receive termination signal
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		if err = client.Agent().ServiceDeregister(serviceID); err != nil {
			zap.S().Info("sign out failed")
		} else {
			zap.S().Info("sign out success")
		}
	}
}
