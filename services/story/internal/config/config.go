package config

import "github.com/spf13/viper"

type Config struct {
	PORT            string `mapstructure:"PORT"`
	SCYLLA_KEYSPACE string `mapstructure:"SCYLLA_KEYSPACE"`
	SCYLLA_HOSTS    string `mapstructure:"SCYLLA_HOSTS"`
	SPACES_ENDPOINT string `mapstructure:"SPACES_ENDPOINT"`
	SPACES_KEY      string `mapstructure:"SPACES_TOKEN"`
	SPACES_SECRET   string `mapstructure:"SPACES_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./services/story/internal/config/envs")
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
