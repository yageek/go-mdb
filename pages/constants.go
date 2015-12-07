package pages

import "errors"

var (
	ErrInvalidPageCode = errors.New("Invalid version constant")
)

const (
	DatabaseDefinitionCode     int = 0x00
	DataPageCode                   = 0x01
	TableDefinitionCode            = 0x02
	IntermediateIndexPagesCode     = 0x03
	LeafIndexPagesCode             = 0x04
	PageUsegeBitmapsCode           = 0x05
)
