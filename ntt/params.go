package ntt

import (
	"math/big"

	"github.com/actuallyachraf/algebra/nt"
)

// NTTParams defines the parameters of the number-theoretic transform.
type NTTParams struct {
	n             int64
	nRev          uint64
	bitlength     uint64
	q             nt.Integer
	qInv          *nt.Integer
	PsiRev        []*nt.Integer
	PsiRevMont    []*nt.Integer
	PsiInvRev     []*nt.Integer
	PsiInvRevMont []*nt.Integer
}

// GenParams generates the parameters for NTT and Inverse NTT.
func GenParams(N int64, Q nt.Integer) *NTTParams {
	// setting up initial parameters
	var nttParams = &NTTParams{
		n:    N,
		nRev: nt.ModInv(nt.FromInt64(int64(N)), &Q).Uint64(),
		q:    Q,
	}
	// computing Psi terms
	g := primitiveRoot(&Q)
	fi := nt.Sub(&Q, nt.One)

	// compute the 2-nth root of unity and its inverse
	// psi = g^(fi/2n) mod q
	// psiInv = psi^-1 mod q
	twoN := nt.FromInt64(2 * N)
	exp := nt.Div(fi, twoN)
	psi := nt.ModExp(g, exp, &Q)
	psiInv := nt.ModInv(psi, &Q)
	// compute the powers of psi and psiInv
	nttParams.PsiRev = make([]*nt.Integer, N)
	nttParams.PsiInvRev = make([]*nt.Integer, N)
	var Nbitlen uint32
	for i := 32 - 1; i >= 0; i-- {
		if N&(1<<uint(i)) != 0 {
			Nbitlen = uint32(i)
			break
		}
	}
	var idxRev uint32
	var i nt.Integer
	var k int64
	for k = 0; k < N; k++ {
		i.SetInt64(k)
		idxRev = bitRev(k, Nbitlen)
		nttParams.PsiRev[idxRev] = nt.ModExp(psi, &i, &Q)
		nttParams.PsiInvRev[idxRev] = nt.ModExp(psiInv, &i, &Q)
	}

	// compute the montgomery rep of psi and psiInv
	nttParams.PsiRevMont = make([]*big.Int, N)
	nttParams.PsiInvRevMont = make([]*big.Int, N)

	Qbitlen := int64(Q.BitLen() + 5)
	R := nt.ModExp(nt.FromInt64(2), nt.FromInt64(Qbitlen), &Q)

	for k = 0; k < N; k++ {
		nttParams.PsiRevMont[k] = nt.ModMul(R, nttParams.PsiRev[k], &Q)
		nttParams.PsiInvRevMont[k] = nt.ModMul(R, nttParams.PsiInvRev[k], &Q)
	}
	nttParams.bitlength = uint64(Qbitlen)
	// computing qInv (montgomery reduction param)
	R = R.Lsh(nt.One, uint(Qbitlen))
	RInv := nt.ModInv(R, &Q)
	nttParams.qInv = nt.Div(nt.Sub(nt.Mul(R, RInv), nt.One), &Q)

	return nttParams
}

// bitRev calculates the bit-reverse index.
func bitRev(index int64, bitLen uint32) uint32 {
	indexReverse := uint32(0)
	for i := uint32(0); i < bitLen; i++ {
		if (index>>i)&1 != 0 {
			indexReverse |= 1 << (bitLen - 1 - i)
		}
	}
	return indexReverse
}

// polynomialPollardsRho calculates x1^2 + c mod x2, and is used in factorizationPollardsRho
func polynomialPollardsRho(x1, x2, c *nt.Integer) *nt.Integer {
	two := nt.FromInt64(2)
	z := new(nt.Integer).Exp(x1, two, x2) // x1^2 mod x2
	z.Add(z, c)                           // (x1^2 mod x2) + 1
	z.Mod(z, x2)                          // (x1^2 + 1) mod x2
	return z
}

// factorizationPollardsRho realizes Pollard's Rho algorithm for fast prime factorization,
// but this function only returns one factor a time
func factorizationPollardsRho(m *nt.Integer) *nt.Integer {
	var x, y, d, c *nt.Integer
	zero := nt.Zero
	one := nt.One
	ten := nt.FromInt64(10)

	// c is to change the polynomial used in Pollard's Rho algorithm,
	// Every time the algorithm fails to get a factor, increasing c to retry,
	// because Pollard's Rho algorithm sometimes will miss some small prime factors.
	for c = nt.FromInt64(1); !nt.Equal(c, ten); c.Add(c, one) {
		x, y, d = nt.FromInt64(2), nt.FromInt64(2), nt.FromInt64(1)
		for d.Cmp(zero) != 0 {
			x = polynomialPollardsRho(x, m, c)
			y = polynomialPollardsRho(polynomialPollardsRho(y, m, c), m, c)
			sub := nt.Sub(x, y)
			d.GCD(nil, nil, sub.Abs(sub), m)
			if d.Cmp(one) == 1 {
				return d
			}
		}
	}
	return d
}

// getFactors returns all the prime factors of m
func getFactors(n *nt.Integer) []*nt.Integer {
	var factor *nt.Integer
	var factors []*nt.Integer
	var m, tmp nt.Integer
	m.Set(n)
	zero := nt.Zero
	one := nt.One

	// first, append small prime factors
	for i := range smallPrimes {
		smallPrime := nt.FromInt64(smallPrimes[i])
		addFactor := false
		for tmp.Mod(&m, smallPrime).Cmp(zero) == 0 {
			m.Div(&m, smallPrime)
			addFactor = true
		}
		if addFactor {
			factors = append(factors, smallPrime)
		}
	}

	if m.Cmp(one) == 0 {
		return factors
	}

	// second, find other prime factors
	for {
		factor = factorizationPollardsRho(&m)
		if factor.Cmp(zero) == 0 {
			factors = append(factors, &m)
			break
		}
		m.Div(&m, factor)
		if len(factors) > 0 && factor.Cmp(factors[len(factors)-1]) == 0 {
			continue
		}
		factors = append(factors, factor)
	}
	return factors
}

// primitiveRoot calculates one primitive root of prime q
func primitiveRoot(q *nt.Integer) *nt.Integer {
	tmp := new(nt.Integer)
	notFoundPrimitiveRoot := true
	qMinusOne := nt.Sub(q, nt.FromInt64(1))
	factors := getFactors(qMinusOne)
	g := nt.FromInt64(2)
	one := nt.One
	for notFoundPrimitiveRoot {
		g.Add(g, one)
		for _, factor := range factors {
			tmp.Div(qMinusOne, factor)
			// once exist g^(q-1)/factor = 1 mod q, g is not a primitive root
			if tmp.Exp(g, tmp, q).Cmp(one) == 0 {
				notFoundPrimitiveRoot = true
				break
			}
			notFoundPrimitiveRoot = false
		}
	}
	return g
}

var smallPrimes = []int64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 73, 79, 83, 89, 97,
	101, 103, 107, 109, 113, 127, 131, 139, 149, 151, 163, 167, 173, 179, 181, 191, 193, 197, 199,
}
