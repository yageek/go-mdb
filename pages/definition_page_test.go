package pages

import (
	"fmt"
	"testing"

	"github.com/yageek/go-mdb/filepage"
)

const (
	jet4DatabasePath = "../test_databases/EPSG_v8_6.mdb"
	jet3DatabasePath = "../test_databases/Books_be.mdb"
)

func TestAllJet4Database(t *testing.T) {

	scanner, err := filepage.NewScanner(jet4DatabasePath, Jet4PageSize)

	if err != nil {
		t.Error(err.Error())
		return
	}
	defer scanner.Close()

	for scanner.ReadPage() {
		page := scanner.Page()
		fmt.Println("------------------------------------------")
		fmt.Printf("Page:%v - Offset:%v \n", scanner.CurrentPageIndex(), scanner.CurrentOffset())

		switch page[0] {
		case DatabaseDefinitionCode:
			definition, err := NewDefinitionPage(page, Jet4)
			if err != nil {
				break
			}
			fmt.Println(definition)

		case DataPageCode:

			header, err := NewDataPageHeader(page, Jet4)
			if err != nil {
				break
			}
			fmt.Println(header)
		case TableDefinitionCode:

			fmt.Println("TDEF Definition Page")
		default:
			fmt.Println("Unknown")
		}
	}

	if err := scanner.Error(); err != nil {
		t.Error(err.Error())
		return
	}

}
