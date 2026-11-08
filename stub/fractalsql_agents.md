## Usage

Sources:

- [v2.0.21 README](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/README.md)
- [Extension control file](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/fractalsql_agents/fractalsql_agents.control)
- [SQL definitions](https://github.com/FractalSQLabs/fractalsql-postgresql/blob/v2.0.21/fractalsql_agents/sql/fractalsql_agents--1.0.sql)
- [v2.0.21 release](https://github.com/FractalSQLabs/fractalsql-postgresql/releases/tag/v2.0.21)

`fractalsql_agents` is the optional dependent extension shipped with FractalSQL v2.0.21. It installs sixteen parameterized PL/pgSQL agent engines over the `fractalsql` search, reasoning and analytics primitives for PostgreSQL 14–18.

### Install

```sql
CREATE EXTENSION fractalsql;
CREATE EXTENSION fractalsql_agents;
```

The control file requires `fractalsql`; install and upgrade the base extension first. The SQL file defines the agent engine functions and their argument contracts. Read those definitions before granting execution to application roles.

### Safety and provider boundaries

Several engines call reasoning or embedding functions, while others are retrieval or analytics-only compositions. Provider credentials and network egress are deployment concerns. Generated SQL and model output are untrusted; use least-privilege roles, explicit data-access grants, statement and provider timeouts, and audit logs. Test partial reasoning failures and transaction rollback before enabling an engine for writes.

### Explicit execution in 2.0.21

`fractal_agent_data_analyst` gains a fifth argument, `auto_execute`, defaulting to false. The default generates and validates SQL without executing it; result data and analysis remain NULL, and the subsequent reasoning call is skipped. Passing true opts into execution and reasoning. The old four-argument definition is removed by the new installation SQL to avoid resolving calls to the previous always-executing body.

The SQL extension version remains 1.0. Existing installations therefore need a reviewed SQL-object migration as well as the matching library; an ordinary version update cannot be assumed to replace the old definition. Inspect dependent objects and grants before migration.
