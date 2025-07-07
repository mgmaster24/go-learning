package config

import "fmt"

type AppConfig struct {
	port int
}

func (appConfing *AppConfig) GetPortString() string {
	return fmt.Sprintf(":%d", appConfing.port)
}

func NewAppConfig(port int) *AppConfig {
	return &AppConfig{
		port,
	}
}
