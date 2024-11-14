package config

import (
	"go-clean/pkg/loader_config"
)

func InitConfig(filename string) (*Config, error) {
	cfg, err := loader_config.Load[Config](filename)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
