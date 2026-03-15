package config

import (
	"fmt"
	"os"
	m "shortify/internal/models"

	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Не удалось загрузить .env:", err)
    }
}

func LoadENV() m.ENV {
	var env m.ENV
	env.DBhost = os.Getenv("POSTGRES_HOST")
	env.DBport = os.Getenv("POSTGRES_PORT")
	env.DBpassword = os.Getenv("POSTGRES_PASSWORD")
	env.DBuser = os.Getenv("POSTGRES_USER")
	env.DBname = os.Getenv("POSTGRES_DB")
	env.DBauth = os.Getenv("POSTGRES_HOST_AUTH_METHOD")
	env.ShortifyPort = os.Getenv("SHORTIFY_PORT")
	return env
}