package schnorr

import (
	"crypto/sha256"
	"math/big"

	"github.com/actuallyachraf/algebra/ec"
)

const (
	bitLength = 512
)

// PublicKey represents a point on the underlying curve
type PublicKey struct {
	P *ec.Point
}

// PrivateKey represents a scalar in Fp
type PrivateKey struct {
	K *big.Int
}

// HashToPoint computes a message fingerprint H(m||P) where m is the message and P the public
// key.
func HashToPoint(message []byte, Pubkey PublicKey) *big.Int {

	var b = make([]byte, 0, len(message))
	copy(b, message)

	b = append(b, Pubkey.P.X.Bytes()...)
	b = append(b, Pubkey.P.Y.Bytes()...)

	hasher := sha256.New()
	hasher.Write(b)

	hash := hasher.Sum(nil)

	r := new(big.Int).SetBytes(hash)

	return r
}

// Params stores the schnorr parameters
type Params struct {
	EC    ec.Curve // Underlying curve group
	Gen   ec.Point // Group generator
	Order int      // Curve order
}
