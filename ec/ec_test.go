package ec

import "testing"

import "github.com/actuallyachraf/algebra/ff"

import "github.com/actuallyachraf/algebra/nt"

func TestCurve(t *testing.T) {

	t.Run("TestIsOnCurve", func(t *testing.T) {

		field, _ := ff.NewFiniteField(nt.FromInt64(29))
		a := field.NewFieldElementFromInt64(4)
		b := field.NewFieldElementFromInt64(20)

		curve := NewEllipticCurve(a, b, field)
		// The commented points are those whose x coordinate is the same
		// but y coordinates are different yet they're both on the same curve.
		points := []*Point{
			//	{nt.FromInt64(2), nt.FromInt64(6)},
			//	{nt.FromInt64(4), nt.FromInt64(19)},
			{nt.FromInt64(5), nt.FromInt64(7)},
			//	{nt.FromInt64(5), nt.FromInt64(22)},
			{nt.FromInt64(2), nt.FromInt64(23)},
			{nt.FromInt64(10), nt.FromInt64(25)},
			//	{nt.FromInt64(13), nt.FromInt64(6)},
			{nt.FromInt64(16), nt.FromInt64(27)},
		}

		for _, point := range points {
			if !curve.IsOnCurve(point) {
				t.Error("IsOnCurve failed")
			}
			actual, err := curve.At(point.X)
			if !actual.Equal(point) || err != nil {
				t.Error("At failed with error : ", err, " expected :", point, "got :", actual)
			}
		}
	})

	t.Run("TestPointAdditionAndDoubling", func(t *testing.T) {
		P := &Point{nt.FromInt64(5), nt.FromInt64(22)}
		Q := &Point{nt.FromInt64(16), nt.FromInt64(27)}

		field, _ := ff.NewFiniteField(nt.FromInt64(29))
		a := field.NewFieldElementFromInt64(4)
		b := field.NewFieldElementFromInt64(20)

		curve := NewEllipticCurve(a, b, field)

		expected := &Point{nt.FromInt64(13), nt.FromInt64(6)}
		actual := curve.Add(P, Q)

		if !actual.Equal(expected) {
			t.Error("TestPointAddition failed expected : ", expected, " got :", actual)
		}

		expected = &Point{nt.FromInt64(14), nt.FromInt64(6)}
		actual = curve.Double(P)

		if !actual.Equal(expected) || !curve.IsOnCurve(actual) || !curve.IsOnCurve(expected) {
			t.Error("TestPointDoubling failed expected : ", expected, "got :", actual)
		}

	})

	t.Run("TestScalarMul", func(t *testing.T) {

		field, _ := ff.NewFiniteField(nt.FromInt64(29))
		a := field.NewFieldElementFromInt64(4)
		b := field.NewFieldElementFromInt64(20)

		curve := NewEllipticCurve(a, b, field)

		P := &Point{X: nt.FromInt64(1), Y: nt.FromInt64(5)}

		k := nt.FromInt64(11)

		actual := curve.ScalarMul(P, k)
		expected := &Point{X: nt.FromInt64(10), Y: nt.FromInt64(25)}

		if !actual.Equal(expected) || !curve.IsOnCurve(actual) || !curve.IsOnCurve(expected) {
			t.Error("TestScalarMul failed expected : ", expected, "got :", actual)
		}
	})

	t.Run("TestMulScalarMult", func(t *testing.T) {
		field, _ := ff.NewFiniteField(nt.FromInt64(29))
		a := field.NewFieldElementFromInt64(4)
		b := field.NewFieldElementFromInt64(20)

		curve := NewEllipticCurve(a, b, field)
		testCases := []*Point{
			{nt.FromInt64(0), nt.FromInt64(0)},
			{nt.FromInt64(1), nt.FromInt64(5)},
			{nt.FromInt64(4), nt.FromInt64(19)},
			{nt.FromInt64(20), nt.FromInt64(3)},
			{nt.FromInt64(15), nt.FromInt64(27)},
			{nt.FromInt64(6), nt.FromInt64(12)},
			{nt.FromInt64(17), nt.FromInt64(19)},
			{nt.FromInt64(24), nt.FromInt64(22)},
			{nt.FromInt64(8), nt.FromInt64(10)},
			{nt.FromInt64(14), nt.FromInt64(23)},
			{nt.FromInt64(13), nt.FromInt64(23)},
			{nt.FromInt64(10), nt.FromInt64(25)},
			{nt.FromInt64(19), nt.FromInt64(13)},
			{nt.FromInt64(16), nt.FromInt64(27)},
			{nt.FromInt64(5), nt.FromInt64(22)},
			{nt.FromInt64(3), nt.FromInt64(1)},
			{nt.FromInt64(0), nt.FromInt64(22)},
			{nt.FromInt64(27), nt.FromInt64(2)},
			{nt.FromInt64(2), nt.FromInt64(23)},
			{nt.FromInt64(2), nt.FromInt64(6)},
			{nt.FromInt64(27), nt.FromInt64(27)},
			{nt.FromInt64(0), nt.FromInt64(7)},
			{nt.FromInt64(3), nt.FromInt64(28)},
			{nt.FromInt64(5), nt.FromInt64(7)},
			{nt.FromInt64(16), nt.FromInt64(2)},
			{nt.FromInt64(19), nt.FromInt64(16)},
			{nt.FromInt64(10), nt.FromInt64(4)},
			{nt.FromInt64(13), nt.FromInt64(6)},
			{nt.FromInt64(14), nt.FromInt64(6)},
			{nt.FromInt64(8), nt.FromInt64(19)},
			{nt.FromInt64(24), nt.FromInt64(7)},
			{nt.FromInt64(17), nt.FromInt64(10)},
			{nt.FromInt64(6), nt.FromInt64(17)},
			{nt.FromInt64(15), nt.FromInt64(2)},
			{nt.FromInt64(20), nt.FromInt64(26)},
			{nt.FromInt64(4), nt.FromInt64(10)},
			{nt.FromInt64(1), nt.FromInt64(24)},
		}

		naiveMulScalarMult := func(P, Q *Point, m, n *nt.Integer) *Point {

			mP := curve.ScalarMul(P, m)
			nQ := curve.ScalarMul(Q, n)

			return curve.Add(mP, nQ)
		}
		for i := 0; i < len(testCases)-1; i++ {

			// expected
			m, _ := field.Rand()
			n, _ := field.Rand()

			P := testCases[i]
			Q := testCases[i+1]

			expected := naiveMulScalarMult(P, Q, m.Big(), n.Big())
			actual := curve.DoubleScalarMult(P, Q, m.Big(), n.Big())

			if !expected.Equal(actual) {
				t.Error("DoubleScalarMult failed")
			}

		}
	})

	t.Run("TestGenerator", func(t *testing.T) {

		field, _ := ff.NewFiniteField(nt.FromInt64(29))
		a := field.NewFieldElementFromInt64(4)
		b := field.NewFieldElementFromInt64(20)

		curve := NewEllipticCurve(a, b, field)

		// generator for E(F29)
		G := &Point{X: nt.FromInt64(1), Y: nt.FromInt64(5)}
		testCases := []*Point{
			{nt.FromInt64(0), nt.FromInt64(0)},
			{nt.FromInt64(1), nt.FromInt64(5)},
			{nt.FromInt64(4), nt.FromInt64(19)},
			{nt.FromInt64(20), nt.FromInt64(3)},
			{nt.FromInt64(15), nt.FromInt64(27)},
			{nt.FromInt64(6), nt.FromInt64(12)},
			{nt.FromInt64(17), nt.FromInt64(19)},
			{nt.FromInt64(24), nt.FromInt64(22)},
			{nt.FromInt64(8), nt.FromInt64(10)},
			{nt.FromInt64(14), nt.FromInt64(23)},
			{nt.FromInt64(13), nt.FromInt64(23)},
			{nt.FromInt64(10), nt.FromInt64(25)},
			{nt.FromInt64(19), nt.FromInt64(13)},
			{nt.FromInt64(16), nt.FromInt64(27)},
			{nt.FromInt64(5), nt.FromInt64(22)},
			{nt.FromInt64(3), nt.FromInt64(1)},
			{nt.FromInt64(0), nt.FromInt64(22)},
			{nt.FromInt64(27), nt.FromInt64(2)},
			{nt.FromInt64(2), nt.FromInt64(23)},
			{nt.FromInt64(2), nt.FromInt64(6)},
			{nt.FromInt64(27), nt.FromInt64(27)},
			{nt.FromInt64(0), nt.FromInt64(7)},
			{nt.FromInt64(3), nt.FromInt64(28)},
			{nt.FromInt64(5), nt.FromInt64(7)},
			{nt.FromInt64(16), nt.FromInt64(2)},
			{nt.FromInt64(19), nt.FromInt64(16)},
			{nt.FromInt64(10), nt.FromInt64(4)},
			{nt.FromInt64(13), nt.FromInt64(6)},
			{nt.FromInt64(14), nt.FromInt64(6)},
			{nt.FromInt64(8), nt.FromInt64(19)},
			{nt.FromInt64(24), nt.FromInt64(7)},
			{nt.FromInt64(17), nt.FromInt64(10)},
			{nt.FromInt64(6), nt.FromInt64(17)},
			{nt.FromInt64(15), nt.FromInt64(2)},
			{nt.FromInt64(20), nt.FromInt64(26)},
			{nt.FromInt64(4), nt.FromInt64(10)},
			{nt.FromInt64(1), nt.FromInt64(24)},
		}

		var i int64

		for i = 1; i < 37; i++ {
			actual := curve.ScalarMul(G, nt.FromInt64(i))
			expected := testCases[i]

			if !actual.Equal(expected) || !curve.IsOnCurve(actual) || !curve.IsOnCurve(expected) {
				t.Error("TestScalarMul failed expected : ", expected, "got :", actual)
			}
		}
	})
}
