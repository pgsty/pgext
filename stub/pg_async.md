## Usage

Sources:

- [Official README](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/README.md)
- [Official extension control file](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async.control)
- [Official extension SQL](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async--1.0.sql)
- [Official hook implementation](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async.c)

`pg_async` 1.0 is a preload module that reimplements PostgreSQL asynchronous notification handling so `LISTEN`, `NOTIFY`, and `UNLISTEN` can operate during physical hot standby recovery. It delegates to PostgreSQL’s normal implementation outside recovery. Notifications remain local to one server instance; they are not replicated between primary and standby.

### Installation and Workflow

Install the extension files and create the extension on the primary so its catalog objects reach the standby. Add the library to every server where it will run, then restart.

```conf
shared_preload_libraries = 'pg_async'
```

```sql
LISTEN replica_events;

NOTIFY replica_events, 'refresh complete';

SELECT pg_listening_channels();
SELECT pg_notification_queue_usage();

UNLISTEN replica_events;
```

During recovery, the process-utility hook intercepts the SQL commands and uses the module’s private queue implementation. Transaction commit still controls when listen/unlisten changes and notifications become visible.

### Installed Functions

- `pg_listen(text)` and `pg_unlisten(text)` provide function forms of channel registration.
- `pg_unlisten_all()` removes all registrations for the current backend.
- `pg_notify(text, text)` queues a channel and payload.
- `pg_listening_channels()` lists channels for the current backend.
- `pg_notification_queue_usage()` reports the fraction of the local queue in use.

These names overlap PostgreSQL built-ins. Schema placement and `search_path` determine which function an unqualified call resolves to; the utility commands are clearer for normal use.

### Recovery and Version Boundaries

The pinned source embeds modified copies of PostgreSQL 13 and 14 `async.c` internals and refuses to load except through `shared_preload_libraries`. It has no implementation branch for PostgreSQL 15 or later. This is a server-internal patch packaged as an extension, so exact-major source compatibility is mandatory.

The queue is node-local and database-scoped. A notification sent on the primary does not arrive on a standby, and standby promotion/restart discards transient listener state. Applications must tolerate duplicates, missed events, disconnects, and promotion; use a durable table or message system when delivery matters. Test hook/signal interaction, queue exhaustion, long transactions, and failover on an isolated matching PostgreSQL build.
