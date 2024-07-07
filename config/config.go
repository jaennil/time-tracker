package config

import "github.com/spf13/viper"

type Config struct {
	Port   int    `mapstructure:"TIMETRACKER_API_PORT"`
	PG_DSN string `mapstructure:"TIMETRACKER_PG_DSN"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
