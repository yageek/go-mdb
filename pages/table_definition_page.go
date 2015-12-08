package pages

const (
	Bool      byte = 0x01
	Byte           = 0x02
	Int            = 0x03
	LongInt        = 0x04
	Money          = 0x05
	Fload          = 0x06
	Double         = 0x07
	DateTime       = 0x08
	Binary         = 0x09
	Text           = 0x0A
	Ole            = 0x0B
	Memo           = 0x0C
	UnknownOD      = 0x0D
	UnknownOE      = 0x0E
	Repid          = 0x0F
	Numeric        = 0x10
)

// TableDefinitionHeader represents a generic
// table definition header
type TableDefinitionHeader struct {
	id                         int16
	nextTableDefinitionAddress int64
}

type TableDefinitionBlock struct {
	dataLength        uint64
	rowsNum           uint64
	autonumber        uint64
	tableType         uint8
	maxColumns        uint16
	varColumnCount    uint16
	columnsCount      uint16
	logicalIndexCount uint64
	indexEntriesCount uint64
	freePagesAddress  uint64
}
