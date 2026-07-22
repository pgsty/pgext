## Usage

Sources:

- [Official README](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/README.md)
- [Extension SQL](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord--0.0.1.sql)
- [Snapshot, worker, and hook implementation](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord.c)
- [Extension control file](https://github.com/rjuju/pg_timelord/blob/309155c3e6b5b5ee72117aa035236b791c47fe4d/pg_timelord.control)

`pg_timelord` is a proof-of-concept that substitutes a historical heap visibility snapshot based on transaction commit timestamps. It deliberately retains old tuple versions by holding back cleanup and is not intended for production, point-in-time recovery, or dependable temporal queries.

### Setup and Core Workflow

The module can only load through `shared_preload_libraries`. Enable commit-timestamp tracking, configure the one database used by its retention worker, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_timelord'
track_commit_timestamp = on
pg_timelord.database = 'postgres'
pg_timelord.keep_xact = 200000000
```

```sql
CREATE EXTENSION pg_timelord;

SELECT * FROM pg_timelord_oldest_xact;
SET pg_timelord.ts = '2016-10-31 23:59:59 GMT';

SELECT * FROM public.example;

RESET pg_timelord.ts;
```

Set or reset `pg_timelord.ts` outside an explicit transaction. The requested timestamp must not predate the oldest retained transaction with an available commit timestamp.

### Objects and Configuration

- `pg_timelord.ts` is a user-settable timestamp string that enables historical visibility in the current session.
- `pg_timelord.database` is a postmaster setting, default `postgres`, selecting the database to which the retention background worker connects.
- `pg_timelord.keep_xact` is a superuser setting, default 200,000,000, controlling how many transactions the worker attempts to keep from cleanup.
- `pg_timelord_oldest_xact()` returns the oldest retained transaction ID.
- The `pg_timelord_oldest_xact` view reports the configured database, retained `xmin`, and its commit timestamp.

### Severe Operational Risks

The worker holds an old transaction horizon so VACUUM cannot remove versions needed by the experiment. The default retention window is extremely large and can cause unbounded table and index bloat, transaction-ID wraparound pressure, disk exhaustion, degraded performance, and availability loss. Monitor age, storage, autovacuum, and wraparound continuously in any disposable test.

Historical visibility is limited to heap tuple versions that have not been removed and whose transaction commit timestamps remain available. It does not reconstruct DDL, sequences, external effects, non-heap state, or a transactionally consistent whole-database image. While active, the planner disables bitmap and index-only scans, and hooks attempt to enforce read-only behavior and reject most utility commands; these proof-of-concept checks are not a security boundary.

The implementation uses obsolete PostgreSQL internal snapshot APIs, and upstream still notes that its behavior needs verification. Version `0.0.1` should be treated as abandoned experimental code despite later copyright-year updates. Do not load it on a valuable cluster; reproduce only on a matching, isolated PostgreSQL build with complete backups and a plan to discard the instance.
