package pages

import "errors"

// Errors relative to pages
var (
	ErrInvalidPageCode        = errors.New("Invalid page constant")
	ErrInvalidVersionConstant = errors.New("Invalid version constant")
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

func IsPageCodeValid(page []byte, code byte) bool {
	return page[0] == code
}

// IsDataPage tells whether the page is a DataCode Page
func IsDataPage(page []byte) bool {
	return IsPageCodeValid(page, DataPageCode)
}
