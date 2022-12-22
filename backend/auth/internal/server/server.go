package server

import (
	context "context"
	"net/mail"

	"github.com/yadunut/CVWO/backend/auth/internal/database"
	"github.com/yadunut/CVWO/backend/auth/internal/proto"
	"go.uber.org/zap"
)

type Server struct {
	DB  database.DB
	log *zap.SugaredLogger
	proto.UnimplementedAuthServiceServer
}

func NewServer(DB database.DB, log *zap.SugaredLogger) *Server {
	return &Server{
		DB:                             DB,
		log:                            log,
		UnimplementedAuthServiceServer: proto.UnimplementedAuthServiceServer{},
	}
}

// Login implements proto.AuthServiceServer
func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	return &proto.LoginResponse{Status: proto.ResponseStatus_SUCCESS, Error: req.String(), Token: ""}, nil
}

// Register implements proto.AuthServiceServer
func (*Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return &proto.RegisterResponse{Status: proto.ResponseStatus_FAILURE, Error: "invalid email address", Token: ""}, nil
	}
	// a validation library should be used here

	return &proto.RegisterResponse{Status: proto.ResponseStatus_SUCCESS, Error: req.String(), Token: ""}, nil
}
