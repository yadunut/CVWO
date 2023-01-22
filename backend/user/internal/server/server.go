package server

import (
	"context"

	"github.com/yadunut/CVWO/backend/database"
	proto "github.com/yadunut/CVWO/backend/proto/user"
	"github.com/yadunut/CVWO/backend/user/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	DB     database.DB
	log    *zap.SugaredLogger
	config config.Config
}

func NewServer(DB database.DB, log *zap.SugaredLogger, c config.Config) *Server {
	return &Server{
		DB:                             DB,
		log:                            log,
		UnimplementedUserServiceServer: proto.UnimplementedUserServiceServer{},
		config:                         c,
	}
}

// GetUser implements proto.UserServiceServer
func (*Server) GetUser(context.Context, *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	panic("unimplemented")
}
