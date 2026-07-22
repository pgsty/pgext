## Usage

Sources:

- [Official README](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/README.md)
- [Official extension SQL](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal--1.0.sql)
- [Official C implementation](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal.c)
- [Official control file](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal.control)

`decimal` is a 2013 experiment exposing compiler `_Decimal32` and `_Decimal64` values as fixed-size PostgreSQL types named `decimal32` and `decimal64`. It was written to compare fixed decimal arithmetic with PostgreSQL `numeric`, not as a complete or production-grade numeric implementation.

### Core Workflow

```sql
CREATE EXTENSION decimal;

SELECT '3.14'::decimal32 + '2.00'::decimal32;
SELECT '12.3456'::decimal64 * '2'::decimal64;
SELECT round('12.3456'::decimal64, 2);

SELECT sum(v)
FROM (VALUES ('1.25'::decimal64), ('2.75'::decimal64)) AS x(v);
```

Arithmetic between two `decimal32` values returns `decimal64`. Arithmetic between two `decimal64` values returns `decimal64`; the only aggregate is `sum(decimal64)`.

### Casts and Missing Surface

Implicit casts accept common integer types, `float4` to `decimal32`, `float8` to `decimal64`, `numeric` to either decimal type, and `decimal32` to `decimal64`. A `decimal64` value can be cast back to `numeric`; narrowing to `decimal32` is explicit.

The extension provides `+`, `-`, `*`, `/`, and two-argument `round`, but no comparison/equality operators, hash support, B-tree operator class, `avg`, typmod precision/scale, or documented NaN/infinity policy. It cannot replace `numeric` in indexed keys or general financial constraints.

### Known Implementation Hazards

The reviewed input functions call `strtod`/`strtold` but do not verify that the entire string was consumed. Inputs with trailing junk, and potentially empty input, can be accepted as a numeric prefix instead of rejected. Validate input before casting.

The `decimal64` multiplication overflow check divides the result by the left operand. When that operand is zero, valid multiplication can fail because the check evaluates a zero-by-zero expression. The arithmetic checks are incomplete and inconsistent; do not rely on them for correctness or overflow protection.

The `decimal32` rounding implementation accepts negative scale while the `decimal64` version rejects it, and both route through binary floating-point `pow`/`round`. Test every required precision and boundary case.

Version `1.0` has a single 2013 source revision and depends on nonportable compiler decimal extensions and their ABI. Use PostgreSQL `numeric` for maintained exact arithmetic; adopt this extension only for isolated research after compiling and testing the exact compiler/architecture/PostgreSQL combination.
