package types

import (
	"crypto/sha256"

	"github.com/obynonwane/blockchain_project/crypto"

	"github.com/obynonwane/blockchain_project/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	// return the signed hash of the txn
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}

	// hash th transaction - which is returned as an array
	hash := sha256.Sum256(b)

	// return the hash as a slice of bytes
	return hash[:]

}
