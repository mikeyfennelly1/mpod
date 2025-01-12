package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Defines the *structure* of a valid command in application logic
// NOTE that the flags - such as name are arbitrary, and validity is only dependent on position thus they are omitted from this struct
type command struct {
	use          string
	argsAndFlags []string // any arbitrary args - this struct doesn't verify validity of args
	short        string
	long         string
	run          func()
}

func main() {
	if len(os.Args) == 1 {
		errors.New("Please provide some arguments\n")
		return
	}
	var containerName string
	flag.StringVar(&containerName, "name", "undefined", "Name of the container you want to start.")

	flag.Parse()
	fmt.Println(containerName)
}
