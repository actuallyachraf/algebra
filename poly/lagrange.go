// Package poly this file implements lagrange polynomial interpolation.
// The unisolvence theorem states succintly that given n+1 distinct Points (x,y)
// there is a unique polynomial of degree at most n s.t p(x_i) = y_i.
// The lagrange interpolant written P(x) = Prod(i,j) (x-ui)/(ui-uj) for 0 <= j <= n.
package poly

import (
	"fmt"

	"github.com/actuallyachraf/algebra/nt"
)

// Point represents a polynomial evaluation
type Point struct {
	x *nt.Integer
	y *nt.Integer
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

// monomial returns a polynomial of the form x-u
func monomial(u *nt.Integer) Polynomial {
	v := new(nt.Integer).Neg(u)
	// The polynomial is represented as -u + x
	return Polynomial{v, nt.FromInt64(1)}
}

// Lagrange implements polynomial interpolation using the Lagrange polynomial
func Lagrange(Points []Point, modulus *nt.Integer) Polynomial {

	if modulus == nil {
		return NewPolynomialInts(0)
	}

	n := len(Points)

	interpolant := NewPolynomialInts(0)

	for i, num, den := 0, new(nt.Integer), new(nt.Integer); i < n; i++ {
		// lx represents a lagrage term
		lx := NewPolynomialInts(1)
		num.Set(Points[i].y)
		// i is fixed and we compute the Lagrange polynomial by iterating over
		// the x terms
		for j := 0; j < n; j++ {
			// skip the case when i = j
			if i == j {
				continue
			}
			lx = lx.Mul(monomial(Points[j].x), modulus)
			den = den.Sub(Points[i].x, Points[j].x)
			den = den.Mod(den, modulus)
			den = den.ModInverse(den, modulus)
			num = num.Mul(num, den)
			num = num.Mod(num, modulus)
		}
		// evaluate the terms for the lagrange interpolant
		for k := 0; k <= lx.Degree(); k++ {
			lx[k].Mul(lx[k], num)
			lx[k].Mod(lx[k], modulus)
		}
		// trim the leading zero coefficient terms
		lx.trim()
		// add the lagrange interpolant to the polynomial
		interpolant = interpolant.Add(lx, modulus)

	}
	interpolant.trim()

	return interpolant

}
