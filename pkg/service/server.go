package service

import (
//	"net"

//	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	"github.com/zulcss/edged/pkg/constants"
	"github.com/zulcss/edged/pkg/config"
)

func Run(configFile string) {
	log.Infof("Running edged-worker version %s", constants.Version)

	config.ReadConfig(configFile)
}
