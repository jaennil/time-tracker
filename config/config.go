package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"TIMETRACKER_DB_HOST"`
	DBPort     int    `mapstructure:"TIMETRACKER_DB_PORT"`
	DBUser     string `mapstructure:"TIMETRACKER_DB_USER"`
	DBName     string `mapstructure:"TIMETRACKER_DB_NAME"`
	DBPassword string `mapstructure:"TIMETRACKER_DB_PASS"`
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
