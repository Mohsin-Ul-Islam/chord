package chord

// Node is an interface for the behaviour of a peer in chord
type Node interface {
	// Join is a method to join another peer in a chord
	Join(peer Node) error

	// Joined is a method to inform peers that a certain peer joined the chord
	Joined(peer Node) error

	// Leave is a method to leave the current chord
	Leave() error

	// Left is a method to inform peers that a certain peer left the chord
	Left(peer Node) error
}

// node is a structure to represent a peer in chord
type node struct {
	// Prev points to the previous peer in the chord clockwise
	Prev Node

	// Next points to the next peer in the chord clockwise
	Next Node

	// Port is the listening port of peer
	Port string

	// Ipv4Addr is the ipv4 listening address of peer
	Ipv4Addr string

	// FingerTable is the data structure to keep track of log(N) peers in the chord
	FingerTable *FingerTable

	// Hasher is an interface to map a peer to chord
	Hasher Hasher
}

// NewNode is a constructor for creating nodes data structures
func NewNode(ipv4Addr string, port string, size uint64, hasher Hasher) *node {
	return &node{
		Port:        port,
		Hasher:      hasher,
		Ipv4Addr:    ipv4Addr,
		FingerTable: NewFingerTable(size),
	}
}

// Compare is the implementation of golang's Comparer interface for a Node
func (n *node) Compare(na *node) int {
	hash := n.Hasher.Hash(n.Ipv4Addr, n.Port)
	peerHash := n.Hasher.Hash(na.Ipv4Addr, na.Port)

	if hash == peerHash {
		return 0
	} else if hash > peerHash {
		return 1
	} else {
		return -1
	}
}

func (n *node) Join(peer Node) error {
	return nil
}

func (n *node) Joined(peer Node) error {
	return nil
}

func (n *node) Left(peer Node) error {

	// ignore if peer is nil or self
	if peer == nil || peer == n {
		return nil
	}

	// remove if peer is your next
	if n.Next == peer {
		n.Next = nil
	}

	// remove if peer is your previous
	if n.Prev == peer {
		n.Prev = nil
	}

	// remove if peer is in your fingertable
	for _, entry := range n.FingerTable.Entries {
		if entry.Node == peer {
			entry.Node = nil
		}
	}

	return nil
}

func (n *node) Leave() error {

	// inform your next peer that you left
	if n.Next != nil {
		err := n.Next.Left(n)
		if err != nil {
			// log error
		}
	}

	// inform your previous peer that you left
	if n.Prev != nil {
		err := n.Prev.Left(n)
		if err != nil {
			// log error
		}
	}

	// inform the rest of the entries in your finger table that you left
	for _, entry := range n.FingerTable.Entries {
		if entry.Node != nil {
			err := entry.Node.Left(n)
			if err != nil {
				// log error
			}
		}
	}

	return nil
}
