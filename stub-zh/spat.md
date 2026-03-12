

## 用法

> [spat: 嵌入 Postgres 的类 Redis 内存数据库](https://github.com/Florents-Tselai/spat)

嵌入在 PostgreSQL 共享内存中的内存键值数据结构服务器。键为字符串；值可以是字符串、列表、集合或哈希。

### 启用

```sql
CREATE EXTENSION spat;
```

### 字符串

```sql
SELECT SPSET('key', 'value');
SELECT SPGET('key');              -- 'value'

-- 带 TTL
SELECT SPSET('temp', 'data', ttl => interval '5 minutes');

-- 将任意类型存储为文本
SELECT SPSET('config', '{"a": 1}'::jsonb);
SELECT SPGET('config')::text::jsonb;
```

### 集合

```sql
SELECT SADD('myset', 'elem1');
SELECT SADD('myset', 'elem2');
SELECT SISMEMBER('myset', 'elem1');  -- true
SELECT SCARD('myset');               -- 2
SELECT SREM('myset', 'elem1');       -- 1
```

### 列表

```sql
SELECT LPUSH('mylist', 'a');
SELECT LPUSH('mylist', 'b');
SELECT LPOP('mylist');     -- 'b'（后进先出）
SELECT LLEN('mylist');     -- 1
```

### 哈希

```sql
SELECT HSET('myhash', 'field1', 'Hello');
SELECT HGET('myhash', 'field1');  -- 'Hello'
```

### 通用操作

```sql
SELECT SPTYPE('key');           -- 'string'、'list'、'set' 或 'hash'
SELECT DEL('key');              -- 如果删除则为 true
SELECT TTL('key');              -- 返回 TTL 间隔
SELECT GETEXPIREAT('key');      -- 返回过期时间戳
SELECT SP_DB_NITEMS();          -- 条目数
SELECT SP_DB_SIZE();            -- 友好的大小显示
```

### 多数据库

```sql
SET spat.db = 'db1';             -- 切换到数据库 'db1'
SET spat.db = 'spat-default';   -- 切换回默认
```

### 重要说明

- 数据存储在 PostgreSQL 共享内存中，**非持久化** —— 重启后丢失
- 操作**非事务性** —— ROLLBACK 不会撤销 spat 变更
- 变更**立即可见**于所有会话（无 MVCC 隔离）
- 按键锁确保并发写入安全
