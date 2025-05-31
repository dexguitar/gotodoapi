package config

import "github.com/spf13/viper"

type Config struct {
	Host      string `mapstructure:"HOST"`
	Port      string `mapstructure:"PORT"`
	DBPrimary string `mapstructure:"DBPRIMARY"`
}

func MustLoad() (config *Config, err error) {
	v := viper.New()

	v.SetConfigFile(".env")

	v.AutomaticEnv()

	v.SetDefault("HOST", "0.0.0.0")
	v.SetDefault("PORT", ":8080")
	v.SetDefault("DBPRIMARY", "postgres://postgres:qwerty@localhost:5432/gotodo?sslmode=disable")

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
