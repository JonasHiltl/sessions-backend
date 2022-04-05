package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	MongoURL       string `mapstructure:"MONGO_URL"`
	NatsCluster    string `mapstructure:"NATS_CLUSTER"`
	SpacesEndpoint string `mapstructure:"SPACES_ENDPOINT"`
	SpacesToken    string `mapstructure:"SPACES_TOKEN"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./services/profile/internal/config/envs")
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
