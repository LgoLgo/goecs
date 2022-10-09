package main

import (
	"E-commerce-system/srvs/user_srv/handler"
	"E-commerce-system/srvs/user_srv/proto/gen"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	flag.Parse()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
