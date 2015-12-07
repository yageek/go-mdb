package filepage

import "testing"

const (
	databaseFile            = "../databases/EPSG_v8_6.mdb"
	databaseLowPageSizeFile = "../databases/Books_be.mdb"
)

func TestNewScanner(t *testing.T) {

	scanner, err := NewScanner(databaseFile, 4096)

	if err != nil {
		t.Error(err.Error())
		return
	}

	defer scanner.Close()

}

func helperValidPageCount(filename string, pageSize int64, t *testing.T) {
	scanner, err := NewScanner(filename, pageSize)

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

	if pageCounter != pageNumber {
		t.Errorf("Unexpected pages number in %v - Counted:%v - Expected:%v\n", filename, pageCounter, pageNumber)
	}

}
func TestDatabaseCount(t *testing.T) {

	helperValidPageCount(databaseFile, 4096, t)
	helperValidPageCount(databaseLowPageSizeFile, 2048, t)

}
