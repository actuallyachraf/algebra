// Package poly provides an implementation of polynomials over fields.
// Proposed are the arithmetic ops, evaluation and interpolation.
// Thanks to @jukworks for the base code.
package poly

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/actuallyachraf/algebra/ff"
	"github.com/actuallyachraf/algebra/nt"
)

// Polynomial implements the polynomial type
// using a vector of integers ordered by decreasing order i.e (lowest degree -> highest degree)
type Polynomial []*nt.Integer

// NewPolynomialInts Helper function for generating a polynomial with given integers
func NewPolynomialInts(coeffs ...int) (p Polynomial) {
	p = make([]*nt.Integer, len(coeffs))
	for i := 0; i < len(coeffs); i++ {
		p[i] = big.NewInt(int64(coeffs[i]))
	}
	p.trim()
	return
}

// NewPolynomialBigInt Helper function for generating a polynomial with given integers
func NewPolynomialBigInt(coeffs ...*big.Int) (p Polynomial) {
	p = make([]*nt.Integer, len(coeffs))
	for i := 0; i < len(coeffs); i++ {
		p[i] = coeffs[i]
	}
	p.trim()
	return
}

// NewPolynomial Helper function for generating a polynomial with field elements
func NewPolynomial(coeffs []ff.FieldElement) (p Polynomial) {
	p = make([]*nt.Integer, len(coeffs))
	for i := 0; i < len(coeffs); i++ {
		p[i] = coeffs[i].Big()
	}
	p.trim()
	return
}

// RandomPolynomial Returns a polynomial with random coefficients given a degree
// and coefficient in [0,2^bits]
func RandomPolynomial(degree, bits int64) (p Polynomial) {
	p = make(Polynomial, degree+1)
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	exp := big.NewInt(2)
	exp.Exp(exp, big.NewInt(bits), nil)
	for i := 0; i <= p.Degree(); i++ {
		p[i] = new(big.Int)
		p[i].Rand(rr, exp)
	}
	p.trim()
	return
}

// trim slices the underlying vector to remove zero coefficients of higher degree
func (p *Polynomial) trim() {
	var last int = 0
	for i := p.Degree(); i > 0; i-- { // why i > 0, not i >=0? do not remove the constant
		if (*p)[i].Sign() != 0 {
			last = i
			break
		}
	}
	*p = (*p)[:(last + 1)]
}

// isZero() checks if P is the zero polynomial
func (p *Polynomial) isZero() bool {
	if p.Degree() == 0 && (*p)[0].Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
}

// Degree returns the degree of the polynomial
func (p Polynomial) Degree() int {
	return len(p) - 1
}

// String implements the printing interface
func (p Polynomial) String() (s string) {
	s = "["
	for i := len(p) - 1; i >= 0; i-- {
		switch p[i].Sign() {
		case -1:
			if i == len(p)-1 {
				s += "-"
			} else {
				s += " - "
			}
			if i == 0 || p[i].Int64() != -1 {
				s += p[i].String()[1:]
			}
		case 0:
			continue
		case 1:
			if i < len(p)-1 {
				s += " + "
			}
			if i == 0 || p[i].Int64() != 1 {
				s += p[i].String()
			}
		}
		if i > 0 {
			s += "x"
			if i > 1 {
				s += "^" + fmt.Sprintf("%d", i)
			}
		}
	}
	if s == "[" {
		s += "0"
	}
	s += "]"
	return
}

// Compare compares two polynomials and returns -1 if P < Q, 0 if P = Q , or 1
func (p *Polynomial) Compare(q *Polynomial) int {
	switch {
	case p.Degree() > q.Degree():
		return 1
	case p.Degree() < q.Degree():
		return -1
	}
	for i := 0; i <= p.Degree(); i++ {
		switch (*p)[i].Cmp((*q)[i]) {
		case 1:
			return 1
		case -1:
			return -1
		}
	}
	return 0
}

// Add adds two polynomials m represents the modulo operator for polynomials
// over finite fields.
func (p Polynomial) Add(q Polynomial, m *nt.Integer) Polynomial {
	if p.Compare(&q) < 0 {
		return q.Add(p, m)
	}
	var r Polynomial = make([]*nt.Integer, len(p))
	for i := 0; i < len(q); i++ {
		a := new(big.Int)
		a.Add(p[i], q[i])
		r[i] = a
	}
	for i := len(q); i < len(p); i++ {
		a := new(big.Int)
		a.Set(p[i])
		r[i] = a
	}
	if m != nil {
		for i := 0; i < len(p); i++ {
			r[i].Mod(r[i], m)
		}
	}
	r.trim()
	return r
}

// Neg returns a polynomial Q = -P
func (p *Polynomial) Neg() Polynomial {
	var q Polynomial = make([]*nt.Integer, len(*p))
	for i := 0; i < len(*p); i++ {
		b := new(big.Int)
		b.Neg((*p)[i])
		q[i] = b
	}
	return q
}

