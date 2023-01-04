package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/config"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/graph"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/middleware"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	log := logger.Sugar()

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("%s", config)

	resolver, err := initResolver(config)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	authHandler := middleware.AuthMiddleware(resolver)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authHandler(srv))

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func initResolver(config config.Config) (*graph.Resolver, error) {
	cc, err := grpc.Dial(config.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &graph.Resolver{AuthClient: proto.NewAuthServiceClient(cc)}, nil
}
