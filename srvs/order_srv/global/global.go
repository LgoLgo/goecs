package global

import (
	"gorm.io/gorm"

	"srvs/order_srv/config"
	"srvs/order_srv/proto/gen"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
)
