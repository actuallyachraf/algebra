package ec

import (
	"bytes"

	"github.com/actuallyachraf/algebra/nt"
)

var (
	// Inf defines the zero point
	Inf = &Point{nt.Zero, nt.Zero}
)

// Point represents a point on an elliptic curve
type Point struct {
	X *nt.Integer
	Y *nt.Integer
}

// Equal checks if two points are equal
func (p *Point) Equal(q *Point) bool {
	if !bytes.Equal(p.X.Bytes(), q.X.Bytes()) {
		return false
	}
	if !bytes.Equal(p.Y.Bytes(), q.Y.Bytes()) {
		return false
	}
	return true
}

// String returns the components of the point in a string
func (p *Point) String() string {
	return "(" + p.X.String() + ", " + p.Y.String() + ")"
}
