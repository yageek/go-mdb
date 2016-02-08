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

func (j *Jet3RowDefinition) ColumnsValue() uint16               { return uint16(j.Columns) }
func (j *Jet3RowDefinition) FixedColumnsLengthValue() []byte    { return j.FixedColumnsLength }
func (j *Jet3RowDefinition) VariableColumnsLengthValue() []byte { return j.VariableColumnsLength }
func (j *Jet3RowDefinition) VarTableValue() []byte              { return j.VarTable }
func (j *Jet3RowDefinition) JumpTableValue() []byte             { return j.JumpTable }
func (j *Jet3RowDefinition) EodValue() uint16                   { return uint16(j.Eod) }
func (j *Jet3RowDefinition) VarLenValue() uint16                { return uint16(j.VarLen) }
func (j *Jet3RowDefinition) NullMaskValue() []byte              { return j.NullMask }

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

func (j *Jet4RowDefinition) ColumnsValue() uint16               { return j.Columns }
func (j *Jet4RowDefinition) FixedColumnsLengthValue() []byte    { return j.FixedColumnsLength }
func (j *Jet4RowDefinition) VariableColumnsLengthValue() []byte { return j.VariableColumnsLength }
func (j *Jet4RowDefinition) VarTableValue() []byte              { return j.VarTable }
func (j *Jet4RowDefinition) JumpTableValue() []byte             { return j.JumpTable }
func (j *Jet4RowDefinition) EodValue() uint16                   { return j.Eod }
func (j *Jet4RowDefinition) VarLenValue() uint16                { return j.VarLen }
func (j *Jet4RowDefinition) NullMaskValue() []byte              { return j.NullMask }

type JetRowDefinition interface {
	ColumnsValue() uint16
	FixedColumnsLengthValue() []byte
	VariableColumnsLengthValue() []byte
	VarTableValue() []byte
	JumpTableValue() []byte
	EodValue() uint16
	VarLenValue() uint16
	NullMaskValue() []byte
}
