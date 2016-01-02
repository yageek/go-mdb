package catalog

// EntryKind represents the different entry types
type EntryKind int

// Constants
const (
	TableKind EntryKind = iota
	FormKind
)

// Entry represents a row in the MSysObjectsTable at 0x02
type Entry struct {
	name             string
	kind             EntryKind
	relatedPageIndex int64
}

// NewEntry creates a new entry
func NewEntry(pageIndex int64, kind EntryKind, name string) *Entry {
	return &Entry{relatedPageIndex: pageIndex, kind: kind, name: name}
}
