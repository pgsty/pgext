## Usage

Sources:

- [Version 1.0 SQL API](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex--1.0.sql)
- [C implementation](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex.c)
- [Extension control file](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex.control)

`complex` defines a 16-byte complex-number type backed by two `float8` components. Version 1.0 provides arithmetic and comparison operators, implicit casts from `int4` and `float8`, `sum`, `min`, `max`, `avg`, and `complex_pwrt(complex,int4)`.

```sql
CREATE EXTENSION complex;
SELECT '(1,2)'::complex + '(3,-4)'::complex;
SELECT avg(z) FROM measurements;
```

The SQL API defines a total ordering for `<`, `<=`, `>`, `>=`, `min`, and `max`, but complex numbers have no natural mathematical order. Inspect and test the implementation's chosen ordering before using it for business logic or sorted indexes. Floating-point equality also inherits NaN, infinity, signed-zero, rounding, and platform concerns.

The repository has no user README or release notes. Establish accepted text syntax, power-function row semantics, divide-by-zero behavior, overflow, binary dump/restore, and PostgreSQL-major compatibility on the exact build. Use core numeric pairs or an application math library when portability and a documented numerical contract matter.
