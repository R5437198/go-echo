package config

import "fmt"

type Server struct {
	Port int
}

type DB struct {
	Driver string
	Url    string
}

type Config struct {
	Server
	DB
}

func New() (*Config, error) {
	env, err := Load()
	if err != nil {
		return nil, err
	}
	port, err := env.ServerPort()
	if err != nil {
		return nil, err
	}

	return &Config{
		Server{
			Port: port,
		},
		DB{
			Driver: "postgres",
			Url:    env.DbUrl(),
		},
	}, nil
}

func (c *Config) Port() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}

func (c *Config) Driver() string {
	return c.DB.Driver
}

func (c *Config) DbUrl() string {
	return c.DB.Url
}
