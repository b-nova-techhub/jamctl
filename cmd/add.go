package cmd

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/pkg/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		Run:   add,
	}
)

func add(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		repo.GetGitRepository(args[0], false)
		fmt.Printf("Repo added.\n")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.WriteConfigAs(home + "/jamctl.yaml")
	} else {
		fmt.Fprintln(os.Stderr, "No repository is specified. Please specify a valid git repository url.")
		return
	}
}

func init() {
	addCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $HOME/.jamctl.yaml)")
	addCmd.Flags().String("absolutePath", "ABSOLUTE PATH", "The absolute path where the git repository is going to cloned to.")
	addCmd.Flags().String("relativePath", "RELATIVE PATH", "The directory within the git repository which contains the markdown files.")
	addCmd.Flags().StringP("branch", "b", "BRANCH", "The branch of the git repository that is to be used.")
	addCmd.Flags().StringP("delimiter", "d", "", "The tag that is being used as the front matter delimiter.")
	addCmd.Flags().BoolP("overwrite", "o", true, "Whether or not to overwrite an existing git repository.")
	viper.BindPFlag("absolutePath", addCmd.Flags().Lookup("absolutePath"))
	viper.BindPFlag("relativePath", addCmd.Flags().Lookup("relativePath"))
	viper.BindPFlag("branch", addCmd.Flags().Lookup("branch"))
	viper.BindPFlag("delimiter", addCmd.Flags().Lookup("delimiter"))
	viper.BindPFlag("overwrite", addCmd.Flags().Lookup("overwrite"))
	viper.SetDefault("absolutePath", "/tmp/jamctl")
	viper.SetDefault("relativePath", "/content")
	viper.SetDefault("branch", "main")
}
