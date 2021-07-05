package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all content as html rendered pages",
		Long:  ``,

		Run: list,
	}
)

func list(ccmd *cobra.Command, args []string) {
	repos, err := ioutil.ReadDir(viper.GetString("absolutePath"))
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range repos {
		fmt.Println(r.Name())
	}
}
