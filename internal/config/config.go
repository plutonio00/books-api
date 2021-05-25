package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type (
	Config struct {
		ServerConfig   ServerConfig
		DatabaseConfig DatabaseConfig
	}

	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	DatabaseConfig struct {
		MySQLConfig MySQLConfig
	}

	MySQLConfig struct {
		URI string
	}
)

func Init(configPath string) (*Config, error) {

	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	if err := parseEnvFile(); err != nil {
		return nil, err
	}

	var conf Config

	if err := unmarshal(&conf); err != nil {
		return nil, err
	}

	setFromEnvFile(&conf)

	return &conf, nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName("env")
	viper.MergeInConfig()

	return nil
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("server", &cfg.ServerConfig); err != nil {
		return err
	}

	return nil
}

func setFromEnvFile(conf *Config) {
	conf.DatabaseConfig.MySQLConfig.URI = os.Getenv("MYSQL_URI")
}

func parseEnvFile() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
