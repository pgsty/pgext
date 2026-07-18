## Usage

Sources:

- [Current database.dev package documentation](https://database.dev/jacobfvanzyl/pgkv)

`pgkv` is a pure-SQL, Redis-inspired key-value layer backed by PostgreSQL tables and `JSONB`. It supports strings and counters, hashes, lists, sets, sorted sets, multi-key operations, glob-style key lookup, and optional time-to-live values while retaining PostgreSQL transactions.

```sql
CREATE EXTENSION pgkv;
SELECT pgkv.set('greeting', 'hello');
SELECT pgkv.get('greeting');
SELECT pgkv.set('session:42', 'active', 60);
SELECT pgkv.ttl('session:42');
```

Values live in the extension's shared store rather than an external Redis server. Expiration is lazy on reads; schedule `pgkv.cleanup_expired()` if expired rows must be reclaimed promptly. `pgkv.keys()` can scan the whole store, and large set/sorted-set operations execute inside PostgreSQL, so benchmark contention, WAL volume, and table growth.

The package also exposes destructive administration such as `pgkv.flushall()`. Restrict function execution by role, especially deletion, cleanup, pattern scans, and global flush. Version 0.1.0 is new provider-distributed SQL; inspect the installed script and test upgrade behavior before making it an application persistence contract.
