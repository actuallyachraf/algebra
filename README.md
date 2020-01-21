# Algebra

This is a small library I write as a parallel of studying algebra
and number theory related to cryptography it's not suitable for production usage
and I doubt it will ever will.

It aims to be literate.

It's a **WIP** and might take time to be finished.

## DO NOT USE FOR ANYTHING BUT LEARNING

## Packages

- ```bf``` package implements binary fields
- ```elliptic``` package implements several elliptic curve primitives and
a handful of cryptographic curves.
- ```nt``` package implements number theoretic algorithms and primitives using
arbitrary precision arithmetic
- ```ff``` package implements generic FiniteFields and Field Elements
- ```group``` package implements some custom groups such as Zp,GF(2),GF(8)...
- ```poly``` package implements polynomials over rings
- ```pairing``` package implements bilinear pairings

## References

- [Handbook Of Applied Cryptography](http://cacr.uwaterloo.ca/hac/)
- [Guide To Elliptic Curve Cryptography](http://cacr.uwaterloo.ca/ecc/)
- [Pairings For Begineers](http://www.craigcostello.com.au/pairings/PairingsForBeginners.pdf)
- [Pairings For Cryptographers](https://eprint.iacr.org/2006/165)
- [Implementing Pairing Based Cryptography](https://crypto.stanford.edu/pbc/thesis.pdf)

## WIP

- ~~Implement operations on arbitrary precision integers and tests.~~
- ~~Implement finite field elements.~~
- ~~Implement polynomial ops.~~
- Implement binary fields.
- Implement elliptic curves.
- Implement pairings.
