package definitionPage

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/yageek/go-mdb/pages"
	"github.com/yageek/go-mdb/util"
	"github.com/yageek/go-mdb/version"
)

// Errors
var (
	ErrInvalidVersionConstant = errors.New("Invalid version constant")
)

// Page offset
const (
	CodeOffset          = 0x3c
	Jet3TextOrderOffset = 0x3a
	Jet4TextOrderOffset = 0x6e
	KeyOffset           = 0x3e
	PasswordOffset      = 0x42
	PasswordMaskOffset  = 0x72

	CodeLength = 2
	KeyLength  = 4

	Jet3PasswordLength  = 20
	Jet3TextOrderLength = 2

	Jet4PasswordLength           = 20
	Jet4PasswordMaskOffsetLength = 8
	Jet4TextOrderLength          = 4
)

// DefinitionPage represents
// the table definition page
type DefinitionPage struct {
	version      version.JetVersion
	password     []byte
	passwordMask []byte
	key          []byte
	code         []byte
	textOrder    []byte
}

func (d *DefinitionPage) String() string {

	s := "Definition Page:\n"
	s += fmt.Sprintf("Version byte: 0x%x\n", d.version)
	s += fmt.Sprintf("Password:\n%v\n", hex.Dump(d.password))
	s += fmt.Sprintf("Password Mask:\n%v\n", hex.Dump(d.passwordMask))
	s += fmt.Sprintf("Key:\n%v\n", hex.Dump(d.key))
	s += fmt.Sprintf("Code:\n%v\n", hex.Dump(d.code))
	s += fmt.Sprintf("Text order:\n%v\n", hex.Dump(d.textOrder))
	return s
}

// NewDefinitionPage creates a new definition page
// from a buffer of bytes
func NewDefinitionPage(page []byte, v version.JetVersion) (*DefinitionPage, error) {

	if !pages.IsPageCodeValid(page, pages.DatabaseDefinitionCode) {
		return nil, pages.ErrInvalidPageCode
	}

	if page[version.VersionOffset] != v.MagicNumber() {
		return nil, ErrInvalidVersionConstant
	}

	definitionPage := new(DefinitionPage)

	definitionPage.version = v

	// Read password
	if v == version.Jet3 {
		util.DecodeBytes(page, &definitionPage.password, PasswordOffset, Jet3PasswordLength)
	} else if v == version.Jet4 {
		util.DecodeBytes(page, &definitionPage.password, PasswordOffset, Jet4PasswordLength)
	}

	// Password Jet4 if needed
	if v == version.Jet4 {
		util.DecodeBytes(page, &definitionPage.passwordMask, PasswordMaskOffset, Jet4PasswordMaskOffsetLength)
	}

	// Code
	util.DecodeBytes(page, &definitionPage.code, CodeOffset, CodeLength)

	// Text Order
	if v == version.Jet3 {
		util.DecodeBytes(page, &definitionPage.textOrder, Jet3TextOrderOffset, Jet3TextOrderLength)
	} else if v == version.Jet4 {
		util.DecodeBytes(page, &definitionPage.textOrder, Jet4TextOrderOffset, Jet4TextOrderLength)
	}

	// Key
	util.DecodeBytes(page, &definitionPage.key, KeyOffset, KeyLength)

	return definitionPage, nil
}
