## Usage

Sources:

- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/README.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/quickstarts/postgres.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/reference/rls.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docs/postgresql/reference/jwt-claims.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/docker/postgresql/cloudsync.control.in)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/src/cloudsync.h)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/src/postgresql/sql_postgresql.c)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/src/postgresql/migrations/README.md)
- [Official source](https://github.com/sqliteai/sqlite-sync/blob/1.2.4/LICENSE)
- [Official source](https://github.com/sqliteai/sqlite-sync/releases/tag/1.2.4)
- [Official source](https://github.com/sqliteai/sqlite-sync/releases/download/1.2.4/cloudsync-postgresql18-macos-arm64-1.2.4.tar.gz)

`cloudsync` is the PostgreSQL side of SQLite AI's CRDT synchronization stack. Release 1.2.4 installs SQL extension version 1.2. The PostgreSQL integration remains beta and requires a separate CloudSync service and SQLite client integration for network synchronization.

### Initialize a Table

Install the official files matching the server platform and PostgreSQL major. Published PostgreSQL artifacts target 15, 17 and 18; Linux binaries require at least 15.14, 17.6 or 18.0 within those majors. Shared preload is not required. Use an administrative installer and table owner for setup:

```sql
CREATE EXTENSION cloudsync;
SELECT cloudsync_version();
CREATE TABLE sync_tasks (
  id text PRIMARY KEY,
  title text NOT NULL DEFAULT ''
);
SELECT cloudsync_init('sync_tasks');
INSERT INTO sync_tasks VALUES ('task-1', 'Review offline changes');
SELECT cloudsync_is_enabled('sync_tasks');
```

The control declares the extension trusted and relocatable, but creation installs C-language functions, and initialization can create a database-wide DDL event trigger. Verify the actual installer privileges and perform setup as an administrator rather than granting untrusted-language creation to application roles. The trigger uses `plpgsql` and tracks schema changes through `public.app_schema_version`.

Create synchronized tables explicitly on PostgreSQL. Use globally unique text/UUID identifiers for independently writable clients and compatible column types. Initialization adds triggers and CRDT shadow metadata; it does not launch a replication worker. The SQLite-side network functions and the separate service's authentication configuration are different components.

### Synchronization Objects and Access

`cloudsync_version()` reports the full library release. `cloudsync_siteid()` and `cloudsync_db_version()` identify synchronization state. `cloudsync_is_enabled`, `cloudsync_enable` and `cloudsync_disable` control table tracking. `cloudsync_payload_encode`, `cloudsync_payload_decode` and `cloudsync_payload_apply` exchange binary change payloads; their counters are not a substitute for checking resulting rows.

The service verifies client JWTs and establishes `request.jwt.claims` and a database role before applying RLS. Direct database clients able to set those claims are outside that authentication boundary. Grant only the required table, sequence, shadow-table and function privileges; reserve setup and cleanup functions for administrators. A denied RLS write raises SQLSTATE `42501` and rolls back that apply statement. Delete markers include primary keys and are delivered to all devices syncing the table, so sensitive primary keys require a separate data-design decision.

### Schema Changes and Recovery

Use `cloudsync_begin_alter` and `cloudsync_commit_alter` around the documented schema-change workflow, and coordinate compatible schemas across clients. Keep application data, shadow tables, site identity, synchronization versions and service state together in recovery planning. Disabling tracking does not undo already delivered changes. Cleanup and termination remove synchronization state and must be treated as deliberate administrative operations.

PostgreSQL text cannot hold NUL characters: the documented path truncates such SQLite text at the first NUL and logs a warning. Store binary values in BYTEA and verify cross-engine type conversions. CRDT convergence does not enforce application-level invariants or replace authorization.

### Upgrade

Keep clients, service and extension artifacts compatible. Patch releases replace the library while retaining the SQL major/minor version. Follow the upstream restart/reconnect procedure so sessions use the new binary. For a major/minor SQL change, install the supplied upgrade scripts and run `ALTER EXTENSION cloudsync UPDATE` per database. The 1.2.4 bundle includes paths from 1.0 through 1.1 to 1.2. Confirm the binary with `cloudsync_version()` and validate synchronized rows and permissions after an upgrade.
