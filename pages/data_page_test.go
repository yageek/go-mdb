package pages

import (
	"os"
	"testing"

	"github.com/yageek/go-mdb/filepage"
	"github.com/yageek/go-mdb/version"
)

func TestDataPage(t *testing.T) {

	file, err := os.Open(jet4DatabasePath)
	if err != nil {
		t.Error(err.Error())
		return
	}

	scanner, err := filepage.NewScanner(file, Jet4PageSize)

	if err != nil {
		t.Error(err.Error())
		return
	}
	defer scanner.Close()

	scanner.ReadPage()

	scanner.ReadPage()

	page := scanner.Page()

	if err := scanner.Error(); err != nil {
		t.Error(err.Error())
		return
	}

	pageData, err := NewDataPageHeader(page, version.Jet4)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if pageData == nil {
		t.Error("Expected non nil page data")
	}

	if pageData.freeSpaceSize != 0x0cc0 {
		t.Errorf("Wrong page file size. Expected 0cc0 - Got:%x\n", pageData.freeSpaceSize)
		return
	}

	if pageData.pageDefinitionAddress != 1 {
		t.Errorf("Wrong page file size. Expected 1 - Got:%d\n", pageData.pageDefinitionAddress)
		return
	}
	if pageData.recordNum != 2 {
		t.Errorf("Wrong page file size. Expected 2 - Got:%d\n", pageData.recordNum)
		return
	}

}
