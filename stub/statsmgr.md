## Usage

Sources:

- [Official README](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/README.statsmgr.md)
- [Extension SQL](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/statsmgr--0.1-alpha.sql)
- [C implementation](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/statsmgr.c)

`statsmgr` 0.1-alpha is a preview extension for taking periodic snapshots of PostgreSQL cumulative archiver, background-writer, checkpointer, I/O, SLRU, and WAL statistics into dynamic shared memory. It targets PostgreSQL 17 or later and keeps only an in-memory rolling history, not a durable monitoring store.

### Setup Modes

For a temporary evaluation, create the extension and start its main dynamic background worker manually:

```sql
CREATE EXTENSION statsmgr;
SELECT statsmgr_start_main_worker();
SELECT statsmgr_snapshot();
SELECT * FROM statsmgr_wal();
```

For automatic startup, preload the module and restart PostgreSQL:

```conf
shared_preload_libraries = 'statsmgr'
```

Each result row adds `tick` and `tick_tz` to the corresponding core statistics fields.

### Function Index

- Management: `statsmgr_start_main_worker()`, `statsmgr_stop_main_worker()`, `statsmgr_snapshot()`, and `statsmgr_reset()`.
- History readers: `statsmgr_archiver()`, `statsmgr_bgwriter()`, `statsmgr_checkpointer()`, `statsmgr_io()`, `statsmgr_slru()`, and `statsmgr_wal()`.

The current implementation uses a 512-slot ring and a fixed 60-second periodic interval. Source variables for interval and per-kind enablement are not registered as PostgreSQL GUCs, and the README's configuration section is commented out, so those values are not runtime settings in this version.

### Security and Operational Notes

The SQL install script does not revoke PUBLIC execution on worker start/stop, reset, or manual snapshot functions, and the inspected C wrappers perform no explicit superuser check. Revoke management functions from PUBLIC immediately and grant them only to an administrative role; grant reader functions separately to a monitoring role after reviewing data exposure. State is cluster-wide shared memory and is lost at restart, while extension objects are database-local, so coordinate one control database. Background workers consume worker slots. This alpha implementation relies on PostgreSQL 17 internals; validate exact minor-version builds, concurrent management calls, ring wraparound, restart behavior, and resource limits before evaluation outside a test cluster.
