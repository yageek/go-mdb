package catalog

import (
	"os"

	"github.com/yageek/go-mdb/filepage"
	"github.com/yageek/go-mdb/pages/definition"
	"github.com/yageek/go-mdb/pages/tabledefinition"
	"github.com/yageek/go-mdb/version"
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

	c.definitionPage, err = definition.NewDefinitionPage(c.scanner.Page(), c.jetVersion)
	if err != nil {
		return nil
	}

	msysObjects, err := tabledefinition.NewTableDefinitionPage(c.scanner.Page(), c.jetVersion)

	if err != nil {
		return nil
	}

	c.mSysObjectsDefinition = msysObjects
	return nil
}

func (c *Catalog) Close() error {

	return c.scanner.Close()
}
