package filepage

import "testing"

const (
	databaseFile            = "../test_databases/EPSG_v8_6.mdb"
	databaseLowPageSizeFile = "../test_databases/Books_be.mdb"
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

	if scanner.Error() != nil {
		t.Error(scanner.Error())
		return
	}

	if pageCounter != pageNumber {
		t.Errorf("Wrong count in %v  Counted:%v - Expected:%v\n", filename, pageCounter, pageNumber)
	}

}
func TestDatabaseCount(t *testing.T) {

	//	helperValidPageCount(databaseFile, 4096, t)
	helperValidPageCount(databaseLowPageSizeFile, 2048, t)

}
