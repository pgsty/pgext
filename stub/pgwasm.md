## Usage

Sources:

- [pgwasm README at the documented revision](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/README.md)
- [pgwasm architecture and SQL lifecycle](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/architecture.md)
- [pgwasm GUC reference](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/guc.md)
- [pgwasm WIT type mapping](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/wit-mapping.md)
- [pgwasm control file](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/pgwasm/pgwasm.control)

`pgwasm` loads WebAssembly components into PostgreSQL and registers WIT exports as typed PostgreSQL functions. Compiled artifacts are stored under the cluster data directory and reused by backend-local instance pools. This document follows pinned revision `535b53363f8208af139e757e508e66c46309ee29`; the source declares version 0.1.0 but does not provide a tagged 0.1.0 release.

### Core Workflow

Create the extension as a superuser. The following file-based workflow must be enabled and confined by an administrator before a loader role can use it:

```sql
CREATE EXTENSION pgwasm;

ALTER SYSTEM SET pgwasm.allow_load_from_file = on;
ALTER SYSTEM SET pgwasm.module_path = '/srv/pgwasm';
ALTER SYSTEM SET pgwasm.allowed_path_prefixes = '/srv/pgwasm';
SELECT pg_reload_conf();

GRANT pgwasm_loader TO app_runtime;

SELECT pgwasm.pgwasm_load(
    'arith',
    '{"path":"arith.component.wasm"}'::json,
    '{}'::json
);

SELECT * FROM pgwasm.pgwasm_functions();
SELECT * FROM pgwasm.pgwasm_modules();

SELECT pgwasm.pgwasm_unload('arith');
```

`pgwasm_load(module_name text, bytes_or_path json, options json)` accepts exactly one `bytes` or `path` source. File loading is off by default. A module name becomes the durable catalog key and the prefix for sanitized generated SQL function names.

### Lifecycle and Type Mapping

- `pgwasm_load` validates, resolves policy, creates required PostgreSQL types and functions, compiles an AOT artifact, and records the module.
- `pgwasm_reload` replaces module bytes while preserving stable identities when signatures remain compatible.
- `pgwasm_reconfigure` narrows or changes policy and resource limits.
- `pgwasm_unload` removes generated functions, types, catalog rows, and artifacts; dependencies block removal unless cascade is explicitly selected.
- WIT records map to composite types, enums to PostgreSQL enums, lists to arrays or `bytea`, and supported variants, flags, options, results, and resources to documented PostgreSQL representations.
- `pgwasm_modules()`, `pgwasm_functions()`, `pgwasm_wit_types()`, `pgwasm_policy_effective()`, and `pgwasm_stats()` provide inspection.

Review the generated function signatures before granting execute privileges or calling a newly loaded component. Reloads with breaking WIT changes require an explicit policy decision and dependency review.

### Sandbox and Privileges

The extension creates `pgwasm_loader` for lifecycle mutation and `pgwasm_reader` for observability. Loading, reloading, reconfiguring, and unloading require superuser or loader-role membership.

WASI filesystem, environment, sockets, HTTP, and SPI host-query access are all disabled by default. An administrator sets the cluster ceiling through `pgwasm.*` GUCs; per-module options can narrow that ceiling but cannot broaden it. Keep `pgwasm.allowed_hosts`, path prefixes, and filesystem preopens explicit and minimal.

### Resource and Operational Boundaries

- The default module-size limit is 32 MiB, invocation memory is 1,024 WebAssembly pages, and the wall-clock deadline is 5 seconds. Fuel metering is available but off by default.
- Artifacts under `$PGDATA/pgwasm/<module_id>/` are derived from module bytes and the Wasmtime build. Recompile them after incompatible Wasmtime or PostgreSQL upgrades rather than copying them as authoritative data.
- Shared counters depend on postmaster-time shared-memory allocation. Preload `pgwasm` when shared metrics are required; otherwise observability can fall back to non-shared counters and reports that state.
- The source exposes build features for PostgreSQL 13 through 18 and defaults to PostgreSQL 17, but the pinned revision has no published support matrix. Validate the exact PostgreSQL-major build and all required WIT mappings before deployment.
- Treat guest code as privileged database-adjacent code even with sandboxing: limit who can load modules, bound every capability and resource, and test traps, cancellation, reload, restart, and rollback behavior.
