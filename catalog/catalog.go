package catalog

import (
	"os"

	"github.com/yageek/go-mdb/filepage"
)

type Catalog struct {
	jetVersion byte
	scanner    *filepage.Scanner
	Entries    []Entry
}

func NewCatalog(filename string) (*Catalog, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

}
