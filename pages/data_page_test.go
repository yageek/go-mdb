package pages

import (
	"fmt"
	"testing"

	"github.com/yageek/go-mdb/filepage"
)

func TestDataPage(t *testing.T) {
	scanner, err := filepage.NewScanner(jet4DatabasePath, Jet4PageSize)

	if err != nil {
		t.Error(err.Error())
		return
	}

	scanner.ReadPage()

	scanner.ReadPage()

	page := scanner.Page()

	if err := scanner.Error(); err != nil {
		t.Error(err.Error())
		return
	}

	pageData, err := NewDataPageHeader(page, Jet4)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if pageData == nil {
		t.Errorf("Expected non nil page data")
	}
	fmt.Println(pageData)

}
