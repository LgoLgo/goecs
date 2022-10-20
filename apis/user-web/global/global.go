package global

import (
	"E-commerce-system/apis/user-web/config"
	pb "E-commerce-system/apis/user-web/proto/gen"
)

var (
	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	UserSrvClient pb.UserClient
)
