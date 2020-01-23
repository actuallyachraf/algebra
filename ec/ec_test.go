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
		points := []*Point{
			{nt.FromInt64(2), nt.FromInt64(6)},
			{nt.FromInt64(4), nt.FromInt64(19)},
			{nt.FromInt64(5), nt.FromInt64(7)},
			{nt.FromInt64(5), nt.FromInt64(22)},
			{nt.FromInt64(2), nt.FromInt64(23)},
			{nt.FromInt64(10), nt.FromInt64(25)},
			{nt.FromInt64(13), nt.FromInt64(6)},
			{nt.FromInt64(16), nt.FromInt64(27)},
		}

		for _, point := range points {
			if !curve.IsOnCurve(point) {
				t.Error("IsOnCurve failed")
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
}
