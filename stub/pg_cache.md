## Usage

Sources:

- [Upstream source 1](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/README.md)
- [Upstream source 2](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/crates/pg-cache/pg_cache.control)

`pg_cache` 0.1.0 is the SQL identity of the experimental Flaro RESP2 cache for PostgreSQL 17.

### Core Workflow

Append the preload, restart, and create the extension as superuser in the configured database. Background workers serve the supported RESP2 commands and persist batches to `cache_store`. `pg_cache.listen_address` controls the listener and requires restart when changed.

```ini
shared_preload_libraries = 'pg_cache'
pg_cache.database = 'postgres'
pg_cache.port = 6380
pg_cache.memory_mb = 64
pg_cache.flush_ms = 100
```

```sql
CREATE EXTENSION pg_cache;
```

```bash
redis-cli -p 6380 SET greeting hello
redis-cli -p 6380 GET greeting
```

### Operational Boundaries

The live cache is in memory. A successful RESP response can precede PostgreSQL persistence, so a crash can lose acknowledged writes. Queue and memory limits can produce backpressure. This is a limited protocol implementation with no RESP authentication; keep it on a trusted network and inspect durability progress separately from command acknowledgments.
