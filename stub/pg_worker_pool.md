## Usage

Sources:

- [Pinned upstream README](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/README.md)
- [Version 0.1 SQL API](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/pg_worker_pool--0.1.sql)
- [Pinned worker and queue implementation](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/pg_worker_pool.c)
- [Pinned extension control file](https://github.com/serpent7776/pg_worker_pool/blob/574258dc96f766a7f4d9ce7023c575dd21ec9bbf/pg_worker_pool.control)

pg_worker_pool version 0.1 creates on-demand named background workers that execute SQL stored in worker_pool.jobs. A worker starts after the transaction that submitted or launched it commits, connects to the submitting database as the submitting PostgreSQL role, drains its named queue, and exits when no waiting row remains.

### Preload and submit

The library must allocate shared memory during postmaster startup. Add it before creating the extension and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_worker_pool'
```

```sql
CREATE EXTENSION pg_worker_pool;

CALL worker_pool.submit(
  'indexes',
  'CREATE INDEX my_table_name_idx ON my_table(name)'
);

CALL worker_pool.launch('indexes');

SELECT id, worker_name, query_text, status
FROM worker_pool.jobs
ORDER BY id;
```

The compiled pool has eight cluster-wide slots. max_worker_processes must reserve capacity for them and other extensions. Changing the pool limit requires editing MAX_WORKERS and rebuilding.

### Injection, identity, and recovery risks

The C code interpolates worker name and query text directly into internal SQL with percent-s formatting. It does not use SPI parameters or quote literals. Ordinary job SQL containing a single quote can fail submission, and crafted worker or query text can alter the internal INSERT or queue-selection statement. Do not expose submit or launch to untrusted callers. Revoke PUBLIC execution, restrict worker_pool tables, and accept only application-generated identifiers and preapproved statements.

Pool identity is keyed only by worker name, not by database or role. If different users or databases reuse a name while one worker is active, work can be stranded or executed under the role that started the existing worker. A query is executed as arbitrary SQL under that role. Dedicate unique cluster-wide names, avoid mixed-trust roles, and isolate this extension to one controlled database.

The dequeue query uses LIMIT 1 without ORDER BY, so source code does not guarantee submission order despite the README's claim. A backend crash can leave a job pending indefinitely, and failures store only a failed status without the error details. There is no retry, cancellation, lease, ownership, timeout, or stale-job recovery API.

Use this project only after fixing parameterization and identity isolation. Test commit and rollback launch behavior, concurrent submitters, role changes, worker exhaustion, restart, crash recovery, long transactions, DDL locks, and idempotent retry. The pinned repository supplies no license and no production support policy.
