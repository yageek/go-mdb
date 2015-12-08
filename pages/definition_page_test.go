package pages

import (
	"testing"

	"github.com/yageek/go-mdb/filepage"
)

const (
	jet4DatabasePath = "../test_databases/EPSG_v8_6.mdb"
	jet3DatabasePath = "../test_databases/Books_be.mdb"
)

func TestJet4Database(t *testing.T) {

	scanner, err := filepage.NewScanner(jet4DatabasePath, Jet4PageSize)

	if err != nil {
		t.Error(err.Error())
		return
	}

	scanner.ReadPage()

	page := scanner.Page()
	if err := scanner.Error(); err != nil {
		t.Error(err.Error())
		return
	}

	defPage, err := NewDefinitionPage(page, Jet4)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if defPage == nil {
		t.Error("Defpage should not be nil")
	}

}
