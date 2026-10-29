## Usage

Sources:

- [v2.0.20 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/README.md)
- [Extension control file](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/fractalsql_agents/fractalsql_agents.control)
- [SQL definitions](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.20/fractalsql_agents/sql/fractalsql_agents--1.0.sql)
- [v2.0.20 release](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.20)

`fractalsql_agents` is the optional dependent extension shipped with FractalSQL v2.0.20. It installs sixteen parameterized PL/pgSQL agent engines over the `fractalsql` search, reasoning and analytics primitives for PostgreSQL 14–18.

### Install

```sql
CREATE EXTENSION fractalsql;
CREATE EXTENSION fractalsql_agents;
```

The control file requires `fractalsql`; install and upgrade the base extension first. The SQL file defines the agent engine functions and their argument contracts. Read those definitions before granting execution to application roles.

### Safety and provider boundaries

Several engines call reasoning or embedding functions, while others are retrieval or analytics-only compositions. Provider credentials and network egress are deployment concerns. Generated SQL and model output are untrusted; use least-privilege roles, explicit data-access grants, statement and provider timeouts, and audit logs. Test partial reasoning failures and transaction rollback before enabling an engine for writes.
