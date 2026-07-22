## Usage

Sources:

- [Extension control file](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/pg/utils/utils.control)
- [Extension SQL API](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/pg/utils/utils--0.0.1.sql)
- [C implementation](https://github.com/9bany/l-learn/blob/49ddb263f4996a1ae227b181a4f16253da7fa010/pg/utils/utils.c)

`utils` 0.0.1 is a two-function C learning example from the `l-learn` monorepo. It adds or multiplies two PostgreSQL `integer` values. Despite its broad name, it contains no general utility collection.

### Available API

```sql
CREATE EXTENSION utils;

SELECT addme(20, 22);
SELECT multiply(6, 7);
```

`addme(integer, integer) RETURNS integer` performs C `int32` addition. `multiply(integer, integer) RETURNS integer` performs C `int32` multiplication. Both functions are `STRICT`, so a NULL argument returns NULL without calling C code.

### Limitations

The SQL declarations do not mark the functions `IMMUTABLE` or `PARALLEL SAFE`, even though the arithmetic has no intended database side effect. More importantly, the implementation uses raw signed C arithmetic without PostgreSQL's normal integer-overflow checks; out-of-range input can invoke undefined C behavior rather than raise the standard `integer out of range` error.

Use PostgreSQL's built-in `+` and `*` operators for real work. The generic extension and function names can collide with unrelated objects, and the monorepo provides no compatibility matrix or extension-specific user documentation. Treat it only as a minimal PGXS/C example after testing the exact server build.
