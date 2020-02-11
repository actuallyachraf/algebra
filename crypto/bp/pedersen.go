package bp

import (
	"crypto/rand"

	"github.com/actuallyachraf/algebra/ec"
	"github.com/actuallyachraf/algebra/nt"
)

// PedersenCommitment implements the pedersen commitment scheme over elliptic curves.
func PedersenCommitment(params *Parameters, value *nt.Integer) (*ec.Point, *nt.Integer) {

	hidingVal, _ := rand.Int(rand.Reader, params.L)
	valueReduced := nt.Mod(value, params.L)
	commitment := params.EC.Add(params.EC.ScalarMul(params.G, valueReduced), params.EC.ScalarMul(params.H, hidingVal))
	return commitment, hidingVal
}
func pedersenCom(params *Parameters, G, H *ec.Point, value, hiding *nt.Integer) *ec.Point {

	com := params.EC.Add(params.EC.ScalarMul(G, value), params.EC.ScalarMul(H, hiding))
	return com
}

// VectorPedersenCommitment implements the pedersen commitment scheme for vectors
func VectorPedersenCommitment(params *Parameters, values Vector) (*ec.Point, Vector) {

	R := NewZeroVector(params.N)

	commitment := &ec.Point{X: nt.FromInt64(0), Y: nt.FromInt64(0)}

	for i := 0; i < values.Len(); i++ {

		hidingVal, _ := rand.Int(rand.Reader, params.L)
		valueReduced := nt.Mod(values[i], params.L)
		com := params.EC.Add(params.EC.ScalarMul(params.GVec[i], valueReduced), params.EC.ScalarMul(params.HVec[i], hidingVal))

		commitment = params.EC.Add(commitment, com)
	}
	return commitment, R
}

// DoubleVectorPedersenCommitment commit to two vectors a and b
// where b acts as the hiding parameter
func DoubleVectorPedersenCommitment(params *Parameters, a, b Vector) *ec.Point {
	// com = G[i]*a[i]+H[i]*b[i]

	commitment := &ec.Point{X: nt.FromInt64(0), Y: nt.FromInt64(0)}

	if a.Len() != b.Len() {
		return commitment
	}

	for i := 0; i < a.Len(); i++ {
		com := params.EC.Add(params.EC.ScalarMul(params.GVec[i], a[i]), params.EC.ScalarMul(params.HVec[i], b[i]))
		commitment = params.EC.Add(commitment, com)
	}

	return commitment
}

// DoubleVectorPedersenCommitmentWithGen commit to two vectors a and b
// where b acts as the hiding parameter with given generators
func DoubleVectorPedersenCommitmentWithGen(params *Parameters, GVec, HVec []*ec.Point, a, b Vector) *ec.Point {
	// com = G[i]*a[i]+H[i]*b[i]

	commitment := &ec.Point{X: nt.FromInt64(0), Y: nt.FromInt64(0)}

	if a.Len() != b.Len() {
		return commitment
	}

	for i := 0; i < a.Len(); i++ {
		com := params.EC.Add(params.EC.ScalarMul(GVec[i], a[i]), params.EC.ScalarMul(HVec[i], b[i]))
		commitment = params.EC.Add(commitment, com)
	}

	return commitment
}
