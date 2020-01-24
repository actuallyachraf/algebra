package poly

import "testing"

import "github.com/actuallyachraf/algebra/ff"

import "github.com/actuallyachraf/algebra/nt"

func TestPolynomial(t *testing.T) {
	// (X+1) * (X+1) == X**2 + 2*X + 1
	t.Run("TestMul", func(t *testing.T) {
		F31, _ := ff.NewFiniteField(nt.FromInt64(31))

		GF7, _ := ff.NewFiniteField(nt.FromInt64(7))

		tests := [][]ff.FieldElement{
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(1)},
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(1)},
			{F31.NewFieldElementFromInt64(1), F31.NewFieldElementFromInt64(2), F31.NewFieldElementFromInt64(1)},

			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(4), GF7.NewFieldElementFromInt64(5)},
			{GF7.NewFieldElementFromInt64(1), GF7.NewFieldElementFromInt64(2)},
			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(2), GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(3)},
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
		GF7, _ := ff.NewFiniteField(nt.FromInt64(7))

		tests := [][]ff.FieldElement{
			{F251.NewFieldElementFromInt64(4), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(123)},
			{F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(225), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(196)},
			{F251.NewFieldElementFromInt64(80), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(50), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(68)},

			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(4), GF7.NewFieldElementFromInt64(5)},
			{GF7.NewFieldElementFromInt64(1), GF7.NewFieldElementFromInt64(2), GF7.NewFieldElementFromInt64(0)},
			{GF7.NewFieldElementFromInt64(0), GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(5)},
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
		r := Add(p, ZeroPolynomial(GF7))
		if !r.Equal(p) {
			t.Errorf("Add error : expected %v got %v", p, r)
		}
	})

	t.Run("TestDiv", func(t *testing.T) {
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

		GF7, _ := ff.NewFiniteField(nt.FromInt64(7))
		GF11, _ := ff.NewFiniteField(nt.FromInt64(11))

		if !ZeroPolynomial(GF7).Equal(ZeroPolynomial(GF7)) {
			t.Error("Zero polynomials aren't equal")
		}

		// GF(11) (3x6 + 7x4 + 4x3 + 5) ÷ (x4 + 3x3 + 4) = 3x2 + 3x + 3 with remainder x3 + 10x2 + 4x +1
		// GF(7) (5x2 + 4x + 6)  ÷ (2x+1) = 6x + 6 with remainder 0
		tests = [][]ff.FieldElement{
			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(4), GF7.NewFieldElementFromInt64(5)},
			{GF7.NewFieldElementFromInt64(1), GF7.NewFieldElementFromInt64(2)},
			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(6)},
			{GF7.NewFieldElementFromInt64(0)},

			{GF11.NewFieldElementFromInt64(5), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(4), GF11.NewFieldElementFromInt64(7), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(3)},
			{GF11.NewFieldElementFromInt64(4), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(3), GF11.NewFieldElementFromInt64(1)},
			{GF11.NewFieldElementFromInt64(3), GF11.NewFieldElementFromInt64(3), GF11.NewFieldElementFromInt64(3)},
			{GF11.NewFieldElementFromInt64(1), GF11.NewFieldElementFromInt64(4), GF11.NewFieldElementFromInt64(10), GF11.NewFieldElementFromInt64(1)},
		}

		p = New(tests[0], 2)
		q = New(tests[1], 1)
		actualQ, actualR = QuoRem(p, q)
		expectedQ, expectedR := New(tests[2], 1), New(tests[3], 0)

		if !actualQ.Equal(expectedQ) || !actualR.Equal(expectedR) {
			t.Log("Expected Deg : ", expectedQ.Degree(), " Actual", actualQ.Degree())
			t.Errorf("Div error : expected quotient %v got quotient %v | expected remainder %v got remainder %v ", expectedQ, actualQ, expectedR, actualR)
		}
		p = New(tests[4], 6)
		q = New(tests[5], 4)
		actualQ, actualR = QuoRem(p, q)
		expectedQ, expectedR = New(tests[6], 2), New(tests[7], 3)
		if !actualQ.Equal(expectedQ) || !actualR.Equal(expectedR) {
			t.Log("Expected Deg : ", expectedQ.Degree(), " Actual", actualQ.Degree())
			t.Errorf("Div error : expected quotient %v got quotient %v | expected remainder %v got remainder %v ", expectedQ, actualQ, expectedR, actualR)
		}

	})

}
