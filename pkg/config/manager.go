package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
)

type Manager struct {
	cfg *Config
}

func NewCM() *Manager {
	return &Manager{cfg: NewConfig()}
}

func (c *Manager) Load() (*Config, error) {
	cfgPath, err := c.buildConfigPath()
	if err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(cfgPath)
	if err := v.ReadInConfig(); err != nil && os.IsNotExist(err) {
		return c.cfg, c.Save()
	} else if err != nil {
		return nil, fmt.Errorf("Error reading config: %s\n", err)
	}

	return c.cfg, v.Unmarshal(c.cfg)
}

func (c *Manager) Save() error {
	cfgPath, err := c.buildConfigPath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", cfgPath, err)
	}

	return toml.NewEncoder(f).Encode(c.cfg)
}

func (c *Manager) GetConfigDir() string {
	cfgPath, _ := c.buildConfigPath()
	return filepath.Dir(cfgPath)
}

func (c *Manager) buildConfigPath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	cfgPath := filepath.Join(cfgDir, "relingo-desktop")
	if _, err := os.Stat(cfgPath); err != nil {
		if err := os.Mkdir(cfgPath, 0755); err != nil {
			return "", err
		}
	}

	return filepath.Join(cfgPath, "config.toml"), nil
}
