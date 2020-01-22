package ec

import "github.com/actuallyachraf/algebra/ff"

import "github.com/actuallyachraf/algebra/nt"

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

	ySquared := field.NewFieldElement(new(nt.Integer).Exp(p.Y, nt.FromInt64(2), nil))
	xCubed := field.NewFieldElement(new(nt.Integer).Exp(p.X, nt.FromInt64(3), nil))

	righthandSide := field.Add(field.Add(xCubed, field.Mul(c.A, field.NewFieldElement(p.X))), c.B)
	if field.Cmp(ySquared, righthandSide) == 0 {
		return true
	}
	return false
}

// Add computes the sum of two points on the curve
func (c *Curve) Add(p, q *Point) *Point {

	// We use the formulas from http://cacr.uwaterloo.ca/ecc/
	field := c.F

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
