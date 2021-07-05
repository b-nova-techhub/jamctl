package commands

import (
	"fmt"
	"github.com/jcelliott/lumber"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	absolutePath string
	relativePath string
	branch       string
	delimiter    string
	overwrite    bool = true

	config   string
	showVers bool

	version string
	commit  string

	JamCtlCmd = &cobra.Command{
		Use:           "jamctl",
		Short:         "jamctl – command-line tool to interact with jamstack",
		Long:          ``,
		SilenceErrors: true,
		SilenceUsage:  true,

		// parse the config if one is provided, or use the defaults
		PersistentPreRunE: readConfig,

		// print version or help, or continue, depending on flag settings
		PreRunE: preFlight,

		// either run demo as a server, or run it as a CLI depending on what flags
		// are provided
		RunE: startCtl,
	}
)

func readConfig(ccmd *cobra.Command, args []string) error {
	// if --config is passed, attempt to parse the config file
	if config != "" {
		filename := filepath.Base(config)
		viper.SetConfigName(filename[:len(filename)-len(filepath.Ext(filename))])
		viper.AddConfigPath(filepath.Dir(config))

		err := viper.ReadInConfig()
		if err != nil {
			return fmt.Errorf("Failed to read config file - %s", err)
		}
	}

	return nil
}

func preFlight(ccmd *cobra.Command, args []string) error {
	// if --version is passed print the version info
	if showVers {
		fmt.Printf("jamctl %s (%s)\n", version, commit)
		return fmt.Errorf("")
	}

	// if --server is not passed, print help
	if !viper.GetBool("server") {
		ccmd.HelpFunc()(ccmd, args)
		return fmt.Errorf("") // no error, just exit
	}

	return nil
}

func startCtl(ccmd *cobra.Command, args []string) error {
	// convert the log level
	logLvl := lumber.LvlInt(viper.GetString("log-level"))

	// configure the logger
	lumber.Prefix("[demo]")
	lumber.Level(logLvl)

	return nil
}

func init() {
	logLevel := "INFO"

	JamCtlCmd.PersistentFlags().String("log-level", logLevel, "Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")

	viper.BindPFlag("log-level", JamCtlCmd.PersistentFlags().Lookup("log-level"))

	JamCtlCmd.Flags().StringVarP(&config, "config", "c", "", "Path to config file (with extension)")
	JamCtlCmd.Flags().BoolVarP(&showVers, "version", "v", false, "Display the current version of this CLI")

	JamCtlCmd.AddCommand(getCmd)
	JamCtlCmd.AddCommand(listCmd)
	JamCtlCmd.AddCommand(updateCmd)
}
