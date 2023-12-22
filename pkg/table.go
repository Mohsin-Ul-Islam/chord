package chord

type FingerTableEntry struct {
	Node Node
}

type FingerTable struct {
	Entries []FingerTableEntry
}

func NewFingerTable(size uint64) *FingerTable {
	return &FingerTable{Entries: make([]FingerTableEntry, size)}
}
