package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("fo bar baz")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))

	// test with invalid messge
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// test with invalid pub & priv key
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, []byte("foo")))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey() // generate private key
	pubKey := privKey.Public()      // generate public key from private key
	address := pubKey.Address()     // generates address from public key

	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
