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
	updateCmd.PersistentFlags().StringVar(&relativePath, "relativePath", "/content", "The directory within the git repository which contains the markdown files.")
	updateCmd.PersistentFlags().StringVar(&targetPath, "targetPath", "/tmp/jamctl", "The absolute path where the git repository is going to cloned to.")
	viper.BindPFlag("relativePath", updateCmd.PersistentFlags().Lookup("relativePath"))
	viper.BindPFlag("targetPath", updateCmd.PersistentFlags().Lookup("targetPath"))
}
