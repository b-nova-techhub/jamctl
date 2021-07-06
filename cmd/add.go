package cmd

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/pkg/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	targetPath   string
	relativePath string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add git repository containing markdown content files",
		Long:  ``,
		Run:   add,
	}

	createCmd = &cobra.Command{
		Hidden: true,

		Use:   "create",
		Short: "Add git repository containing markdown content files",
		Long:  ``,

		Run: add,
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
	includeAddFlags(addCmd)
	includeAddFlags(createCmd)
}

func includeAddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&relativePath, "relativePath", "/content", "The directory within the git repository which contains the markdown files.")
	cmd.PersistentFlags().StringVar(&targetPath, "targetPath", "/tmp/jamctl", "The absolute path where the git repository is going to cloned to.")
	viper.BindPFlag("relativePath", cmd.PersistentFlags().Lookup("relativePath"))
	viper.BindPFlag("targetPath", cmd.PersistentFlags().Lookup("targetPath"))
}
