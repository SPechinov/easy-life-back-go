package config

import (
	"fmt"
	"go-clean/pkg/loader_config"
)

func InitConfig(filename string) (*Config, error) {
	fmt.Println("Config initializing...")
	cfg, err := loader_config.Load[Config](filename)
	if err != nil {
		fmt.Printf("Config not initialized: %s\n", err)
		return nil, err
	}
	fmt.Println("Config initialized")
	return cfg, nil
}
