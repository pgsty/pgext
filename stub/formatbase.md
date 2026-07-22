## Usage

Sources:

- [Pinned upstream README](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/README.md)
- [Pinned version 2.0.0 installation SQL](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase--2.0.0.sql)
- [Pinned C implementation](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase.c)
- [Pinned extension control file](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase.control)

`formatbase` version `2.0.0` converts signed integers to text in bases 2 through 64 and parses that text back into `smallint`, `integer`, or `bigint`. It is useful for compact identifiers and explicit radix conversion; it is not arbitrary-precision arithmetic.

### Core Workflow

```sql
CREATE EXTENSION formatbase;

SELECT text(255::bigint, 16);       -- ff
SELECT text(-32602::bigint, 36);    -- -p5m
SELECT int4('ff', 16);              -- 255
SELECT int8('-p5m', 36);            -- -32602
SELECT int2(text(7::smallint, 2), 2);
```

Bases 2 through 36 use digits followed by lower-case letters and parse letters case-insensitively. Bases 37 through 64 are case-sensitive and extend the alphabet with upper-case letters, `_`, and `` ` ``.

### Important Objects

- `text(bigint, integer)`, `text(integer, integer)`, and `text(smallint, integer)`: encode a signed integer using the requested base.
- `int8(text, integer)`, `int4(text, integer)`, and `int2(text, integer)`: decode text and return the named integer type.

### Operational Notes

All functions are `IMMUTABLE` and `STRICT`, so they can be used in generated expressions and indexes and return `NULL` when any input is `NULL`. A base outside 2–64, an empty string, an invalid digit, a bare minus sign, or a value outside the result type raises an error. The function names `text`, `int2`, `int4`, and `int8` are intentionally terse and can resemble casts; schema-qualify them where search-path ambiguity matters. Preserve the chosen base alongside encoded values because the text itself does not record its radix.
