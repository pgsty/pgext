## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.6.0/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.6.0/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.6.0/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.6.0/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.6.0/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.6.0/CHANGELOG.md)

`pg_pure` provides the trusted Pure procedural language and SPI database access. Distribution 0.6.0 retains SQL extension version 0.1. Since 0.2.0, trusted functions execute in a libpure safe-mode interpreter; the earlier execution refusal no longer applies.

### Core Workflow

```sql
CREATE EXTENSION pg_pure;
CREATE FUNCTION pure_factorial(n int) RETURNS numeric
LANGUAGE pg_pure AS $$ foldl (*) 1L (1..n) $$;
SELECT pure_factorial(30);
```

The extension is trusted and non-relocatable. Ordinary roles with the required database and schema privileges may install it and define functions in the trusted language. PostgreSQL 17–19 and the upstream embeddable libpure fork at revision ab12de13 or later are required.

### Types and Database Access

Scalar values include integer, floating point, numeric, boolean and text. SQL NULL is represented by `sql::null`. Exact `numeric` values use Pure big integers or rationals; use big integers for intermediate arithmetic that can exceed machine range.

`float8[]` maps to a dense Pure double matrix: a one-dimensional array becomes a row vector and a two-dimensional array becomes a matrix. NULL elements and arrays of more than two dimensions are rejected. `sql::query`, `sql::exec`, `sql::notice` and `sql::warning` provide database access and messages; database failures become `sql::error`. Parameterized SPI, set-returning functions, triggers and composite types are not implemented. Do not interpolate untrusted values into SQL strings.

### Sandbox and Modules

Safe mode refuses `extern`, inline C, `eval`/`evalcmd`, arbitrary native or non-allowlisted `using`, and dangerous prelude operations. Trusted and untrusted languages use separate backend-local interpreters. The untrusted sibling `pg_pureu` retains native-code capability and must remain restricted to superusers.

The DBA chooses modules through `pg_pure.using`, whose default is `math`. For the compiled GSL Cholesky solver in 0.6.0, install the matching Pure module and preload/allowlist it through the setting:

```ini
pg_pure.using = 'math gsl::linalg'
```

A trusted body may call `gsl::linalg::cholesky_solve` directly after this setup. A namespaced `using gsl::linalg` inside a trusted body is still rejected by libpure's module-path check; use the DBA-managed setting.

### Runtime and Upgrade

Both languages load the `pg_pure` library. Optional `shared_preload_libraries` loading warms the interpreter and requires a restart; without it, initialization happens per backend. The required libpure fork hides its LLVM symbols, so the earlier instruction to disable PostgreSQL JIT is obsolete. Deploy matching extension and libpure files, and replace processes that loaded older libraries. Distribution 0.6.0 does not change SQL version 0.1, so that version alone cannot identify the installed library.

### Distribution 0.6.0

Version 0.6.0 adds parameter arrays to `sql::query` and `sql::exec`, binding Pure values through SPI instead of interpolating them into SQL text. It also supports `RETURNS SETOF` scalar or array results from a Pure list or a body-bounded lazy stream. A `LIMIT` outside the function does not bound an unbounded stream, so bound the stream in the Pure body. The trusted language remains non-relocatable and uses the DBA module allowlist.
