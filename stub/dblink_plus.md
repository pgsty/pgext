## Usage

Sources:

- [Official README](https://github.com/ossc-db/dblink_plus/blob/cb0efffe7f9098e352ccf5da65738630cf40dc5a/README.md)
- [Official documentation](https://ossc-db.github.io/dblink_plus/index.html)
- [Official extension SQL template](https://github.com/ossc-db/dblink_plus/blob/cb0efffe7f9098e352ccf5da65738630cf40dc5a/dblink_plus.sql.in)

`dblink_plus` 1.0.11 executes SQL on external PostgreSQL, Oracle, MySQL, or SQLite databases through connector plugins. Unlike core `dblink`, it can coordinate a remote transaction with the local transaction using two-phase commit when `use_xa` is enabled.

### Core Workflow

Create a foreign-data-wrapper definition, server, and user mapping for the compiled connector, then open a named connection:

```sql
CREATE EXTENSION dblink_plus;

CREATE FOREIGN DATA WRAPPER remote_postgres
  VALIDATOR dblink.postgres;

CREATE SERVER reporting_db
  FOREIGN DATA WRAPPER remote_postgres
  OPTIONS (dbname 'reporting');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER reporting_db
  OPTIONS (user 'report_reader');

BEGIN;
SELECT dblink.connect('reporting_conn', 'reporting_db');

SELECT *
FROM dblink.query(
  'reporting_conn',
  'SELECT id, name FROM customers ORDER BY id'
) AS t(id bigint, name text);

SELECT dblink.exec(
  'reporting_conn',
  'UPDATE job_state SET checked_at = CURRENT_TIMESTAMP WHERE id = 1'
);
COMMIT;

SELECT dblink.disconnect('reporting_conn');
```

Because `dblink.query` and `dblink.call` return `SETOF record`, callers must provide the expected column names and types. Passing a server name instead of a named connection creates an anonymous connection that is closed at transaction end.

### Important Objects

- `dblink.connect(...)` and `dblink.disconnect(text)` manage persistent named connections.
- `dblink.query(...)` returns rows, while `dblink.exec(text, text)` returns the affected-row count for a remote command.
- `dblink.open(...)`, `dblink.fetch(...)`, and `dblink.close(...)` provide transaction-scoped cursor access.
- `dblink.call(...)` invokes a stored function and returns records.
- `dblink.connections` reports connection name, server OID, status, `use_xa`, and whether the connection is kept.
- `dblink.postgres`, `dblink.oracle`, `dblink.mysql`, and `dblink.sqlite3` are connector validators present in the source tree.

### Transaction and Configuration Boundaries

`dblink_plus.use_xa` controls automatic transaction management for anonymous connections and defaults to on. An explicit `use_xa` argument to `dblink.connect` takes precedence; omitting that argument starts an XA-enabled connection even if the GUC is off. PostgreSQL remote servers need `max_prepared_transactions` greater than zero for two-phase commit. Commands forbidden inside a transaction block, such as PostgreSQL `VACUUM` or many Oracle DDL commands, require a connection created with `use_xa = false`.

No preload is required. The connector libraries and client dependencies are selected at build time, so verify which of PostgreSQL, Oracle, MySQL, and SQLite is actually present in the installed binary.

### Operational Notes

The XA path uses a deferred delete trigger in `dblink.atcommit`; it cannot run in a read-only or hot-standby transaction, where `use_xa` must be disabled. Large remote results are materialized in the local backend and can consume substantial memory. Oracle `max_value_len` and fetch-size settings directly affect buffer size, and the official documentation lists unsupported Oracle LOB/raw types. Grant non-superusers only the required `USAGE` on schema `dblink`, function `EXECUTE`, and the documented `DELETE` privilege on the internal table. Test rollback, prepared-transaction recovery, logging, credential storage, and remote failure behavior before production use.
