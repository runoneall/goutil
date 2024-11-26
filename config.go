package goutil

import (
	configutil "github.com/spf13/viper"
)

func Config_Read(filepath string) map[string]interface{} {
	configutil.SetConfigFile(filepath)
	configutil.ReadInConfig()
	var configMap map[string]interface{}
	configutil.Unmarshal(&configMap)
	return configMap
}

func Config_Write(filepath string, configMap map[string]interface{}) {
	configutil.SetConfigFile(filepath)
	for key, value := range configMap {
		configutil.Set(key, value)
	}
	configutil.WriteConfig()
}
