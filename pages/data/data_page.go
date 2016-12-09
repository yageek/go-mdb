package data

import (
	"bytes"
	"encoding/binary"
	"github.com/yageek/go-mdb/version"
)

type recordOffset struct {
	value    uint16
	indirect bool
}

func newRecordOffset(offset uint16) recordOffset {

	o := recordOffset{}
	o.value = offset & 0x00FF
	flag := (offset & 0xFF00) >> 8
	o.indirect = flag == 0x40

	return o
}

type Page struct {
	Header         DatapageHeader
	RowDefinition  JetRowDefinition
	recordsOffsets []recordOffset
}

func NewPage(page []byte, v version.JetVersion) (*Page, error) {

	header, err := NewDataPageHeader(page, v)
	if err != nil {
		return nil, err
	}
	offsets := []recordOffset{}

	buff := bytes.NewBuffer(page)

	var i uint16
	var offset uint16 = 0
	for i = 0; i < header.RowsCount(); i++ {
		err := binary.Read(buff, binary.LittleEndian, &offset)
		if err != nil {
			return nil, err
		}
		offsets = append(offsets, newRecordOffset(offset))
	}

	return &Page{Header: header, recordsOffsets: offsets}, nil
}
