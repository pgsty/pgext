## Usage

Sources:

- [Official upstream README](https://github.com/sanidhyavijay24/pgtime/blob/d5815a9780dc3f943e67ffbc9cc97d92c9db53e7/README.md)
- [Official extension control file (pgtime.control)](https://github.com/sanidhyavijay24/pgtime/blob/d5815a9780dc3f943e67ffbc9cc97d92c9db53e7/extension/pgtime.control)
- [Official extension SQL (pgtime--0.1.sql)](https://github.com/sanidhyavijay24/pgtime/blob/d5815a9780dc3f943e67ffbc9cc97d92c9db53e7/extension/pgtime--0.1.sql)

`pgtime` — Temporal tables extension for PostgreSQL. Automatically tracks Transaction Time (system time) on any table using range types and a high-performance C-based AFTER ROW trigger. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgtime;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgtime.as_of(target_table TEXT, ts TIMESTAMPTZ)` is an extension function and returns `SETOF`.
- `pgtime.attach(target_table TEXT)` is an extension function and returns `VOID`.
- `pgtime.detach(target_table TEXT)` is an extension function and returns `VOID`.
- `pgtime.diff(target_table TEXT, t1 TIMESTAMPTZ, t2 TIMESTAMPTZ)` is an extension function and returns `SETOF`.
- `pgtime.history(target_table TEXT, row_id anyelement)` is an extension function and returns `SETOF`.
- `pgtime.pgtime_trigger_fn()` is an extension function and returns `TRIGGER`.
- `pgtime.versions(target_table TEXT, row_id anyelement)` is an extension function and returns `BIGINT`.
- `pgtime` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
