// Package gauss implements discrete gaussian sampling from :
// ref https://eprint.iacr.org/2013/383.pdf
package gauss

import (
	"crypto/rand"
	"encoding/binary"
)

var lsOne float64 = 1 << 31

// Bernoulli returns a random 1/0 drawn from Bernoulli Distribution.
func Bernoulli(p float64) uint {
	// discretize the p value by doing a multiple by left shift
	pInt := uint32(p * lsOne)
	x := RandInt(32)
	if x < pInt {
		return 1
	}
	return 0
}

// RandInt generates a random integer with bounded bitlength.
func RandInt(bitlength uint32) uint32 {
	// we can't use Go's default random function
	// because we can't specify the bitlength
	// generate a mask for the given bitlength
	var mask uint32 = 1<<bitlength - 1
	// generate a 4 byte integer
	b := make([]byte, 4)
	n, err := rand.Read(b)
	if n != 4 || err != nil {
		panic(err)
	}
	// store it as a little endian uint32
	x := binary.LittleEndian.Uint32(b)
	return mask & x
}
