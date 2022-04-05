package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	ScyllaKeyspace string `mapstructure:"SCYLLA_KEYSPACE"`
	ScyllaHosts    string `mapstructure:"SCYLLA_HOSTS"`
	NatsCluster    string `mapstructure:"NATS_CLUSTER"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./services/party/internal/config/envs")
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
