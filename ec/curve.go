package ec

import (
	"errors"
	"math/big"

	"github.com/actuallyachraf/algebra/ff"
	"github.com/actuallyachraf/algebra/nt"
)

// Curve represents an elliptic curve by it's parameters.
// An elliptic curve is defined by the Weirstrass equation:
// (E): y^2 + a1xy + a3y = x^3 +a2x^2 + a4x + a6
// (E) can be simplified using changes of variables
// There are two separate cases :
// (E) is defined over a field K with characteristic different than 2 and 3
// (E) is defined over a field K with characteristic 2 or 3
// We treat the first case for now .
// The change of variable used is :
// Phi : (X,Y) -> ((x-3a1^2-12a2)/36,((y-3a1x)/216)-((a1^3 + 4a1a2 - 12a3)/24)
// Applying Phi to (E) gives us the simplified Weirstreass equation :
// (E)s : y^2 = x^3 + ax + b
type Curve struct {
	A ff.FieldElement
	B ff.FieldElement
	F ff.FiniteField
}

// NewEllipticCurve creates an instance of an elliptic curve
func NewEllipticCurve(a, b ff.FieldElement, q ff.FiniteField) *Curve {
	return &Curve{
		A: a,
		B: b,
		F: q,
	}
}

// IsOnCurve checks if a given point is on the curve
func (c *Curve) IsOnCurve(p *Point) bool {

	// get the field we're operating in
	field := c.F

	ySquared := field.NewFieldElement(p.Y).Square()
	xCubed := field.NewFieldElement(p.X).Exp(nt.FromInt64(3))

	righthandSide := field.Add(field.Add(xCubed, field.Mul(c.A, field.NewFieldElement(p.X))), c.B)
	if righthandSide.Equal(ySquared) {
		return true
	}
	return false
}

// Add computes the sum of two points on the curve
func (c *Curve) Add(p, q *Point) *Point {

	field := c.F

	if p.Equal(Inf) && q.Equal(Inf) {
		return Inf
	} else if !p.Equal(Inf) && q.Equal(Inf) {
		return p
	} else if p.Equal(Inf) && !q.Equal(Inf) {
		return q
	} else if q.Equal(c.Neg(p)) || p.Equal(c.Neg(q)) {
		return Inf
	} else if p.Equal(q) {
		x1 := field.NewFieldElement(p.X)
		y1 := field.NewFieldElement(p.Y)

		x3 := field.Sub(field.Div(field.Add(field.Mul(x1.Square(), field.NewFieldElementFromInt64(3)), c.A), field.Mul(y1, field.NewFieldElementFromInt64(2))).Square(), field.Mul(x1, field.NewFieldElementFromInt64(2)))

		y3 := field.Sub(field.Mul(field.Sub(x1, x3), field.Div(field.Add(field.Mul(x1.Square(), field.NewFieldElementFromInt64(3)), c.A), field.Mul(y1, field.NewFieldElementFromInt64(2)))), y1)

		return &Point{x3.Big(), y3.Big()}
	}

	// We use the formulas from http://cacr.uwaterloo.ca/ecc/
	x1 := field.NewFieldElement(p.X)
	x2 := field.NewFieldElement(q.X)

	y1 := field.NewFieldElement(p.Y)
	y2 := field.NewFieldElement(q.Y)

	// x3 = ((y2-y1)/(x2-x1))^2 - x1 - x2
	x3 := field.Sub(field.Div(field.Sub(y2, y1), field.Sub(x2, x1)).Square(), field.Add(x1, x2))
	// y3 = ((y2-y1)/(x2-x1))(x1-x3)-y1
	y3 := field.Sub(field.Mul(field.Div(field.Sub(y2, y1), field.Sub(x2, x1)), field.Sub(x1, x3)), y1)

	return &Point{x3.Big(), y3.Big()}
}

// Double computes 2P
func (c *Curve) Double(p *Point) *Point {
	return c.Add(p, p)
}

// Neg gives you the inverse of (X,Y) which is (X,-Y).
func (c *Curve) Neg(p *Point) *Point {

	return &Point{X: p.X, Y: nt.Sub(c.F.Modulus(), p.Y)}
}

// ScalarMul computes multiplication of curve points by scalars
func (c *Curve) ScalarMul(p *Point, s *nt.Integer) *Point {
	k := new(big.Int).Set(s)
	// the algorithm uses the double and square methods
	q := &Point{X: nt.Zero, Y: nt.Zero}
	for nt.Zero.Cmp(k) == -1 {
		// get the rightmost bit
		b := new(nt.Integer).And(k, nt.One)
		// check if it's a one
		// then add it
		if b.Cmp(nt.One) == 0 {
			q = c.Add(q, p)
		}
		// right shift the scalar bits by one
		s = k.Rsh(k, 1)
		p = c.Double(p)
	}

	return q

}

// Order returns smallest n where nG = O (point at zero)
func (c *Curve) Order(g *Point) (*nt.Integer, error) {
	// loop from i:=1 to i<ec.Q+1
	start := nt.One
	end := c.F.Modulus()
	for i := new(big.Int).Set(start); i.Cmp(end) <= 0; i.Add(i, nt.One) {
		iCopy := new(big.Int).SetBytes(i.Bytes())
		mPoint := c.ScalarMul(g, iCopy)

		if mPoint.Equal(Inf) {
			return i, nil
		}
	}
	return nt.Zero, errors.New("invalid order")
}
