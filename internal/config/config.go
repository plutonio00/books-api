package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		Token    TokenConfig
	}

	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	DatabaseConfig struct {
		MySQL MySQLConfig
	}

	MySQLConfig struct {
		URI string
	}

	TokenConfig struct {
		JWT                    JWTConfig
		VerificationCodeLength int `mapstructure:"verificationCodeLength"`
	}

	JWTConfig struct {
		SigningKey      string
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

	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("token", &cfg.Token); err != nil {
		return err
	}

	return nil
}

func setFromEnvFile(conf *Config) {
	conf.Database.MySQL.URI = os.Getenv("MYSQL_URI")
	conf.Token.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

func parseEnvFile() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
