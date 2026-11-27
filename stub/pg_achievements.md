## Usage

Sources:

- [CHANGELOG.md](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.2.2/CHANGELOG.md)

- [README.md](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.2.2/README.md)
- [pg_achievements.control](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.2.2/pg_achievements.control)
- [sql/pg_achievements--1.2.2.sql](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.2.2/sql/pg_achievements--1.2.2.sql)
- [LICENSE](https://api.pgxn.org/src/pg_achievements/pg_achievements-1.2.2/LICENSE)

`pg_achievements` awards achievements from database activity and displays notices in the originating session. Release 1.2.2 installs SQL version 1.2.2. It requires PostgreSQL 14–19 on Linux, a superuser installer, and cluster-wide preload; the fixed `achievements` schema lives in one configured database.

### Enable and Inspect

Add the library to the existing preload list, choose the storage database, and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_achievements'
pg_achievements.database = 'postgres'
```

Create the extension in that database. Events are collected across databases in the cluster.

```sql
CREATE EXTENSION pg_achievements;
SELECT tier, title, points, awarded_at FROM achievements.my;
SELECT code, current, threshold, unlocked FROM achievements.progress;
SELECT * FROM achievements.leaderboard;
SELECT achievements.counter('stmt_select');
SELECT achievements.sync();
```

`achievements.my` and `achievements.progress` show the current role's awards and progress. `achievements.catalog_public` hides secret details until appropriate; the leaderboard exposes scores of roles that have not opted out. `achievements.sync(timeout_ms)` waits for the worker and returns false on timeout; its default is 5000 milliseconds.

### Tracking Controls

```sql
ALTER ROLE robot SET pg_achievements.opt_out = on;
SET pg_achievements.notify = off;
SET pg_achievements.notice_style = ascii;
```

The role must exist before running the role-setting example. Opting out stops tracking, while disabling notices only silences notifications. `pg_achievements.track_nested` is superuser-controlled and defaults off. The worker flush interval defaults to 1000 milliseconds; the database and maximum tracked-role capacity require a restart to change.

`achievements.reset()` clears the caller's progress; a superuser can reset other roles. `achievements.define()` creates custom awards from existing counters, and `achievements.grant()` awards one manually. These administrative functions are restricted. Dropping a role retains its data until a superuser calls `achievements.forget_role()`.

### Operational Boundaries

Tracking is independent of the user's transaction: rollback does not revoke an award. The worker persists counters asynchronously, so a crash can lose recent unflushed progress. The extension records counters and awards, not SQL text, parameters or object names. Base tables are not readable by PUBLIC.

Parse-time features of prepared statements count only when analyzed, and client meta-commands are not directly visible. After restoring progress into a running server, restart PostgreSQL to refresh shared memory. Preserve the other libraries already present in the preload setting.

### Version 1.2 Announcements and Upgrade

Version 1.2 sends new awards through NOTIFY on `pg_achievements.notify_channel`, whose default is `achievements`; an empty channel disables it. Opted-out roles are not announced. `achievements.rarity` reports how many players hold each award. Version 1.1 consolidated platinum into one collection trophy, added counters and awards, and repaired counter/award restoration races.

Install matching files, restart PostgreSQL, then run `ALTER EXTENSION pg_achievements UPDATE TO '1.2.2';`. The update preserves counters and awards. Custom platinum awards from 1.0 survive, while new custom platinum definitions are rejected.

### Version 1.2.2

Install the matching library, restart PostgreSQL and run the SQL upgrade above. PostgreSQL 14 and 15 are now supported. On 14, the unavailable `merge` counter is hidden from progress, catalog and rarity views; it is excluded from the collection trophy and cannot be used by `achievements.define()`. It becomes available after upgrading PostgreSQL to 15 or later.

The earlier 1.2.1 library fixes are included: every execution of an extended-protocol prepared statement counts when query IDs are enabled. Malformed statements, custom-award text, role renames and worker recovery no longer trigger the documented tracking failures. Attempt counters still count at parse analysis; `compute_query_id = off` limits prepared-statement tracking across transactions.
