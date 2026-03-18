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

	database := db.New(ctx, ENV, log)
	err := database.InitPool()
	if err != nil {
		log.Error("Ошибка инициализации пула соединений", zap.Error(err))
		return err
	}
	
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		//c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/", h.GetFullAddres())
	r.GET("/last_id", h.GetLastID(&database))
	r.POST("/encode", h.GetEncodedAddress())

	r.Run(":"+ENV.ShortifyPort)
	return nil
}