// Clone does deep-copy of the polynomial, given adjust != 0 it raises
// the polynomial to a higher degree.
// for example, P = x + 1 and adjust = 2, Clone() returns x^3 + x^2
func (p Polynomial) Clone(adjust int) Polynomial {
	var q Polynomial = make([]*nt.Integer, len(p)+adjust)
	if adjust < 0 {
		return NewPolynomialInts(0)
	}
	for i := 0; i < adjust; i++ {
		q[i] = big.NewInt(0)
	}
	for i := adjust; i < len(p)+adjust; i++ {
		b := new(big.Int)
		b.Set(p[i-adjust])
		q[i] = b
	}
	return q
}

// reduce does modular arithmetic over modulus m
func (p *Polynomial) reduce(m *nt.Integer) {
	if m == nil {
		return
	}
	for i := 0; i <= (*p).Degree(); i++ {
		(*p)[i].Mod((*p)[i], m)
	}
	p.trim()
}

// Sub subtracts P from Q by simply P + (Neg(Q))
func (p Polynomial) Sub(q Polynomial, m *nt.Integer) Polynomial {
	r := q.Neg()
	return p.Add(r, m)
}

// Mul computes P * Q
func (p Polynomial) Mul(q Polynomial, m *nt.Integer) Polynomial {
	if m != nil {
		p.reduce(m)
		q.reduce(m)
	}
	var r Polynomial = make([]*nt.Integer, p.Degree()+q.Degree()+1)
	for i := 0; i < len(r); i++ {
		r[i] = big.NewInt(0)
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(q); j++ {
			a := new(big.Int)
			a.Mul(p[i], q[j])
			a.Add(a, r[i+j])
			if m != nil {
				a.Mod(a, m)
			}
			r[i+j] = a
		}
	}
	r.trim()
	return r
}

// Div returns (P / Q, P % Q)
func (p Polynomial) Div(q Polynomial, m *nt.Integer) (quo, rem Polynomial) {
	if m != nil {
		p.reduce(m)
		q.reduce(m)
	}
	if p.Degree() < q.Degree() || q.isZero() {
		quo = NewPolynomialInts(0)
		rem = p.Clone(0)
		return
	}
	quo = make([]*nt.Integer, p.Degree()-q.Degree()+1)
	rem = p.Clone(0)
	for i := 0; i < len(quo); i++ {
		quo[i] = big.NewInt(0)
	}
	t := p.Clone(0)
	qd := q.Degree()
	for {
		td := t.Degree()
		rd := td - qd
		if rd < 0 || t.isZero() {
			rem = t
			break
		}
		r := new(big.Int)
		if m != nil {
			r.ModInverse(q[qd], m)
			r.Mul(r, t[td])
			r.Mod(r, m)
		} else {
			r.Div(t[td], q[qd])
		}
		// if r == 0, it means that the highest coefficient of the result is not an integer
		// this polynomial library handles integer coefficients
		if r.Cmp(big.NewInt(0)) == 0 {
			quo = NewPolynomialInts(0)
			rem = p.Clone(0)
			return
		}
		u := q.Clone(rd)
		for i := rd; i < len(u); i++ {
			u[i].Mul(u[i], r)
			if m != nil {
				u[i].Mod(u[i], m)
			}
		}
		t = t.Sub(u, m)
		t.trim()
		quo[rd] = r
	}
	quo.trim()
	rem.trim()
	return
}

// GCD returns the greatest common divisor(GCD) of P and Q (Euclidean algorithm)
func (p Polynomial) GCD(q Polynomial, m *nt.Integer) Polynomial {
	if p.Compare(&q) < 0 {
		return q.GCD(p, m)
	}
	if q.isZero() {
		return p
	}
	_, rem := p.Div(q, m)
	return q.GCD(rem, m)

}

// Mod reduces a polynomial modulo another polynomial
func (p Polynomial) Mod(q Polynomial, m *nt.Integer) Polynomial {

	_, rem := p.Div(q, m)

	return rem
}

// Quo returns the quotient of two polynomials
func (p Polynomial) Quo(q Polynomial, m *nt.Integer) Polynomial {
	quo, _ := p.Div(q, m)
	return quo
}

// Eval returns p(v) where v is the given big integer
func (p Polynomial) Eval(x *nt.Integer, m *nt.Integer) (y *nt.Integer) {
	y = big.NewInt(0)
	accx := big.NewInt(1)
	xd := new(big.Int)
	for i := 0; i <= p.Degree(); i++ {
		xd.Mul(accx, p[i])
		y.Add(y, xd)
		accx.Mul(accx, x)
		if m != nil {
			y.Mod(y, m)
			accx.Mod(accx, m)
		}
	}
	return y
}

// Compose returns p(q(x))
func (p Polynomial) Compose(q Polynomial, m *nt.Integer) Polynomial {

	r := NewPolynomialInts(0)

	other := q.Clone(0)

	for _, coeff := range p.Reverse() {
		r = other.Mul(r, m).Add(NewPolynomialBigInt(coeff), m)
	}
	return r
}

// Reverse the order of the polynomial coefficients
func (p Polynomial) Reverse() Polynomial {
	a := p.Clone(0)
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
