## Usage

Sources:

- [pgcel.control](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/pgcel.control)
- [README.md](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/README.md)
- [Cargo.toml](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/Cargo.toml)
- [src/lib.rs](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/src/lib.rs)
- [tests/session_contracts.sh](https://github.com/krashanoff/postgres-cel/blob/908363804df284f82a086526ab7e9fc98f98029f/tests/session_contracts.sh)

`pgcel` stores CEL expression source in the `celprogram` type and evaluates it against a JSONB context. Use it for reusable predicates or calculations inside SQL. The implementation registers shared memory during initialization, so preload the library and restart before use.

### Core Workflow

```ini
shared_preload_libraries = 'pgcel'
```

```sql
CREATE EXTENSION pgcel;
SELECT cel_eval(cel_compile('ctx.age >= 18'), '{"age": 21}'::jsonb);
SELECT cel_eval_json(cel_compile('ctx.a + ctx.b'), '{"a": 1, "b": 2}'::jsonb);
SET pgcel.program_cache_size = 2048;
```

### Operational Boundaries

`cel_compile` validates source, `cel_eval` requires a boolean result, and `cel_eval_json` returns other results as JSONB. The input object is available as `ctx`; invalid expressions, missing values or incompatible result types can raise errors. PostgreSQL build features cover 14–17. The control file is non-relocatable and sets superuser=false; normal database/object privileges still apply. `pgcel.program_cache_size` defaults to 1024 compiled programs per backend and accepts 1–65536. Shared memory holds only validation hashes, not executable programs; every backend compiles its own executable. Size connection pools and caches accordingly. This early release is not a replacement for PostgreSQL role privileges or RLS.
