package main

import (
	"fmt"
	"math/big"

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

	var x = new(nt.Integer).SetInt64(13)
	var y = new(nt.Integer).SetInt64(16)
	var modulus = new(nt.Integer).SetInt64(25)
	fmt.Println(nt.ModMul(x, y, modulus))

	// Let's generate the multiplicative group Z modulo 25
	// this is the set of coprimes less than 25
	Zmod25 := make([]*nt.Integer, 20)
	// Zn = {g^i mod n | 0 <= i <= phi(n)-1}
	// phi is the euler totient function
	// the generator of this group is 2
	var generator = new(nt.Integer).SetInt64(2)
	var i int64
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
	modulus.SetInt64(1234)

	fmt.Println(nt.ModExp(a, k, modulus))
}
