## 用法

来源：

- [当前 database.dev 软件包文档](https://database.dev/jacobfvanzyl/pgkv)

`pgkv` 0.1.0 是通过 `jacobfvanzyl@pgkv` 分发的 Redis 风格 PostgreSQL 可信语言扩展。它使用 `PL/pgSQL` 实现，将带类型的值以 `JSONB` 形式存储在 `pgkv.store` 中，无需运行 Redis 服务器即可提供事务性键值命令。

### 核心工作流

```sql
CREATE EXTENSION pgkv;

SELECT pgkv.set('user:42', '{"name":"Ada","active":true}');
SELECT pgkv.get('user:42')->>'name';

SELECT pgkv.set('session:42', 'active', 60);
SELECT pgkv.ttl('session:42');

SELECT pgkv.incr('jobs:completed');
```

需要 PostgreSQL 12 或更高版本以及 `PL/pgSQL`。键会被标记为扩展定义的字符串、哈希、列表、集合或有序集合类型。对错误类型调用命令会抛出 `WRONGTYPE`；数值计数器以 JSONB 数字存储。

### 命令分组

- 基本键：`pgkv.set`、`pgkv.get`、`pgkv.del`、`pgkv.exists`、`pgkv.type` 和 `pgkv.dbsize`。
- 过期：`pgkv.expire`、`pgkv.ttl` 和 `pgkv.cleanup_expired`。
- 字符串与计数器：`pgkv.incr`、`pgkv.decr`、`pgkv.incrby`、`pgkv.decrby`、`pgkv.append`、`pgkv.strlen`、`pgkv.getrange` 和 `pgkv.setrange`。
- 哈希：`pgkv.hset`、`pgkv.hget`、`pgkv.hmget`、`pgkv.hgetall`、`pgkv.hdel`、`pgkv.hkeys` 和 `pgkv.hvals`。
- 列表：`pgkv.lpush`、`pgkv.rpush`、`pgkv.lpop`、`pgkv.rpop`、`pgkv.lrange`、`pgkv.lindex`、`pgkv.ltrim` 和 `pgkv.lrem`。
- 集合：`pgkv.sadd`、`pgkv.srem`、`pgkv.smembers`、`pgkv.sinter`、`pgkv.sunion` 和 `pgkv.sdiff`。
- 有序集合：`pgkv.zadd`、`pgkv.zrem`、`pgkv.zrange`、`pgkv.zrevrange`、`pgkv.zscore`、`pgkv.zrank` 和 `pgkv.zrangebyscore`。
- 多键与检查操作包括 `pgkv.mset`、`pgkv.mget` 和 `pgkv.keys`。

### 过期、性能与权限

过期是惰性的：读取操作检测到过期键后才删除它。若需在键被访问之前回收过期行，应调度 `pgkv.cleanup_expired()`。`pgkv.keys()` 接受 Redis 风格的 `*`、`?` 和 `[...]` 模式，但可能扫描整个共享存储；字符类使用正则表达式，更简单的模式使用 `LIKE`。

所有值和集合操作都由 PostgreSQL 执行。大型集合交集与差集、有序集合排序、争用、WAL、vacuum 和表增长因此可能占用大量数据库资源。应使用真实键与集合规模测试，不能假设它具有 Redis 的延迟或内存行为。

该存储在数据库内共享，软件包不提供应用或租户授权层。应只授予每个角色所需命令的 `EXECUTE`。应将 `pgkv.flushall()` 和 `pgkv.cleanup_expired()` 保留为管理操作，并严格控制删除、模式扫描和多键变更。database.dev 仅显示 0.1.0 一个已发布版本；将其作为持久化契约前，应审查已安装 SQL，并测试转储、恢复、权限、并发与未来升级。
