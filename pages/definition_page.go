package pages

import "errors"

var (
	ErrInvalidVersionConstant = errors.New("Invalid version constant")
)

// Page offset
const (
	VersionOffset      int64 = 0x14
	CodeOffset               = 0x3c
	TextOrderOffset          = 0x3a
	KeyOffset                = 0x3e
	PasswordOffset           = 0x42
	PasswordMaskOffset       = 0x72

	Jet3 int = 0x00
	Jet4     = 0x01
)

// DefinitionPage represents
// the table definition page
type DefinitionPage struct {
	version      int
	password     []byte
	passwordMask []byte
	key          []byte
	code         []byte
	textOrder    []byte
}

// NewDefinitionPage creates a new definition page
// from a buffer of bytes
func NewDefinitionPage(page []byte) (*DefinitionPage, error) {

	pageCode := int(page[0])
	if pageCode != DatabaseDefinitionCode {
		return nil, ErrInvalidPageCode
	}

	version := int(page[VersionOffset])

	if version != Jet3 && version != Jet4 {
		return nil, ErrInvalidVersionConstant
	}

	return &DefinitionPage{version: version}, nil
}
