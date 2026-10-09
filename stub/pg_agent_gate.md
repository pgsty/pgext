## Usage

Sources:

- [README.md](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/README.md)
- [Cargo.toml](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/Cargo.toml)
- [src/lib.rs](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/src/lib.rs)
- [sql/pg_agent_gate--0.2.0--0.2.1.sql](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/sql/pg_agent_gate--0.2.0--0.2.1.sql)
- [pg_agent_gate.control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/pg_agent_gate.control)

`pg_agent_gate` routes a registered agent role through proposal, dry-run and guarded execution functions on PostgreSQL 18+. The library must load before that role executes its first statement.

### Core Workflow

```sql
CREATE EXTENSION pg_agent_gate;
```

```sql
SELECT agent_gate.whoami();
SELECT agent_gate.discover('customers');
SELECT agent_gate.propose('SELECT 1', 'Check connectivity');
```

### Operational Boundaries

As an administrator, merge the library into `shared_preload_libraries` and restart, or let `agent_gate.register_agent` set `session_preload_libraries` on the role; reconnect afterward. Install as a superuser. Register an ordinary role, grant only its required table privileges, and set its row limit and DDL policy. `agent_gate.discover`, `agent_gate.propose`, `agent_gate.dry_run`, `agent_gate.commit`, `agent_gate.acts` and `agent_gate.whoami` are the agent surface. Optional bound assertions use `pg_living_assertions`. A dry run rolls back database changes but cannot undo sequence advances, session locks or external effects. Existing privileged functions, view ownership and fast-path function calls remain security boundaries; this is not a complete sandbox. Default `agent_gate.attempt_durability = fast` can lose attempt records on a crash, and an outer transaction rollback also rolls back the record. Version 0.2.1 closes writing-CTE row-limit and direct parameter-change bypasses.
