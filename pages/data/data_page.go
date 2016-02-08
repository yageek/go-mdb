package data

import (
	"github.com/yageek/go-mdb/version"
)

type Page struct {
	Header        DatapageHeader
	RowDefinition JetRowDefinition
}

func NewPage(page []byte, v version.JetVersion) (*Page, error) {

	header, err := NewDataPageHeader(page, v)
	if err != nil {
		return nil, err
	}

	return &Page{Header: header}, nil
}
