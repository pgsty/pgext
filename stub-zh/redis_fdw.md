

## 用法

> [redis_fdw: 查询 Redis 服务器的外部数据包装器](https://github.com/pg-redis-fdw/redis_fdw)

### 创建服务器

```sql
CREATE EXTENSION redis_fdw;

CREATE SERVER redis_server FOREIGN DATA WRAPPER redis_fdw
  OPTIONS (address '127.0.0.1', port '6379');
```

**服务器选项：** `address`（默认 `127.0.0.1`）、`port`（默认 `6379`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR pguser SERVER redis_server
  OPTIONS (password 'secret');
```

### 标量键值对

```sql
CREATE FOREIGN TABLE redis_db0 (
  key text,
  val text
)
SERVER redis_server
OPTIONS (database '0');

SELECT * FROM redis_db0;
```

### 哈希表（带键前缀）

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

### 哈希表（单键）

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

### 表选项

| 选项 | 描述 |
|--------|-------------|
| `database` | Redis 数据库编号（默认 `0`） |
| `tabletype` | `hash`、`list`、`set` 或 `zset`（省略则为标量键值） |
| `tablekeyprefix` | 按键前缀过滤项目 |
| `tablekeyset` | 从特定 Redis 集合中获取键 |
| `singleton_key` | 从单个 Redis 键访问所有值 |

只能使用 `tablekeyset` 或 `tablekeyprefix` 其中之一。不要与 `singleton_key` 组合使用。

哈希值以交替键值对的形式在 `text[]` 数组中返回。列表、集合和有序集合也以数组形式返回值。
