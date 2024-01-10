package main

import (
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		println("Please provide a argv")
		os.Exit(123)
	case 2:
		// flag parse

		println("hello", os.Args[1])
	default:
		os.Exit(309)
		println("Too many arguments")
	}
}
