package main

import (
	"fmt"
	"github.com/yageek/go-mdb/version"
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
	versionBuf := make([]byte, 1)

	_, err = file.ReadAt(versionBuf, version.VersionOffset)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	v, err := version.NewJetVersion(versionBuf[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(v)
	defer file.Close()
}
