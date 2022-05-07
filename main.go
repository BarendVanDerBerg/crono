package main

import (
	"fmt"
	"os"

	"github.com/barendvanderberg/crono/pkg/descriptor"
	"github.com/barendvanderberg/crono/pkg/parser"
)

func printHelp() {
	fmt.Println("")
	fmt.Println("crono takes a cron expression and command, and prints the output of the cron schedule and command")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\tcrono \"*/15 0 1,15 * 1-5 /usr/bin/find -r passwords.txt\"")
}

func main() {
	descriptors, err := parser.ParseArguments(os.Args)
	if err != nil {
		fmt.Printf("error: %v", err)
		printHelp()
		os.Exit(1)
	}

	for _, desc := range descriptors {
		fmt.Printf("%s\n", descriptor.Describe(desc))
	}
}
