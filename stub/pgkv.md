## Usage

Sources:

- [Current database.dev package documentation](https://database.dev/jacobfvanzyl/pgkv)

`pgkv` 0.1.0 is a Redis-inspired PostgreSQL Trusted Language Extension distributed as `jacobfvanzyl@pgkv`. It is implemented in `PL/pgSQL`, stores typed values in `JSONB` inside `pgkv.store`, and provides transactional key-value commands without running a Redis server.

### Core Workflow

```sql
CREATE EXTENSION pgkv;

SELECT pgkv.set('user:42', '{"name":"Ada","active":true}');
SELECT pgkv.get('user:42')->>'name';

SELECT pgkv.set('session:42', 'active', 60);
SELECT pgkv.ttl('session:42');

SELECT pgkv.incr('jobs:completed');
```

PostgreSQL 12 or later with `PL/pgSQL` is required. Keys have one of the extension's logical types: string, hash, list, set, or sorted set. Calling a command for the wrong type raises `WRONGTYPE`; numeric counters are stored as JSONB numbers.

### Command Groups

- Basic keys: `pgkv.set`, `pgkv.get`, `pgkv.del`, `pgkv.exists`, `pgkv.type`, and `pgkv.dbsize`.
- Expiration: `pgkv.expire`, `pgkv.ttl`, and `pgkv.cleanup_expired`.
- Strings and counters: `pgkv.incr`, `pgkv.decr`, `pgkv.incrby`, `pgkv.decrby`, `pgkv.append`, `pgkv.strlen`, `pgkv.getrange`, and `pgkv.setrange`.
- Hashes: `pgkv.hset`, `pgkv.hget`, `pgkv.hmget`, `pgkv.hgetall`, `pgkv.hdel`, `pgkv.hkeys`, and `pgkv.hvals`.
- Lists: `pgkv.lpush`, `pgkv.rpush`, `pgkv.lpop`, `pgkv.rpop`, `pgkv.lrange`, `pgkv.lindex`, `pgkv.ltrim`, and `pgkv.lrem`.
- Sets: `pgkv.sadd`, `pgkv.srem`, `pgkv.smembers`, `pgkv.sinter`, `pgkv.sunion`, and `pgkv.sdiff`.
- Sorted sets: `pgkv.zadd`, `pgkv.zrem`, `pgkv.zrange`, `pgkv.zrevrange`, `pgkv.zscore`, `pgkv.zrank`, and `pgkv.zrangebyscore`.
- Multi-key and inspection operations include `pgkv.mset`, `pgkv.mget`, and `pgkv.keys`.

### Expiration, Performance, and Privileges

Expiration is lazy: reads detect an expired key and delete it. Schedule `pgkv.cleanup_expired()` when expired rows must be reclaimed without waiting for access. `pgkv.keys()` accepts Redis-style `*`, `?`, and `[...]` patterns, but it can scan the whole shared store; character classes use regular expressions while simpler patterns use `LIKE`.

All values and collection operations remain PostgreSQL work. Large set intersections and differences, sorted-set sorting, contention, WAL, vacuum, and table growth can therefore dominate database resources. Benchmark realistic key and collection sizes instead of assuming Redis latency or memory behavior.

The store is shared within the database and the package does not provide an application or tenant authorization layer. Grant `EXECUTE` only on the commands each role needs. Keep `pgkv.flushall()` and `pgkv.cleanup_expired()` administrative, and tightly control deletion, pattern scans, and multi-key mutation. Version 0.1.0 is the only published package version shown by database.dev; review the installed SQL and test dump, restore, permissions, concurrency, and future upgrades before making it a persistence contract.
