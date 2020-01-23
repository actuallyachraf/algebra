package poly

import "testing"

import "github.com/actuallyachraf/algebra/ff"

import "github.com/actuallyachraf/algebra/nt"

func TestPolynomial(t *testing.T) {
	// (X+1) * (X+1) == X**2 + 2*X + 1
	t.Run("TestMul", func(t *testing.T) {
		F31, _ := ff.NewFiniteField(nt.FromInt64(31))

		F7, _ := ff.NewFiniteField(nt.FromInt64(7))

		tests := [][]ff.FieldElement{
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(1)},
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(1)},
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(2), F31.NewFieldElementFromInt64(1)},

			{F7.NewFieldElementFromInt64(6), F7.NewFieldElementFromInt64(4), F7.NewFieldElementFromInt64(5)},
			{F7.NewFieldElementFromInt64(1), F7.NewFieldElementFromInt64(2)},
			{F7.NewFieldElementFromInt64(6), F7.NewFieldElementFromInt64(2), F7.NewFieldElementFromInt64(6), F7.NewFieldElementFromInt64(3)},
		}

		p := New(tests[0], 1)
		q := New(tests[1], 1)
		actual := Mul(p, q)
		expected := New(tests[2], 2)
		if !actual.Equal(expected) {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Mul error : expected %v got %v :", expected, actual)
		}

		p = New(tests[3], 2)
		q = New(tests[4], 1)
		actual = Mul(p, q)
		expected = New(tests[5], 3)
		if !actual.Equal(expected) {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Mul error : expected %v got %v :", expected, actual)
		}
	})
	t.Run("TestAdd", func(t *testing.T) {
		F251, _ := ff.NewFiniteField(nt.FromInt64(251))
		F7, _ := ff.NewFiniteField(nt.FromInt64(7))

		tests := [][]ff.FieldElement{
			{F251.NewFieldElementFromInt64(4), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(123)},
			{F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(225), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(196)},
			{F251.NewFieldElementFromInt64(80), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(50), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(68)},

			{F7.NewFieldElementFromInt64(6), F7.NewFieldElementFromInt64(4), F7.NewFieldElementFromInt64(5)},
			{F7.NewFieldElementFromInt64(1), F7.NewFieldElementFromInt64(2), F7.NewFieldElementFromInt64(0)},
			{F7.NewFieldElementFromInt64(0), F7.NewFieldElementFromInt64(6), F7.NewFieldElementFromInt64(5)},
		}
		p := New(tests[0], 4)
		q := New(tests[1], 4)
		actual := Add(p, q)
		expected := New(tests[2], 4)
		if !actual.Equal(expected) {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Add error : expected %v got %v :", expected, actual)
		}

		p = New(tests[3], 2)
		q = New(tests[4], 1)
		actual = Add(p, q)
		expected = New(tests[5], 2)
		if !actual.Equal(expected) {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Add error : expected %v got %v :", expected, actual)
		}
	})

	t.Run("TestQuoRem", func(t *testing.T) {
		F251, _ := ff.NewFiniteField(nt.FromInt64(251))
		tests := [][]ff.FieldElement{
			{F251.NewFieldElementFromInt64(5), F251.NewFieldElementFromInt64(1), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(2), F251.NewFieldElementFromInt64(3)},
			{F251.NewFieldElementFromInt64(3), F251.NewFieldElementFromInt64(2), F251.NewFieldElementFromInt64(1)},
		}

		p := New(tests[0], 4)
		q := New(tests[1], 2)
		actualQ, actualR := QuoRem(p, q)

		t.Log("Quotient : ", actualQ)
		t.Log("Remainder : ", actualR)
	})

}
