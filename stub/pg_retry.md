


## Usage

> [pg_retry: Retry SQL statements on transient errors with exponential backoff](https://github.com/Agent-Hellboy/pg_retry)

### Function Signature

```sql
retry.retry(
  sql TEXT,                          -- SQL statement to run (exactly one)
  max_tries INT DEFAULT 3,           -- total attempts (1 + retries), >= 1
  base_delay_ms INT DEFAULT 50,      -- initial backoff delay in ms
  max_delay_ms INT DEFAULT 1000,     -- cap for exponential backoff
  retry_sqlstates TEXT[] DEFAULT ARRAY['40001','40P01','55P03','57014']
) RETURNS INT                       -- number of rows processed
```

Default retryable SQLSTATEs: `40001` (serialization_failure), `40P01` (deadlock_detected), `55P03` (lock_not_available), `57014` (query_canceled).

### Examples

Basic retry with defaults:

```sql
SELECT retry.retry('UPDATE accounts SET balance = balance - 100 WHERE id = 1');
```

Custom retry parameters:

```sql
SELECT retry.retry(
    'INSERT INTO audit_log (event) VALUES (''test'')',
    5,      -- max_tries
    100,    -- base_delay_ms
    5000    -- max_delay_ms
);
```

### GUC Configuration

```sql
ALTER SYSTEM SET pg_retry.default_max_tries = 5;
ALTER SYSTEM SET pg_retry.default_base_delay_ms = 100;
ALTER SYSTEM SET pg_retry.default_max_delay_ms = 5000;
ALTER SYSTEM SET pg_retry.default_sqlstates = '40001,40P01,55P03,57014';
SELECT pg_reload_conf();
```

### Safety Rules

- Only one SQL statement per call (multi-statement fails)
- Transaction control statements (BEGIN, COMMIT, ROLLBACK) are prohibited
- Parameters are validated (max_tries >= 1, non-negative delays, base <= max delay)
