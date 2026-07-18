## Usage

Sources:

- [Upstream README](https://github.com/tomhebbron/pg_bitstring_helpers/blob/43011651dd5c600e7b9cb413d0ed3ea32683a1bd/README.rst)
- [Version 1.0 install SQL](https://github.com/tomhebbron/pg_bitstring_helpers/blob/43011651dd5c600e7b9cb413d0ed3ea32683a1bd/bitstring_helpers--1.0.sql)
- [C implementation](https://github.com/tomhebbron/pg_bitstring_helpers/blob/43011651dd5c600e7b9cb413d0ed3ea32683a1bd/bitstring_helpers.c)

bitstring_helpers 1.0 adds bit-counting, Hamming-distance, one-bit-neighbour, concatenation, conversion, and integer-shuffling helpers.

### Bit-string operations

```sql
CREATE EXTENSION bitstring_helpers;
SELECT popcount(B'101101'::bit varying);
SELECT hamming_distance(B'101101'::bit varying, B'100001'::bit varying);
SELECT * FROM neighbours(B'101'::bit varying);
```

Use equal-length operands for Hamming distance. The neighbours function returns one value for every possible single-bit flip.

### Caveats

- shuffled_ints produces values from zero through limit minus one, despite the README describing a range starting at one.
- A negative shuffle limit is used in an allocation before validation, and a large limit can exhaust backend memory. Never pass untrusted or unbounded values.
- The shuffle uses a modulo-based process-global random generator and is biased; it is unsuitable for cryptographic, security, sampling, or fairness-sensitive work.
- This 2013 code has no current PostgreSQL compatibility matrix, tests, or published license. Its legacy polymorphic aggregate declarations may require changes on modern PostgreSQL releases.
