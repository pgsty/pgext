## Usage

Sources:

- [Official README](https://github.com/Florents-Tselai/spat/blob/bf7d47b6f6bf0b65bd0d0704a08edadea60942b9/README.md)
- [Official extension SQL](https://github.com/Florents-Tselai/spat/blob/bf7d47b6f6bf0b65bd0d0704a08edadea60942b9/sql/spat--0.1.0a5.sql)
- [Official extension control file](https://github.com/Florents-Tselai/spat/blob/bf7d47b6f6bf0b65bd0d0704a08edadea60942b9/spat.control)

`spat` is an alpha, Redis-like key/value store embedded in PostgreSQL dynamic shared memory. It provides strings, sets, lists, hashes, TTLs, and named in-memory namespaces for caches or cross-session transient state; it is not a transactional or durable table replacement.

### Core Workflow

```sql
CREATE EXTENSION spat;

SELECT spset('session:42', 'ready', ttl => interval '5 minutes');
SELECT spget('session:42');
SELECT ttl('session:42');

SELECT sadd('roles:42', 'reader');
SELECT sadd('roles:42', 'writer');
SELECT sismember('roles:42', 'writer');

SELECT lpush('jobs', 'job-1');
SELECT rpush('jobs', 'job-2');
SELECT lpop('jobs');

SELECT hset('user:42', 'name', 'Alice');
SELECT hget('user:42', 'name');
```

`spset` accepts `text` and integer values directly, plus a generic `anyelement` path that stores PostgreSQL's textual representation. It returns the `spvalue` output type; cast retrieved text again when reconstructing values such as JSONB.

### Main Objects

- `spget`, `sptype`, and `del` read, inspect, and remove a key.
- `sadd`, `sismember`, `scard`, `srem`, and `sinter` operate on sets.
- `lpush`, `rpush`, `lpop`, `rpop`, and `llen` operate on lists.
- `hset` and `hget` operate on hash fields.
- `getexpireat` and `ttl` inspect expiration; `sp_db_nitems`, `sp_db_size_bytes`, and `sp_db_size` inspect the current namespace.
- `spat.db` selects a named shared-memory namespace for the session.

```sql
SET spat.db = 'analytics-cache';
SELECT spat_db_name(), spat_db_created_at(), sp_db_nitems();
SET spat.db = 'spat-default';
```

### ACID and Version Boundary

Changes take effect immediately, survive transaction rollback, and are visible to other sessions without MVCC isolation. They are not WAL-logged and all data disappears on PostgreSQL restart or crash. Per-key locks protect in-memory structures, but they do not add PostgreSQL transaction semantics.

The upstream README labels the project alpha, not production-ready, and requires PostgreSQL `17+` for the dynamic shared-memory registry. The catalog target is `0.1.0a4`, while the reviewed upstream control and SQL now declare `0.1.0a5`; verify the packaged a4 function signatures before using examples based on a5.
