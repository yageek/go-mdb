//go:generate stringer -type=JetVersion
package version

import "errors"

type JetVersion byte

const (
	VersionOffset = 0x14

	Jet3 JetVersion = 0x00
	Jet4 JetVersion = 0x01
)

var (
	ErrUnknownVersion = errors.New("unknown version")
)

func NewJetVersion(b byte) (JetVersion, error) {

	v := JetVersion(b)
	switch v {
	case Jet3:
		fallthrough
	case Jet4:
		return v, nil
	default:
		return v, ErrUnknownVersion
	}
}
func (v JetVersion) MagicNumber() byte {
	return byte(v)
}

func (v JetVersion) PageSize() int64 {
	switch v {
	case Jet3:
		return 2048
	default:
		return 4096
	}
}

type JetBinaryStructure interface {
	MagicNumber() byte
	PageSize() int64
}
