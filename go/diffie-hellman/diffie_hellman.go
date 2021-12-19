package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey generates a private key
func PrivateKey(p *big.Int) *big.Int {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	return p.Rand(r, p)
}

// PublicKey generates public key
func PublicKey(a, p *big.Int, g int64) *big.Int {

}

// SecretKey returns secret key
func SecretKey(a, B, p *big.Int) *big.Int {
	// return math.Pow(B, a) % p
	return nil
}

// NewPair returns a new pair
func NewPair(p *big.Int, g int64) {
	return
}
