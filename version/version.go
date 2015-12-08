package version

type JetVersion byte

const (
	Jet3 JetVersion = 0x00
	Jet4            = 0x01
)

func (v JetVersion) MagicNumber() byte {
	return byte(v)
}
