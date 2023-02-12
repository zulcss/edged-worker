package config

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	server
}

func ReadConfig(config string) {
	log.Info("in ReadConfig")
}
