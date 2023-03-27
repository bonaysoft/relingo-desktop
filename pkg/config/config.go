package config

import (
	"github.com/bonaysoft/relingo-desktop/pkg/relingo"
	"github.com/bonaysoft/relingo-desktop/pkg/system"
)

type Config struct {
	Version string          `toml:"version"`
	System  *system.Config  `toml:"system"`
	Relingo *relingo.Config `toml:"relingo"`
}

func NewConfig() *Config {
	return &Config{Version: "v1beta1", System: system.NewConfig(), Relingo: relingo.NewConfig()}
}
