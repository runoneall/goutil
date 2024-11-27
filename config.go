package goutil

import (
	"log"
	"strings"

	configutil "github.com/spf13/viper"
)

func Config_Read_File(filepath string) map[string]interface{} {
	configutil.SetConfigFile(filepath)
	err := configutil.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	var configMap map[string]interface{}
	err = configutil.Unmarshal(&configMap)
	if err != nil {
		log.Fatal(err)
	}
	return configMap
}

func Config_Write_File(filepath string, configMap map[string]interface{}) {
	configutil.SetConfigFile(filepath)
	for key, value := range configMap {
		configutil.Set(key, value)
	}
	err := configutil.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func Config_Format(content string, targetFormat string) map[string]interface{} {
	configutil.SetConfigType(targetFormat)
	err := configutil.ReadConfig(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	var configMap map[string]interface{}
	err = configutil.Unmarshal(&configMap)
	if err != nil {
		log.Fatal(err)
	}
	return configMap
}
