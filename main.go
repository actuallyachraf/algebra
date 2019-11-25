package main

import (
	"fmt"
	"math/big"

	"github.com/actuallyachraf/algebra/ff"

	"github.com/actuallyachraf/algebra/nt"
)

func main() {

	var a = new(nt.Integer).SetInt64(7382508725235)
	var b = new(nt.Integer).SetInt64(-92098532)
	fmt.Println(nt.Add(a, b))
	var c = new(nt.Integer).SetInt64(4864)
	var d = new(nt.Integer).SetInt64(3458)
	fmt.Println(nt.XGCD(c, d))
	fmt.Println(nt.XGCD2(c, d))

	// Let's generate the multiplicative group Z modulo 25
	// this is the set of coprimes less than 25
	Zmod25 := make([]*nt.Integer, 25)
	// Zn = {g^i mod n | 0 <= i <= phi(n)-1}
	// phi is the euler totient function
	// the generator of this group is 2
	var generator = new(nt.Integer).SetInt64(2)
	var i int64
	modulus := nt.FromInt64(2234)

	for i = 0; i <= 19; i++ {
		Zmod25[i] = nt.ModExp(generator, new(nt.Integer).SetInt64(i), modulus)
	}
	fmt.Println(Zmod25)
	for i = 0; i <= 19; i++ {
		Zmod25[i] = nt.ModExp2(generator, new(nt.Integer).SetInt64(i), modulus)
	}
	fmt.Println(Zmod25)

	a.SetInt64(5)
	var k = new(big.Int).SetInt64(596)

	fmt.Println(nt.ModExp(a, k, modulus))

	Z17 := ff.NewFiniteField(new(big.Int).SetInt64(17))
	x := Z17.NewFieldElement(nt.FromInt64(4))
	y := Z17.NewFieldElement(nt.FromInt64(5))
	// what's 4 modinv ? 4/5 = 11
	// because 5*11 mod 17 = 4
	z := Z17.Div(x, y)

	fmt.Println(z.String())

}
