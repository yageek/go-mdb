package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage: mdb-ver <file>")
}
func main() {

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(-1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Could not open file:", err)
		os.Exit(-1)
	}

	defer file.Close()

}
