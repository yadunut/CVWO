package server

import (
	"github.com/yadunut/CVWO/backend/database"
	"github.com/yadunut/CVWO/backend/proto"
	"github.com/yadunut/CVWO/backend/thread/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	proto.UnimplementedThreadServiceServer
	DB     database.DB
	log    *zap.SugaredLogger
	config config.Config
}

func NewServer(DB database.DB, log *zap.SugaredLogger, c config.Config) *Server {
	return &Server{
		DB:                               DB,
		log:                              log,
		UnimplementedThreadServiceServer: proto.UnimplementedThreadServiceServer{},
		config:                           c,
	}
}
