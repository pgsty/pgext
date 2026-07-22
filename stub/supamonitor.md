## Usage

Sources:

- [Official 0.0.6 Rust package manifest](https://github.com/supabase/supamonitor/blob/497f1ecc66e297f04b32954efdfabf0ac21b52ff/Cargo.toml)
- [Official 0.0.6 hook implementation](https://github.com/supabase/supamonitor/blob/497f1ecc66e297f04b32954efdfabf0ac21b52ff/src/lib.rs)
- [Official current development manifest](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/Cargo.toml)

`supamonitor` 0.0.6 is a Supabase pgrx preload module that logs per-statement planning and execution timing as JSON records in the PostgreSQL server log. It is a log producer, not a SQL statistics view or persistent aggregate. The reviewed 0.0.6 release commit supports PostgreSQL 13 through 18 build features.

### Preload and Install

The module refuses ordinary session loading. Add it to preload configuration, restart the server, and then register the extension:

```conf
shared_preload_libraries = 'supamonitor'
```

```sql
CREATE EXTENSION supamonitor;
```

There are no documented SQL functions or GUCs. Removing or upgrading the module also requires a controlled server restart.

### Log Record Shape

Planner, executor, and utility hooks emit lines whose message begins with a versioned prefix followed by JSON:

```text
supamonitor_0.0.6_log:{"query_id":123,"query":"SELECT 1","calls":1,"total_plan_time":0.0,"total_exec_time":0.1}
```

Planning records use `calls` equal to zero and populate `total_plan_time`; completed execution or utility records use `calls` equal to one and populate `total_exec_time`. These are individual events, despite the field names, and must be collected and aggregated by an external logging pipeline.

### Security and Overhead

Records contain cleaned but not normalized query text, so SQL literals can expose credentials, personal data, tokens, or tenant identifiers to logs. Apply log access control, encryption, redaction downstream, bounded retention, and ingestion backpressure. The executor hook requests full instrumentation when a query has no timing state, and all major hooks add timing and JSON serialization work; benchmark overhead on representative query rates and utility-heavy workloads.

The project is active but not API-stable. The cataloged release commit reports 0.0.6, while the current default-branch manifest at the reviewed commit reports 0.0.0 after a source rewrite with materially different behavior. Pin the exact release source and PostgreSQL major rather than building an unqualified branch. Test hook chaining with other preload extensions, prepared statements, nested queries, errors, log rotation, and collector outages before production use.
