package data

// Jet3RowDefinition as described in documentation
type Jet3RowDefinition struct {
	Columns               uint8
	FixedColumnsLength    []byte
	VariableColumnsLength []byte
	VarTable              []byte
	JumpTable             []byte
	Eod                   uint8
	VarLen                uint8
	NullMask              []byte
}

// Jet4RowDefinition as described in documentation
type Jet4RowDefinition struct {
	Columns               uint16
	FixedColumnsLength    []byte
	VariableColumnsLength []byte
	VarTable              []byte
	JumpTable             []byte
	Eod                   uint16
	VarLen                uint16
	NullMask              []byte
}
