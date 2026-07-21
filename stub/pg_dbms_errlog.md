## Usage

Sources:

- [Official v2.4 README](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/README.md)
- [v2.4 release changelog](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/ChangeLog)
- [v2.4 control file](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/pg_dbms_errlog.control)
- [v2.4 extension SQL](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/sql/pg_dbms_errlog--2.4.sql)

`pg_dbms_errlog` provides Oracle-style DML error logging for PostgreSQL. It queues an error from a failed `INSERT`, `UPDATE`, or `DELETE`, writes it to a registered `ERR$_...` table through background workers, and lets the surrounding script continue after rolling back to a savepoint. It requires either `pg_statement_rollback` or explicit savepoint management by the caller.

### Enable the Extension

Add the library to `shared_preload_libraries`, ensure `max_worker_processes` can accommodate `pg_dbms_errlog.max_workers` plus the fixed worker, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_dbms_errlog'
```

```sql
CREATE EXTENSION pg_dbms_errlog;
```

Create and register an error table for each DML target:

```sql
CREATE TABLE raises (
    employee_id integer,
    salary integer CHECK (salary > 8000)
);

CALL dbms_errlog.create_error_log('raises');
-- Creates and registers public."ERR$_raises" by default.
```

### Log and Continue after an Error

```sql
LOAD 'pg_statement_rollback';

SET pg_statement_rollback.enabled = on;
SET pg_dbms_errlog.enabled = on;
SET pg_dbms_errlog.query_tag = 'daily_load';
SET pg_dbms_errlog.reject_limit = 10;

BEGIN;
INSERT INTO raises VALUES (145, 15400);
INSERT INTO raises VALUES (161, 7700);  -- logged failure
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
INSERT INTO raises VALUES (175, 9680);
COMMIT;

SELECT * FROM "ERR$_raises";
```

The error table contains `pg_err_number$`, `pg_err_mesg$`, `pg_err_optyp$`, `pg_err_tag$`, `pg_err_query$`, and `pg_err_detail$`.

### API and Configuration Index

- `dbms_errlog.create_error_log(dml_table_name, err_log_table_name, err_log_table_owner, err_log_table_space)`: creates and registers an error table.
- `dbms_errlog.publish_queue(wait_for_completion)`: asks workers to process queued errors; execution is not granted to `PUBLIC` by default.
- `dbms_errlog.queue_size()`: reports queued errors.
- `pg_dbms_errlog.synchronous`: `transaction` by default, `query`, or `off`. Transaction mode guarantees that only errors from committed transactions are logged.
- `pg_dbms_errlog.reject_limit`: transaction-wide error limit; `-1` is unlimited and `0` logs nothing and rolls back.
- `pg_dbms_errlog.no_client_error`: suppresses client error messages while retaining server logging; enabled by default.
- `pg_dbms_errlog.frequency` and `pg_dbms_errlog.max_workers`: asynchronous worker timing and concurrency.

### Caveats

- A caller needs DML privileges on the target and error tables; creating an error table also requires execution and registration-table privileges described upstream.
- `INSERT INTO ... SELECT ...` is one PostgreSQL statement and cannot preserve only successful rows in the Oracle manner.
- Syntax and other parse-time errors are not logged. Stored query text must remain below PostgreSQL's 1 GB value limit.
- Version `2.4` changes no SQL API; it fixes worker shutdown loops and a dynamic-background-worker crash.
