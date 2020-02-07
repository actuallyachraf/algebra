package bp

import "github.com/actuallyachraf/algebra/ec"

import "github.com/actuallyachraf/algebra/nt"

// PedersenCommitment implements the pedersen commitment scheme over elliptic curves.
func PedersenCommitment(params Parameters, value *nt.Integer) (*ec.Point, *nt.Integer) {

	hidingVal, _ := params.EC.F.Rand()
	commitment := params.EC.DoubleScalarMult(&params.G, &params.H, value, hidingVal.Big())

	return commitment, hidingVal.Big()
}
