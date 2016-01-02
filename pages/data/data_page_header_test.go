package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/yageek/go-mdb/filepage"
	"github.com/yageek/go-mdb/test_databases"
	"github.com/yageek/go-mdb/version"
)

func TestDataPage(t *testing.T) {

	file, err := os.Open(testdatabases.JET4DatabasePath())
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	scanner, err := filepage.NewScanner(file, 4096)

	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	defer scanner.Close()

	scanner.ReadPage()

	scanner.ReadPage()

	page := scanner.Page()

	if err := scanner.Error(); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	pageDataHeader, err := NewDataPageHeader(page, version.Jet4)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if pageDataHeader == nil {
		t.Error("Expected non nil page data")
	}

	if pageDataHeader.FreeSpace() != 0x0cc0 {
		t.Errorf("Wrongfree space. Expected 0cc0 - Got:%x\n", pageDataHeader.FreeSpace())
		t.FailNow()
	}

	if pageDataHeader.PagePointer() != 1 {
		t.Errorf("Wrong pointer value. Expected 1 - Got:%d\n", pageDataHeader.PagePointer())
		t.FailNow()
	}
	if pageDataHeader.RowsCount() != 2 {
		t.Errorf("Wrong rows count file size. Expected 2 - Got:%d\n", pageDataHeader.RowsCount())
		t.FailNow()
	}

	fmt.Printf("Page Data header:%+v\n\n", pageDataHeader)
	scanner.ReadPage()

	if scanner.Error() != nil {
		t.Error(err)
		t.FailNow()
	}

}
