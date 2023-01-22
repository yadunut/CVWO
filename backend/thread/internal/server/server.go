package server

import (
	context "context"

	"github.com/yadunut/CVWO/backend/database"
	proto "github.com/yadunut/CVWO/backend/proto/thread"
	"github.com/yadunut/CVWO/backend/thread/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	proto.UnimplementedThreadServiceServer
	DB     database.DB
	log    *zap.SugaredLogger
	config config.Config
}

// CreateThread implements thread.ThreadServiceServer
func (*Server) CreateThread(context.Context, *proto.Thread) (*proto.ThreadResponse, error) {
	panic("unimplemented")
}

// GetThread implements thread.ThreadServiceServer
func (*Server) GetThread(context.Context, *proto.GetThreadRequest) (*proto.ThreadResponse, error) {
	panic("unimplemented")
}

// GetThreads implements thread.ThreadServiceServer
func (*Server) GetThreads(context.Context, *proto.GetThreadsRequest) (*proto.ThreadsResponse, error) {
	panic("unimplemented")
}

func NewServer(DB database.DB, log *zap.SugaredLogger, c config.Config) *Server {
	return &Server{
		DB:                               DB,
		log:                              log,
		UnimplementedThreadServiceServer: proto.UnimplementedThreadServiceServer{},
		config:                           c,
	}
}
