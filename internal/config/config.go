package config

import (
	"github.com/spf13/viper"
	"strings"
	"fmt"
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

    fmt.Println(viper.GetString("uri"))
	conf.DatabaseConfig.MySQLConfig.URI = viper.GetString("uri")
}

func parseEnvFile() error {

	if err := parseMySQLEnvVariables(); err != nil {
		return err
	}

	return nil
}

func parseMySQLEnvVariables() error {
	viper.SetEnvPrefix("mysql")
	if err := viper.BindEnv("uri"); err != nil {
		return err
	}

	return nil
}
