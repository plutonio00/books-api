package config

import (
	"github.com/spf13/viper"
	"strings"
)

type (
	Config struct {
		ServerConfig ServerConfig
	}

	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}
)

func Init(configPath string) (*Config, error) {
	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var conf Config

	if err := unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("server", &cfg.ServerConfig); err != nil {
		return err
	}

	return nil
}
