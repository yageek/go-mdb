package filepage

import (
	"os"
	"testing"

	"github.com/yageek/go-mdb/pages"
)

const (
	jet4DatabasePath = "../test_databases/EPSG_v8_9.mdb"
	jet3DatabasePath = "../test_databases/Books_be.mdb"
)

func helperValidPageCount(filename string, pageSize int64, t *testing.T) {

	file, err := os.Open(filename)

	if err != nil {
		t.Error(err.Error())
		return
	}
	scanner, err := NewScanner(file, pageSize)

	if err != nil {
		t.Error(err.Error())
		return
	}

	defer scanner.Close()

	pageNumber := scanner.PagesNumber()
	var pageCounter int64

	for scanner.ReadPage() {
		pageCounter++
	}

	if scanner.Error() != nil {
		t.Error(scanner.Error())
		return
	}

	if pageCounter != pageNumber {
		t.Errorf("Wrong count in %v  Counted:%v - Expected:%v\n", filename, pageCounter, pageNumber)
	}

}
func TestDatabaseCount(t *testing.T) {

	helperValidPageCount(jet4DatabasePath, pages.Jet4PageSize, t)
	helperValidPageCount(jet3DatabasePath, pages.Jet3PageSize, t)

}
