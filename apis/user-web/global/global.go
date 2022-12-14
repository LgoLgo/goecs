package global

import (
	"apis/user-web/config"
	"apis/user-web/proto/gen"
)

var (
	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	UserSrvClient proto.UserClient
)
