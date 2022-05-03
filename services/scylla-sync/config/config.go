package config

import "github.com/spf13/viper"

type Config struct {
	SCYLLA_KEYSPACE string   `mapstructure:"SCYLLA_KEYSPACE"`
	SCYLLA_HOSTS    string   `mapstructure:"SCYLLA_HOSTS"`
	NATS_CLUSTER    string   `mapstructure:"NATS_CLUSTER"`
	PROGRESS_TABLE  string   `mapstructure:"PROGRESS_TABLE"`
	SCYLLA_TABLES   []string `mapstructure:"SCYLLA_TABLES"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	config.SCYLLA_TABLES = viper.GetStringSlice("SCYLLA_TABLES")
	err = viper.Unmarshal(&config)
	return
}
