package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

type Config struct {
	Port  int
	Test  string
	Redis RedisConfig `yaml:"redis"`
}

func GetConfig(serviceName string, configPath string) *Config {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(strings.ReplaceAll(serviceName, "-", "_"))
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if configPath != "" {
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			logrus.Panicln("Can't load config from file")
		}
	}

	viper.SetDefault("test", "local")

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		logrus.Panicln("Can't unmarshal config from file")
	}

	return &config
}
