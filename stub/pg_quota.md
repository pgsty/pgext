## Usage

Sources:

- [Official README](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/README.md)
- [Version 1.0 extension SQL](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/pg_quota--1.0.sql)
- [Background worker and GUC implementation](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/pg_quota.c)

`pg_quota` tracks relation-file ownership and applies soft disk-space quotas per database role. It is suitable for detecting and blocking some new writes after a role has exceeded its configured quota; it is not a hard storage boundary.

### Preload and Activation

Add the library and target databases to `postgresql.conf`, then restart PostgreSQL. The database list is a postmaster-level setting, and one background worker is registered for each listed database.

```conf
shared_preload_libraries = 'pg_quota'
pg_quota.databases = 'appdb'
pg_quota.refresh_naptime = '10 s'
```

After restart, create the extension in each configured database and assign quotas in bytes:

```sql
CREATE EXTENSION pg_quota;

INSERT INTO quota.config(roleid, quota)
VALUES ('app_writer'::regrole, pg_size_bytes('10 GB'))
ON CONFLICT (roleid)
DO UPDATE SET quota = EXCLUDED.quota;

SELECT rolname,
       pg_size_pretty(space_used) AS used,
       pg_size_pretty(quota) AS quota
FROM quota.status;
```

`quota.config` is the persistent configuration table. `quota.status` reports every tracked role's current relation space and configured quota; a NULL quota means the role is tracked but unrestricted.

### GUCs

- `pg_quota.databases` selects databases and requires restart to change.
- `pg_quota.refresh_naptime` controls the delay between full data-directory scans and can be reloaded.
- `pg_quota.restart_interval` controls worker restart delay and is a postmaster-level setting.

Ensure `max_worker_processes` leaves capacity for one worker per configured database plus other extensions and parallel work.

### Soft-Quota Limitations

Space is attributed to each relation's direct owner; group quotas are not supported. Temporary files, temporary relations, and catalog objects are ignored. Ownership changes and newly created relations are reflected only after commit and a later scan, so status is delayed and scanning can be expensive with many relations or partitions.

Enforcement checks occur only at the beginning of `INSERT` and `COPY`. A statement that starts below quota may exceed it, and `UPDATE`, `CREATE INDEX`, and other utility commands are not blocked by the quota hook. Test failure behavior and recovery procedures carefully, and do not use `pg_quota` as a security or billing-grade hard limit.
