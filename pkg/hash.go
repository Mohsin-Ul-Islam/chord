package chord

// HashValue is the type of value returned by the hash function
type HashValue string

// Hasher is an interface for a deterministic hash function of a distributed hash table
type Hasher interface {

	// Hash is a deterministic function to map a peer's ipv4 address
	// and port to some value in a chord's distributed hash table
	//
	// The hash function implementation should be hermetic
	// i.e it should return the same HashValue for the same
	// ipv4Addr and port all the time
	Hash(ipv4Addr string, port string) HashValue
}

// simpleHasher
type simpleHasher struct {
}

// NewSimpleHasher
func NewSimpleHasher() *simpleHasher {
    return &simpleHasher{}
}

func (h *simpleHasher) Hash(ipv4Addr string, port string) HashValue {
    return HashValue(ipv4Addr)
}
