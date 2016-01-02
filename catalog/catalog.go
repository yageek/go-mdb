package catalog

import (
	"os"

	"github.com/yageek/go-mdb/filepage"
	"github.com/yageek/go-mdb/pages"
	"github.com/yageek/go-mdb/version"
)

// Catalog represents the structure of the access files
type Catalog struct {
	jetVersion     version.JetVersion
	scanner        *filepage.Scanner
	definitionPage *pages.DefinitionPage
	entries        []*Entry
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

	return &Catalog{scanner: scanner}, nil
}

func (c *Catalog) AddEntry(entry *Entry) {
	c.entries = append(c.entries, entry)
}

func (c *Catalog) Read() error {

	// Read definition page
	c.scanner.ReadPage()
	err := c.scanner.Error()

	if err := c.scanner.Error(); err != nil {
		return err
	}

	c.definitionPage, err = pages.NewDefinitionPage(c.scanner.Page(), c.jetVersion)
	if err != nil {
		return nil
	}

	// Add the MSysObjects entry
	entry := NewEntry(2, TableKind, "MSysObjects")
	c.AddEntry(entry)

	return nil
}

func (c *Catalog) Close() error {

	return c.scanner.Close()
}
