package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	MongoURL       string `mapstructure:"MONGO_URL"`
	NatsCluster    string `mapstructure:"NATS_CLUSTER"`
	TokenSecret    string `mapstructure:"TOKEN_SECRET"`
	GoogleClientID string `mapstructure:"GOOGLE_CLIENTID"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./services/auth/internal/config/envs")
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
