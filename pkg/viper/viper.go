package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

func refreshConfig() {
	initViper()
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/cobra-demo/")
	viper.AddConfigPath("$HOME/.cobra-demo")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
