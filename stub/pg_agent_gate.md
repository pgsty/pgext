## Usage

Sources:

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/sql/pg_agent_gate--0.4.1--0.4.3.sql)

`pg_agent_gate` 0.4.3 routes registered agent roles through proposals, checks, dry runs and guarded execution on PostgreSQL 18/19, with optional human approval before commit. The library must be loaded before the first agent statement, through shared preload or a role-level session preload followed by reconnection.

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

### Write Budgets and Role Membership

Version 0.4.0 introduced a default cumulative write budget: the agent's per-statement row limit per hour. `agent_gate.set_budget` changes the allowance/window; `agent_gate.reset_budget` starts a new allowance without rewriting the record. Kept writes count actual row effects, including cascades and triggers; concurrent commits serialize the budget check.

Agent roles must have no members. Registration refuses such roles, and upgrades warn about existing memberships that an administrator must revoke. Connect each agent as its own login; a broker must not switch into the agent role.

### Function and Session Checks in 0.4.3

The 0.4.3 library judges hidden function calls when they execute, including type support functions, aggregates, defaults and constraints. It rejects SQL passed as text to known execution functions, agent-alterable functions, protected-context changes and unaccounted state changes. DDL row effects also count toward row limits. These checks deliberately reject some statements that ordinary PostgreSQL roles can execute.

Protected context includes role identity, row-security settings, search path and unregistered custom application parameters. A changed context causes refusal rather than automatic reset. Native code attached to agent-owned objects is restricted. Optional `pg_living_assertions` 0.5.5 can leak context and cause refusals; 0.5.6 fixes that issue. Keep agent-controlled schemas out of other users' search paths, since DDL can create code that later runs under their rights.

### Library and SQL Upgrades

After replacing the library, restart or reconnect all processes that preload it, then update each database:

```sql
ALTER EXTENSION pg_agent_gate UPDATE TO '0.4.3';
```

The 0.4.1 to 0.4.3 migration changes no schema objects. Older upgrade paths still add budget and approval objects; inspect existing agents' budgets after upgrading. Version 0.4.2 was not released separately. Superusers, the operating system and trusted DBA code remain outside this boundary; an authorized statement can still violate human intent.
