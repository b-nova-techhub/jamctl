package cmd

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/pkg/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update git repository containing markdown content files",
		Long:  ``,
		Run:   update,
	}
)

func update(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		repo.GetGitRepository(args[0], true)
		fmt.Printf("Repo updated.\n")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.WriteConfigAs(home + "/jamctl.yaml")
	} else {
		fmt.Fprintln(os.Stderr, "No repository is specified. Please specify a valid git repository url.")
		return
	}
}

func init() {
	updateCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $HOME/.jamctl.yaml)")
	updateCmd.Flags().String("absolutePath", "ABSOLUTE PATH", "The absolute path where the git repository is going to cloned to.")
	updateCmd.Flags().String("relativePath", "RELATIVE PATH", "The directory within the git repository which contains the markdown files.")
	updateCmd.Flags().StringP("branch", "b", "BRANCH", "The branch of the git repository that is to be used.")
	updateCmd.Flags().StringP("delimiter", "d", "", "The tag that is being used as the front matter delimiter.")
	updateCmd.Flags().BoolP("overwrite", "o", true, "Whether or not to overwrite an existing git repository.")
	viper.BindPFlag("absolutePath", updateCmd.Flags().Lookup("absolutePath"))
	viper.BindPFlag("relativePath", updateCmd.Flags().Lookup("relativePath"))
	viper.BindPFlag("branch", updateCmd.Flags().Lookup("branch"))
	viper.BindPFlag("delimiter", updateCmd.Flags().Lookup("delimiter"))
	viper.BindPFlag("overwrite", updateCmd.Flags().Lookup("overwrite"))
	viper.SetDefault("absolutePath", "/tmp/jamctl")
	viper.SetDefault("relativePath", "/content")
	viper.SetDefault("branch", "main")
}
