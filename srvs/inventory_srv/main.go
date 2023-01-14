package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"srvs/inventory_srv/global"
	"srvs/inventory_srv/handler"
	"srvs/inventory_srv/initialize"
	"srvs/user_srv/utils"
	"syscall"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	proto "srvs/inventory_srv/proto/gen"
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
	}

	zap.S().Info("port: ", *Port)

	server := grpc.NewServer()
	proto.RegisterInventoryServer(server, &handler.InventoryServer{})
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
	registration.Tags = []string{"L2ncE", "inventory", "srv"}
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

	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"10.17.105.123:9876"}),
		consumer.WithGroupName("LgoECS-inventory"),
	)

	if err := c.Subscribe("order_reback", consumer.MessageSelector{}, handler.AutoReback); err != nil {
		zap.S().Info("Failed to read message")
	}
	_ = c.Start()

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
