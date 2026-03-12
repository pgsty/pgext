


## Usage

> [redis_fdw: Foreign data wrapper for querying a Redis server](https://github.com/pg-redis-fdw/redis_fdw)

### Create Server

```sql
CREATE EXTENSION redis_fdw;

CREATE SERVER redis_server FOREIGN DATA WRAPPER redis_fdw
  OPTIONS (address '127.0.0.1', port '6379');
```

**Server Options:** `address` (default `127.0.0.1`), `port` (default `6379`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR pguser SERVER redis_server
  OPTIONS (password 'secret');
```

### Scalar Key-Value Pairs

```sql
CREATE FOREIGN TABLE redis_db0 (
  key text,
  val text
)
SERVER redis_server
OPTIONS (database '0');

SELECT * FROM redis_db0;
```

### Hash Tables (with Key Prefix)

```sql
CREATE FOREIGN TABLE redis_hash (
  key text,
  val text[]
)
SERVER redis_server
OPTIONS (database '0', tabletype 'hash', tablekeyprefix 'mytable:');

INSERT INTO redis_hash VALUES ('mytable:r1', '{prop1,val1,prop2,val2}');
UPDATE redis_hash SET val = '{prop3,val3}' WHERE key = 'mytable:r1';
DELETE FROM redis_hash WHERE key = 'mytable:r1';
SELECT * FROM redis_hash;
```

### Hash Tables (Singleton Key)

```sql
CREATE FOREIGN TABLE redis_singleton (
  key text,
  val text
)
SERVER redis_server
OPTIONS (database '0', tabletype 'hash', singleton_key 'myhash');

INSERT INTO redis_singleton VALUES ('field1', 'value1');
UPDATE redis_singleton SET val = 'newvalue' WHERE key = 'field1';
DELETE FROM redis_singleton WHERE key = 'field1';
```

### Table Options

| Option | Description |
|--------|-------------|
| `database` | Redis database number (default `0`) |
| `tabletype` | `hash`, `list`, `set`, or `zset` (omit for scalar key-value) |
| `tablekeyprefix` | Filter items by key prefix |
| `tablekeyset` | Fetch keys from a specific Redis set |
| `singleton_key` | Access all values from a single Redis key |

Use only one of `tablekeyset` or `tablekeyprefix`. Do not combine them with `singleton_key`.

Hash values are returned as alternating key-value pairs in a `text[]` array. Lists, sets, and sorted sets also return values as arrays.
