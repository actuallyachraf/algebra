package bp

import (
	"github.com/actuallyachraf/algebra/ec"
	"github.com/actuallyachraf/algebra/ff"
	"github.com/actuallyachraf/algebra/nt"
	"golang.org/x/crypto/sha3"
)

// Bulletproofs don't have a trusted setup 	but a set of public parameters
// shared by all proofs.
var (
	// a string used to supplement hashing to curve points
	algebraBulletProofParameter = []byte("algebra-does-bulletproofs")
)

// Parameters define the public parameters of the proof system
// L : Order of the subgroup of points on the defined elliptic curve
// N : Bitlength of the values we want to commit to i.e v in [0,2^n]
// M : Bounded number of aggregate proofs
// G : Generator of the subgroup of curve points
// H : The nothing up my sleeve second generator whose discrete log w.r.t to G is unkown
type Parameters struct {
	EC   *ec.Curve
	G    *ec.Point
	H    *ec.Point
	U    *ec.Point
	L    *nt.Integer
	M    int
	N    int
	GVec []*ec.Point
	HVec []*ec.Point
}

// GenParametersSecp256k1 generates bulletproof parameters using the curve
// secp256k1 it takes as parameters the bitlength of the integer range we want
// to construct proofs for.
func GenParametersSecp256k1(bitlength int) *Parameters {

	var (
		secp256k1GeneratorX, _     = new(nt.Integer).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
		secp256k1GeneratorY, _     = new(nt.Integer).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
		secp256k1GeneratorOrder, _ = new(nt.Integer).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
		fieldOrder, _              = new(nt.Integer).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)
		secp256k1A                 = new(nt.Integer).SetInt64(0)
		secp256k1B                 = new(nt.Integer).SetInt64(7)
	)

	// Given the first parameters we build the elliptic curve
	Fq, _ := ff.NewFiniteField(fieldOrder)
	secp256k1 := ec.NewEllipticCurve(Fq.NewFieldElement(secp256k1A), Fq.NewFieldElement(secp256k1B), Fq)
	secp256k1BasePoint := ec.Point{X: secp256k1GeneratorX, Y: secp256k1GeneratorY}
	// The subgroup of elliptic curve points is of prime order we pick H by Hashing the base point
	// Algorithm is the naive hash2point other specific algorithms that depend
	// on the curve properties can be used ref : https://eprint.iacr.org/2009/226.pdf
	hash2Point := func(hash []byte) *ec.Point {
		var success = false
		var point = new(ec.Point)
		// SetBytes interprets the hash as a big-endian integer
		var x = new(nt.Integer).SetBytes(hash)
		for !success {
			x = nt.Mod(x, fieldOrder)
			// Use x to compute y if y doesn't have a square root add one and repeat.
			p, err := secp256k1.At(x)
			if err != nil {
				x = nt.Add(x, nt.One)
			} else {
				point.X = p.X
				point.Y = p.Y
				success = true
			}
		}

		return point
	}
	// useful callbacks
	hashedG := sha3.Sum256(secp256k1GeneratorX.Bytes())
	concat := func(a ...[]byte) []byte {
		c := make([]byte, 0, len(a)*2)

		for _, v := range a {
			c = append(c, v...)
		}
		return c
	}

	H := hash2Point(hashedG[:])

	// Constructing the generator vectors
	// We will support aggregating a maximum of 16 proofs
	M := 16
	VectorG := make([]*ec.Point, M*bitlength)
	VectorH := make([]*ec.Point, M*bitlength)

	var i int64
	for i = 0; i < int64(M*bitlength); i++ {
		// Indexes offer a bit of impredictability
		i1AsBig := nt.FromInt64(i * 2)
		i2AsBig := nt.FromInt64(i*2 + 1)

		hashedH := sha3.Sum256(concat(H.Bytes(), algebraBulletProofParameter, i1AsBig.Bytes()))
		hashedG := sha3.Sum256(concat(H.Bytes(), algebraBulletProofParameter, i2AsBig.Bytes()))

		VectorG[i] = hash2Point(hashedG[:])
		VectorH[i] = hash2Point(hashedH[:])
	}
	var U = &ec.Point{
		X: new(nt.Integer).SetBytes(H.X.Bytes()),
		Y: new(nt.Integer).SetBytes(H.Y.Bytes()),
	}
	return &Parameters{
		EC:   secp256k1,
		G:    &secp256k1BasePoint,
		L:    secp256k1GeneratorOrder,
		H:    H,
		U:    U,
		N:    64,
		GVec: VectorG,
		HVec: VectorH,
	}
}
