package config

import "github.com/spf13/viper"

var Instance Config

type Config struct {
	Addr string
}

func ReadConfig() error {
	configer := viper.New()
	configer.AddConfigPath("./")
	configer.AddConfigPath("../")
	configer.SetConfigType("yaml")
	configer.SetConfigName("config")
	err := configer.ReadInConfig()
	if err != nil {
		return err
	}
	configer.SetDefault("addr", ":4300")

	Instance = Config{
		Addr: configer.GetString("addr"),
	}
	return nil
}
