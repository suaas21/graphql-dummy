package config

import (
	"github.com/spf13/viper"
	"sync"
)

// Arango holds arango config
type Arango struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

var arangoOnce = sync.Once{}
var arangoConfig *Arango

// loadArango loads config from path
func loadArango(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	arangoConfig = &Arango{
		Host:       viper.GetString("arango.host"),
		Port:       viper.GetString("arango.port"),
		DBUser:     viper.GetString("arango.db_user"),
		DBPassword: viper.GetString("arango.db_password"),
		DBName:     viper.GetString("arango.db_name"),
	}

	return nil
}

// GetArango returns arango config
func GetArango(fileName string) *Arango {
	arangoOnce.Do(func() {
		loadArango(fileName)
	})

	return arangoConfig
}
