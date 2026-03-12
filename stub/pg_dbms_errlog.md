

## Usage

> [pg_dbms_errlog: Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table](https://github.com/HexaCluster/pg_dbms_errlog)

Enables DML operations to continue after encountering errors by logging failures to an error table, rather than aborting the transaction.

### Enabling

Add to `shared_preload_libraries` in `postgresql.conf`:

```ini
shared_preload_libraries = 'pg_dbms_errlog'
```

```sql
CREATE EXTENSION pg_dbms_errlog;
LOAD 'pg_dbms_errlog';
```

### Create Error Log Table

```sql
BEGIN;
CALL dbms_errlog.create_error_log('employees');
END;
-- Creates table "ERR$_employees" with error logging columns

-- With custom name and schema:
BEGIN;
CALL dbms_errlog.create_error_log('hr.employees', '"ERRORS"."ERR$_EMPTABLE"');
END;
```

### Configuration

```sql
SET pg_dbms_errlog.enabled TO true;       -- enable error logging
SET pg_dbms_errlog.query_tag TO 'daily_load';  -- tag for identifying statements
SET pg_dbms_errlog.reject_limit TO 10;    -- max errors before rollback (-1=unlimited)
SET pg_dbms_errlog.synchronous TO 'transaction'; -- 'transaction', 'query', or 'off'
SET pg_dbms_errlog.no_client_error TO true;      -- suppress client error messages
```

### Usage with pg_statement_rollback

```sql
LOAD 'pg_dbms_errlog';
LOAD 'pg_statement_rollback';

CREATE TABLE hr.raises (emp_id integer, sal integer CHECK(sal > 8000));

BEGIN;
CALL dbms_errlog.create_error_log('hr.raises');
END;

SET pg_dbms_errlog.query_tag TO 'daily_load';
SET pg_dbms_errlog.reject_limit TO 10;
SET pg_dbms_errlog.enabled TO true;

BEGIN;
SET pg_statement_rollback.enabled TO on;
INSERT INTO hr.raises VALUES (145, 15400);  -- Success
INSERT INTO hr.raises VALUES (161, 7700);   -- Failure (logged)
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
INSERT INTO hr.raises VALUES (175, 9680);   -- Success
COMMIT;
```

### Viewing Error Logs

```sql
SELECT * FROM "ERR$_raises";
-- pg_err_number$  | 23514
-- pg_err_mesg$    | new row for relation "raises" violates check constraint
-- pg_err_optyp$   | I
-- pg_err_tag$     | daily_load
-- pg_err_query$   | INSERT INTO hr.raises VALUES (161, 7700);
```

### Flush Queued Errors

```sql
SELECT dbms_errlog.publish_queue();
```
