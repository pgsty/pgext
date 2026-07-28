## Usage

Sources:

- [Official upstream README](https://github.com/mizcausevic-dev/pg-audit-stream-extension/blob/73a456241b0c6fa0ef4f46ba84c97bc0e39cc126/README.md)
- [Official extension control file (audit_stream.control)](https://github.com/mizcausevic-dev/pg-audit-stream-extension/blob/73a456241b0c6fa0ef4f46ba84c97bc0e39cc126/audit_stream.control)
- [Official extension SQL (audit_stream--0.1.0.sql)](https://github.com/mizcausevic-dev/pg-audit-stream-extension/blob/73a456241b0c6fa0ef4f46ba84c97bc0e39cc126/audit_stream--0.1.0.sql)

`audit_stream` — > A Postgres extension that turns any table-level CRUD into an audit-stream-py-compatible governance event — without writing one line of application code. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION audit_stream;
SELECT audit_stream.watch('decisions', 'decision_card_status_changed', 'procurement-api');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `list_watches()` is an extension function and returns `TABLE`.
- `unwatch(p_table TEXT)` is an extension function and returns `BOOLEAN`.
- `watch(p_table TEXT, p_event_kind TEXT, p_source TEXT DEFAULT NULL)` is an extension function and returns `TEXT`.
- `watches` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
