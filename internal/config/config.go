package config

import (
	"fmt"

	"github.com/irede-interview/cinema-api/internal/envs"
)

type Config struct {
	DSN  string
	Port string
}

func New() Config {
	envs.Load()
	port := envs.Get(envs.SERVER_PORT)

	dbUsername := envs.Get(envs.DB_USERNAME)
	dbPassword := envs.Get(envs.DB_PASSWORD)
	dbHost := envs.Get(envs.DB_HOST)
	dbPort := envs.Get(envs.DB_PORT)
	dbName := envs.Get(envs.DB_NAME)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)

	return Config{
		DSN:  dsn,
		Port: port,
	}
}
