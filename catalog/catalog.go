package catalog

import (
	"fmt"
	"github.com/yageek/go-mdb/filepage"
	"github.com/yageek/go-mdb/pages"
	"github.com/yageek/go-mdb/pages/data"
	"github.com/yageek/go-mdb/pages/definition"
	"github.com/yageek/go-mdb/pages/tabledefinition"
	"github.com/yageek/go-mdb/version"
	"os"
)

// Catalog represents the structure of the access files
type Catalog struct {
	jetVersion            version.JetVersion
	scanner               *filepage.Scanner
	definitionPage        *definition.DefinitionPage
	mSysObjectsDefinition *tabledefinition.TableDefinitionPage
}

// NewCatalog returns a new catalog object
func NewCatalog(filename string) (*Catalog, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	versionBuf := make([]byte, 1)

	_, err = file.ReadAt(versionBuf, version.VersionOffset)

	if err != nil {
		return nil, err
	}

	v, err := version.NewJetVersion(versionBuf[0])

	if err != nil {
		return nil, err
	}

	scanner, err := filepage.NewScanner(file, v.PageSize())

	if err != nil {
		return nil, err
	}

	return &Catalog{scanner: scanner, jetVersion: v}, nil
}

func (c *Catalog) Read() error {

	// Read definition page
	c.scanner.ReadPage()
	err := c.scanner.Error()

	if err := c.scanner.Error(); err != nil {
		return err
	}

	//Read MDB file header
	c.definitionPage, err = definition.NewDefinitionPage(c.scanner.Page(), c.jetVersion)
	if err != nil {
		return nil
	}

	// Read MDB MSysObjects tabledefinition
	c.scanner.ReadPageAtIndex(2)

	msysObjects, err := tabledefinition.NewTableDefinitionPage(c.scanner.Page(), c.jetVersion)

	if err != nil {
		return err
	}

	c.mSysObjectsDefinition = msysObjects
	// Looks for the data pages for the MSysObjects

	for c.nextDataPageForTDEF(0x02) {
		row, err := data.NewPage(c.scanner.Page(), c.jetVersion)
		if err != nil {
			return err
		}
		fmt.Println("Data page for TDEF found:", row)
	}

	return nil
}

func (c *Catalog) nextDataPage() bool {
	for c.scanner.ReadPage() {
		page := c.scanner.Page()
		if pages.IsDataPage(page) {
			return true
		}
	}
	return false
}

func (c *Catalog) nextDataPageForTDEF(pointer int64) bool {

	for c.nextDataPage() {
		header, err := data.NewDataPageHeader(c.scanner.Page(), c.jetVersion)
		if err != nil {
			fmt.Println("Could not go on")
			break
		}

		if int64(header.PagePointer()) == pointer {
			return true
		}
	}
	return false
}

// Close closes the catalog
func (c *Catalog) Close() error {

	return c.scanner.Close()
}

// Version returns the jet version
func (c *Catalog) Version() version.JetVersion {
	return c.jetVersion
}
