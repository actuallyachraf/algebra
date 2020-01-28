package poly

import (
	"math/big"
	"testing"
)

// Generates a dataset
func genTestCase(p Polynomial, start *big.Int, nhints int, q *big.Int) []Point {
	var pts []Point = make([]Point, nhints)
	for i := 0; i < nhints; i++ {
		x := big.NewInt(int64(i))
		x.Add(x, start)
		pts[i] = Point{x, p.Eval(x, q)}
	}
	return pts
}

func TestMonomial(t *testing.T) {
	cases := []struct {
		a   *big.Int
		ans Polynomial
	}{
		{
			big.NewInt(1),
			NewPolynomialInts(-1, 1),
		},
		{
			big.NewInt(-8),
			NewPolynomialInts(8, 1),
		},
		{
			big.NewInt(54321),
			NewPolynomialInts(-54321, 1),
		},
	}
	for _, c := range cases {
		res := monomial(c.a)
		if res.Compare(&c.ans) != 0 {
			t.Errorf("(x-a) from %v != %v (your answer was %v)", c.a, c.ans, res)
		}
	}
}

func TestLagrange(t *testing.T) {
	cases := []struct {
		ps  []Point
		m   *big.Int
		ans Polynomial
	}{
		{
			[]Point{
				Point{big.NewInt(1), big.NewInt(1)},
				Point{big.NewInt(2), big.NewInt(4)},
			},
			nil,
			NewPolynomialInts(0),
		},
		{
			[]Point{
				Point{big.NewInt(1), big.NewInt(1)},
				Point{big.NewInt(2), big.NewInt(4)},
				Point{big.NewInt(3), big.NewInt(9)},
			},
			big.NewInt(13),
			NewPolynomialInts(0, 0, 1),
		},
		{
			genTestCase(NewPolynomialInts(43, 53, 45, 63, 43, 55, 75), big.NewInt(11), 8, big.NewInt(311)),
			big.NewInt(311),
			NewPolynomialInts(43, 53, 45, 63, 43, 55, 75),
		},
		{
			genTestCase(NewPolynomialInts(43, 53, 45, 63, 43, 55, 75), big.NewInt(111), 8, big.NewInt(311)),
			big.NewInt(311),
			NewPolynomialInts(43, 53, 45, 63, 43, 55, 75),
		},
		{
			genTestCase(NewPolynomialInts(43, 53, 45, 63, 43, 55, 75), big.NewInt(111), 10, big.NewInt(311)),
			big.NewInt(311),
			NewPolynomialInts(43, 53, 45, 63, 43, 55, 75),
		},
		{
			genTestCase(NewPolynomialInts(1234561, 1234562, 1234563, 1234564, 1234565, 1234566, 1234567, 1234568, 1234569), big.NewInt(1234561), 11, big.NewInt(16769023)),
			big.NewInt(16769023),
			NewPolynomialInts(1234561, 1234562, 1234563, 1234564, 1234565, 1234566, 1234567, 1234568, 1234569),
		},
	}

	for _, c := range cases {
		res := Lagrange(c.ps, c.m)
		if res.Compare(&c.ans) != 0 {
			t.Errorf("Lagrange interpolation error : expected %v [modulo: %v] != %v | got %v", c.ps, c.m, c.ans, res)
		}
	}
}
