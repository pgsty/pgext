## Usage

Sources:

- [Pinned official README](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/README.md)
- [Pinned extension SQL](https://github.com/DanielJDufour/safecast/blob/0ad39b4f99d373f309817ba803374b5b2c9229a3/safecast--0.0.1.sql)

`safecast` supplies small SQL functions that return `NULL` for some text values that do not match a numeric-looking regular expression before casting. Despite its name, version 0.0.1 does not catch cast exceptions and is not safe for arbitrary untrusted input.

### Core Workflow

```sql
CREATE EXTENSION safecast;

SELECT to_integer('42');
SELECT to_integer('not-a-number');
SELECT to_float('-12.5');
SELECT to_bigint('9000000000');
SELECT to_double_precision('123');
```

The second call returns `NULL`; accepted strings are passed to PostgreSQL's normal cast implementation.

### Functions and Accepted Shapes

- `to_integer(text)`: accepts only one or more ASCII digits, then casts to `integer`.
- `to_bigint(text)`: accepts only one or more ASCII digits, then casts to `bigint`.
- `to_float(text)`: accepts any nonempty combination of `-`, digits, and `.`, then casts to `float` (`double precision`).
- `to_double_precision(text)`: accepts only one or more ASCII digits, then casts to `double precision`.

### Caveats

Range overflow still raises an error. The `to_float` regular expression is permissive enough to accept malformed values such as `.` or `--1`, which then raise cast errors. The integer functions reject leading signs, whitespace, decimal points, and exponent notation; `to_double_precision` rejects signs, fractions, and exponents as well. Locale-specific numbers and `NaN`/`Infinity` are not accepted.

The functions are not declared `STRICT` or with volatility/parallel attributes, and the abandoned project has a very small test surface. Use explicit validation or an exception-catching function when failure must reliably become `NULL`. The extension is SQL-only and needs no preload or restart.
