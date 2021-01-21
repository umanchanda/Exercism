package diffiehellman

import (
	"math/big"
	"math/rand"
	"math"
)

// PrivateKey generates a private key
func PrivateKey(p *big.Int) *big.Int {
	return rand.Intn(p)
}

// PublicKey generates public key
func PublicKey(a, p *big.Int, g int64) *big.Int {
	return math.Pow(g, a) % p
}

// SecretKey returns secret key
func SecretKey(a, B, p *big.Int) *big.Int {
	return math.Pow(B, a) % p
}

// NewPair returns a new pair
func NewPair(p *big.Int, g int64) {

}