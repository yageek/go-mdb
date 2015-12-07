package mdb

// Version of the database
type JetVersion uint8

// Database magic
const (
	Jet3       JetVersion = 0x00
	Jet4                  = 0x01
	JetUnknown            = 0xff
)

func (v JetVersion) String() string {
	switch v {
	case Jet3:
		return "JET3"
	case Jet4:
		return "JET4"
	default:
		return "Unknown version"

	}
}

// Offset
const (
	PasswordOffset     int64 = 0x42
	PasswordMaskOffset       = 0x72
	KeyOffset                = 0x3e
	CodeOffset               = 0x3c
	TextOrderOffset          = 0x3a
)

// PageType reprensents the Page type
type PageType int16

const (
	DatabaseDefinitionPageType PageType = 0x00
	DataPageType                        = 0x01
	TableDefinitionType                 = 0x02
	IntermediateIndexPagesType          = 0x03
	LeafIndexPagesType                  = 0x04
	PageUsegeBitmapsType                = 0x05

	VersionOffset int = 0x14
)
