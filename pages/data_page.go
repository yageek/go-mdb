package pages

import (
	"fmt"

	"github.com/yageek/go-mdb/util"
)

// Datapage holds information
// about a data page
type Datapage struct {
	index  int64
	header *DatapageHeader
}

const (
	HeaderFreeSpaceOffset int = 0x02

	HeaderTableDefinitionPointerOffset = 0x04

	Jet3HeaderNumRowsOffset = 0x08
	Jet4HeaderNumRowsOffset = 0x0C

	HeaderNumRowsLength = 2
)

// DatapageHeader represents a generic
// data page header
type DatapageHeader struct {
	freeSpaceSize         uint16
	pageDefinitionAddress uint64
	recordNum             uint16
}

// NewDataPageHeader returns a new datapage header from page
func NewDataPageHeader(page []byte, version byte) (*DatapageHeader, error) {
	if page[0] != DataPageCode {
		return nil, ErrInvalidPageCode
	}

	header := new(DatapageHeader)
	err := header.readHeaderValues(page, version)

	if err != nil {
		return nil, err
	}

	return header, err
}

func (h *DatapageHeader) String() string {
	s := "Data Page Header:\n"
	s += fmt.Sprintf("\tFree Space:%d KB\n", h.freeSpaceSize/1024)
	s += fmt.Sprintf("\tPage Adress:0x%x\n", h.pageDefinitionAddress)
	s += fmt.Sprintf("\tRecord num:%d\n", h.recordNum)
	return s
}

func (h *DatapageHeader) readHeaderValues(page []byte, version byte) error {

	var headerNumOffset int
	if version == Jet3 {
		headerNumOffset = Jet3HeaderNumRowsOffset
	} else {
		headerNumOffset = Jet4HeaderNumRowsOffset
	}

	lookupValues := map[int]interface{}{
		HeaderFreeSpaceOffset:              &h.freeSpaceSize,
		HeaderTableDefinitionPointerOffset: &h.pageDefinitionAddress,
		headerNumOffset:                    &h.recordNum,
	}

	for k, v := range lookupValues {
		err := util.DecodeValue(page, k, v)

		if err != nil {
			return err
		}
	}
	return nil
}
