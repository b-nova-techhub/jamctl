package main

import (
	"fmt"
	"github.com/b-nova-techhub/jamctl/cmd"
)

func main() {
	err := commands.JamCtlCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
