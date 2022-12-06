package jgconf

import (
	"github.com/Jinglever/go-config/option"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Loads a yaml file into a struct
func LoadYamlConfig(path string, cfg interface{}, opts ...option.Option) error {
	for _, o := range opts {
		o()
	}
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Failed to read config file: %s", err)
		return err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		logrus.Errorf("unable to decode env config: %v", err)
		return err
	}
	return nil
}
