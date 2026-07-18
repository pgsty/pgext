## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/README.md)
- [Background-worker implementation](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/src/worker.rs)
- [SQL-callable functions](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/src/functions.rs)
- [Extension control file](https://github.com/willibrandon/pg_walrus/blob/99e3e183893002090e996a424b6ad7bcd43388a1/pg_walrus.control)

`pg_walrus` monitors requested checkpoints and automatically increases or decreases `max_wal_size`. It requires `shared_preload_libraries` and a restart; create the extension in the database selected by `walrus.database` to store adjustment history.

```sql
CREATE EXTENSION pg_walrus;
ALTER SYSTEM SET walrus.dry_run = true;
SELECT pg_reload_conf();
SELECT jsonb_pretty(walrus.status());
SELECT walrus.recommendation();
```

Start with `walrus.dry_run = true`. The worker normally executes `ALTER SYSTEM`, writes `postgresql.auto.conf`, and sends SIGHUP when its threshold, cap, shrink, cooldown, and hourly-rate rules select a change. `walrus.history` records decisions; `walrus.analyze(apply := true)` can apply one immediately and requires superuser.

Changing WAL sizing affects checkpoint cadence, disk usage, crash recovery, replication retention, and backup operations. Validate the algorithm under peak WAL rates, set a safe `walrus.max` and `walrus.min_size`, monitor free space, and coordinate with existing configuration management so automated rewrites do not fight declarative settings.
