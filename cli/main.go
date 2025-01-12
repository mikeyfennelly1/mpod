package main

import (
	"fmt"
	"os"
)

func main() {
	firstSubCommand := os.Args[1]
	handleFirstSubCommand(firstSubCommand)
}

func handleFirstSubCommand(command string) {
	switch command {
	case "ps":
		fmt.Printf("ps ran\n")
		break
	case "start":
		fmt.Printf("start ran\n")
		break
	default:
		fmt.Printf("mpod: '%s' is not an mpod command.\n See 'mpod --help'\n", command)
		os.Exit(1)
	}
}
