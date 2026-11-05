## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.4.0/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.4.0/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.4.0/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.4.0/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.4.0/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.4.0/CHANGELOG.md)

`pg_pureu` executes untrusted Pure scalar functions inside PostgreSQL and exposes SPI database access. It belongs to distribution `pg_pure` 0.4.0, with SQL version 0.1. Only superusers may create functions in this language: Pure can call arbitrary native functions through `extern`.

### Core Workflow

```sql
CREATE EXTENSION pg_pureu;
CREATE FUNCTION pure_factorial(n int) RETURNS numeric
LANGUAGE pg_pureu AS $$ foldl (*) 1L (1..n) $$;
SELECT pure_factorial(30);
```

The extension is non-relocatable. It requires PostgreSQL 17–19, the embeddable libpure fork and its prelude. `numeric` values use exact Pure big integers or rationals; a fractional result is rounded to PostgreSQL's numeric display scale on return. Machine integer arithmetic can wrap before the final range-checked conversion, so use big integers when intermediate values can grow.

### Objects and Database Access

The language installs call, inline and validation handlers. Scalar SQL types include smallint/integer/bigint, floating point, numeric, boolean and text. SQL NULL is represented by `sql::null`; a text result must be a Pure string.

`sql::query` and `sql::exec` access the database; `sql::notice` and `sql::warning` report messages. A database error becomes `sql::error` carrying SQLSTATE and message, and an uncaught exception becomes a SQL error. Distribution 0.4.0 supports dense one- and two-dimensional `float8[]` arrays as Pure double matrices; NULL elements and higher dimensions are rejected. Parameterized SPI, set-returning functions, triggers and composites remain unimplemented. Do not interpolate untrusted values into SQL strings.

### Configuration and Failure Boundaries

The shared library is `pg_pure`, including for this sibling. Optional `shared_preload_libraries` loading warms the interpreter in the postmaster and needs a restart; otherwise initialization happens in each backend. `pg_pure.using` defaults to `math`. Query cancellation and stack-overflow hooks are implemented; global definitions made with `eval` are not transactional and survive rollback.

The embeddable libpure fork at revision ab12de13 or later hides its LLVM symbols and supports normal PostgreSQL JIT. The old JIT incompatibility guidance no longer applies. The trusted `pg_pure` sibling now executes in a separate safe-mode interpreter; this does not restrict the native-code powers of the untrusted language. The compiled `gsl::linalg::cholesky_solve` solver in 0.4.0 needs the matching Pure module. Untrusted bodies may import installed modules with `using`; trusted bodies rely on the DBA's allowlist. Deploy matching libraries and replace sessions that loaded an older version; SQL version 0.1 alone does not identify the distribution release.
