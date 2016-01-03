package catalog

import (
	"fmt"
	"testing"

	"github.com/yageek/go-mdb/test_databases"
)

func TestMSysObjects(t *testing.T) {

	c, err := NewCatalog(testdatabases.JET4DatabasePath())
	if err != nil {
		t.Error("Could not open file:", err.Error())
		t.FailNow()
	}

	if c == nil {
		t.Error("Invalid catalog file")
		t.FailNow()
	}

	err = c.Read()
	if err != nil {
		t.Error("Failed reading:", err.Error())
		t.FailNow()
	}

	fmt.Println("Catalog:", c.mSysObjectsDefinition)
}
