package database

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func Init(url string, zLog *zap.SugaredLogger) (DB, error) {
	var log logger.Interface

	if zLog == nil {
		log = logger.Default
	} else {
		log = GormLogger{zLog}
	}

	gormDB, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		return DB{}, err
	}

	return DB{gormDB}, nil
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
