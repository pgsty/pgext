## Usage

Sources:

- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/README.md)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate/percolate.control)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate/sql/percolate--0.2.30.sql)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/dev/image-init/01-percolate.sh)

`percolate` 0.2.30 is a pure SQL workflow and agent-memory system for PostgreSQL 19. Workflows, runs, identity, documents and graph data live in database schemas. The complete service also uses separate workers and HTTP services; creating the extension does not start those processes.

### Bootstrap and Ownership

Follow the official bootstrap to create the application roles and database privileges first. It makes database-wide privilege changes, including access to the built-in configuration setter, so review it on a dedicated database. The prerequisites are `pgcrypto`, `pg_trgm`, `vector` and `percolate_parser`; a superuser must preinstall the untrusted compiled prerequisites.

After the bootstrap has prepared the roles and grants:

```sql
SET ROLE app_owner;
CREATE EXTENSION percolate CASCADE;
RESET ROLE;
SELECT extname, extversion FROM pg_extension WHERE extname = 'percolate';
```

The installer explicitly rejects a superuser caller: application-owned objects and row-level security are part of its security contract. It also checks required roles, membership and privilege restrictions. It is not a generic schema package that can be installed with arbitrary role names.

### Workflows and Maintenance

The runtime executes stored workflow tasks through workers; schedules are rows in `workflow.schedules`. `workflow.tick()` advances schedules, `workflow.reap_stale_tasks()` recovers abandoned work and `workflow.promote_due_timers()` advances timer tasks. The standard deployment runs these through optional `pg_cron`, which requires preload and the correct database configuration. Without it, scheduling, the reaper and timers do not run automatically.

The extension is not relocatable. Use the matching parser and upgrade scripts, preserve application roles and grants in backups, and test restore with actual application data. SQL installation, worker deployment and scheduled execution are separate steps.
