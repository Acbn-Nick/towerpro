package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type configuration struct {
	Toggle string `mapstructure:"Toggle"`
}

func NewConfiguration() *configuration {
	return setDefaults()
}

func setDefaults() *configuration {
	var config configuration

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("Toggle", "`")
	viper.Unmarshal(&config)

	return &config
}

func (c *configuration) loadConfig() error {
	log.Info("loading config from: %s", viper.ConfigFileUsed())

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("failed to find file, using defaults")
			return err
		} else {
			log.Info("failed to access config file (are permissions correct?")
			return err
		}
	}

	if err := viper.UnmarshalExact(c); err != nil {
		log.Info("failed to parse config, using defaults")
		log.Info(err.Error())
		return err
	}

	log.Info("using settings:")
	log.Info("Toggle: " + c.Toggle)

	return nil
}
