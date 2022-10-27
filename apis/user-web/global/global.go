package global

import (
	"apis/user-web/config"
	pb "apis/user-web/proto/gen"
)

var (
	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	UserSrvClient pb.UserClient
)
