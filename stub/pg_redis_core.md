## Usage

Sources:

- [README 0.4.1](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/README.md)
- [Configuration and SQL API](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/docs/getting-started/configuration.md)
- [Control](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/pg_redis_core.control)
- [Worker initialization](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/src/lib.rs)
- [Release with generated SQL](https://github.com/MegacoderKim/pg2redis/releases/tag/v0.4.1)

`pg_redis_core` 0.4.1, distributed as pg2redis, implements a Redis-compatible command layer backed by PostgreSQL tables. The reviewed release supplies PostgreSQL 15 and 16 artifacts. Its SQL API works without a background listener; RESP access requires a preloaded worker.

### SQL Workflow

Install the matching library and create the extension as a superuser in the intended database:

```sql
CREATE EXTENSION pg_redis_core;
SELECT redis_dispatcher('SET', 'hello', 'world', NULL);
SELECT redis_dispatcher('GET', 'hello', NULL, NULL);
SELECT key, convert_from(value, 'UTF8'), expires_at
FROM pg_redis_strings;
```

`redis_dispatcher` returns RESP2-encoded text. The storage tables include `pg_redis_strings`, `pg_redis_hashes`, `pg_redis_lists` and `pg_redis_zsets`.

### RESP Listener

To enable network clients, append the library to the existing preload setting and restart PostgreSQL. This example restricts the listener to the local host:

```ini
shared_preload_libraries = 'pg_redis_core'
pg_redis.bind_address = '127.0.0.1'
pg_redis.port = 6379
pg_redis.database = 'postgres'
pg_redis.unlogged = off
```

Create the extension in the database named by `pg_redis.database`. The upstream default listener address is `0.0.0.0`, so review network access before enabling it.

```bash
redis-cli -h 127.0.0.1 -p 6379 SET hello world
redis-cli -h 127.0.0.1 -p 6379 GET hello
```

### Administration and Limits

`pg_redis_evict_expired()` runs a bounded expiration pass. `pg_redis_ensure_schema(false)` ensures storage exists. `pg_redis_set_unlogged(true)` switches storage away from ordinary WAL durability; retain logged storage when crash recovery is required.

`pg_redis.cache_enabled`, `pg_redis.cache_ttl_ms` and `pg_redis.eviction_interval_ms` control caching and expiration. Listener address, port and target database are startup settings.

The project implements a specific Redis command subset, including strings, hashes, lists, sorted sets, streams and Pub/Sub; do not assume full Redis compatibility. The worker executes through its own PostgreSQL session, so RESP access is a separate access boundary from SQL login. This early release should be evaluated with the application's actual client commands and durability requirements.

### Authentication and Bounds in 0.4

`pg_redis.requirepass` enables RESP password checks through `AUTH` and `HELLO`; an empty value disables authentication. Store the password in protected server configuration. This shared RESP credential does not create per-client PostgreSQL roles.

`pg_redis.lua_time_limit_ms` limits sandboxed Lua calls, while `pg_redis.max_clients`, `pg_redis.max_frame_bytes` and `pg_redis.pubsub_queue_size` bound connections, request frames and subscriber queues. Expiration now covers all six supported data types through `pg_redis_key_expiry`. Existing installations must apply the supplied SQL upgrade chain after deploying matching files.
