package schnorr

import (
	"testing"

	"github.com/actuallyachraf/algebra/ec"
	"github.com/actuallyachraf/algebra/ff"
	"github.com/actuallyachraf/algebra/nt"
)

// Test parameters are those of secp256k1 curve
var (
	secp256k1GeneratorX, _     = new(nt.Integer).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	secp256k1GeneratorY, _     = new(nt.Integer).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	secp256k1GeneratorOrder, _ = new(nt.Integer).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	fieldOrder, _              = new(nt.Integer).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	secp256k1A                 = new(nt.Integer).SetInt64(0)
	secp256k1B                 = new(nt.Integer).SetInt64(7)
)

func TestSchnorr(t *testing.T) {
	Fq, err := ff.NewFiniteField(fieldOrder)
	if err != nil {
		t.Fatal("bad field order for secp256k1")
	}
	secp256k1 := ec.NewEllipticCurve(Fq.NewFieldElement(secp256k1A), Fq.NewFieldElement(secp256k1B), Fq)
	secp256k1Generator := ec.Point{X: secp256k1GeneratorX, Y: secp256k1GeneratorY}
	parameters := Params{
		EC:    *secp256k1,
		Gen:   secp256k1Generator,
		Order: secp256k1GeneratorOrder,
	}
	t.Run("TestGenerator", func(t *testing.T) {
		// Test that the generator is on the curve
		if !secp256k1.IsOnCurve(&secp256k1Generator) {
			t.Fatal("bad generator for secp256k1")
		}
		// Test whether a scalar multiple is on the curve
		s, _ := Fq.Rand()
		if !secp256k1.IsOnCurve(secp256k1.ScalarMul(&secp256k1Generator, s.Big())) {
			t.Fatal("generated points should be on curve")
		}
	})
	t.Run("TestSignVerify", func(t *testing.T) {

		msg := []byte("helloworld")
		kp := GenerateKeypair(parameters)
		t.Log("Public Key :", kp.P, " | Private Key :", kp.K)
		t.Log("Check Public Key on Curve")
		if !secp256k1.IsOnCurve(kp.PublicKey.P) {
			t.Fatal("bad public key for secp256k1")
		}
		sig := Sign(msg, &parameters, kp)

		if !Verify(msg, sig, kp, parameters) {
			t.Fatal("bad signature")
		}
	})

}
