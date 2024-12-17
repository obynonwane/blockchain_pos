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

func TestNewPrivateKeyfromString(t *testing.T) {
	// seed := make([]byte, 32)
	// io.ReadFull(rand.Reader, seed)
	// fmt.Println(hex.EncodeToString(seed))

	var (
		seed       = "e2cfe6e22dea6dbbd7c2ca84e4798b8f7f5bfe3bacc68a90805a805858e8f63f"
		privKey    = NewPrivateKeyfromString(seed)
		addressStr = "22de4300626fb92ae5f5eafa67b8ad59b83fc64d"
	)

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())

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
	privKey := GeneratePrivateKey() // generate private key from seed
	pubKey := privKey.Public()      // generate public key from private key
	address := pubKey.Address()     // generates address from public key

	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
