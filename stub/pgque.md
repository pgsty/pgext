## Usage

Sources:

- [Extension control file](https://github.com/scttnlsn/pgque/blob/6d0af3b0bb95d02d98fa316051040ff5a6a8409d/pgque.control)
- [Version 0.1.0 SQL implementation](https://github.com/scttnlsn/pgque/blob/6d0af3b0bb95d02d98fa316051040ff5a6a8409d/pgque--0.1.0.sql)
- [Basic exercise](https://github.com/scttnlsn/pgque/blob/6d0af3b0bb95d02d98fa316051040ff5a6a8409d/test/basic.sql)

`pgque` 0.1.0 is a small pure-SQL job queue. Jobs are rows in `pgque_jobs`, ordered by ascending `priority` and `id`. Workers call `pgque_lock()` to acquire a session-level advisory lock, then explicitly destroy and unlock the returned job.

### Minimal worker flow

```sql
CREATE EXTENSION pgque;

SELECT (pgque_enqueue(
  'mail',
  10,
  '{"message_id":123}'::jsonb
)).*;

SELECT * FROM pgque_lock('mail');

-- After processing the returned job in this same session:
SELECT pgque_destroy(42);
SELECT pgque_unlock(42);
```

Replace `42` with the returned `id`. The advisory lock belongs to the database session, survives transaction rollback, and is released when the session ends. `pgque_destroy()` deletes the row but does not release the lock, so every worker path must call `pgque_unlock()` or close the session.

### Delivery and coordination limits

This design provides neither exactly-once delivery nor a visibility timeout. A worker crash releases the lock and leaves an existing row available again; an external side effect may already have occurred. Deleting before or after an external effect creates different loss and duplication windows. Consumers must use idempotency keys and a durable application protocol.

The bigint advisory-lock namespace is shared with other application locks, so job IDs can collide with unrelated conventions. Version 0.1.0 has no retry count, scheduled time, dead-letter state, lease ownership, timeout recovery, payload schema, or retention policy. It also has no current compatibility or concurrency test suite. Use it only after stress-testing queue starvation, transactions, pooling mode, worker cancellation, and crash recovery; a maintained queue implementation is preferable for production work.
