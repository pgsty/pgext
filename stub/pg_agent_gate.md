## Usage

Sources:

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/sql/pg_agent_gate--0.2.15--0.3.0.sql)

`pg_agent_gate` 0.3.0 routes registered agent roles through proposals, checks, dry runs and guarded execution on PostgreSQL 18/19, with optional human approval before commit. The library must be loaded before the first agent statement, through shared preload or a role-level session preload followed by reconnection.

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

### Human Approval

Set `p_requires_approval` when calling `agent_gate.register_agent` to require consent for that agent's proposals. A member of `agent_gate_approver` records consent with `agent_gate.approve`; the agent then commits its own proposal. Approval executes no SQL and does not bypass row limits, assertions or commit-time checks. The agent cannot approve itself, and `propose_and_commit` cannot bypass the approval requirement.

The fixed `agent_gate_internal` schema stores proposals, executions and approvals; agents do not access those tables directly. The execution record identifies the approver in `approved_by`.

### Library and SQL Upgrades

The library retains the 0.2.15 protections against function-running argument casts and lone-semicolon crashes. Replace the library and restart or reconnect every process that preloads it, then update each database to 0.3.0 using the supplied SQL upgrade. The 0.3.0 migration adds approval objects and is not a no-op. Treat DDL permission as schema power that row-count limits do not bound.
