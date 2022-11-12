package global

import (
	"apis/goods-web/config"
	"apis/goods-web/proto/gen"
)

var (
	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient
)
