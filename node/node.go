package node

import (
	"context"
	"fmt"

	"github.com/obynonwane/blockchain_project/proto"
	"google.golang.org/grpc/peer"
)

// instance of Node struct
type Node struct {
	proto.UnimplementedNodeServer
}

// constructor function
func NewNode() *Node {
	return &Node{}
}
func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {

	peer, _ := peer.FromContext(ctx) // extract information about the client making the call - ie peer or node
	fmt.Println("received tx from", peer)

	return &proto.Ack{}, nil
}
