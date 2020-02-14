package bp

import (
	"testing"

	"github.com/actuallyachraf/algebra/nt"
)

func TestBulletProofs(t *testing.T) {
	t.Run("TestParams", func(t *testing.T) {
		GenParametersSecp256k1(64)
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

	t.Run("TestPedersenCommitment", func(t *testing.T) {
		params := GenParametersSecp256k1(64)

		A := nt.FromInt64(12345)
		B := nt.FromInt64(98765)
		C := nt.Add(A, B)

		comA, r1 := PedersenCommitment(params, A)
		comB, r2 := PedersenCommitment(params, B)
		comC := pedersenCom(params, params.G, params.H, C, nt.Mod(nt.Add(r1, r2), params.L))

		if !params.EC.Add(comA, comB).Equal(comC) {
			t.Error("homomorphism property non-preserved")
		}

	})
	t.Run("TestInnerProductArgument", func(t *testing.T) {
		curveParams := GenParametersSecp256k1(64)
		a := NewVector([]*nt.Integer{nt.FromInt64(1), nt.FromInt64(2), nt.FromInt64(3), nt.FromInt64(4)})
		b := NewVector([]*nt.Integer{nt.FromInt64(8), nt.FromInt64(7), nt.FromInt64(6), nt.FromInt64(5)})
		c, _ := a.InnerProdMod(b, curveParams.L)
		P := DoubleVectorPedersenCommitment(curveParams, a, b)
		arg := ProveInnerProdArg(curveParams, a, b, c, P, curveParams.U, curveParams.GVec, curveParams.HVec)

		ok, err := VerifyInnerProdArg(curveParams, c, P, curveParams.U, curveParams.GVec, curveParams.HVec, *arg)
		if !ok || err != nil {
			t.Error("failed to verify inner product argument")
		}
	})
}
