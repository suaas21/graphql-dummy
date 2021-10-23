package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

// Application holds application configurations
type Application struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	GracefulTimeout int    `yaml:"graceful_timeout"`
}

var appOnce = sync.Once{}
var appConfig *Application

// loadApp loads config from path
func loadApp(fileName string) error {
	readConfig(fileName)
	viper.AutomaticEnv()

	appConfig = &Application{
		Host:            viper.GetString("app.host"),
		GracefulTimeout: viper.GetInt("app.graceful_timeout"),
		Port:            viper.GetInt("app.port"),
	}

	log.Println("app config ", appConfig)

	return nil
}

// GetApp returns application config
func GetApp(fileName string) *Application {
	appOnce.Do(func() {
		loadApp(fileName)
	})

	return appConfig
}
