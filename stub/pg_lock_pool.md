## Usage

Sources:

- [Official README](https://github.com/adjust/pg_lock_pool/blob/4d8a59ad82b6170b4fe827e0fa5ce4c8bc4b4695/Readme.md)
- [Extension control file](https://github.com/adjust/pg_lock_pool/blob/4d8a59ad82b6170b4fe827e0fa5ce4c8bc4b4695/pg_lock_pool.control)
- [Version 0.0.1 extension SQL](https://github.com/adjust/pg_lock_pool/blob/4d8a59ad82b6170b4fe827e0fa5ce4c8bc4b4695/pg_lock_pool--0.0.1.sql)

`pg_lock_pool` 0.0.1 provides PL/pgSQL helpers that wait for any free slot in a pool represented by two-key PostgreSQL advisory locks. Use it to cap cooperating work to `poolsize` concurrent holders without creating a lock table.

### Session-Level Pool

```sql
CREATE EXTENSION pg_lock_pool;

SELECT get_lock_pool(999, 3, 30) AS slot;
SELECT pg_advisory_unlock(999, 1);
```

`get_lock_pool(key, poolsize, timeout)` probes slots `1..poolsize` with `pg_try_advisory_lock(key, slot)`. On success it returns the positive slot number; the caller must pass the same key and returned slot to `pg_advisory_unlock`. A session-level lock otherwise remains until the database session ends, including across transaction rollback.

### Transaction-Level Pool

```sql
BEGIN;
SELECT get_xact_lock_pool(999, 3, 30) AS slot;
COMMIT;
```

`get_xact_lock_pool` uses `pg_try_advisory_xact_lock` and releases the chosen slot automatically at transaction end. Both functions default `timeout` to 10 seconds and return `-1` when no slot is acquired before the loop reaches the timeout.

### Operational Semantics

Each unsuccessful pass tries every slot, checks its integer retry counter, then sleeps for one second. The timeout is therefore an approximate retry duration rather than a precise wall-clock deadline, and each waiting caller occupies a PostgreSQL backend.

Advisory locks are cooperative: unrelated SQL can ignore them, and every participant must use the same `key`, `poolsize`, and session-versus-transaction convention. Validate that `poolsize` and `timeout` are positive in application code. Always retain the returned session slot and release it in error paths; connection pooling can otherwise keep the lock far longer than the logical request.
