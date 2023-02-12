package service

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	pb "github.com/zulcss/edged/proto"
	"github.com/zulcss/edged/pkg/constants"
	"github.com/zulcss/edged/pkg/config"
)

type server struct {
	pb.UnimplementedEdgeServer
}
	

// Run GPRC worker
func Run(configFile string) {
	log.Infof("Running edged-worker version %s", constants.Version)

	config.ReadConfig(configFile)
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")

	// Starting GPRC worker
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEdgeServer(s, &server{})
	log.Infof("Server listening at %s on %s", host, port)

	if err := s.Serve(lis); err != nil {
                log.Fatalf("Worker failed: %v", err)
        }

}
