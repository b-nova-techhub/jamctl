package main

import (
	commands "b-nova-techub/cobra-demo/commands"
	"fmt"
)

func main() {
	err := commands.DemoCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
