package types

import (
	"crypto/sha256"

	"github.com/obynonwane/blockchain_project/proto"
	pb "google.golang.org/protobuf/proto"
)

// hashing a block means hashing only the header without the
// transaction
func HashBlock(block *proto.Block) []byte {
	// marshal the supplied block
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}

	// convert to hash - returns an array
	hash := sha256.Sum256(b)

	//converts the array to slice
	return hash[:]

}
