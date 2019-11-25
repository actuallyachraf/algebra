// Package poly provides an implementation of polynomials over fields.
// Proposed are the arithmetic ops, evaluation and interpolation.
package poly

import (
	"github.com/actuallyachraf/algebra/ff"
)

// Polynomial implements the polynomial type using runtime slices.
// What is a polynomial ?
// P(X) = \sigma_{i=1..n} pow(X,i)*Coeff_i
// We appreciate polynomials to be sorted by their powers
// the index represents the powers, the item at p[i] is the coefficient.
// A Polynomial is used for implementing binary fields (package bf).
type Polynomial struct {
	degree int
	poly   []ff.FieldElement
}

// New instantiates a new polynomial
func New(coeffs []ff.FieldElement, degree int) Polynomial {

	var p Polynomial
	p.poly = make([]ff.FieldElement, len(coeffs))
	for idx, e := range coeffs {
		p.poly[idx] = e
	}

	p.degree = len(coeffs)

	return p
}

// Add returns r = p + q it applies the underlying ring arithmetic.
func Add(p, q Polynomial) Polynomial {

	var r Polynomial

	return r
}
