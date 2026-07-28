## Usage

Sources:

- [Official upstream README](https://github.com/gdiazlo/pg_sexp/blob/6bf80b142756273278610fa9ff66472db4cc7f98/README.md)
- [Official extension control file (pg_sexp_rs.control)](https://github.com/gdiazlo/pg_sexp/blob/6bf80b142756273278610fa9ff66472db4cc7f98/rs/pg_sexp_rs.control)
- [Official implementation source](https://github.com/gdiazlo/pg_sexp/blob/6bf80b142756273278610fa9ff66472db4cc7f98/rs/src/lib.rs)

`pg_sexp_rs` — PostgreSQL extension to add support for s-expressions store, query, index. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_sexp_rs;

-- Atoms
SELECT 'hello'::sexp;           -- symbol
SELECT '42'::sexp;              -- integer
SELECT '3.14'::sexp;            -- float
SELECT '"hello world"'::sexp;   -- string
SELECT '()'::sexp;              -- nil

-- Lists
SELECT '(a b c)'::sexp;
SELECT '(define x 10)'::sexp;
SELECT '(lambda (x) (* x x))'::sexp;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `sexp_car` is an extension function.
- `sexp_cdr` is an extension function.
- `sexp_contains` is an extension function.
- `sexp_contains_key` is an extension function.
- `sexp_eq` is an extension function.
- `sexp_extract_keys` is an extension function.
- `sexp_extract_query_keys` is an extension function.
- `sexp_find` is an extension function.
- `sexp_gin_consistent_fn` is an extension function.
- `sexp_gin_extract_query_fn` is an extension function.
- `sexp_gin_extract_value_fn` is an extension function.
- `sexp_gin_triconsistent_fn` is an extension function.
- `sexp_hash` is an extension function.
- `sexp_hash_extended` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
