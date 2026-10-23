## 用法

来源：

- [官方资料](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/README.md)
- [官方资料](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/pg_kv.control)
- [官方资料](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/Cargo.toml)
- [官方资料](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/src/lib.rs)

`pg_kv` 0.1.0 提供共享内存缓存、计数器、令牌桶限流与有序集合。该早期版本面向 PostgreSQL 15–18，尚不适合生产使用。

### 核心流程

将库加入现有预加载列表并重启 PostgreSQL，然后以超级用户安装扩展：

```ini
shared_preload_libraries = 'pg_kv'
```

```sql
CREATE EXTENSION pg_kv;
SELECT kv.set('app:session:42', '{"user":1}', interval '30 minutes');
SELECT kv.get('app:session:42');
SELECT kv.ttl('app:session:42');
SELECT kv.incr('app:views:7');
SELECT kv.rate_limit('app:api:42', 100, interval '1 minute');
SELECT kv.zadd('app:leaderboard', 'player:9', 1520);
SELECT * FROM kv.zrange('app:leaderboard', 0, 9, reverse => true);
SELECT * FROM kv.stats();
```

没有预加载库时，安装会失败。仅向预期调用者授予 `kv` 模式的使用权限。集群中的所有数据库共享同一个键空间；前缀便于组织键，但不能提供租户隔离。

### 操作与容量

`kv.ttl` 返回毫秒数，-1 表示永不过期，-2 表示键不存在。`kv.expire` 修改有效期，`kv.del` 删除键，`kv.rate_limit_info` 返回是否放行及重试时间。有序集合还提供 `kv.zscore`、`kv.zrank`、`kv.zrevrank`、`kv.zincrby`、`kv.zcard` 和 `kv.zrem`。

需要重启的参数包括 `pg_kv.max_entries`、`pg_kv.max_key_bytes`、`pg_kv.max_value_bytes`、`pg_kv.shards`、`pg_kv.sweep_interval_ms`、`pg_kv.max_zset_members` 和 `pg_kv.max_member_bytes`。存储容量固定，超大值会报错，容量压力可能触发缓存淘汰。`kv.flush` 是管理操作，默认不向公共角色开放。

### 持久性边界

缓存修改不会随事务回滚撤销，但服务器重启或崩溃后会丢失。它不写入 WAL，也不复制；备库拥有独立缓存。适合可丢弃状态，不适合作为权威数据，或跨独立服务器协调的全局限流器。
