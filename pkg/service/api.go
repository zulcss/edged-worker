package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "github.com/zulcss/edged/proto"
)

func (s *server) Ping(ctx context.Context, in *pb.VersionRequest) (*pb.VersionReply, error) {
        log.Debug("in Ping")
        return &pb.VersionReply{Message: "pong"}, nil
}
