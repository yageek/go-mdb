package main

import (
	"fmt"
	"os"

	"github.com/yageek/go-mdb/catalog"
)

func main() {

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	} else {

		catalog, err := catalog.NewCatalog(os.Args[1])
		if err != nil {
			fmt.Printf("Could not open file as MDB file:%v\n", err.Error())
			os.Exit(1)
		}

		defer catalog.Close()

		err = catalog.Read()
		if err != nil {
			fmt.Printf("Could not read catalog information:%v\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%v\n", catalog.Version())
	}

}

func printUsage() {
	fmt.Println("Usage:mdb-ver <file>")
}
