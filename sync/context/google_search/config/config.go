package config

import "fmt"

const (
	_host = "localhost"
	_port = 8080
)

type HttpConfig struct {
	Host string
	Port int32
}

type Config struct {
	Http HttpConfig
}

func New() *Config {
	return &Config{
		Http: HttpConfig{
			Host: _host,
			Port: _port,
		},
	}
}

func (h *HttpConfig) GetUrl() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
