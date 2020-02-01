# Bulletproofs : NIZK Range Proofs

[Bulletproofs](https://eprint.iacr.org/2017/1066.pdf) are zero-knowledge arguments
of knowledge that are non-interactive, transparent, succint and efficent.

Bulletproofs can prove that a value *v* is in range [0,2^n] usinng 2+log(n) Group
and Field elements. Moreover proofs can be aggregated and batch-verified which
is computationally efficient.

Bulletproofs also provide zero-knowledge proofs for general arithmetic circuits (the
general case for zk-SNARKs).

This implementation **is not production oriented** there might be bugs and doesn't
apply optimizations.
