package global

import (
	"E-commerce-system/apis/user-web/config"
	pb "E-commerce-system/apis/user-web/proto/gen"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator

	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	UserSrvClient pb.UserClient
)
