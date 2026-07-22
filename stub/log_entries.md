## Usage

Sources:

- [Extension SQL for version 1.0](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/log_entries--1.0.sql)
- [Trigger implementation](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/log_entries.c)
- [Upstream regression example](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/sql/log_entries.sql)

`log_entries` is a historical C trigger that writes the current time and PostgreSQL role into two named columns during row inserts and updates. Version 1.0 requires the first target column to use the obsolete `abstime` type and the second to use `text`, which confines it to old PostgreSQL releases.

### Historical Workflow

On a PostgreSQL version that still provides `abstime`, attach the function with exactly two column-name arguments:

```sql
CREATE EXTENSION log_entries;

CREATE TABLE legacy_records (
  record_id integer PRIMARY KEY,
  payload text,
  updated_at abstime,
  updated_by text
);

CREATE TRIGGER legacy_records_log_entries
BEFORE INSERT OR UPDATE OR DELETE ON legacy_records
FOR EACH ROW
EXECUTE FUNCTION log_entries(updated_at, updated_by);
```

For insert and update events, the trigger writes the current absolute time and `current_user`. For delete events it returns the old tuple without changing audit columns. Direct calls and statement-level triggers are rejected, and the implementation requires exactly two arguments with the exact legacy types.

### Compatibility Boundary

The `abstime` type was removed from modern PostgreSQL, so the upstream SQL and C code cannot be used as written on supported current releases. The source is a single 2016 commit with no README, upgrade scripts, or compatibility matrix. Its C implementation also uses old trigger and tuple APIs whose behavior must not be assumed safe on newer server headers.

Do not replace `abstime` with `timestamptz` in the table alone: the binary function checks for the old type OID and will reject it. A migration requires a reviewed source-code port and new regression tests, or, preferably, a small PL/pgSQL trigger using `timestamptz` and an explicit actor policy.

### Audit Caveats

This trigger stores the effective database role, not an application user, original session role, client identity, statement text, or immutable history. Users who can update the audit columns or disable the trigger can alter the result. Row deletion leaves no separate audit record.

Version 1.0 is relocatable and declares no preload or extension dependency. Treat it as source archaeology for legacy databases. For a current audit requirement, use supported types, protect audit data and trigger ownership, define actor semantics, and test updates, deletes, replication, bulk loading, and restore behavior explicitly.
