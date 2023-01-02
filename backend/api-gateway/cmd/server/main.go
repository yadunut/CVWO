package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/graph"
	"github.com/yadunut/CVWO/backend/api-gateway/internal/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Port           string `default:"8080"`
	AuthServiceUrl string `split_words:"true"`
}


func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()
	log := logger.Sugar()

	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("%s", config)

	loadConfig()

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

func initResolver(config Config) (*graph.Resolver, error) {
	cc, err := grpc.Dial(config.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &graph.Resolver{AuthClient: proto.NewAuthServiceClient(cc)}, nil
}


func loadConfig() (c Config, err error) {
	// only call load if .env exists
	if _, err = os.Stat(".env"); !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			return
		}
	}

	err = envconfig.Process("CVWO", &c)
	if err != nil {
		return
	}

	cRef := reflect.ValueOf(&c).Elem()
	for i := 0; i < cRef.NumField(); i++ {
		field := cRef.Field(i)
		if field.IsZero() {
			err = fmt.Errorf("%s cannot be empty", cRef.Type().Field(i).Name)
			return
		}
	}
	return
}
