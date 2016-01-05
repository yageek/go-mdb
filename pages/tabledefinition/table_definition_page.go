package tabledefinition

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/yageek/go-mdb/pages"
	"github.com/yageek/go-mdb/version"
)

// Column type
const (
	Bool      byte = 0x01
	Byte           = 0x02
	Int            = 0x03
	LongInt        = 0x04
	Money          = 0x05
	Float          = 0x06
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

func (j *Jet3TableDefinitionBlock) String() string {
	s := fmt.Sprintf("Length:%d\n", j.Length)
	s += fmt.Sprintf("Rows:%d\n", j.Rows)
	s += fmt.Sprintf("Autonumber:%d\n", j.Autonumber)

	return s
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

func (j *Jet3ColumnProperties) ColumnKind() byte                { return j.Kind }
func (j *Jet3ColumnProperties) ColumnDeletedID() uint16         { return j.IDDeleted }
func (j *Jet3ColumnProperties) VarLengthColumnOffset() uint16   { return j.OffsetV }
func (j *Jet3ColumnProperties) ColumnID() uint16                { return j.ID }
func (j *Jet3ColumnProperties) MiscValue() [2]byte              { return [2]byte{0, 0} }
func (j *Jet3ColumnProperties) MiscExtValue() [2]byte           { return [2]byte{0, 0} }
func (j *Jet3ColumnProperties) Bitmask() byte                   { return j.Mask }
func (j *Jet3ColumnProperties) Flags() byte                     { return 0 }
func (j *Jet3ColumnProperties) FixedLengthColumnOffset() uint16 { return j.OffsetF }
func (j *Jet3ColumnProperties) ColumnLength() uint16            { return j.Length }

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
	_      [4]byte
	NumIdx uint32
	_      [4]byte
}

func (j *Jet4RealIndex) Count() uint32 { return j.NumIdx }

// Jet3RealIndex as defined in documentation
type Jet3RealIndex struct {
	_      [4]byte
	NumIdx uint32
}

func (j *Jet3RealIndex) Count() uint32 { return j.NumIdx }

type JetRealIndex interface {
	Count() uint32
}

// NewRealIndex creates a new index
func NewRealIndex(page []byte, v version.JetVersion) (JetRealIndex, error) {

	buff := bytes.NewBuffer(page)
	var index JetRealIndex

	if v == version.Jet3 {
		index = new(Jet4RealIndex)
	} else {
		index = new(Jet4RealIndex)
	}
	err := binary.Read(buff, binary.LittleEndian, index)
	return index, err
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

func (j *Jet4ColumnProperties) ColumnKind() byte                { return j.Kind }
func (j *Jet4ColumnProperties) ColumnDeletedID() uint16         { return j.IDDeleted }
func (j *Jet4ColumnProperties) VarLengthColumnOffset() uint16   { return j.OffsetV }
func (j *Jet4ColumnProperties) ColumnID() uint16                { return j.ID }
func (j *Jet4ColumnProperties) MiscValue() [2]byte              { return j.Misc }
func (j *Jet4ColumnProperties) MiscExtValue() [2]byte           { return j.MiscExt }
func (j *Jet4ColumnProperties) Bitmask() byte                   { return j.Mask }
func (j *Jet4ColumnProperties) Flags() byte                     { return j.MiscFlags }
func (j *Jet4ColumnProperties) FixedLengthColumnOffset() uint16 { return j.OffsetF }
func (j *Jet4ColumnProperties) ColumnLength() uint16            { return j.Length }

type ColumnProperties interface {
	ColumnKind() byte
	ColumnDeletedID() uint16
	VarLengthColumnOffset() uint16
	ColumnID() uint16
	MiscValue() [2]byte
	MiscExtValue() [2]byte
	Bitmask() byte
	Flags() byte
	FixedLengthColumnOffset() uint16
	ColumnLength() uint16
}

func NewColumnProperty(page []byte, v version.JetVersion) (ColumnProperties, error) {

	buff := bytes.NewBuffer(page)
	var columnProperty ColumnProperties

	if v == version.Jet3 {
		columnProperty = new(Jet3ColumnProperties)
	} else {
		columnProperty = new(Jet4ColumnProperties)
	}
	err := binary.Read(buff, binary.LittleEndian, columnProperty)
	return columnProperty, err
}

// Column Order represents
// the bytes fields after the column properties
type ColumnOrder struct {
	Num   uint16
	Order byte
}

type Jet3ColumnUsage struct {
	UsedPages uint32
	FirstDp   uint32
	Flags     byte
}

func (j *Jet3ColumnUsage) ColumnUsedPages() uint32  { return j.UsedPages }
func (j *Jet3ColumnUsage) FirstPagePointer() uint32 { return j.FirstDp }
func (j *Jet3ColumnUsage) ColumnsFlags() byte       { return j.Flags }

type Jet4ColumnUsage struct {
	UsedPages uint32
	FirstDp   uint32
	Flags     byte
	_         [9]byte
}

func (j *Jet4ColumnUsage) ColumnUsedPages() uint32  { return j.UsedPages }
func (j *Jet4ColumnUsage) FirstPagePointer() uint32 { return j.FirstDp }
func (j *Jet4ColumnUsage) ColumnsFlags() byte       { return j.Flags }

type ColumnUsage interface {
	ColumnUsedPages() uint32
	FirstPagePointer() uint32
	ColumnsFlags() byte
}

// NewDefinitionBlock creates a new definition block
func NewDefinitionBlock(page []byte, v version.JetVersion) (TableDefinitionBlock, error) {

	buff := bytes.NewBuffer(page)
	var definitionBlock TableDefinitionBlock = nil

	if v == version.Jet4 {
		definitionBlock = new(Jet4TableDefinitionBlock)
	} else if v == version.Jet3 {
		definitionBlock = new(Jet3TableDefinitionBlock)
	} else {
		return nil, pages.ErrInvalidVersionConstant
	}

	err := binary.Read(buff, binary.LittleEndian, definitionBlock)
	return definitionBlock, err

}

type TableDefinitionPage struct {
	definitionHeader *TableDefinitionHeader
	definitionBlock  TableDefinitionBlock
	indexes          []JetRealIndex
	columnsProperty  []ColumnProperties
	columnNames      []string
	columnOrders     []ColumnOrder
	columnUsages     []ColumnUsage
}

func NewTableDefinitionPage(page []byte, v version.JetVersion) (*TableDefinitionPage, error) {

	// Creates a buffer
	buffer := bytes.NewBuffer(page)

	// Read page header
	headerRaw := make([]byte, 8)
	_, err := buffer.Read(headerRaw)

	if err != nil {
		return nil, err
	}

	header := new(TableDefinitionHeader)
	buff := bytes.NewBuffer(headerRaw)

	err = binary.Read(buff, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	// Read definition block
	blockRaw := make([]byte, 55)
	_, err = buffer.Read(blockRaw)
	if err != nil {
		return nil, err
	}

	definitionBlock, err := NewDefinitionBlock(blockRaw, v)

	if err != nil {
		return nil, err
	}

	// Read indexes
	var indexes []JetRealIndex = make([]JetRealIndex, 0)

	indexBuff := make([]byte, 12)
	for i := uint32(0); i < definitionBlock.LogicalIndexesCount(); i++ {

		_, err := buffer.Read(indexBuff)
		if err != nil {
			return nil, err
		}

		index, err := NewRealIndex(indexBuff, v)
		if err != nil {
			return nil, err
		}

		indexes = append(indexes, index)
	}

	// Read Columns properties
	var columProperties []ColumnProperties = make([]ColumnProperties, 0)
	columnBuff := make([]byte, 25)

	for i := uint16(0); i < definitionBlock.ColumnsCount(); i++ {
		_, err := buffer.Read(columnBuff)
		if err != nil {
			return nil, err
		}

		column, err := NewColumnProperty(columnBuff, v)
		if err != nil {
			return nil, err
		}
		columProperties = append(columProperties, column)
	}

	// Read columns name
	var columnNames []string = make([]string, definitionBlock.ColumnsCount())
	var lengthBuffer []byte = make([]byte, 2)

	for i := uint16(0); i < definitionBlock.ColumnsCount(); i++ {

		//Read length
		_, err := buffer.Read(lengthBuffer)
		if err != nil {
			return nil, err
		}

		buff := bytes.NewBuffer(lengthBuffer)
		var length uint16
		err = binary.Read(buff, binary.LittleEndian, &length)

		if err != nil {
			return nil, err
		}

		bufferString := make([]byte, length)

		_, err = buffer.Read(bufferString)

		if err != nil {
			return nil, err
		}

		columnNames[i] = string(bufferString)

	}

	// Pass unknown byte
	buffer.Next(int(definitionBlock.IndexEntriesCount()))

	columnsOrder := make([]ColumnOrder, 10)
	for i := range columnsOrder {
		order := ColumnOrder{}
		err := binary.Read(buffer, binary.LittleEndian, &order)
		if err != nil {
			return nil, err
		}

		columnsOrder[i] = order
	}

	return &TableDefinitionPage{
		definitionHeader: header,
		definitionBlock:  definitionBlock,
		indexes:          indexes,
		columnsProperty:  columProperties,
		columnNames:      columnNames,
		columnOrders:     columnsOrder,
	}, nil
}
