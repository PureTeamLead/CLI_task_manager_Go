package main

import (
	"cli/TaskManager/cmd"
	"fmt"
	"os"
)


func main() {
	
	args := os.Args

	err := cmd.HandleActions(args)
	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
}
