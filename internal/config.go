package client

import (
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
	viper.AddConfigPath("./cmd")
	viper.SetDefault("Toggle", "`")
	viper.Unmarshal(&config)

	return &config
}

func (p *Performer) loadConfig() {
	p.log.Info("loading config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			p.log.Warn("failed to find file, using defaults")
		} else {
			p.log.Fatal("failed to access config file (are permissions correct?")
		}
	}

	if err := viper.UnmarshalExact(p.config); err != nil {
		p.log.Warn("failed to parse config, using defaults")
		p.log.Warn(err.Error())
	}

	p.log.Info("using settings:")
	p.log.Info("Toggle: " + p.config.Toggle)
}
