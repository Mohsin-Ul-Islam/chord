package v1

import (
    "chord/pkg"
	chordpbv1 "chord/proto/v1"
)

type nodeService struct {
	chordpbv1.UnimplementedNodeServiceServer

    node chord.Node
}

func NewNodeService(node chord.Node) *nodeService {
    return &nodeService{node: node}
}

