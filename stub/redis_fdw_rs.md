## Usage

Sources:

- [Official README](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/README.md)
- [FDW setup example](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/examples/00_setup.sql)
- [Option and credential loading](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/src/utils/helpers.rs)
- [FDW option validator](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/src/core/validator.rs)
- [Connection-pool implementation](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/src/core/pool_manager.rs)
- [Extension control file](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/redis_fdw_rs.control)

`redis_fdw_rs` 0.6.0 is a Rust foreign data wrapper for reading and modifying Redis strings, hashes, lists, sets, sorted sets, and streams through PostgreSQL foreign tables. It supports PostgreSQL 14 through 18, Redis Cluster, TLS, connection pooling, predicate pushdown, batch inserts, TTLs, schema import, and selected join pushdown.

### Create the Wrapper and a Table

Installing the extension creates the handler and validator functions, but not a named foreign-data wrapper or server. Create those explicitly, then map each table to one Redis key or key pattern:

```sql
CREATE EXTENSION redis_fdw_rs;

CREATE FOREIGN DATA WRAPPER redis_wrapper
HANDLER redis_fdw_handler
VALIDATOR redis_fdw_validator;

CREATE SERVER redis_server
FOREIGN DATA WRAPPER redis_wrapper
OPTIONS (host_port '127.0.0.1:6379');

CREATE FOREIGN TABLE user_profiles (
    field text,
    value text
)
SERVER redis_server
OPTIONS (
    database '0',
    table_type 'hash',
    table_key_prefix 'user:profiles'
);

INSERT INTO user_profiles VALUES
    ('name', 'Ada'),
    ('email', 'ada@example.com');

SELECT value FROM user_profiles WHERE field = 'email';
```

The hash predicate can be pushed to `HGET`. Use `EXPLAIN (VERBOSE)` to confirm pushdown instead of assuming it. Column shapes are position-based: string has one value column, hash has field/value, list has element or index/element, set has member, sorted set has member/score, and stream has an ID plus field/value columns.

### Redis Types and Options

- `table_type`: one of `string`, `hash`, `list`, `set`, `zset`, or `stream`; streams are append-only for updates.
- `table_key_prefix`: an exact key or a glob pattern. Pattern tables add a leading key column and may require a Redis `SCAN`.
- `database`: logical Redis database 0 through 15; Redis Cluster effectively uses database 0.
- `ttl`: default expiration in seconds. An optional `ttl bigint` column supplies a per-row value; `-1` persists a key.
- `batch_size` and `join_batch_size`: bound pipelined insert and parameterized-join batches.
- `IMPORT FOREIGN SCHEMA`: samples up to 10,000 Redis keys, groups prefixes, and generates foreign-table definitions.

The FDW implements `SELECT`, `INSERT`, `DELETE`, and most `UPDATE` operations for all six types, except stream `UPDATE`. Same-server single-column equality joins have limited pushdown; multi-key patterns and streams have additional restrictions. Check the official matrix for the exact operation before designing a write path.

### TLS, Credentials, and Pooling

Use `rediss://` in `host_port` to require TLS certificate verification. The documented `#insecure` suffix disables verification and is appropriate only for isolated testing.

At the reviewed commit, the runtime combines wrapper, server, and table options but has PostgreSQL user-mapping option loading commented out. Consequently `username` and `password` must be server options to take effect, where privileged catalog readers and DDL viewers may see them. Restrict ownership and visibility, prefer a narrowly scoped Redis ACL identity, and test a credential rotation procedure. The per-backend pool cache key distinguishes the username or the presence of a password but does not include the password value, so a changed secret can reuse an old pool until connections or backends are recycled.

### Consistency and Destructive Operations

Redis commands take effect remotely and are not rolled back as part of a PostgreSQL transaction. Design writes to be idempotent and expect partial effects after statement errors, client cancellation, failover, or retries.

`TRUNCATE` unlinks the exact Redis key. On a pattern table it performs `SCAN` plus batched `UNLINK` and can remove every matching key; multi-key `TRUNCATE` is unsupported in cluster mode. Review `table_key_prefix` and Redis ACL scope before granting truncation. Full pattern scans, schema import, and pushed-down same-server joins can consume substantial Redis, network, or PostgreSQL memory, so test production key cardinalities and enforce query timeouts.
