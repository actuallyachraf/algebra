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
	degree    int
	poly      []ff.FieldElement
	baseField ff.FiniteField
}

// New instantiates a new polynomial
func New(coeffs []ff.FieldElement, degree int) *Polynomial {

	var p = &Polynomial{}

	p.poly = coeffs
	p.baseField = coeffs[0].Field()
	return p
}
func (p *Polynomial) trim() {
	var last int = 0
	for i := p.Degree(); i >= 0; i-- { // why i > 0, not i >=0? do not remove the constant
		if (p.poly)[i].Big().Sign() != 0 {
			last = i
			break
		}
	}
	p.poly = p.poly[:(last + 1)]

}

// ZeroPolynomial returns the zeroth monomial for a given field
func ZeroPolynomial(field ff.FiniteField) *Polynomial {
	m := New([]ff.FieldElement{field.NewFieldElementFromInt64(0)}, 0)
	return m
}

// String implements pretty printing for polynomials
func (p *Polynomial) String() string {
	var s = ""
	for idx, coeff := range p.poly {
		s += fmt.Sprintf("%sX^%d+", coeff.String(), idx)
	}
	return s
}

// Evaluate a polynomial at some point in F
func (p *Polynomial) Evaluate(x ff.FieldElement) ff.FieldElement {

	if p.Degree() == 0 {
		r, _ := ff.New(nt.Zero, nt.One)
		return r
	}

	field := p.baseField
	eval, _ := ff.New(nt.Zero, field.Modulus())

	for idx, coeff := range p.poly {
		field.Add(eval, field.Mul(coeff, x.Exp(nt.FromInt64(int64(idx)))))
	}

	return eval
}

// Degree returns the polynomial degree
func (p *Polynomial) Degree() int {
	return len(p.poly) - 1
}

// LeadingCoeff returns the leading coefficient of the polynomial
func (p *Polynomial) LeadingCoeff() ff.FieldElement {
	return p.poly[p.Degree()]
}

// Equal are two polynomials equal ?
func (p *Polynomial) Equal(q *Polynomial) bool {

	if p.Degree() != q.Degree() {
		return false
	}

	if len(p.poly) != len(q.poly) {
		return false
	}

	fieldP := p.baseField
	fieldQ := q.baseField

	if fieldP.Char().Cmp(fieldQ.Char()) != 0 {
		return false
	}

	for idx := 0; idx < len(p.poly); idx++ {
		if fieldP.Cmp(p.poly[idx], q.poly[idx]) != 0 {
			return false
		}
	}

	return true
}

// Add returns r = p + q it applies the underlying ring arithmetic.
func Add(p, q *Polynomial) *Polynomial {

	m, n := p.Degree(), q.Degree()
	field := q.baseField

	var d int

	if q.Equal(ZeroPolynomial(field)) {
		return p
	} else if p.Equal(ZeroPolynomial(field)) {
		return q
	} else if m > n {
		d = m
	} else if m <= n {
		d = n
	}
	var r = make([]ff.FieldElement, d+1)

	for i := 0; i < d+1; i++ {
		r[i] = field.Add(p.poly[i], q.poly[i])
	}

	return New(r, d)
}

// Sub substracts two polynomials
func Sub(p, q *Polynomial) *Polynomial {

	m, n := p.Degree(), q.Degree()
	// the degree of r is the max(deg(p),deg(q))
	d := int(math.Max(float64(m), float64(n)))

	r := make([]ff.FieldElement, d+1)

	copy(r, p.poly)

	field := p.baseField

	for idx := 0; idx < d+1; idx++ {
		r[idx] = field.Sub(r[idx], q.poly[idx])
	}

	return New(r, d)
}

// Mul multiply two polynomials
func Mul(p, q *Polynomial) *Polynomial {

	d := p.Degree() + q.Degree()
	r := make([]ff.FieldElement, d+1)

	field := p.baseField

	for idx := 0; idx < len(r); idx++ {
		r[idx] = field.NewFieldElementFromInt64(0)
	}

	for i := 0; i < len(p.poly); i++ {
		for j := 0; j < len(q.poly); j++ {
			r[i+j] = field.Add(r[i+j], field.Mul(p.poly[i], q.poly[j]))
		}
	}
	return New(r, d)
}

// Copy reproduces a copy of the polynomial
func (p *Polynomial) Copy() *Polynomial {

	q := make([]ff.FieldElement, len(p.poly))

	copy(q, p.poly)

	return New(q, p.Degree())
}

// QuoRem computes polynomial long division
func QuoRem(p, q *Polynomial) (*Polynomial, *Polynomial) {

	r := p.Copy()

	field := p.baseField

	n := p.Degree()
	m := q.Degree()

	diff := n - m
	out := make([]ff.FieldElement, diff+1)

	for idx := diff; diff >= 0; idx-- {
		quot := field.Div(p.poly[n], q.poly[m])

		out[idx] = quot

		for i := m; i >= 0; i-- {
			r.poly[diff+i] = field.Sub(r.poly[diff+i], field.Mul(q.poly[i], quot))
		}
		diff--
		n--
		r.trim()
	}
	quo := New(out, len(out)-1)
	quo.trim()
	r.trim()
	return quo, r
}
