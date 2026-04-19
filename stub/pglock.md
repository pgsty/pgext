## Usage

- Source: [README](https://github.com/fraruiz/pglock/blob/master/README.md)

`pglock` is a lightweight distributed lock service implemented inside PostgreSQL. It stores lock state in `pglock.locks` and uses TTL-based cleanup for stale rows.

### Create The Extension

```sql
CREATE EXTENSION pglock;
```

The upstream README lists PostgreSQL 9.1+ and `plpgsql` as requirements.

### Acquire And Release Locks

```sql
SELECT pglock.lock('worker-1', 'users');
SELECT pglock.unlock('worker-1', 'users');
```

`pglock.lock(id, resource)` returns `true` if the lock is acquired and `false` if another process already holds it. `pglock.unlock(id, resource)` is documented as idempotent.

### Isolation And Expiration

The README recommends serializable isolation for correctness under concurrency:

```sql
SELECT pglock.set_serializable();

BEGIN ISOLATION LEVEL SERIALIZABLE;
SELECT pglock.lock('my-id', 'my-resource');
SELECT pglock.unlock('my-id', 'my-resource');
COMMIT;
```

Stale locks are expired with:

```sql
SELECT pglock.ttl();
```

The documented default TTL is 5 minutes. If `pg_cron` is available, the README says `pglock.ttl()` can be scheduled every minute.

### Lock Table

The upstream schema is `pglock.locks` with columns `id`, `resource`, `locked`, `ttl`, `created_at`, and `updated_at`. The primary key is `(id, resource)`.
