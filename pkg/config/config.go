package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func NewConfig() (*Config, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	cfgPath := filepath.Join(cfgDir, "relingo-desktop")
	if _, err := os.Stat(cfgPath); err != nil {
		if err := os.Mkdir(cfgPath, 0755); err != nil {
			return nil, err
		}
	}

	v := viper.New()
	v.SetConfigFile(filepath.Join(cfgPath, "config.toml"))
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("config not exist:", err)
	}

	return &Config{
		v: v,
	}, nil
}

func (c *Config) Load() error {
	return c.v.ReadInConfig()
}

func (c *Config) GetToken() string {
	return c.v.GetString("relingo.token")
}

func (c *Config) SaveToken(token string) {
	c.v.Set("relingo.token", token)
	_ = c.v.WriteConfig()
}

func (c *Config) GetDBPath() string {
	return filepath.Join(filepath.Dir(c.v.ConfigFileUsed()), "relingo.db")
}
