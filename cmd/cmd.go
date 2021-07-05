package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string

	jamctlCmd = &cobra.Command{
		Use:           "jamctl",
		Short:         "jamctl – command-line tool to interact with jamstack",
		Long:          ``,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return jamctlCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	jamctlCmd.AddCommand(addCmd)
	jamctlCmd.AddCommand(getCmd)
	jamctlCmd.AddCommand(listCmd)
	jamctlCmd.AddCommand(updateCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".jamctl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for jamctl: ", viper.ConfigFileUsed())
	}
}
