package nt

// PollardRho implements the PollarRho factorization algorithm.
// Returns zero on failure.
func PollardRho(n *Integer) *Integer {
	a := FromInt64(2)
	b := FromInt64(2)

	d := FromInt64(0)

	for {
		a = ModAdd(ModExp(a, FromInt64(2), nil), One, n)
		b = ModAdd(ModExp(b, FromInt64(2), nil), One, n)

		d = GCD(Sub(a, b), n)

		if d.Cmp(One) == 1 && d.Cmp(n) == -1 {
			return d
		}
		if d.Cmp(n) == 0 {
			return Zero
		}
	}

}
