package config

import (
	"os"
	"strconv"
)

type Env struct {
	ServerPORT string
	DbURL      string
}

func Load() (*Env, error) {
	return &Env{
		ServerPORT: os.Getenv("REST_CONTAINER_PORT"),
		DbURL:      os.Getenv("DB_URL"),
	}, nil
}

func (e *Env) ServerPort() (int, error) {
	port, err := strconv.Atoi(e.ServerPORT)
	return port, err
}

func (e *Env) DbUrl() string {
	return e.DbURL
}
