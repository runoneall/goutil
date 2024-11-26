package goutil

import (
	configutil "github.com/spf13/viper"
)

func ReadConfig(filepath string) map[string]interface{} {
	configutil.SetConfigFile(filepath)
	configutil.ReadInConfig()
	var configMap map[string]interface{}
	configutil.Unmarshal(&configMap)
	return configMap
}
