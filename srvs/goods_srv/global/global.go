package global

import (
	"gorm.io/gorm"
	"srvs/goods_srv/config"

	"github.com/olivere/elastic/v7"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	EsClient *elastic.Client
)
