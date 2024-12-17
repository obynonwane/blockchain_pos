package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

const (
	privKeyLen = 64 // private key & public key 32 each, public key attached to private key
	pubKeyLen  = 32
	seedLen    = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

// generates private key
func GeneratePrivateKey() *PrivateKey {

	// make a buffer of seed of length 32
	seed := make([]byte, seedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

// returns private key slice of bytes
func (p *PrivateKey) Bytes() []byte {
	return p.key
}

// signs transaction
func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

// generate public key from private key
func (p *PrivateKey) Public() *PublicKey {

	// make b (buffer) of slice of length pubKeyLen = 32
	b := make([]byte, pubKeyLen)

	// copy from p.key from the 33rd byte into b
	copy(b, p.key[32:])

	return &PublicKey{
		key: b,
	}
}

// returns bublic key slice of bytes
func (p *PublicKey) Bytes() []byte {
	return p.key
}

type PublicKey struct {
	key ed25519.PublicKey
}
