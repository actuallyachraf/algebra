package bp

// The goal of a range proof is for a verifier to prove that a certain
// value v is in range (0,2^n) without revealing any information about v.
// Range proofs are possible due to Bootle et al Inner Product argument
// The following is a construction of the proof from the Bulletproofs
// paper it assumes the reader has skimmed the paper especially the first
// sections on the inner product argument.
// Range Proof :
// Prover starts with a pair (v,Com(v)) the value and a Pedersen commitment of v.
// 1. Let a = binary_representation(v) s.t v = <a,2^n>
// where 2^n is the vector of successive powers of two.
// Our proof system will require the verifier to satisfy certain constraints
// that constitute a valid proof.
// 2. C(1) : a o (a - 1^n) = 0 , where (o : Hadamard product)
// and 1^n the identity vector.
// Satisfying C(1) is a valid proof that a is a bitvector
// 3. C(2) : aL o aR = 0 s.t aL = a and aR = a - 1^n
// For any v the last constraint are considered a valid range proof (NOT in ZK).
// To make the protocol an actual proof system we add a challenge (y,z) a pair
// of scalars and collapse the 3 constraints into a single one.
// Let y a scalar different than 0, then y^n is the vector of successive powers
// of y.
// For any vector b, b = 0 i.i.f <b,y^n> = 0
// Let z a scalar different than 0 then :
// 4. C(4) : z^2*v = z^2*<aL,2^n> + z*<aL - 1 - aR,y^n> + <aL,aR o y^n>
// Where * denotes vector scalar product (component wise).
// C(4) is a constraint system (basically a circuit) that outputs true if v is
// indeed in (0,2^n) it's a valid non-interactive proof system.
// The protocol requires (a,b) to be verified (THE proof) and that's O(2n).
// the other issue is that the verifier needs to reveal v since a = bin_rep(v).
// First we have to make this protocol zero-knowledge and second we have
// to make it succint.
// The equality C(4) can be turned into a single inner product which allows
// us to leverage the improved Inner product argument to construct a zk
// range proof.
// Let gamma(y,z) = (z-z^2)*<1,y^n>-z^3*<1^n,2^n> s.t y,z are scalars.
// We rewrite C(4) as :
// C(5) : <aL-z*1^n,y^n o (aR + z*1^n) + z^2*2^n> = z^2*v + gamma(y,z)
// Now that we collapsed our constraint system into a single inner product
// we first make it zero knowledge then construct the argument to prove
// that the inner product relation holds.
// Protocol :
// Prover given inputs v,r (r : random scalar)
// The prover computes two pedersen commitments
// * : Group operation (EC Addition in this package)
// ^ : scalar multiplication
// g ,h are vectors of generators
// sL,sR : blinding terms for the pedersen commitment
// Input(v,r)
// aL = bin_rep(v) s.t <aL,2^n> = v
// aR = aL - 1^n
// alpha <-$- Zp
// A = h^alpha * g^aL * h^aR : Pedersen commitment to aL,aR
// sL,sR <-$- Zp
// rho <-$- Zp
// S = h^rho * g^sLh^sR : Pedersen commitment to sL,sR
// Prover sends A,S to the verifier
// The verifier samples y,z uniformly from Zp and sends them to the prover.
// Let's define the arguments of our proof
// We define multivariate polynomials over Zp l(X),r(X) and a quadratic
// polynomial t(X) such as :
// l(X) = (aL - z*1^n) + sL*X
// r(X) = y^n o (aR+sR*X+z*1^n) + z^2*2^n
// t(X) = <l(X),r(X)> = t0 + t1*X + t2*X^2
// Prover now executes the following, his goal is to hide t1,t2 s.t t0 = v*z^2+gamma(y,z)
// Tau1,Tau2 <-$- Zp
// T1 = PedersenComm(t1,Tau1)
// T2 = PedersenCom(t2,Tau2)
// Verifier samples a random element
// x <-$- Zp
// Verifier sends x to the Prover
// Prover computes
//
// l = l(x) = aL - z*1^n + sL*x
// r = r(x) = y^n o (aR+sR*x+z*1^n) + z^2*2^n
// t' =  <l,r>
// Tau' = Tau2*x^2 + Tau1*x + z^2*gamma
// mu = alpha + rho*x where alpha,rho blind A,S
// Prover sends Tau',mu,t',l,r to the Verifier
// Verifier checks
// l,r,t are correct
// Verifies the equality t' = t0 + t1*x + t2*x^2
// g^t' * h^Tau' = V^(z^2) * g^gamma(y,z) * T1^x + T2^x
// Compute commitments to l,r
// P = A * S^x * g^(-z) * (h')^(z*y^n + z^2*y^n) : h' is the generators vector
// The proof has now completness, and honest verifier zero-knowledge.
