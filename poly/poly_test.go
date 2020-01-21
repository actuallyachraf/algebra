package poly

import "testing"

import "github.com/actuallyachraf/algebra/ff"

import "github.com/actuallyachraf/algebra/nt"

func TestPolynomial(t *testing.T) {
	// (X+1) * (X+1) == X**2 + 2*X + 1
	t.Run("TestMul", func(t *testing.T) {
		F31 := ff.NewFiniteField(nt.FromInt64(31))
		F251 := ff.NewFiniteField(nt.FromInt64(251))

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
	t.Run("TestAdd", func(t *testing.T) {
		F251 := ff.NewFiniteField(nt.FromInt64(251))
		tests := [][]ff.FieldElement{
			{F251.NewFieldElementFromInt64(4), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(123)},
			{F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(225), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(196)},
			{F251.NewFieldElementFromInt64(80), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(50), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(68)},
		}
		p := New(tests[0], 4)
		q := New(tests[1], 4)
		actual := Add(p, q)
		expected := New(tests[2], 4)
		if !actual.Equal(expected) {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Add error : expected %v got %v :", expected, actual)
		}
	})

}
