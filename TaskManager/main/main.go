package main

import (
	"cli/TaskManager/cmd"
	"fmt"
	"os"
)


func main() {
	
	args := os.Args

	//Starting the CLI app
	err := cmd.HandleActions(args)
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
}
