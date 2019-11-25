package nt

import (
	"crypto/rand"
	"math/big"
)

// Integer type wraps bigint internally represents arbitrary precision
// numbers in Z
type Integer = big.Int

// One as an arbitrary precision integer
var One = new(Integer).SetInt64(1)

// Zero as an arbitrary precision integer
var Zero = new(Integer).SetInt64(0)

// Cmp compars two numbers returns 0 if equal , -1 if the first lesser than the
// second and 1 otherwise.
func Cmp(a, b *Integer) int {
	return a.Cmp(b)
}

// FromInt64 returns an Integer from int64
func FromInt64(x int64) *Integer {
	return new(Integer).SetInt64(x)
}

// Add sums two Integers
func Add(a, b *Integer) *Integer {

	return new(big.Int).Add(a, b)
}

// Sub substracts two Integers
func Sub(a, b *Integer) *Integer {
	return new(big.Int).Sub(a, b)
}

// Mul multiplies two Integers
func Mul(a, b *Integer) *Integer {
	return new(big.Int).Mul(a, b)
}

// Div divides two Integers
func Div(a, b *Integer) *Integer {
	return new(big.Int).Div(a, b)
}

// GCD computes the greatest common divsior of two Integer
func GCD(a, b *Integer) *Integer {
	return new(big.Int).GCD(nil, nil, a, b)
}

// XGCD computes the extended GCD of two Integers
func XGCD(a, b *Integer) (d, x, y *Integer) {

	x = new(big.Int)
	y = new(big.Int)
	d = new(big.Int).GCD(x, y, a, b)
	return
}

// Mod computes the module division
func Mod(a, b *Integer) *Integer {
	return new(big.Int).Mod(a, b)
}

// ModAdd computes modular addition
func ModAdd(a, b, m *Integer) *Integer {
	return new(big.Int).Mod(Add(a, b), m)
}

// ModSub computes modular substraction
func ModSub(a, b, m *Integer) *Integer {
	return new(big.Int).Mod(Sub(a, b), m)
}

// ModMul computes modular multiplication
func ModMul(a, b, m *Integer) *Integer {
	return new(big.Int).Mod(Mul(a, b), m)
}

// ModInv computes the multiplicative inverse
func ModInv(a, m *Integer) *Integer {
	return new(big.Int).ModInverse(a, m)
}

// ModDiv computes the division modulo
func ModDiv(a, b, m *Integer) *Integer {
	// Modular division is equivalent to multiplication by the modular
	// inverse, b must be invertible i.e gcd(b,m) = 1.
	bInv := ModInv(b, m)
	if b == nil {
		return nil
	}
	return ModMul(a, bInv, m)
}

// ModExp computes the modular power
func ModExp(a, b, m *Integer) *Integer {
	return new(big.Int).Exp(a, b, m)
}

// Jacobi returns the Jacobi symbol
// a useful tool for keeping track whether an integer is a quadratic residue
// modulo n :
// 0 if n | q
// 1 if n is a quadratic residue
// -1 otherwise
func Jacobi(a, n *Integer) int {
	return big.Jacobi(a, n)
}

// GenPrime generates a random prime using a crypto/rand
func GenPrime(numbits int) *Integer {

	p, err := rand.Prime(rand.Reader, numbits)
	if err != nil {
		return nil
	}
	return p

}

// IsPrime runs a probabilistic check for whether p is a prime
func IsPrime(p *Integer) bool {
	return p.ProbablyPrime(100)
}

// XGCD2 is an alternative implementation to the extended GCD by math/big
func XGCD2(a, b *Integer) (d, x, y *Integer) {

	if b.Cmp(Zero) == 0 {
		d = a
		x = One
		y = Zero
	}

	var x2 = new(big.Int).Set(One)
	var x1 = new(big.Int).Set(Zero)
	var y2 = new(big.Int).Set(Zero)
	var y1 = new(big.Int).Set(One)

	for b.Cmp(Zero) == 1 {
		q := Div(a, b)
		r := Sub(a, Mul(q, b))

		x = Sub(x2, Mul(q, x1))
		y = Sub(y2, Mul(y1, q))

		a.Set(b)
		b.Set(r)

		x2.Set(x1)
		x1.Set(x)

		y2.Set(y1)
		y1.Set(y)
	}

	d = a
	x = x2
	y = y2

	return
}

// ModInv2 computes Multiplicate Inverse
func ModInv2(a, m *Integer) *Integer {
	d, x, _ := XGCD2(a, m)

	if d.Cmp(One) == 1 {
		return nil
	}
	return x

}

// ModExp2 implements modular exponentiation using fast square and multiply
func ModExp2(a, k, m *Integer) *Integer {

	var b = new(big.Int).Set(One)
	if k.Cmp(Zero) == 0 {
		return b
	}
	var A = new(big.Int).Set(a)

	if k.Bit(0) == 1 {
		b.Set(a)
	}

	for i := 1; i < k.BitLen(); i++ {
		A.Set(ModMul(A, A, m))
		if k.Bit(i) == 1 {
			b.Set(ModMul(A, b, m))
		}
	}
	return b

}
