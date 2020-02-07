package bp

import (
	"fmt"
	"testing"

	"github.com/actuallyachraf/algebra/nt"
)

func TestBulletProofs(t *testing.T) {
	t.Run("TestParams", func(t *testing.T) {
		params := GenParametersSecp256k1(64)
		fmt.Println(params)
	})
	t.Run("TestVector", func(t *testing.T) {

		testCases := [][]*nt.Integer{
			{nt.FromInt64(1), nt.FromInt64(1), nt.FromInt64(1)},
			{nt.FromInt64(0), nt.FromInt64(3), nt.FromInt64(7)},
			{nt.FromInt64(3), nt.FromInt64(8), nt.FromInt64(5)},
			{nt.FromInt64(1), nt.FromInt64(4), nt.FromInt64(8)},
		}
		v1 := NewVector(testCases[0])
		v2 := NewVector(testCases[1])

		w, err := v1.Add(v2)
		if !w.Equal(testCases[3]) || err != nil {
			t.Error("add failed with error", err)
		}

		u, err := v1.InnerProdMod(v2, nt.FromInt64(50))
		if u.Cmp(nt.FromInt64(10)) != 0 || err != nil {
			t.Error("inner product failed with expected", 10, "got", u, "error", err)
		}
	})
}
