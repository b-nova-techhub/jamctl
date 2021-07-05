package cmd

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/pkg/gen"
	"github.com/b-nova-techhub/jamctl/pkg/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get content as a html rendered page",
		Long:  ``,

		Run: get,
	}
)

func init() {
	getCmd.Flags().String("absolutePath", "ABSOLUTE PATH", "The absolute path where the git repository is going to cloned to.")
	getCmd.Flags().String("relativePath", "RELATIVE PATH", "The directory within the git repository which contains the markdown files.")
	getCmd.Flags().StringP("delimiter", "d", "", "The tag that is being used as the front matter delimiter.")
	viper.SetDefault("absolutePath", "/tmp/jamctl")
	viper.SetDefault("relativePath", "/content")
	viper.SetDefault("delimiter", "b-nova-content-header")
}

func get(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		fmt.Print(gen.Generate(repo.ReadRepoContents(args[0])))
	} else {
		fmt.Fprintln(os.Stderr, "No repository is specified. Please specify a valid git repository url.")
		return
	}
}
