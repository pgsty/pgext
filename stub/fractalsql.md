## Usage

Sources:

- [Current upstream](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.22/sql/fractalsql--1.0.sql)
- [Current upstream](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.22/docs/api-analytics.md)
- [Current upstream](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.22/docs/mcp-plugin.md)

- [v2.0.21 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/README.md)
- [Extension control file](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/fractalsql.control)
- [SQL definitions](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/sql/fractalsql--1.0.sql)
- [v2.0.21 release](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.21)

`fractalsql` is the base PostgreSQL extension in the upstream v2.0.21 release. Its SQL version is 1.0 and its C functions provide vector search, reasoning, embedding and agent-oriented primitives for PostgreSQL 14–18.

### Install and inspect

```sql
CREATE EXTENSION fractalsql;
SELECT fractal_search(ARRAY[0.1, 0.2, 0.3]::float8[]);
```

The exact function surface is defined by the versioned SQL file. The optional `fractalsql_agents` extension depends on this base extension and installs parameterized PL/pgSQL agent engines.

### Operational boundaries

Reasoning and embedding functions can call configured external model providers. Treat provider credentials, model inputs, generated SQL and model outputs as untrusted application data; grant execution narrowly and review network egress. Keep generated plans and decisions auditable, and test provider failures, timeouts, retries and transaction rollback. The upstream release ships platform packages for PostgreSQL 14–18, but those upstream artifacts do not make the extension part of the local packaged baseline.

### Release 2.0.21 changes

The feature store records writer provenance and rejects overwriting another role's vector unless the caller belongs to that role or is a superuser. `fractal_store_morphology` accepts an optional source label. Public execution of `fractal_audit_log` is revoked in the installation SQL; existing databases need an explicit privilege review. `fractal_lock_destructive()` irreversibly disables ledger mutation and audit writes for the current connection. Completed vectorizer jobs are pruned after one hour and failed jobs after one day when the caller can delete queue rows.

The extension SQL version remains 1.0 while installation definitions changed. Replacing the library alone does not update existing SQL objects or privileges; review and back up the installed definitions before applying the upstream migration guidance.

### Upstream 2.0.22 API boundary

The upstream distribution is 2.0.22 while SQL extension version remains 1.0. `fractal_tda_persistence_diagram` now requires explicit `max_dim` and `max_thresh`; use 0 or 1 for the former and a positive threshold for the latter. The reported graph cycle rank is not a full simplicial homology computation.

The installation SQL renames the configuration wrappers to `fractal_config_set` and `fractal_config_get`; MCP names are aligned with the documented `fractal_version` and `fractal_configure_reasoning` interface. Existing databases do not acquire changed SQL definitions just from replacing a library under unchanged version 1.0. Review the official installation SQL and dependency-aware migration procedure before changing a live installation; do not assume ALTER EXTENSION will apply a missing versioned upgrade. The companion agents installation SQL is unchanged in this release.
