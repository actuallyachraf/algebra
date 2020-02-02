package schnorr

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/actuallyachraf/algebra/ec"
	"github.com/actuallyachraf/algebra/nt"
)

// PublicKey represents a point on the underlying curve
type PublicKey struct {
	P *ec.Point
}

// PrivateKey represents a scalar in Fp
type PrivateKey struct {
	K *nt.Integer
}

// Keypair represents a private key and public key
type Keypair struct {
	PublicKey
	PrivateKey
}

// Signature represents a schnorr signature which is a curve point and a scalar
type Signature struct {
	R *nt.Integer
	S *nt.Integer
}

// Params stores the schnorr parameters
// An elliptic curve over a field of order q generates a group of unknown order
// the number of points i.e the order of the group can be computed trough
// point counting algorithms, there exist an upperbound given by Hasse's theorem.
type Params struct {
	EC    ec.Curve    // Underlying curve group
	Gen   ec.Point    // Group generator
	Order *nt.Integer // Group order
}

// GenerateKeypair generates a keypair
func GenerateKeypair(params Params) Keypair {

	order := new(nt.Integer).Set(params.Order)

	sk, _ := rand.Int(rand.Reader, order)
	q := params.EC.ScalarMul(&params.Gen, sk)
	return Keypair{
		PublicKey:  PublicKey{P: q},
		PrivateKey: PrivateKey{sk},
	}
}

// HashToPoint computes a message fingerprint H(m||P.X) where m is the message and P the public
// key.
func HashToPoint(message []byte, P *ec.Point) *nt.Integer {

	var b = make([]byte, 0)

	b = append(b, message...)
	b = append(b, P.X.Bytes()...)

	hasher := sha256.New()
	hasher.Write(b)

	hash := hasher.Sum(nil)

	r := new(nt.Integer).SetBytes(hash)

	return r
}

// Sign a message given a keypair
func Sign(message []byte, params *Params, kp Keypair) Signature {
	// underlying curve field
	order := new(nt.Integer).Set(params.Order)
	s := new(nt.Integer).SetInt64(0)
	R := new(nt.Integer).SetInt64(0)
	for s.Cmp(nt.Zero) == 0 {
		k, _ := rand.Int(rand.Reader, order)
		Q := params.EC.ScalarMul(&params.Gen, k)
		R = HashToPoint(message, Q)
		rk := new(nt.Integer).Mul(R, kp.K)
		s = new(nt.Integer).Sub(k, rk)
		s = new(nt.Integer).Mod(s, order)
	}
	return Signature{R, s}
}

// Verify a message given a signature, message and a public key
func Verify(message []byte, sig Signature, kp Keypair, params Params) bool {

	rP := params.EC.ScalarMul(kp.P, sig.R)
	sG := params.EC.ScalarMul(&params.Gen, sig.S)
	Q := params.EC.Add(rP, sG)
	v := HashToPoint(message, Q)
	if v.Cmp(sig.R) != 0 {
		fmt.Println(v, sig.R)
		return false
	}
	return true
}
