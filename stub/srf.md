## Usage

Sources:

- [Version 1.0 extension SQL](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf--1.0.sql)
- [C implementation](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf.c)
- [Extension control file](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf.control)

`srf` is an abandoned demonstration extension containing one C set-returning function. `srf_c` emits an inclusive sequence of 32-bit integers and is useful only for studying PostgreSQL's multi-call C function API.

### Core Workflow

```sql
CREATE EXTENSION srf;

SELECT * FROM srf_c(3, 7);
```

The query returns `3`, `4`, `5`, `6`, and `7`. When `start` is greater than `end`, it returns no rows.

### Object Index

- `srf_c(start integer, "end" integer) RETURNS SETOF integer` emits each integer from `start` through `end`, inclusive.

### Operational Notes

Version `1.0` is relocatable and declares no dependency or preload requirement. The SQL function has no volatility, strictness, or parallel-safety annotation, so PostgreSQL's defaults apply.

This repository has no user documentation or release support and is cataloged as abandoned. Do not use it as a production replacement for `generate_series`. Avoid an `end` value of `2147483647`: the C implementation increments its signed 32-bit counter after returning the endpoint and does not guard overflow. Prefer PostgreSQL's built-in `generate_series(integer, integer)` for application queries.
