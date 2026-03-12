

## Usage

> [pg_statement_rollback: Server side rollback at statement level for PostgreSQL like Oracle or DB2](https://github.com/lzlabs/pg_statement_rollback)

Provides automatic server-side savepoints before each statement, allowing individual statement failures without aborting the entire transaction.

### Enabling

```sql
LOAD 'pg_statement_rollback.so';
SET pg_statement_rollback.enabled TO on;
```

Or in `postgresql.conf` for all sessions:

```ini
session_preload_libraries = 'pg_statement_rollback'
pg_statement_rollback.enabled = on
```

### Basic Usage

```sql
BEGIN;
CREATE TABLE test(id integer);
INSERT INTO test SELECT 1;
SELECT COUNT(*) FROM test;         -- returns 1

INSERT INTO test SELECT 'wrong';   -- ERROR: invalid input syntax
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";  -- rollback only the failed statement

SELECT COUNT(*) FROM test;         -- still returns 1
COMMIT;
```

Without this extension, the error would abort the entire transaction and all subsequent statements would fail with "current transaction is aborted".

### Configuration

```sql
-- Enable/disable at any time in a session
SET pg_statement_rollback.enabled TO off;

-- Change the savepoint name (superuser only)
SET pg_statement_rollback.savepoint_name TO 'my_savepoint';

-- Limit savepoints to write-only statements (default: on)
SET pg_statement_rollback.enable_writeonly TO off;
```

### Key Behaviors

- Automatic savepoints are created server-side with minimal performance overhead
- By default, savepoints are only created after write statements (INSERT, UPDATE, DELETE, etc.)
- When `enable_writeonly` is on, SELECT statements do not trigger automatic savepoints
- The client must still call `ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt"` when handling errors
