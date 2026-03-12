

## Usage

> [spat: Redis-like In-Memory DB Embedded in Postgres](https://github.com/Florents-Tselai/spat)

An in-memory key-value data structure server embedded in PostgreSQL shared memory. Keys are strings; values can be strings, lists, sets, or hashes.

### Enabling

```sql
CREATE EXTENSION spat;
```

### Strings

```sql
SELECT SPSET('key', 'value');
SELECT SPGET('key');              -- 'value'

-- With TTL
SELECT SPSET('temp', 'data', ttl => interval '5 minutes');

-- Store any type as text
SELECT SPSET('config', '{"a": 1}'::jsonb);
SELECT SPGET('config')::text::jsonb;
```

### Sets

```sql
SELECT SADD('myset', 'elem1');
SELECT SADD('myset', 'elem2');
SELECT SISMEMBER('myset', 'elem1');  -- true
SELECT SCARD('myset');               -- 2
SELECT SREM('myset', 'elem1');       -- 1
```

### Lists

```sql
SELECT LPUSH('mylist', 'a');
SELECT LPUSH('mylist', 'b');
SELECT LPOP('mylist');     -- 'b' (LIFO)
SELECT LLEN('mylist');     -- 1
```

### Hashes

```sql
SELECT HSET('myhash', 'field1', 'Hello');
SELECT HGET('myhash', 'field1');  -- 'Hello'
```

### Generic Operations

```sql
SELECT SPTYPE('key');           -- 'string', 'list', 'set', or 'hash'
SELECT DEL('key');              -- true if removed
SELECT TTL('key');              -- returns TTL interval
SELECT GETEXPIREAT('key');      -- returns expiration timestamp
SELECT SP_DB_NITEMS();          -- number of entries
SELECT SP_DB_SIZE();            -- human-friendly size
```

### Multiple Databases

```sql
SET spat.db = 'db1';             -- switch to database 'db1'
SET spat.db = 'spat-default';   -- switch back to default
```

### Important Notes

- Data is stored in PostgreSQL shared memory and is **not durable** -- lost on restart
- Operations are **not transactional** -- ROLLBACK does not undo spat changes
- Changes are **immediately visible** across all sessions (no MVCC isolation)
- Per-key locks ensure concurrent write safety
