package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	c "shortify/internal/config"
	db "shortify/internal/db"
	h "shortify/internal/handlers"
	l "shortify/internal/logger"
	m "shortify/internal/models"
)

func App() error {	

    l.Init()
    defer l.Sync() // гарантируем запись всех буферизованных сообщений

    log := l.Logger() // *zap.Logger

	ctx := context.Background()
	var ENV m.ENV = c.LoadENV()

	database := db.New(ctx, ENV)
	err := database.InitPool()
	if err != nil {
		log.Error("Ошибка инициализации пула соединений", zap.Error(err))
		return err
	}
	
	r := gin.Default()
	r.GET("/", h.GetFullAddres())
	r.POST("/encode", h.GetEncodedAddress())

	r.Run(":"+ENV.DBport)
	return nil
}