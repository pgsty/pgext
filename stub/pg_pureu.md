## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/CHANGELOG.md)

`pg_pureu` executes untrusted Pure scalar functions inside PostgreSQL and exposes SPI database access. It belongs to distribution `pg_pure` 0.1.1, with SQL version 0.1. Only superusers may create functions in this language: Pure can call arbitrary native functions through `extern`.

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

`sql::query` and `sql::exec` access the database; `sql::notice` and `sql::warning` report messages. A database error becomes `sql::error` carrying SQLSTATE and message, and an uncaught exception becomes a SQL error. Parameterized SPI, set-returning functions, triggers, arrays and composites are not implemented in 0.1; some README examples are explicitly roadmap examples. Do not interpolate untrusted values into SQL strings.

### Configuration and Failure Boundaries

The shared library is `pg_pure`, including for this sibling. Optional `shared_preload_libraries` loading warms the interpreter in the postmaster and needs a restart; otherwise initialization happens in each backend. `pg_pure.using` defaults to `math`. Query cancellation and stack-overflow hooks are implemented; global definitions made with `eval` are not transactional and survive rollback.

Avoid mixing different LLVM major versions between libpure and PostgreSQL JIT in one backend: upstream documents a backend-exit crash with cluster-wide consequences. Its supported 0.1 setup uses PostgreSQL without LLVM JIT. The trusted `pg_pure` sibling remains an execution-refusing stub pending a sandbox; installing it does not make this language safe for untrusted authors.
