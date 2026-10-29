## Usage

Sources:

- [v2.0.20 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/README.md)
- [Extension control file](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/fractalsql.control)
- [SQL definitions](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/sql/fractalsql--1.0.sql)
- [v2.0.20 release](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.20)

`fractalsql` is the base PostgreSQL extension in the upstream v2.0.20 release. Its SQL version is 1.0 and its C functions provide vector search, reasoning, embedding and agent-oriented primitives for PostgreSQL 14–18.

### Install and inspect

```sql
CREATE EXTENSION fractalsql;
SELECT fractal_search(ARRAY[0.1, 0.2, 0.3]::float8[]);
```

The exact function surface is defined by the versioned SQL file. The optional `fractalsql_agents` extension depends on this base extension and installs parameterized PL/pgSQL agent engines.

### Operational boundaries

Reasoning and embedding functions can call configured external model providers. Treat provider credentials, model inputs, generated SQL and model outputs as untrusted application data; grant execution narrowly and review network egress. Keep generated plans and decisions auditable, and test provider failures, timeouts, retries and transaction rollback. The upstream release ships platform packages for PostgreSQL 14–18, but those upstream artifacts do not make the extension part of the local packaged baseline.
