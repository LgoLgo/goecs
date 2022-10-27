package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"srvs/user_srv/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("ECS_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./user_srv/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("./user_srv/%s-debug.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("Config Info: %v", global.ServerConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("config file changed: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("Config Info: %v", global.ServerConfig)
	})
}
