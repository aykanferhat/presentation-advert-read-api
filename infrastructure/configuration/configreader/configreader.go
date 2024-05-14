package configreader

import (
	"github.com/spf13/viper"
	"os"
	"presentation-advert-read-api/infrastructure/configuration/elastic"
	"presentation-advert-read-api/infrastructure/configuration/log"
	"presentation-advert-read-api/infrastructure/configuration/server"
)

const (
	configPath     = "./configs"
	yamlConfigType = "yaml"
)

func ReadServerConf(serverConfigPath string) *server.Config {
	var conf server.Config
	err := readFile(&conf, serverConfigPath)
	if err != nil {
		log.Panic("Server Config file couldn't read")
	}
	return &conf
}

func ReadLogConfig(logConfigPath string) *log.Config {
	var conf log.Config
	err := readFile(&conf, logConfigPath)
	if err != nil {
		log.Panic("Log Config file couldn't read")
	}
	return &conf
}

func ReadElasticConfig(elasticConfigPath string) elastic.ConfigMap {
	var conf map[string]*elastic.Config
	err := readFile(&conf, elasticConfigPath)
	if err != nil {
		log.Panic("Elastic Config file couldn't read")
	}
	return conf
}

func GetProfile(envName string, defaultValue string) string {
	profile := os.Getenv(envName)
	if profile == "" {
		profile = defaultValue
	}
	return profile
}

func readFile(conf interface{}, filePath string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigType(yamlConfigType)
	viper.SetConfigName(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&conf); err != nil {
		return err
	}
	return nil
}
