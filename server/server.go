package server

import (
	"context"

	"github.com/obynonwane/blockchain_project/proto"
)

type Node struct {
	proto.UnimplementedNodeServer
}

func HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.None, error) {
	return nil, nil
}
