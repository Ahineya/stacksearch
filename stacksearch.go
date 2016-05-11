package main

import (
	"fmt"
	"github.com/Ahineya/stacksearch/api"
	"os"
)

const (
	gotest_version = "0.0.1"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Usage: gotest [-v | query]\n")
		os.Exit(0)
	}

	if args[0] == "-v" {
		fmt.Printf("Gotest version: %s\n", gotest_version)
		os.Exit(0)
	}

	answer, err := api.GetAnswer(args[0])
	processErr(err)

	fmt.Printf("%s", answer)
}

func processErr(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}