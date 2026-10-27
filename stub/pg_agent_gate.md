## Usage

Sources:

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/sql/pg_agent_gate--0.2.14--0.2.15.sql)

`pg_agent_gate` 0.2.15 routes registered agent roles through proposals, checks, dry runs and guarded execution on PostgreSQL 18/19. The library must be loaded before the first agent statement, through shared preload or a role-level session preload followed by reconnection.

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

### 0.2.15 Security Corrections

0.2.15 rejects explicit and I/O casts around gate arguments that could run an agent-selected function before verification. It also refuses a lone semicolon instead of crashing the backend. These fixes live in the library; the SQL update is a no-op. Replace the library and restart or reconnect every process that preloads it, then update each database to 0.2.15. Treat DDL permission as schema power that row-count limits do not bound.
