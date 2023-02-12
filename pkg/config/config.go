package config

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)


func ReadConfig(CconfigFile string) {
	log.Debug("in ReadConfig")

	viper.SetConfigName("config")
	viper.AddConfigPath("etc/edged")
	err := viper.ReadInConfig() 
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v\n", err)
	}
}
