## Usage

Sources:

- [Pinned upstream README](https://github.com/scottjustin5000/pg-gcd/blob/20286b76a2e6d0304d19fcc6083aaea3fc2683df/README.md)
- [Version 0.0.1 installation SQL](https://github.com/scottjustin5000/pg-gcd/blob/20286b76a2e6d0304d19fcc6083aaea3fc2683df/gcd--0.0.1.sql)
- [Pinned C implementation](https://github.com/scottjustin5000/pg-gcd/blob/20286b76a2e6d0304d19fcc6083aaea3fc2683df/gcd.c)

gcd version 0.0.1 installs one C function that computes the greatest common divisor of an integer array. The SQL declaration accepts any array, but the implementation handles only smallint, integer, and bigint element types.

### Example

```sql
CREATE EXTENSION gcd;

SELECT gcd(ARRAY[14, 21, 35, 42]::integer[]);
```

The result is 7. Because the function is declared immutable and strict, PostgreSQL may fold calls with constant arrays and returns null for a null array argument.

### Input and compatibility limits

Use only nonempty arrays without null elements. The C implementation initializes its calculation from the first element and does not safely handle an empty array. It also ignores the array's null bitmap.

Although bigint arrays are accepted, every value is narrowed to a C int during calculation, so values outside the signed 32-bit range can produce incorrect results. The implementation also returns an internal 64-bit datum while SQL declares integer. Prefer smallint or integer values in the signed 32-bit range and verify edge cases.

The project provides no maintained PostgreSQL compatibility matrix or recent release history; its last source update was in 2019 and the repository is not archived. Treat it as a small legacy helper and test it against the exact PostgreSQL major version and compiler used for deployment.
