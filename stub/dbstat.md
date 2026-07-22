## Usage

Sources:

- [Official README](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/README.md)
- [Official extension SQL](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/sql/dbstat--0.1.sql)
- [Official collector source](https://github.com/spd010273/dbstat/blob/aa56284e40a428393985b687fa6ad8199712bd51/dbstat.c)

`dbstat` is an experimental table-row counter built from database triggers, `NOTIFY`, and a separate libpq collector process. It is a learning project rather than a durable statistics service: version `0.1` observes only tables present in the `public` schema at installation time, and missed notifications make its counts inaccurate.

### Core Workflow

Creating the extension installs the fixed `dbstat` schema, scans ordinary `public` tables, stores an initial exact count, and adds an `AFTER INSERT OR DELETE OR UPDATE` row trigger to each one:

```sql
CREATE EXTENSION dbstat;

SELECT schema_name, table_name, row_count
FROM dbstat.tb_catalog_table
ORDER BY schema_name, table_name;
```

Run the separately built collector continuously for the same database. It listens on one channel per registered table and updates the stored count when a trigger sends an operation:

```console
dbstat -h localhost -p 5432 -U postgres -d appdb
```

Recording is on when `dbstat.enable_logging` is unset. A session can suppress notifications with `SET dbstat.enable_logging = false`, but those changes are then permanently absent from the derived count.

### Object Index

- `dbstat.tb_catalog_table` records table OIDs, names, and estimated current row counts.
- `dbstat.tb_catalog_table_modification` is an unlogged event table written by the collector.
- `dbstat.fn_install_triggers()` registers a table and computes its initial exact count.
- `dbstat.fn_issue_notify()` is the row-trigger function.
- `dbstat.fn_get_table_oid()` and `dbstat.fn_log_action()` support registration and event logging.

### Operational Boundaries

Installation is superuser-only and mutates every existing ordinary table in `public`; inspect those trigger changes before use. Tables created later are not registered automatically. The trigger sends only the operation, so collector downtime, transaction-notification coalescing, disabled logging, crashes, and unlogged-table resets can lose history with no recovery mechanism. The registration function builds SQL from unquoted schema and table strings, and the code targets old PostgreSQL APIs. Use `COUNT(*)` reconciliation, least-privilege collector credentials, and a process supervisor if evaluating it; do not treat the result as audit evidence or authoritative cardinality.
