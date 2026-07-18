## Usage

Sources:

- [Upstream README](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/README.md)
- [Extension control file](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/pg_typescript.control)
- [Rust package manifest](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/Cargo.toml)
- [Permission enforcement](https://github.com/isaacd9/pg_typescript/blob/e03ed5460decab37526e129f5577907b972b4898/src/permissions.rs)

`pg_typescript` version `0.1.0` is an alpha procedural language that runs TypeScript with Deno/V8 in PostgreSQL 16–18. It supports ordinary functions, remote or npm imports, and database access through `_pg.execute`. A runtime is created per backend; adding the library to `shared_preload_libraries` is optional prewarming rather than a requirement.

### Create a dependency-free function

```sql
LOAD 'pg_typescript';
CREATE EXTENSION pg_typescript;

CREATE FUNCTION slugify(title text)
RETURNS text
LANGUAGE typescript
IMMUTABLE STRICT
AS $$
  return title.trim().toLowerCase().replace(/\s+/g, '-');
$$;

SELECT slugify('Hello TypeScript');
```

The control file marks the language trusted, but extension installation is superuser-only and function owners can request capabilities through user-set `typescript.allow_*` settings. The superuser `typescript.max_*` ceilings are the real upper boundary; upstream defaults permit imports and `_pg.execute`, while several other ceilings are effectively uncapped. Set restrictive maximum import, network, file, environment, process, FFI, system, and database-execution policies before granting language use. Remote modules are fetched when functions are created and cached locally, so pin and audit every dependency. Treat untrusted TypeScript as database code, and test memory, timeout, backend-crash, cold-start, and supply-chain behavior before production use.
