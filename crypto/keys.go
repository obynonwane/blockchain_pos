package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	privKeyLen = 64 // private key & public key 32 each, public key attached to private key
	pubKeyLen  = 32
	seedLen    = 32
	addressLen = 20
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
func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
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

// returns a reference to the address struct
func (p *PublicKey) Address() Address {
	// extracts the last 20 bytes of the publick key and returns it
	return Address{
		value: p.key[len(p.key)-addressLen:],
	}
}

type Signature struct {
	value []byte
}

// return signature bytes
func (s *Signature) Bytes() []byte {
	return s.value
}

// verify if signature is valid of a message
func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.key, msg, s.value)
}

// address struct
type Address struct {
	value []byte
}

// returns the address as bytes from the struct
func (a Address) Bytes() []byte {
	return a.value
}

// returns the address as a string after converting bytes to string
func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
