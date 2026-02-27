package config

import (
	"os"
	m "shortify/internal/models"	
)

func LoadENV() m.ENV {
	var env m.ENV
	env.DBhost = os.Getenv("POSTGRES_HOST")
	env.DBport = os.Getenv("POSTGRES_PORT")
	env.DBpassword = os.Getenv("POSTGRES_PASSWORD")
	env.DBuser = os.Getenv("POSTGRES_USER")
	env.DBname = os.Getenv("POSTGRES_DB")
	env.DBauth = os.Getenv("POSTGRES_HOST_AUTH_METOD")
	return env
}