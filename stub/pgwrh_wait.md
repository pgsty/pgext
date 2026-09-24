## Usage

Sources:

- [Logical apply visibility guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/lsn-wait.md)
- [Control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_wait/pgwrh_wait.control)
- [Extension SQL](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_wait/pgwrh_wait--1.0.0-alpha1.sql)
- [Alpha release boundaries](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/releases/1.0.0-alpha1.md)

`pgwrh_wait` 1.0.0-alpha1 waits until a logical-replication subscriber has applied a publisher watermark before a read takes its snapshot. It targets PostgreSQL 18 with the built-in `pgoutput` plugin and works independently of the other pgwrh extensions. This alpha has fresh-install SQL only.

### Enable on the Subscriber

Append the library to `shared_preload_libraries`, preserving other entries, and restart PostgreSQL. Then install the SQL API as an administrator in each database that serves guarded reads:

```ini
shared_preload_libraries = 'pgwrh_wait'
```

```sql
CREATE EXTENSION pgwrh_wait;
SELECT pgwrh.applied_lsn('pgwrh_replica_subscription');
```

Run the diagnostic call in a separate health-check transaction before relying on settings. Creating the extension without preloading is allowed, but the API reports a prerequisite error. A successful custom setting assignment does not prove the library is active. `pgwrh.applied_lsn` can return NULL until monitoring initializes and does not itself certify initial table synchronization.

### Wait Before Taking a Snapshot

Use the actual subscriber connection that will execute the read. Replace the example LSN with a watermark obtained from the correct publisher after the relevant write committed, and use the actual subscription and table names:

```sql
BEGIN ISOLATION LEVEL REPEATABLE READ READ ONLY;
SET LOCAL pgwrh.read_after_subscription = 'pgwrh_replica_subscription';
SET LOCAL pgwrh.wait_timeout_ms = '5s';
SET LOCAL pgwrh.read_after_lsn = '0/12345678';
SELECT * FROM my_table;
COMMIT;
```

Send the watermark before any query or snapshot, including `SELECT 1` or `set_config`. The setting must be applied directly in the explicit top-level transaction, not inside a function or savepoint. Read Committed, Repeatable Read, and Serializable are supported. Timeout or cancellation aborts the operation; roll back instead of proceeding with a stale read.

`pgwrh.wait_for_lsn(text, pg_lsn, integer)` is an alternative SQL function with a default 10,000 ms timeout. It protects only a separate, subsequent Read Committed statement; a wait and read in one statement share the old snapshot. The function rejects Repeatable Read and Serializable. Both wait APIs are available to ordinary users, subject to normal table privileges.

### Replication and Recovery Requirements

All subscription tables must have completed initial synchronization. Two-phase subscriptions and a pending transaction skip are rejected. An LSN belongs to one publisher history; routing must preserve that identity. Progress advances on applied commits, not received WAL or keepalives. A watermark sampled after commit can exceed the last published commit, so an idle or filtered subscription may require a later replicated heartbeat transaction.

`pgwrh.max_tracked_subscriptions` defaults to 256 and requires restart to change. Dropped subscriptions still consume entries until restart. With foreign tables, each actual remote transaction must apply a barrier before its snapshot. The guarantee is a visibility lower bound, not a cluster-wide snapshot or a failover durability guarantee. Re-establish a valid baseline after skipped changes, divergence, or recovery.
