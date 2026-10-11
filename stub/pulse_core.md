## Usage

Sources:

- [README.md](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/README.md)
- [extensions/pulse_core/pulse_core--0.1.0.sql](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/pulse_core--0.1.0.sql)
- [extensions/pulse_core/src/pulse_core.c](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/src/pulse_core.c)
- [extensions/pulse_core/src/sqlapi.c](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/src/sqlapi.c)
- [extensions/pulse_core/test/pulse_core.conf](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/test/pulse_core.conf)
- [extensions/pulse_core/pulse_core.control](https://github.com/AndrewJesse/pg_pulse/blob/ddd8034881bd2b466b489883e103e597aa1375f7/extensions/pulse_core/pulse_core.control)

`pulse_core` is the implemented query-statistics component of the pg_pulse prototype. Executor hooks collect per-query calls, timing, row counts and last-seen timestamps in shared memory.

### Core Workflow

```ini
shared_preload_libraries = 'pulse_core'
compute_query_id = on
```

```sql
CREATE EXTENSION pulse_core;
SELECT * FROM pulse.queries() ORDER BY total_ms DESC;
```

### Operational Boundaries

This project targets PostgreSQL 18. Add `pulse_core` to `shared_preload_libraries` and restart before administrator installation; late loading is rejected. Enable `compute_query_id` as in the upstream regression configuration. `pulse.max_entries` is a postmaster setting with default 1000; `pulse.track` accepts none, top or all and defaults to top. `pulse.queries()` reports the current statistics, while `pulse.reset()` clears them. Both functions are revoked from PUBLIC and granted to `pg_monitor`; review membership because that role can reset statistics too. Shared-memory counters do not establish durable history. The separate health-worker component is still a scaffold and is not represented by this entry.
