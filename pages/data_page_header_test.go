package pages

import (
	"fmt"
	"os"
	"testing"

	"github.com/yageek/go-mdb/filepage"
	"github.com/yageek/go-mdb/version"
)

func TestDataPage(t *testing.T) {

	file, err := os.Open(jet4DatabasePath)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	scanner, err := filepage.NewScanner(file, Jet4PageSize)

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

	pageData, err := NewDataPageHeader(page, version.Jet4)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if pageData == nil {
		t.Error("Expected non nil page data")
	}

	if pageData.FreeSpace() != 0x0cc0 {
		t.Errorf("Wrongfree space. Expected 0cc0 - Got:%x\n", pageData.FreeSpace())
		t.FailNow()
	}

	if pageData.PagePointer() != 1 {
		t.Errorf("Wrong pointer value. Expected 1 - Got:%d\n", pageData.PagePointer())
		t.FailNow()
	}
	if pageData.RowsCount() != 2 {
		t.Errorf("Wrong rows count file size. Expected 2 - Got:%d\n", pageData.RowsCount())
		t.FailNow()
	}

	scanner.ReadPage()

	if scanner.Error() != nil {
		t.Error(err)
		t.FailNow()
	}

	rawBlock := scanner.Page()[8:]

	blockDefinition, err := NewDefinitionBlock(rawBlock, version.Jet4)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_ = "breakpoint"
	fmt.Println("Block:", blockDefinition)
}
