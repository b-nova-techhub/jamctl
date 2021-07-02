package commands

import (
	"b-nova-techub/cobra-demo/pkg/gen"
	"fmt"
	"github.com/spf13/cobra"
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
	fmt.Print(gen.GeneratedPages)
}
