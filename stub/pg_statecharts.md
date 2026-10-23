## Usage

Sources:

- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/README.md)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/pg_statecharts.control)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/sql/function/start_machine_with_latest_statechart.sql)
- [Official source](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/sql/table/state_machine_event.sql)

`pg_statecharts` 0.1.0 runs hierarchical state machines and their events transactionally inside PostgreSQL. This release is pure SQL, tested on PostgreSQL 16–18, and needs no preload or restart.

### Core Workflow

Install `ltree` in `public`, then install the extension with a role that has the required object-creation privileges:

```sql
CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;
CREATE EXTENSION pg_statecharts;
```

Deploy a statechart definition through SQL migrations or the development companion before starting a machine. For an already deployed chart named order, this starts its latest version and returns the machine row:

```sql
SELECT * FROM fsm.start_machine_with_latest_statechart(1, 'order', '{}');
```

Events are inserted into `fsm.state_machine_event` with the shard, machine identifier, event name and JSON data. Its triggers process the event in the inserting transaction, including entry and exit callbacks. Callers must have the appropriate table and function privileges; chart callbacks are application code that needs its own review.

### Catalog and Versions

Definitions and running machines live in `fsm`. `fsm.to_semver` parses numeric chart versions and `fsm.semver_text` renders them; stored versions are integer arrays so their native ordering remains numeric. Do not sort a textual rendering when version precedence matters.

### Backups and Legacy Migration

The reviewed 0.1.0 source registers state tables and sequences for database dumps. The event inbox is intentionally omitted to avoid replaying callbacks during restoration. Backups made with 0.0.0 omitted machine state; take a fresh, restore-tested backup after upgrading.

An earlier 0.1.0 installation may predate dump registration, so check the installed catalog instead of trusting its version string. Do not drop a populated extension merely to refresh registration. Legacy Rust or Sqitch migrations require superuser work and the old semver dependency during conversion; follow the migration guide and rehearse on a restored copy. The old 0.0.0 path is no longer independently exercised upstream.
