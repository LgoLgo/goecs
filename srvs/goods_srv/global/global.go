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

//func init() {
//	dsn := "root:root@tcp(192.168.0.104:3306)/LgoECS_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
//
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold: time.Second,   // Slow SQL Threshold
//			LogLevel:      logger.Info, // Log level
//			Colorful:      true,         // Disable color printing
//		},
//	)
//
//	// global mode
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(err)
//	}
//}
