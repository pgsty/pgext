## Usage

Sources:

- [Official upstream README](https://github.com/ryrobes/rvbbit-sql/blob/6c82cb49a85937ca1ebc0361d703101c1300ae52/README.md)
- [Official extension control file (pg_rvbbit.control)](https://github.com/ryrobes/rvbbit-sql/blob/6c82cb49a85937ca1ebc0361d703101c1300ae52/crates/pg_rvbbit/pg_rvbbit.control)
- [Official implementation source](https://github.com/ryrobes/rvbbit-sql/blob/6c82cb49a85937ca1ebc0361d703101c1300ae52/crates/pg_rvbbit/src/lib.rs)

`pg_rvbbit` — Semantic SQL operators backed by configurable LLMs, with caching, receipts, and routing across multiple execution engines. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_rvbbit;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `backend_probe` is an extension function.
- `backend_probe_with_input` is an extension function.
- `env_present` is an extension function.
- `reload_backends()` is an extension function.
- `rvbbit_build_info()` is an extension function.
- `rvbbit_version()` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `4.1.4`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
