package global

import (
	"apis/userop-web/config"
	"apis/userop-web/proto/gen"
)

var (
	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	MessageClient proto.MessageClient
	AddressClient proto.AddressClient
	UserFavClient proto.UserFavClient
)
