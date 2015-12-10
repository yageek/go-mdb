package pages

import (
	"bytes"
	"encoding/binary"

	"github.com/yageek/go-mdb/version"
)

//Offset
const (
	HeaderFreeSpaceOffset int = 0x02

	HeaderTableDefinitionPointerOffset = 0x04

	Jet3HeaderNumRowsOffset = 0x08
	Jet4HeaderNumRowsOffset = 0x0c
)

type Jet3DatapageHeader struct {
	Kind    byte
	_       byte
	Space   uint16
	Pointer uint32
	Count   uint16
}

func (j *Jet3DatapageHeader) PageKind() byte      { return j.Kind }
func (j *Jet3DatapageHeader) FreeSpace() uint16   { return j.Space }
func (j *Jet3DatapageHeader) PagePointer() uint32 { return j.Pointer }
func (j *Jet3DatapageHeader) RowsCount() uint16   { return j.Count }

type Jet4DatapageHeader struct {
	Kind    byte
	_       byte
	Space   uint16
	Pointer uint32
	_       [4]byte
	Count   uint16
}

func (j *Jet4DatapageHeader) PageKind() byte      { return j.Kind }
func (j *Jet4DatapageHeader) FreeSpace() uint16   { return j.Space }
func (j *Jet4DatapageHeader) PagePointer() uint32 { return j.Pointer }
func (j *Jet4DatapageHeader) RowsCount() uint16   { return j.Count }

type DatapageHeader interface {
	PageKind() byte
	FreeSpace() uint16
	PagePointer() uint32
	RowsCount() uint16
}

// NewDataPageHeader returns a new datapage header from page
func NewDataPageHeader(page []byte, v version.JetVersion) (DatapageHeader, error) {

	buff := bytes.NewReader(page)
	var header DatapageHeader

	if v == version.Jet3 {
		header = new(Jet3DatapageHeader)
	} else {
		header = new(Jet4DatapageHeader)
	}
	err := binary.Read(buff, binary.LittleEndian, header)

	if err != nil {
		return nil, err
	}

	if header.PageKind() != v.MagicNumber() {
		return nil, ErrInvalidPageCode
	}
	return header, nil
}
