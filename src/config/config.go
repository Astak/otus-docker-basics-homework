package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	fmtDbUrl    string = "postgres://%s:%s@%s:%s/%s"
	fmtDbDevUrl string = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
)

type Config struct {
	Port       int    `mapstructure:"PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(configPath *string) (c Config, err error) {
	viper.SetConfigFile(*configPath)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal().Msgf("unable to decode into struct, %v", err)
	}
	return
}

func (config *Config) GetDbUrl() string {
	return config.getUrl(fmtDbUrl)
}

func (config *Config) GetDbDevUrl() string {
	return config.getUrl(fmtDbDevUrl)
}

func (config *Config) getUrl(format string) string {
	return fmt.Sprintf(format, config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
}
