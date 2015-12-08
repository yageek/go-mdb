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

func NewEntry(pageIndex int64, kind EntryKind, name string) *Entry {
	return &Entry{relatedPageIndex: pageIndex, kind: kind, name: name}
}
