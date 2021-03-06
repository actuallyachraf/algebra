package bp

import (
	"errors"

	"github.com/actuallyachraf/algebra/nt"
)

// Vector defines a vector type
type Vector []*nt.Integer

// NewZeroVector creates a new vector of  zeros of size x
func NewZeroVector(size int) Vector {
	v := make(Vector, size)

	for i := range v {
		v[i] = nt.FromInt64(0)
	}
	return v
}

// NewVector creates a vector of elements
func NewVector(elems []*nt.Integer) Vector {

	v := NewZeroVector(len(elems))

	for i := range v {
		v[i] = elems[i]
	}

	return v
}

// Equal compares two vectors component wise
func (v Vector) Equal(w Vector) bool {

	if v.Len() != w.Len() {
		return false
	}
	for i := range v {
		if v[i].Cmp(w[i]) != 0 {
			return false
		}
	}
	return true
}

// Len returns the count of elements in the vector
func (v Vector) Len() int {
	return len(v)
}

// Add sums two vectors and returns the result
func (v Vector) Add(w Vector) (Vector, error) {
	if v.Len() != w.Len() {
		return Vector{}, errors.New("vectors are of different sizes")
	}

	u := NewZeroVector(v.Len())

	for i := range u {
		u[i] = nt.Add(v[i], w[i])
	}

	return u, nil
}

// AddMod sums two vectors modulo order and returns the result
func (v Vector) AddMod(w Vector, order *nt.Integer) (Vector, error) {
	if v.Len() != w.Len() {
		return Vector{}, errors.New("vectors are of different sizes")
	}

	u := NewZeroVector(v.Len())

	for i := range u {
		u[i] = nt.Mod(nt.Add(v[i], w[i]), order)
	}

	return u, nil
}

// InnerProdMod computes the inner product of two vectors modulo order and returns the result
func (v Vector) InnerProdMod(w Vector, order *nt.Integer) (*nt.Integer, error) {

	res := nt.FromInt64(0)
	if v.Len() != w.Len() {
		return res, errors.New("vectors are of different size")
	}

	for i := range v {
		res = nt.Add(res, nt.Mod(nt.Mul(v[i], w[i]), order))
		res = nt.Mod(res, order)
	}

	return res, nil
}

// HadamardProdMod computes the Hadamard (component-wise) product of two vectors modulo order and returns the result
func (v Vector) HadamardProdMod(w Vector, order *nt.Integer) (Vector, error) {

	if v.Len() != w.Len() {
		return NewZeroVector(0), errors.New("vectors are of different size")
	}
	u := NewZeroVector(v.Len())

	for i := range v {
		u[i] = nt.Mod(nt.Mul(v[i], w[i]), order)
	}

	return u, nil
}

// ScalarAddMod computes component wise addition with scalars reduced modulo order
func (v Vector) ScalarAddMod(s *nt.Integer, order *nt.Integer) (Vector, error) {

	u := NewZeroVector(v.Len())

	for i := range v {
		u[i] = nt.Mod(nt.Add(v[i], s), order)
	}

	return u, nil
}

// ScalarMulMod computes component wise multiplication with scalars reduced modulo order
func (v Vector) ScalarMulMod(s *nt.Integer, order *nt.Integer) (Vector, error) {

	u := NewZeroVector(v.Len())

	for i := range v {
		u[i] = nt.Mod(nt.Mul(v[i], s), order)
	}

	return u, nil
}
