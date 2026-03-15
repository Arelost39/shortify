package db

import (
	"context"
	"fmt"
	"time"

	m "shortify/internal/models"

	"go.uber.org/zap"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	ctx 	context.Context
	env 	m.ENV
	log		*zap.Logger
	pool	*pgxpool.Pool
}

func New(context context.Context, ENV m.ENV, log *zap.Logger) Database {
	var DBconn Database

	DBconn.ctx = context
	DBconn.env = ENV
	DBconn.log = log
	
	return DBconn
}

func (db *Database) InitPool() error {

	// db.log.Info(db.env.DBport)


	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db.env.DBuser,
		db.env.DBpassword,
		db.env.DBhost,
		db.env.DBport,
		db.env.DBname,
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		db.log.Error("ParseConfig error", zap.Error(err))
		return err
	}

	// конфигурация пула подкючений
	// 
	cfg.MaxConns = 4								// максимум подключений
	cfg.MinConns = 1								// минимум всегда открытых
	cfg.MaxConnLifetime = 1 * time.Hour			// переподключаемся раз в час
	cfg.MaxConnIdleTime = 10 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	cfg.ConnConfig.ConnectTimeout = 10 * time.Second

	db.pool, err = pgxpool.NewWithConfig(db.ctx, cfg)
	if err != nil {
		db.log.Error("не удалось создать пул:", zap.Error(err))
		return err
	} else {
		db.log.Info("Пул создан", 
		zap.Int32("Max conn", cfg.MaxConns), 
		zap.Int32("Min conn", cfg.MinConns))
	}

	// пингуем
	if err := db.pool.Ping(db.ctx); err != nil {
		db.log.Error("ping не прошёл:", zap.Error(err))
		return err
	} else {
		db.log.Info("Успешный пинг бд")
	}
	return nil
}

// graceful shutdown
func (db *Database) Close() {
	if db.pool != nil {
		db.pool.Close()
		db.log.Info("Пул подключений закрыт")
	}
}

// Stat — для мониторинга
func (db *Database) Stat() pgxpool.Stat {
	if db.pool == nil { return pgxpool.Stat{} }
	return *db.pool.Stat()
}