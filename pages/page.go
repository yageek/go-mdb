package pages

import "errors"

// Errors relative to pages
var (
	ErrInvalidPageCode = errors.New("Invalid version constant")
)

// Page code
const (
	DatabaseDefinitionCode     byte = 0x00
	DataPageCode                    = 0x01
	TableDefinitionCode             = 0x02
	IntermediateIndexPagesCode      = 0x03
	LeafIndexPagesCode              = 0x04
	PageUsegeBitmapsCode            = 0x05

	Jet3PageSize = 2048
	Jet4PageSize = 4096
)

func isPageCodeValid(page []byte, code byte) bool {
	return page[0] == code
}

// Page is a generic interface to a page
type Page interface {
	PageIndex() int64
	PageType() byte
}

type basePage struct {
	index int64
}

func (p *basePage) PageIndex() int64 {
	return p.index
}
