## Usage

Sources:

- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/README.md)
- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/pg_kv.control)
- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/Cargo.toml)
- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/src/lib.rs)

`pg_kv` 0.1.0 provides a shared-memory cache, counters, token-bucket rate limiting and sorted sets. This early release targets PostgreSQL 15–18 and is not production ready.

### Core Workflow

Add the library to the existing preload list and restart PostgreSQL before installing the extension as a superuser:

```ini
shared_preload_libraries = 'pg_kv'
```

```sql
CREATE EXTENSION pg_kv;
SELECT kv.set('app:session:42', '{"user":1}', interval '30 minutes');
SELECT kv.get('app:session:42');
SELECT kv.ttl('app:session:42');
SELECT kv.incr('app:views:7');
SELECT kv.rate_limit('app:api:42', 100, interval '1 minute');
SELECT kv.zadd('app:leaderboard', 'player:9', 1520);
SELECT * FROM kv.zrange('app:leaderboard', 0, 9, reverse => true);
SELECT * FROM kv.stats();
```

Installation fails if the library was not preloaded. Grant schema `kv` usage only to intended callers. The keyspace is shared across every database in the cluster; prefixes organize keys but do not enforce tenant isolation.

### Operations and Capacity

`kv.ttl` returns milliseconds, -1 for a permanent key and -2 for a missing key. `kv.expire` changes expiry, `kv.del` removes a key, and `kv.rate_limit_info` reports allowance and retry timing. Sorted sets also support `kv.zscore`, `kv.zrank`, `kv.zrevrank`, `kv.zincrby`, `kv.zcard` and `kv.zrem`.

Restart settings include `pg_kv.max_entries`, `pg_kv.max_key_bytes`, `pg_kv.max_value_bytes`, `pg_kv.shards`, `pg_kv.sweep_interval_ms`, `pg_kv.max_zset_members` and `pg_kv.max_member_bytes`. Storage is bounded; oversize values fail, and pressure can evict cached keys. `kv.flush` is an administrative operation unavailable to the public role by default.

### Durability Boundary

Cache changes survive transaction rollback but disappear on a server restart or crash. They are not WAL logged or replicated; a standby has its own cache. Use this for disposable state, not authoritative records or a globally coordinated limiter across independent servers.
