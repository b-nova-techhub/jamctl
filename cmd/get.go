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
	delimiter string

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get content as a html rendered page",
		Long:  ``,
		Run:   get,
	}
)

func get(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		fmt.Printf("%+v\n", gen.Generate(repo.ReadRepoContents(args[0])))
	} else {
		fmt.Fprintln(os.Stderr, "No repository is specified. Please specify a valid git repository url.")
		return
	}
}

func init() {
	getCmd.PersistentFlags().StringVarP(&delimiter, "delimiter", "d", "content-header", "The tag that is being used as the front matter delimiter.")
	viper.BindPFlag("delimiter", getCmd.PersistentFlags().Lookup("delimiter"))
}
