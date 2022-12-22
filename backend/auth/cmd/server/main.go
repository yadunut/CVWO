package main

import (
	"fmt"
	"net"

	"github.com/yadunut/CVWO/backend/auth/internal/database"
	"github.com/yadunut/CVWO/backend/auth/internal/proto"
	"github.com/yadunut/CVWO/backend/auth/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	log := logger.Sugar()

	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("%s", config)

	db, err := database.InitDB(config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	server := server.NewServer(db, log)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", config.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, server)
	grpcServer.Serve(lis)

}
