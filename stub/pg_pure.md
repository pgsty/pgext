## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.1.1/CHANGELOG.md)

`pg_pure` is the installable trusted PL/Pure language entry. Distribution 0.1.1 ships SQL version 0.1. Its handler deliberately refuses to execute functions until libpure has a sandbox. Trusted interpretation is therefore unavailable in this release. The operational sibling is `pg_pureu`.

### Enablement and Current Boundary

```sql
CREATE EXTENSION pg_pure;
SELECT extversion FROM pg_extension WHERE extname = 'pg_pure';
```

The control marks the extension trusted, non-relocatable and installable without superuser-only extension creation. That catalog setting does not override the handler's execution refusal. Do not grant ordinary roles the untrusted sibling to work around it.

### Runtime Requirements

PostgreSQL 17–19 and the upstream embeddable libpure fork are required. Both siblings use the same `pg_pure` shared library. libpure permits native `extern` calls; the missing sandbox must eventually restrict native calls, inline C, evaluation and arbitrary module loading before trusted execution is safe.

Preloading `pg_pure` is optional and requires a server restart when configured through `shared_preload_libraries`. Its `pg_pure.using` setting selects modules, defaulting to `math`. The ordinary interpreter is created per backend without preload.

### Deployment Limits

libpure and PostgreSQL's LLVM JIT must not load different LLVM major versions into the same backend; upstream reports a backend-exit crash, which can trigger a cluster-wide restart. The supported 0.1 deployment uses PostgreSQL without LLVM JIT; a matching LLVM major is another documented configuration. This distribution does not establish production readiness or a trusted sandbox.
