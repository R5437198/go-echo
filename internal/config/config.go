package config

import "fmt"

type Server struct {
	Port int `default:"80"`
}

type Config struct {
	Server
}

func New() *Config {
	return &Config{
		Server{Port: 80},
	}
}

func (c *Config) Port() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}
