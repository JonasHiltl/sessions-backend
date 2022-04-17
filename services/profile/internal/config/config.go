package config

import "github.com/spf13/viper"

type Config struct {
	PORT            string `mapstructure:"PORT"`
	MONGO_URL       string `mapstructure:"MONGO_URL"`
	NATS_CLUSTER    string `mapstructure:"NATS_CLUSTER"`
	SPACES_ENDPOINT string `mapstructure:"SPACES_ENDPOINT"`
	SPACES_TOKEN    string `mapstructure:"SPACES_TOKEN"`
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
