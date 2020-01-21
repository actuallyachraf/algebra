// Package poly provides an implementation of polynomials over fields.
// Proposed are the arithmetic ops, evaluation and interpolation.
package poly

import (
	"fmt"
	"math"

	"github.com/actuallyachraf/algebra/ff"
	"github.com/actuallyachraf/algebra/nt"
)

// Polynomial implements the polynomial type using runtime slices.
// What is a polynomial ?
// P(X) = \sigma_{i=1..n} pow(X,i)*Coeff_i
// We appreciate polynomials to be sorted in decreasing by their powers
// the index represents the powers, the item at p[i] is the coefficient.
// A Polynomial is used for implementing binary fields (package bf).
type Polynomial struct {
	degree int
	poly   []ff.FieldElement
}

// String implements pretty printing for polynomials
func (p *Polynomial) String() string {
	var s = ""
	for idx, coeff := range p.poly {
		s += fmt.Sprintf("%sX^%d", coeff.String(), idx)
	}
	return s
}

// Evaluate a polynomial at some point in F
func (p *Polynomial) Evaluate(x ff.FieldElement) ff.FieldElement {

	if p.degree == 0 {
		r, _ := ff.New(nt.Zero, nt.One)
		return r
	}

	field := p.poly[0].Field()
	eval, _ := ff.New(nt.Zero, field.Modulus())

	for idx, coeff := range p.poly {
		field.Add(eval, field.Mul(coeff, x.Exp(nt.FromInt64(int64(idx)))))
	}

	return eval
}

// Degree returns the polynomial degree
func (p *Polynomial) Degree() int {
	return p.degree
}

// Equal are two polynomials equal ?
func (p *Polynomial) Equal(q *Polynomial) bool {

	if p.Degree() != q.Degree() {
		return false
	}

	fieldP := p.poly[0].Field()

	for idx := range p.poly {
		if fieldP.Cmp(p.poly[idx], q.poly[idx]) != 0 {
			return false
		}
	}

	return true
}

// New instantiates a new polynomial
func New(coeffs []ff.FieldElement, degree int) *Polynomial {

	var p = &Polynomial{}

	if degree <= 0 {
		return p
	}
	p.poly = make([]ff.FieldElement, degree+1)
	for idx, e := range coeffs {
		p.poly[idx] = e
	}

	p.degree = degree

	return p
}

// Add returns r = p + q it applies the underlying ring arithmetic.
func Add(p, q *Polynomial) *Polynomial {

	m, n := p.Degree(), q.Degree()
	// the degree of r is the max(deg(p),deg(q))
	d := int(math.Max(float64(m), float64(n)))

	r := make([]ff.FieldElement, d)

	copy(r, p.poly)

	field := p.poly[0].Field()

	for idx := 0; idx < n; idx++ {
		r[idx] = field.Add(r[idx], q.poly[idx])
	}

	return New(r, d)
}

// Mul multiply two polynomials
func Mul(p, q *Polynomial) *Polynomial {

	d := len(p.poly) + len(q.poly)
	r := make([]ff.FieldElement, d-1)

	field := p.poly[0].Field()

	for idx := range r {
		r[idx] = field.NewFieldElement(nt.Zero)
	}

	for i := 0; i < len(p.poly); i++ {
		for j := 0; j < len(q.poly); j++ {
			r[i+j] = field.Add(r[i+j], field.Mul(p.poly[i], q.poly[j]))
		}
	}
	return New(r, p.Degree()+q.Degree())
}
