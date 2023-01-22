package main

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/yadunut/CVWO/backend/database"
	proto "github.com/yadunut/CVWO/backend/proto/user"
	"github.com/yadunut/CVWO/backend/user/internal/config"
	"github.com/yadunut/CVWO/backend/user/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	log := logger.Sugar()

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("%s", config)

	db, err := database.Init(config.DatabaseUrl, log)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(db, log, config)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_zap.StreamServerInterceptor(logger))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_zap.UnaryServerInterceptor(logger))),
	)
	proto.RegisterUserServiceServer(grpcServer, server)
	log.Infof("Serving grpcServer")
	grpcServer.Serve(lis)

}
