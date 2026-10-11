## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_agent_gate/pg_agent_gate-0.2.14/README.md)
- [pg_agent_gate.control](https://api.pgxn.org/src/pg_agent_gate/pg_agent_gate-0.2.14/pg_agent_gate.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_agent_gate/pg_agent_gate-0.2.14/CHANGELOG.md)

`pg_agent_gate` 0.2.14 routes registered agent roles through proposals, checks, dry runs and guarded execution on PostgreSQL 18/19. The library must be loaded before the first agent statement, through shared preload or a role-level session preload followed by reconnection.

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

Install as a superuser and register ordinary roles with narrow table privileges, row limits and a DDL allow-list. The API includes `agent_gate.whoami`, `agent_gate.discover`, `agent_gate.propose`, `agent_gate.dry_run`, `agent_gate.commit` and `agent_gate.acts`. Verbs require READ COMMITTED. Proposal-time session settings are checked again before dry-run/commit, and reads that invoke known state-changing built-ins are rejected. Re-registering a name does not inherit old proposals. `agent_gate.max_proposal_bytes` and `agent_gate.max_intent_bytes` bound input sizes. Existing duplicate committed proposals can block the upgrade and require inspection. Optional bound checks use `pg_living_assertions`. Rollback cannot undo every sequence or external effect; existing privileged code must remain trusted. The server log can contain literal SQL. Attempt records participate in the outer transaction, and the default fast durability mode can lose recent records on a crash.
