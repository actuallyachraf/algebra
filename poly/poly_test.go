package poly

import (
	"testing"

	"github.com/actuallyachraf/algebra/ff"
	"github.com/actuallyachraf/algebra/nt"
	"github.com/stretchr/testify/assert"
)

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

		p := NewPolynomial(tests[0])
		q := NewPolynomial(tests[1])
		actual := p.Mul(q, F31.Modulus())
		expected := NewPolynomial(tests[2])
		if actual.Compare(&expected) != 0 {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Mul error : expected %v got %v :", expected, actual)
		}

		p = NewPolynomial(tests[3])
		q = NewPolynomial(tests[4])
		actual = p.Mul(q, GF7.Modulus())
		expected = NewPolynomial(tests[5])
		if actual.Compare(&expected) != 0 {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Mul error : expected %v got %v :", expected, actual)
		}
	})
	t.Run("TestAddSub", func(t *testing.T) {
		F251, _ := ff.NewFiniteField(nt.FromInt64(251))
		GF7, _ := ff.NewFiniteField(nt.FromInt64(7))

		tests := [][]ff.FieldElement{
			{F251.NewFieldElementFromInt64(4), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(123)},
			{F251.NewFieldElementFromInt64(76), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(225), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(196)},
			{F251.NewFieldElementFromInt64(80), F251.NewFieldElementFromInt64(7), F251.NewFieldElementFromInt64(50), F251.NewFieldElementFromInt64(12), F251.NewFieldElementFromInt64(68)},

			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(4), GF7.NewFieldElementFromInt64(5)},
			{GF7.NewFieldElementFromInt64(1), GF7.NewFieldElementFromInt64(2), GF7.NewFieldElementFromInt64(0)},
			{GF7.NewFieldElementFromInt64(0), GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(5)},
			{GF7.NewFieldElementFromInt64(5), GF7.NewFieldElementFromInt64(2), GF7.NewFieldElementFromInt64(5)},
		}
		p := NewPolynomial(tests[0])
		q := NewPolynomial(tests[1])
		actual := p.Add(q, F251.Modulus())
		expected := NewPolynomial(tests[2])
		if actual.Compare(&expected) != 0 {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Add error : expected %v got %v :", expected, actual)
		}

		p = NewPolynomial(tests[3])
		q = NewPolynomial(tests[4])
		actual = p.Add(q, GF7.Modulus())
		expected = NewPolynomial(tests[5])
		if actual.Compare(&expected) != 0 {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Add error : expected %v got %v :", expected, actual)
		}

		actual = p.Sub(q, GF7.Modulus())
		expected = NewPolynomial(tests[6])
		if actual.Compare(&expected) != 0 {
			t.Log("Expected Deg : ", expected.Degree(), " Actual", actual.Degree())
			t.Errorf("Add error : expected %v got %v :", expected, actual)
		}

	})

	t.Run("TestDiv", func(t *testing.T) {
		F251, _ := ff.NewFiniteField(nt.FromInt64(251))
		tests := [][]ff.FieldElement{
			{F251.NewFieldElementFromInt64(5), F251.NewFieldElementFromInt64(1), F251.NewFieldElementFromInt64(0), F251.NewFieldElementFromInt64(2), F251.NewFieldElementFromInt64(3)},
			{F251.NewFieldElementFromInt64(3), F251.NewFieldElementFromInt64(2), F251.NewFieldElementFromInt64(1)},
		}

		p := NewPolynomial(tests[0])
		q := NewPolynomial(tests[1])
		actualQ, actualR := p.Div(q, F251.Modulus())

		t.Log("Quotient : ", actualQ)
		t.Log("Remainder : ", actualR)

		GF7, _ := ff.NewFiniteField(nt.FromInt64(7))
		GF11, _ := ff.NewFiniteField(nt.FromInt64(11))

		/*
			if !ZeroPolynomial(GF7).Equal(ZeroPolynomial(GF7)) {
				t.Error("Zero polynomials aren't equal")
			}
		*/
		// GF(11) (3x6 + 7x4 + 4x3 + 5) ÷ (x4 + 3x3 + 4) = 3x2 + 3x + 3 with remainder x3 + 10x2 + 4x +1
		// GF(7) (5x2 + 4x + 6)  ÷ (2x+1) = 6x + 6 with remainder 0
		tests = [][]ff.FieldElement{
			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(4), GF7.NewFieldElementFromInt64(5)},
			{GF7.NewFieldElementFromInt64(1), GF7.NewFieldElementFromInt64(2)},
			{GF7.NewFieldElementFromInt64(6), GF7.NewFieldElementFromInt64(6)},
			{GF7.NewFieldElementFromInt64(0)},

			{GF11.NewFieldElementFromInt64(5), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(4), GF11.NewFieldElementFromInt64(7), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(3)},
			{GF11.NewFieldElementFromInt64(4), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(0), GF11.NewFieldElementFromInt64(3), GF11.NewFieldElementFromInt64(1)},
			{GF11.NewFieldElementFromInt64(1), GF11.NewFieldElementFromInt64(2), GF11.NewFieldElementFromInt64(3)},
			{GF11.NewFieldElementFromInt64(1), GF11.NewFieldElementFromInt64(3), GF11.NewFieldElementFromInt64(10), GF11.NewFieldElementFromInt64(1)},
		}

		p = NewPolynomial(tests[0])
		q = NewPolynomial(tests[1])
		actualQ, actualR = p.Div(q, GF7.Modulus())
		expectedQ, expectedR := NewPolynomial(tests[2]), NewPolynomial(tests[3])

		if actualQ.Compare(&expectedQ) != 0 || actualR.Compare(&expectedR) != 0 {
			t.Log("Expected Deg : ", expectedQ.Degree(), " Actual", actualQ.Degree())
			t.Errorf("Div error : expected quotient %v got quotient %v | expected remainder %v got remainder %v ", expectedQ, actualQ, expectedR, actualR)
		}
		p = NewPolynomial(tests[4])
		q = NewPolynomial(tests[5])
		actualQ, actualR = p.Div(q, GF11.Modulus())
		expectedQ, expectedR = NewPolynomial(tests[6]), NewPolynomial(tests[7])
		if actualQ.Compare(&expectedQ) != 0 || actualR.Compare(&expectedR) != 0 {
			t.Log("Expected Deg : ", expectedQ.Degree(), " Actual", actualQ.Degree())
			t.Errorf("Div error : expected quotient %v got quotient %v | expected remainder %v got remainder %v ", expectedQ, actualQ, expectedR, actualR)
		}

	})
	t.Run("TestCompose", func(t *testing.T) {
		f := NewPolynomialInts(0, 1, 1)
		g := NewPolynomialInts(1, 1)

		assert.Equal(t, f.Compose(g, nt.FromInt64(233)), NewPolynomialInts(2, 3, 1))
	})
	t.Run("TestPow", func(t *testing.T) {
		f := NewPolynomialInts(1, 1)
		fSquared := f.Pow(nt.FromInt64(2), nt.FromInt64(233))
		assert.Equal(t, fSquared, NewPolynomialInts(1, 2, 1))
	})
	t.Run("TestMonomial", func(t *testing.T) {
		f := NewPolynomialInts(0, 1)
		f = f.Clone(1023).Add(NewPolynomialInts(1).Neg(), nil)
		t.Log(f)
	})
}
