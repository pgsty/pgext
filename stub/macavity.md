## Usage

Sources:

- [README.md](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/README.md)
- [CHANGELOG.md](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/CHANGELOG.md)
- [macavity.control](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/macavity.control)
- [sql/macavity--0.2.0.sql](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/sql/macavity--0.2.0.sql)
- [sql/macavity--0.1.0--0.2.0.sql](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/sql/macavity--0.1.0--0.2.0.sql)

`macavity` 0.2.1 provides deterministic fault injection for disposable PostgreSQL test clusters. Upstream supports PostgreSQL 16–19 on Linux; other platforms are not validated, and Windows crash injection is unsupported. The crash action can disconnect every session and force crash recovery.

### Event workflow

```sql
CREATE EXTENSION macavity;
SELECT * FROM macavity_points();
SELECT macavity_arm('executor_start', 'error', 1);
SELECT 1; -- expected injected error
SELECT * FROM macavity_status();
SELECT macavity_disarm();
SELECT macavity_reset();
```

### Registry and upgrade

Each session keeps multiple events with stable IDs, separate counters and states `armed`, `completed` or `disarmed`. `macavity_arm` returns an integer event ID; its single-ID overload re-arms a completed or disarmed event. `macavity_status` returns zero or more event rows. `macavity_disarm` retains event history; `macavity_reset` clears it and restarts IDs.

Points are `executor_start`, `executor_end`, `before_commit` and `before_abort`; actions are `error`, `delay` and `crash`. Delay lasts one second. Abort-time error injection is rejected. At a shared hit, delay precedes crash, then error, with IDs breaking ties.

Creation requires a superuser; no shared preload is needed. Only point enumeration is granted to `PUBLIC`. Upgrading from 0.1.0 with `ALTER EXTENSION macavity UPDATE` recreates functions with changed signatures: restore explicit grants, update callers and reconnect old sessions. Current Pigsty package fields still refer to 0.1.0.

### Version 0.2.1

The 0.2.0 to 0.2.1 SQL upgrade is a no-op; this release adds PostgreSQL 19 compatibility without changing the SQL API. Install the matching library and SQL files, run `ALTER EXTENSION macavity UPDATE TO '0.2.1';`, and reconnect existing sessions before testing.
