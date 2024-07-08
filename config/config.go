package config

import "github.com/spf13/viper"

type Config struct {
	Port       int    `mapstructure:"TIMETRACKER_API_PORT"`
	DBHost     string `mapstructure:"TIMETRACKER_DB_HOST"`
	DBPort     int    `mapstructure:"TIMETRACKER_DB_PORT"`
	DBUser     string `mapstructure:"TIMETRACKER_DB_USER"`
	DBName     string `mapstructure:"TIMETRACKER_DB_NAME"`
	DBPassword string `mapstructure:"TIMETRACKER_DB_PASS"`
	DBSSLMode  string `mapstructure:"TIMETRACKER_DB_SSLMODE"`
	UserApiUrl string `mapstructure:"TIMETRACKER_EXTERNAL_API"`
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
