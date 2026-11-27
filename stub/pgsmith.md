## Usage

Sources:

- [README.md](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/README.md)
- [pgsmith.control](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/extension/pgsmith.control)
- [build.py](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/extension/build.py)
- [install.md](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/install.md)
- [sql-reference.md](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/sql-reference.md)
- [sync.md](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/sync.md)
- [operations.md](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/docs/operations.md)
- [SECURITY.md](https://github.com/SriPrarabdha/pgsmith/blob/75bddc282d216f2154dd30f20cd1b67a08ebe4a1/SECURITY.md)

`pgsmith` 0.1.0 mirrors Flagsmith configuration into PostgreSQL and evaluates feature flags in SQL. Evaluation uses local snapshots; a separate sync service or optional database-side integrations obtain new configurations. This is an early source release with SQL/PL/pgSQL code and no compiled PostgreSQL library.

### Enable and Synchronize

Upstream targets PostgreSQL 14–18 and UTF-8 databases. On a server with the extension files installed:

```sql
CREATE EXTENSION pgsmith;
GRANT flagsmith_reader, flagsmith_tracker TO app_user;
GRANT flagsmith_loader TO sync_user;
```

The control marks it trusted and nonrelocatable. It creates the `flagsmith` and `flagsmith_c` schemas; standard installation needs no preload or restart. The four helper roles are created only when the installer has permission to create roles, so verify their existence and grants.

The `pgsmith` sync command uses a database connection for a role with `flagsmith_loader`, the Flagsmith API address, and server-side environment keys supplied through private configuration:

```bash
pgsmith sync --once
```

The alternative `pg_tle` installer requires that extension and its administrative permission. RDS/Aurora require `pg_tle` in their preload configuration; that requirement belongs to the installation route. Optional in-database synchronization needs Supabase Vault, `http` or `pg_net`, and `pg_cron` for scheduling. These are not dependencies of local flag evaluation.

### Evaluate Flags

After at least one environment has synchronized:

```sql
SELECT * FROM flagsmith.flags('user-42', '{"plan":"pro","country":"IN"}'::jsonb);
SELECT flagsmith.is_enabled('new_checkout', 'user-42', '{"plan":"pro"}'::jsonb);
SELECT * FROM flagsmith.health();
```

`flagsmith.flags` returns `feature`, `enabled`, `value`, `variant`, `reason`, and `metadata`. `flagsmith.flag` retrieves one flag; `flagsmith.value` and `flagsmith.variant` return its value or experiment arm. `flagsmith.get_flags` evaluates a specified environment or snapshot through the interpreter. Use a `MATERIALIZED` CTE when one flag evaluation should be shared across a query.

### Permissions and Operations

- `flagsmith_reader` evaluates flags, `flagsmith_tracker` records exposures/events, `flagsmith_planner` adds review access, and `flagsmith_loader` changes served configuration. Nothing is granted to `PUBLIC` by default.
- Readers can evaluate arbitrary identities and traits. Do not directly grant that role to Supabase `anon` or `authenticated` for row-security decisions; expose a restricted wrapper bound to the authenticated caller.
- `flagsmith.pin` and `flagsmith.unpin` control the served snapshot. `flagsmith.health`, `flagsmith.sync_status`, and `flagsmith.event_status` report stale configuration and sync/event problems. A sync failure leaves the last accepted configuration serving.
- Backups omit derived snapshots and compiled functions. After restore, run `pgsmith sync --once` or `flagsmith.sync_tick` to rebuild them. Upstream 0.1.0 has no upgrade path yet; the documented reinstall requires resynchronization.
- Only the owner may add remote sync sources, because these determine where server keys are sent. Review `pg_net` queue permissions before choosing that transport; the queue can contain credentials.
