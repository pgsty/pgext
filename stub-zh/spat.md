## 用法

来源：

- [官方 README](https://github.com/Florents-Tselai/spat/blob/bf7d47b6f6bf0b65bd0d0704a08edadea60942b9/README.md)
- [官方扩展 SQL](https://github.com/Florents-Tselai/spat/blob/bf7d47b6f6bf0b65bd0d0704a08edadea60942b9/sql/spat--0.1.0a5.sql)
- [官方扩展控制文件](https://github.com/Florents-Tselai/spat/blob/bf7d47b6f6bf0b65bd0d0704a08edadea60942b9/spat.control)

`spat` 是嵌入 PostgreSQL 动态共享内存的 alpha 阶段类 Redis 键值存储。它为缓存或跨会话临时状态提供字符串、集合、列表、哈希、TTL 与命名内存命名空间；不能替代事务性持久表。

### 核心流程

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

`spset` 可直接接受 `text` 和整数值，还提供通用 `anyelement` 路径，存储 PostgreSQL 的文本表示。它返回 `spvalue` 输出类型；重建 JSONB 等值时，需要再次转换取回的文本。

### 主要对象

- `spget`、`sptype` 与 `del` 分别读取、检查和删除键。
- `sadd`、`sismember`、`scard`、`srem` 与 `sinter` 操作集合。
- `lpush`、`rpush`、`lpop`、`rpop` 与 `llen` 操作列表。
- `hset` 与 `hget` 操作哈希字段。
- `getexpireat` 与 `ttl` 检查过期时间；`sp_db_nitems`、`sp_db_size_bytes` 与 `sp_db_size` 检查当前命名空间。
- `spat.db` 为当前会话选择一个命名共享内存命名空间。

```sql
SET spat.db = 'analytics-cache';
SELECT spat_db_name(), spat_db_created_at(), sp_db_nitems();
SET spat.db = 'spat-default';
```

### ACID 与版本边界

变更立即生效，不受事务回滚影响，并且没有 MVCC 隔离就能被其他会话看到。数据不写 WAL，PostgreSQL 重启或崩溃后全部消失。按键锁能保护内存结构，但不会增加 PostgreSQL 事务语义。

上游 README 把项目标记为 alpha、不可用于生产，并且因动态共享内存注册表要求 PostgreSQL `17+`。目录目标版本为 `0.1.0a4`，而本次复核的上游 control 与 SQL 已声明 `0.1.0a5`；使用基于 a5 的示例前，应确认打包的 a4 函数签名。
