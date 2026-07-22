## Usage

Sources:

- [Official README](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/README.md)
- [Version 1.0 extension SQL](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/pg_testgen--1.0.sql)
- [C implementation](https://github.com/yuesong-feng/pg_testgen/blob/5a6c127a1596eb8fa678fc0df6608ee9d83cd115/pg_testgen.c)

`pg_testgen` generates random integers and alphanumeric text, either as scalar values or set-returning functions. It is intended to populate development and test tables quickly, not to generate credentials, security tokens, or statistically rigorous samples.

### Core Workflow

```sql
CREATE EXTENSION pg_testgen;

CREATE TABLE test_rows (id integer, txt text);

INSERT INTO test_rows
SELECT rows_int(5000, 1, 100),
       rows_text(5000, 20, 30);
```

The example inserts 5,000 rows: integer values are in the inclusive range `[1, 100]`, and strings contain 20 to 30 ASCII letters or digits.

### Object Index

- `rand_int()` returns a nonnegative 32-bit-range random integer.
- `rand_int(a, b)` returns an integer in inclusive range `[a, b]`.
- `rand_text()`, `rand_text(length)`, and `rand_text(min_length, max_length)` return random ASCII alphanumeric text.
- `rows_int(count [, min, max])` returns `count` random integers.
- `rows_text(count [, length])` and `rows_text(count, min_length, max_length)` return `count` random strings.

### Operational Notes

Version `1.0` declares no extension dependency or preload requirement. The implementation uses the C library `rand()` with modulo arithmetic; it is not cryptographically secure, does not promise independent backend seeding, and can be biased. Never use it for passwords, keys, identifiers that require uniqueness, or benchmark-quality distributions.

Validate arguments in the caller. The C code does not safely reject a reversed/overflowing integer range, negative text length, or invalid row count; such inputs can cause errors or unsafe allocation behavior. For large data sets, batch inserts and monitor WAL, table size, indexes, and transaction duration.
