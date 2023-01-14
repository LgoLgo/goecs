package global

import (
	"srvs/goods_srv/config"

	"gorm.io/gorm"

	"github.com/olivere/elastic/v7"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
	EsClient     *elastic.Client
)
