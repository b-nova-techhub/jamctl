package jamctl

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/cmd"
)

func main() {
	err := commands.DemoCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
