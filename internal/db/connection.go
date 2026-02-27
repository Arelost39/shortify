package db

import (
	"context"

	"strconv"
	"time"

	m "shortify/internal/models"

	"go.uber.org/zap"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool	*pgxpool.Pool
	ctx 	context.Context
	env 	m.ENV
	log		*zap.Logger
}

func New(context context.Context, ENV m.ENV) Database {
	var DBconn Database

	DBconn.ctx = context
	DBconn.env = ENV
	
	return DBconn
}

func (db *Database) InitPool() error {

	port, err := strconv.ParseUint(db.env.DBport, 10, 32)
	if err != nil {
		db.log.Error("Ошибка парсинга порта", zap.Error(err))
		return err
	}

    cfg := &pgxpool.Config{}
    cfg.ConnConfig.Host = db.env.DBhost

    cfg.ConnConfig.Port = uint16(port)
    cfg.ConnConfig.Database = db.env.DBname
    cfg.ConnConfig.User = db.env.DBuser
    cfg.ConnConfig.Password = db.env.DBpassword
    cfg.ConnConfig.Config.TLSConfig = nil

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