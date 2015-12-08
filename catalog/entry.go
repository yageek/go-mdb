package catalog

type EntryKind int

const (
	TableKind EntryKind = iota
	FormKind
)

type Entry struct {
	name             string
	kind             EntryKind
	relatedPageIndex int64
}
