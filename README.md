# Algebra

This is a small library I write as a parallel of studying algebra
and number theory related to cryptography it's not suitable for production usage
and I doubt it ever will.

It aims to be literate.

It's a **WIP** and might take time to be finished.

## DO NOT USE FOR ANYTHING BUT LEARNING

## Packages

### Cryptography Implementations

- ```crypto/schnorr``` package implements Vanilla EC-Schnorr.
- ```crypto/bp``` package implements [Bulletproofs](https://eprint.iacr.org/2017/1066).

### Algebraic Tools Implementations

- ```bf``` package implements binary fields.
- ```ec``` package implements elliptic curve primitives and
a few cryptographic curves.
- ```nt``` package implements number theoretic algorithms and primitives using
arbitrary precision arithmetic.
- ```ff``` package implements generic finite fields and field elements.
- ```group``` package implements some custom groups such as Zp,GF(2),GF(8)...
- ```poly``` package implements polynomials over rings.
- ```pairing``` package implements bilinear pairings.

## References

- [Handbook Of Applied Cryptography](http://cacr.uwaterloo.ca/hac/)
- [Guide To Elliptic Curve Cryptography](http://cacr.uwaterloo.ca/ecc/)
- [Pairings For Begineers](http://www.craigcostello.com.au/pairings/PairingsForBeginners.pdf)
- [Pairings For Cryptographers](https://eprint.iacr.org/2006/165)
- [Implementing Pairing Based Cryptography](https://crypto.stanford.edu/pbc/thesis.pdf)
- [Modern Computer Algebra](https://www.cambridge.org/core/books/modern-computer-algebra/DB3563D4013401734851CF683D2F03F0)

## WIP

- ~~Implement operations on arbitrary precision integers and tests.~~
- ~~Implement finite field elements.~~
  - Optimized version instead of wrapping bigint
- ~~Implement polynomial ops.~~
  - Optimized FFT instead of naive Eval/Mul algorithms
- ~~Implement elliptic curves.~~
  - Add projective coordinates support
  - Support typed curves (Weirstrass,Edwards)
  - Implement optimized formulas for Weirstrass curves
- Implement binary fields.
- Implement groups for char 2 fields.
- Implement pairings.
