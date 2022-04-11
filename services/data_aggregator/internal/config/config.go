package config

import "github.com/spf13/viper"

type Config struct {
	PORT                    string `mapstructure:"PORT"`
	AUTH_SERVICE_ADDRESS    string `mapstructure:"AUTH_SERVICE_ADDRESS"`
	PROFILE_SERVICE_ADDRESS string `mapstructure:"PROFILE_SERVICE_ADDRESS"`
	PARTY_SERVICE_ADDRESS   string `mapstructure:"PARTY_SERVICE_ADDRESS"`
	STORY_SERVICE_ADDRESS   string `mapstructure:"STORY_SERVICE_ADDRESS"`
	TOKEN_SECRET            string `mapstructure:"TOKEN_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
