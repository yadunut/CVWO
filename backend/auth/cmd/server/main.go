package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/yadunut/CVWO/backend/auth/internal/database"
	"github.com/yadunut/CVWO/backend/auth/internal/proto"
	"github.com/yadunut/CVWO/backend/auth/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm/logger"
)

type Config struct {
	DatabaseUrl string `split_words:"true"`
	Port        string `default:"8080"`
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

	db, err := database.InitDB(config.DatabaseUrl, GormLogger{log})
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

type GormLogger struct {
	log *zap.SugaredLogger
}

// Error implements logger.Interface
func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.log.Error(msg, "data", data)
}

// Info implements logger.Interface
func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.log.Info(msg, "data", data)
}

// LogMode implements logger.Interface
func (l GormLogger) LogMode(logger.LogLevel) logger.Interface {
	return nil
	// do nothing
}

// Trace implements logger.Interface
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)
	l.log.Infof("elapsed: %v err: %v sql: %v rows affected: %v", float64(elapsed.Nanoseconds()), err, sql, rows)
}

// Warn implements logger.Interface
func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.log.Warn(msg, "data", data)
}
