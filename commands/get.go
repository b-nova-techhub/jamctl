package commands

import (
	"b-nova-techub/cobra-demo/pkg/gen"
	"fmt"
	"github.com/spf13/cobra"
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
	includeGetFlags(getCmd)
}

func includeGetFlags(cmd *cobra.Command) {
}

func get(ccmd *cobra.Command, args []string) {
	pages := gen.GeneratedPages
	page := getPageById(args[1], pages)
	fmt.Print(page)
}

func getPageById(id string, pages []gen.StaticPage) *gen.StaticPage {
	var page *gen.StaticPage
	for _, p := range pages {
		if p.Permalink == id {
			page = &p
		}
	}
	return page
}
