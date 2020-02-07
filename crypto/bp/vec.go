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

// InnerProd computes the inner product of two vectors and return the result
func (v Vector) InnerProd(w Vector, order *nt.Integer) (*nt.Integer, error) {

	res := nt.FromInt64(0)
	ord := new(big.Int)
	if order == nil {
		ord.SetInt64(1)
	}
	if v.Len() != w.Len() {
		return res, errors.New("vectors are of different size")
	}


	for i := range v {
		res = nt.Add(res, nt.Mod(nt.Mul(v[i], w[i]),ord))
		res = nt.Mod(res, ord)
	}

	return res,nil
}
