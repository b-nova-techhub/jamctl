package commands

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/pkg/gen"
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
