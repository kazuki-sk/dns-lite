package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ListenPort string `toml:ListenPort`
	DB         struct {
		Type string `toml:Type`
		DNS  string `toml:DNS`
	} `toml:db`
}

func ReadConf(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	var c Config
	toml.Unmarshal(b, &c)
	return &c, nil
}
