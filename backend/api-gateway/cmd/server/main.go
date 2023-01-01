package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/graph"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func initResolver(config Config) (*graph.Resolver, error) {
	cc, err := grpc.Dial(config.AuthServiceUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &graph.Resolver{AuthClient: proto.NewAuthServiceClient(cc)}, nil
}

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	log := logger.Sugar()

	config, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("%s", config)

	LoadConfig()

	resolver, err := initResolver(config)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
