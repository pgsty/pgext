## Usage

Sources:

- [Official PGXN documentation](https://pgxn.org/dist/bignum/doc/bignum.html)
- [Pinned extension SQL](https://github.com/beargiles/pg-bignum/blob/79950064ca96595dd0a26c81915ea2d7396e6986/sql/bignum.sql)

`bignum` adds an arbitrary-precision integer type backed by OpenSSL's `BIGNUM` implementation. Use it when values must exceed PostgreSQL's built-in `bigint` range and integer arithmetic is sufficient; it is not a replacement for `numeric` or `decimal`.

### Core Workflow

Create the extension, cast or write integer literals as `bignum`, and use ordinary arithmetic and comparison operators:

```sql
CREATE EXTENSION bignum;

SELECT '170141183460469231731687303715884105727'::bignum + 1;
SELECT 12::bignum * 12::bigint;
SELECT gcd(48::bignum, 18::bignum);
```

Assignment casts from `int4` and `int8` are provided, and mixed `bignum`/`int8` forms are defined for the arithmetic and comparison operators.

### Important Objects

- Type: `bignum`.
- Arithmetic: unary `-`, `+`, `-`, `*`, `/`, and `%`.
- Comparison: `=`, `<>`, `<`, `<=`, `>=`, and `>`.
- Functions: `abs(bignum)`, `gcd(bignum, bignum)`, and mixed `gcd` overloads.
- Indexing: the default `bignum_ops` B-tree operator class supports ordering and B-tree indexes.

### Operational Notes

The module is a C extension linked against OpenSSL; the server package must provide the matching shared library. It is relocatable and does not require `shared_preload_libraries` or a restart.

The upstream implementation does not provide casts to or from `numeric`/`decimal`, cryptographic functions, hash-index support, or `min`, `max`, `sum`, and `avg` aggregates for `bignum`. Division is integer division. Test mixed-type expressions explicitly because only the documented `int4`/`int8` conversion paths are defined.
