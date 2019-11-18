package ff

import (
	"testing"

	"github.com/actuallyachraf/algebra/nt"
)

func TestFF(t *testing.T) {

	var two = new(nt.Integer).SetInt64(2)

	// GF(2)
	var gf2 = NewFiniteField(two)

	var x = gf2.NewFieldElement(nt.One)
	var y = gf2.NewFieldElement(nt.One)

	// XOR
	if gf2.Cmp(gf2.Add(x, y), gf2.Zero()) != 0 {
		t.Error("failed to add two field elements")
	}
	if gf2.Cmp(gf2.Mul(x, gf2.Zero()), gf2.Zero()) != 0 {
		t.Error("failed to mul two field elements")
	}

}
