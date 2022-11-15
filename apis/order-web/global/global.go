package global

import (
	"apis/order-web/config"
	"apis/order-web/proto/gen"
)

var (
	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	OrderSrvClient proto.OrderClient

	InventorySrvClient proto.InventoryClient
)
