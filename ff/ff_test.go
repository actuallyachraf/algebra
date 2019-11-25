package ff

import (
	"math/big"
	"testing"

	"github.com/actuallyachraf/algebra/nt"
)

var F31 = NewFiniteField(nt.FromInt64(31))

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

func TestField(t *testing.T) {

	t.Run("TestAdd", func(t *testing.T) {
		tests := [][]int64{
			{2, 15, 17},
			{17, 21, 7},
		}
		for _, test := range tests {
			a := F31.newFieldElementFromInt64(test[0])
			b := F31.newFieldElementFromInt64(test[1])
			actual := F31.Add(a, b)
			expected := F31.newFieldElementFromInt64(test[2])
			assertEqual(actual, expected, t)
		}
	})

	t.Run("TestSub", func(t *testing.T) {
		tests := [][]int64{
			{29, 4, 25},
			{15, 30, 16},
		}
		for _, test := range tests {
			a := F31.newFieldElementFromInt64(test[0])
			b := F31.newFieldElementFromInt64(test[1])
			actual := F31.Sub(a, b)
			expected := F31.newFieldElementFromInt64(test[2])
			assertEqual(actual, expected, t)
		}
	})

	t.Run("TestMul", func(t *testing.T) {
		tests := [][]int64{
			{24, 19, 22},
		}
		for _, test := range tests {
			a := F31.newFieldElementFromInt64(test[0])
			b := F31.newFieldElementFromInt64(test[1])
			actual := F31.Mul(a, b)
			expected := F31.newFieldElementFromInt64(test[2])
			assertEqual(actual, expected, t)
		}
	})

	t.Run("TestDiv", func(t *testing.T) {
		tests := [][]int64{
			{3, 24, 1, 1, 4},
			{17, 1, -3, 1, 29},
			{4, 1, -4, 11, 13},
		}
		for _, test := range tests {
			a := F31.newFieldElementFromInt64(test[0])
			b := F31.newFieldElementFromInt64(test[1])
			c := big.NewInt(test[2])
			d := F31.newFieldElementFromInt64(test[3])
			expected := F31.newFieldElementFromInt64(test[4])
			actual := F31.Div(a, b)
			actual = actual.Exp(c)
			actual = F31.Mul(actual, d)
			assertEqual(actual, expected, t)
		}
	})
}

func assertEqual(a FieldElement, b FieldElement, t *testing.T) {
	if a.p.Cmp(a, b) != 0 {
		t.Errorf("%v is not equal to %v\n", a, b)
	}
}
