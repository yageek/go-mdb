package catalog

import (
	"errors"
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

	dataPageIndex, err := c.nextDataPageForTDEF(0x02)
	if err != nil {
		return err
	}

	if !c.scanner.ReadPageAtIndex(dataPageIndex) {
		return c.scanner.Error()
	}

	row, err := data.NewPage(c.scanner.Page(), c.jetVersion)

	if err != nil {
		return err
	}

	fmt.Println("Row:", row)

	return nil
}

func (c *Catalog) nextDataPageForTDEF(pointer int64) (int64, error) {
	for c.scanner.ReadPage() {
		page := c.scanner.Page()
		if pages.IsDataPage(page) {

			header, err := data.NewDataPageHeader(page, c.jetVersion)

			if err != nil {
				fmt.Println("Could not go on")
				break
			}

			if int64(header.PagePointer()) == pointer {
				return c.scanner.CurrentPageIndex(), nil
			}
		}
	}

	msg := fmt.Sprintf("Can not find datapage for TDEF:%v\n", pointer)
	return -1, errors.New(msg)
}

// Close closes the catalog
func (c *Catalog) Close() error {

	return c.scanner.Close()
}

// Version returns the jet version
func (c *Catalog) Version() version.JetVersion {
	return c.jetVersion
}
