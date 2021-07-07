package ntt

import (
	"testing"

	"github.com/actuallyachraf/algebra/nt"
)

type argFactors struct {
	q       *nt.Integer
	factors []*nt.Integer
}

var factorVec = []argFactors{
	{nt.FromInt64(7680), []*nt.Integer{nt.FromInt64(2), nt.FromInt64(3), nt.FromInt64(5)}},
}

func TestGetFactors(t *testing.T) {
	for i, testPair := range factorVec {
		factors := getFactors(testPair.q)
		t.Log("factors found !", len(factors))
		t.Log("factors :", factors)
		for j := range factors {
			t.Log("test-factors :", testPair.factors[j])
			t.Log("found-factors :", factors[j])
			if !nt.Equal(factors[j], testPair.factors[j]) {
				t.Errorf("factor not match in test pair %v expected %d got %v", i, j, testPair.factors[j])
			}
		}
	}
}

type argRoots struct {
	q, root *nt.Integer
}

var rootsVec = []argRoots{
	{nt.FromInt64(7681), nt.FromInt64(17)},
	{nt.FromInt64(1152921504382476289), nt.FromInt64(11)},
}

func TestPrimitiveRoot(t *testing.T) {
	for i, testPair := range rootsVec {
		root := primitiveRoot(testPair.q)
		if !nt.Equal(root, testPair.root) {
			t.Errorf("primitive root not match %v", i)
		}
	}
}
