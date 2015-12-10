package pages

import (
	"bytes"
	"encoding/binary"

	"github.com/yageek/go-mdb/version"
)

// Column type
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

// Table system
const (
	User   byte = 0x4e
	System      = 0x53
)

// TableDefinitionHeader represents a generic
// table definition header
type TableDefinitionHeader struct {
	Kind            uint8
	_               byte
	ID              uint16
	NextPagePointer uint32
}

//TableDefinitionBlock interface
type TableDefinitionBlock interface {
	DataLength() uint32
	RowsCount() uint32
	AutonumberColumn() uint32
	PageKind() byte
	MaxColumnsCount() uint16
	VarColumnsCount() uint16
	ColumnsCount() uint16
	LogicalIndexesCount() uint32
	IndexEntriesCount() uint32
	UsedPagesPointer() uint32
	FreePagesPointer() uint32
}

// Jet3TableDefinitionBlock as defined in documentation
type Jet3TableDefinitionBlock struct {
	Length     uint32
	Rows       uint32
	Autonumber uint32

	Kind           byte
	MaxColumns     uint16
	VarCount       uint16
	Columns        uint16
	LogicalIndexes uint32
	IndexEntries   uint32
	UsedPages      uint32
	FreePages      uint32
}

// Implementation of TableDefinitionBlock
func (j *Jet3TableDefinitionBlock) DataLength() uint32          { return j.Length }
func (j *Jet3TableDefinitionBlock) RowsCount() uint32           { return j.Rows }
func (j *Jet3TableDefinitionBlock) AutonumberColumn() uint32    { return j.Autonumber }
func (j *Jet3TableDefinitionBlock) PageKind() byte              { return j.Kind }
func (j *Jet3TableDefinitionBlock) MaxColumnsCount() uint16     { return j.MaxColumns }
func (j *Jet3TableDefinitionBlock) VarColumnsCount() uint16     { return j.VarCount }
func (j *Jet3TableDefinitionBlock) ColumnsCount() uint16        { return j.Columns }
func (j *Jet3TableDefinitionBlock) LogicalIndexesCount() uint32 { return j.LogicalIndexes }
func (j *Jet3TableDefinitionBlock) IndexEntriesCount() uint32   { return j.IndexEntries }
func (j *Jet3TableDefinitionBlock) UsedPagesPointer() uint32    { return j.UsedPages }
func (j *Jet3TableDefinitionBlock) FreePagesPointer() uint32    { return j.FreePages }

// Jet3RealIndex as defined in documentation
type Jet3RealIndex struct {
	_     [4]byte
	Count uint32
}

// Jet3ColumnProperties as defined in documentation
type Jet3ColumnProperties struct {
	Kind      byte
	IDDeleted uint16
	OffsetV   uint16
	ID        uint16
	Order     uint16

	_       [2]byte
	Mask    byte
	OffsetF uint16
	Length  uint16
}

// Jet4TableDefinitionBlock as defined in documentation
type Jet4TableDefinitionBlock struct {
	Length     uint32
	_          [4]byte
	Rows       uint32
	Autonumber uint32

	AutonumberFlag    byte
	_                 [3]byte
	AutonumberComplex uint32

	_ [8]byte

	Kind byte

	MaxColumns     uint16
	VarCount       uint16
	Columns        uint16
	LogicalIndexes uint32
	IndexEntries   uint32
	UsedPages      uint32
	FreePages      uint32
}

// Implementation of TableDefinitionBlock
func (j *Jet4TableDefinitionBlock) DataLength() uint32          { return j.Length }
func (j *Jet4TableDefinitionBlock) RowsCount() uint32           { return j.Rows }
func (j *Jet4TableDefinitionBlock) AutonumberColumn() uint32    { return j.Autonumber }
func (j *Jet4TableDefinitionBlock) PageKind() byte              { return j.Kind }
func (j *Jet4TableDefinitionBlock) MaxColumnsCount() uint16     { return j.MaxColumns }
func (j *Jet4TableDefinitionBlock) VarColumnsCount() uint16     { return j.VarCount }
func (j *Jet4TableDefinitionBlock) ColumnsCount() uint16        { return j.Columns }
func (j *Jet4TableDefinitionBlock) LogicalIndexesCount() uint32 { return j.LogicalIndexes }
func (j *Jet4TableDefinitionBlock) IndexEntriesCount() uint32   { return j.IndexEntries }
func (j *Jet4TableDefinitionBlock) UsedPagesPointer() uint32    { return j.UsedPages }
func (j *Jet4TableDefinitionBlock) FreePagesPointer() uint32    { return j.FreePages }

// Jet4RealIndex as defined in documentation
type Jet4RealIndex struct {
	_     [4]byte
	Count uint32
	_     [4]byte
}

// Jet4ColumnProperties as defined in documentation
type Jet4ColumnProperties struct {
	Kind byte
	_    [4]byte

	IDDeleted uint16
	OffsetV   uint16
	ID        uint16
	Misc      [2]byte
	MiscExt   [2]byte

	Mask      byte
	MiscFlags byte
	_         [4]byte
	OffsetF   uint16
	Length    uint16
}

// NewDefinitionBlock creates a new definition block
func NewDefinitionBlock(page []byte, v version.JetVersion) (TableDefinitionBlock, error) {

	if v == version.Jet4 {

		definitionBlock := new(Jet4TableDefinitionBlock)
		buff := bytes.NewBuffer(page)
		err := binary.Read(buff, binary.LittleEndian, definitionBlock)

		return definitionBlock, err
	}

	return nil, ErrInvalidVersionConstant

}
