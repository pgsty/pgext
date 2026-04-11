
## Usage

> Syntax:
>
> ```sql
> SELECT pglock.lock('b3d8a762-3a0e-495b-b6a1-dc8609839f7b', 'users');
> SELECT pglock.unlock('b3d8a762-3a0e-495b-b6a1-dc8609839f7b', 'users');
> SELECT pglock.ttl();
> ```
>
> Source: [README](https://github.com/fraruiz/pglock)

`pglock` is a lightweight distributed lock service implemented inside PostgreSQL. It stores locks in a table and supports TTL-based expiration for stale locks.

## Basic Functions

The README documents four core functions:

- `pglock.lock(id, resource)` to acquire a lock
- `pglock.unlock(id, resource)` to release a lock
- `pglock.ttl()` to expire stale locks
- `pglock.set_serializable()` to switch to serializable isolation

Acquire a lock:

```sql
SELECT pglock.lock('worker-1', 'users');
```

Release it:

```sql
SELECT pglock.unlock('worker-1', 'users');
```

## Isolation

The upstream docs recommend serializable isolation for correctness under concurrency:

```sql
SELECT pglock.set_serializable();
```

or:

```sql
BEGIN ISOLATION LEVEL SERIALIZABLE;
SELECT pglock.lock('my-id', 'my-resource');
SELECT pglock.unlock('my-id', 'my-resource');
COMMIT;
```

## TTL Expiration

Locks have a configurable TTL with a documented default of 5 minutes. `pglock.ttl()` unlocks records whose `updated_at` is older than the TTL:

```sql
SELECT pglock.ttl();
```

If `pg_cron` is installed, the README says a cron job can run `pglock.ttl()` every minute.

## Schema

The lock table is `pglock.locks` with columns:

- `id`
- `resource`
- `locked`
- `ttl`
- `created_at`
- `updated_at`

The primary key is `(id, resource)`.
