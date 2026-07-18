## Usage

Sources:

- [Current upstream type and operator documentation](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/README.md)
- [Version 0.1.0 SQL definitions](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/sql/pg_sexp--0.1.0.sql)
- [Extension control file](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/pg_sexp.control)

`pg_sexp` adds a compact binary `sexp` type for Lisp-style S-expressions. It provides list accessors, type inspection, structural and key-based containment, wildcard matching, equality hashing, and GIN operator classes.

```sql
CREATE EXTENSION pg_sexp;
SELECT car('(a b c)'::sexp), cdr('(a b c)'::sexp);
SELECT '(user (name "alice") (age 30))'::sexp @>> '(name "alice")';
CREATE INDEX expr_gin ON documents USING gin (expr sexp_gin_ops);
```

Upstream declares PostgreSQL 14 or later. Before relying on indexes, test parser canonicalization, numeric/string distinctions, duplicate keys, nesting depth, wildcard complexity, and operator/index agreement. Treat deeply nested or adversarial expressions as potentially expensive and pin the extension version because the format is pre-1.0.
