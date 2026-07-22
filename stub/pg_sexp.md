## Usage

Sources:

- [Official type and operator documentation](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/README.md)
- [Version 0.1.0 SQL definitions](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/sql/pg_sexp--0.1.0.sql)
- [Extension control file](https://github.com/gdiazlo/pg_sexp/blob/17e9ef42c856f04981e30a094b275fa3c0f3edee/pg_sexp.control)

`pg_sexp` adds a compact binary `sexp` type for Lisp-style S-expressions. Use it when PostgreSQL must store, inspect, match, and index symbolic lists without flattening them into text.

### Core Workflow

```sql
CREATE EXTENSION pg_sexp;

CREATE TABLE document (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    expr sexp NOT NULL
);

INSERT INTO document (expr)
VALUES ('(user (name "alice") (age 30))');

SELECT id, car(expr), cdr(expr)
FROM document
WHERE expr @>> '(name "alice")';

CREATE INDEX document_expr_gin
ON document USING gin (expr sexp_gin_ops);
```

### Objects and Operators

- `car(sexp)`, `cdr(sexp)`, `nth(sexp, integer)`, `head(sexp)`, and `sexp_length(sexp)` inspect lists.
- `sexp_typeof(sexp)` and the `is_list`, `is_atom`, `is_symbol`, `is_string`, `is_number`, and `is_nil` predicates inspect value kinds.
- `@>` performs structural containment; `@>>` performs key-based, order-independent containment; `@~` and `sexp_find` support wildcard, rest, and capture patterns.
- `sexp_gin_ops` supports GIN containment searches, while equality has hash-index support.

### Compatibility and Caveats

Upstream requires PostgreSQL 14 or later. Version `0.1.0` is pre-1.0, so pin the binary format and extension version before storing durable data. Test parser canonicalization, symbol versus string and integer versus float distinctions, duplicate keys, nesting depth, wildcard complexity, and agreement between indexed and sequential results. Deeply nested or adversarial expressions and broad wildcard searches may be expensive; apply input and query limits at the application boundary.
