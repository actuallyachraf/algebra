package poly

import "testing"

import "github.com/actuallyachraf/algebra/ff"

import "github.com/actuallyachraf/algebra/nt"

func TestPolynomial(t *testing.T) {
	// (X+1) * (X+1) == X**2 + 2*X + 1
	t.Run("TestMul", func(t *testing.T) {
		F31 := ff.NewFiniteField(nt.FromInt64(31))
		tests := [][]ff.FieldElement{
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(1)},
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(1)},
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(2), F31.NewFieldElementFromInt64(1)},
		}

		p := New(tests[0], 1)
		q := New(tests[0], 1)
		actual := Mul(p, q)
		expected := New(tests[2], 2)
		if !actual.Equal(expected) {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Mul error : expected %v got %v :", expected, actual)
		}

	})
	/*
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
	*/
}
