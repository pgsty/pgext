## Usage

Sources:

- [pgmp 1.0.6 README](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/README.rst)
- [pgmp 1.0.6 release notes](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/NEWS.rst)
- [pgmp 1.0.6 metadata](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/META.json)
- [pgmp control file](https://github.com/dvarrazzo/pgmp/blob/rel-1.0.6/pgmp.control)
- [Official pgmp documentation](https://dvarrazzo.github.io/pgmp/)

`pgmp` exposes GNU MP arithmetic inside PostgreSQL. It adds arbitrary-size integer values through `mpz` and exact rational values through `mpq`, together with casts, arithmetic, comparison, aggregate, number-theory, bit, and random-number functions.

### Core Workflow

```sql
CREATE EXTENSION pgmp;

SELECT '123456789012345678901234567890'::mpz * 2;
SELECT mpq(1::mpz, 3::mpz) + mpq(1::mpz, 6::mpz);
SELECT gcd(48::mpz, 18::mpz);
SELECT nextprime(100000000000000000000::mpz);
```

`mpz` is an arbitrary-size integer type, subject to PostgreSQL's value-size limits. `mpq` stores a canonical numerator and denominator so fractional arithmetic remains exact until explicitly converted to an approximate type.

### Important Objects

- `mpz(text)` and casts construct integers in decimal or supported base-prefixed forms.
- `mpq(text)` and `mpq(mpz, mpz)` construct rational values.
- Both types support ordinary comparisons and btree or hash indexes.
- Integer helpers include division with explicit rounding modes, powers, roots, primality tests, `gcd`, `lcm`, factorials, Fibonacci and Lucas numbers, bit operations, and random-state functions.
- Rational helpers include numerator and denominator access, inversion, denominator limiting, arithmetic, comparison, and aggregates.
- `gmp_version()` and `gmp_max_bitcnt()` expose library information.

Do not use floating-point input when exact decimal or rational meaning matters; construct values from text, integers, or explicit numerator and denominator values.

### Version 1.0.6 Notes

The 1.0.6 distribution adds PostgreSQL 19 build compatibility, sets PostgreSQL 14 as the supported runtime floor in its metadata, and adds missing unsigned-long range checks for the power, Fibonacci, and Lucas-number paths.

The upstream distribution version is 1.0.6, while its tagged `pgmp.control` currently declares SQL extension version `1.1`. Create the extension without forcing a version and inspect the database-reported value before designing an upgrade:

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmp';
```

pgmp requires the GMP shared library. GMP 4.1 lacks a few functions documented by upstream, including some root, bit, and random-state helpers; use a current GMP release when those objects are required. Large operands can consume substantial backend memory and CPU, so apply statement timeouts and input limits to untrusted arithmetic workloads.
