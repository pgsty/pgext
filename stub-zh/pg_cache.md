## 用法

来源：

- [上游来源 1](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/README.md)
- [上游来源 2](https://github.com/justusepostgres/pg_flaro/blob/d2cc08a462e46f3072c2fdc148ebb5b80ab3b0c2/crates/pg-cache/pg_cache.control)

`pg_cache` 0.1.0 是 PostgreSQL 17 实验性 Flaro RESP2 缓存的 SQL 扩展名。

### 核心流程

加入预加载项并重启，在配置指定的数据库中以超级用户创建扩展。后台工作进程提供受支持的 RESP2 命令，并将批次持久化到 `cache_store`。`pg_cache.listen_address` 控制监听地址，变更后需要重启。

```ini
shared_preload_libraries = 'pg_cache'
pg_cache.database = 'postgres'
pg_cache.port = 6380
pg_cache.memory_mb = 64
pg_cache.flush_ms = 100
```

```sql
CREATE EXTENSION pg_cache;
```

```bash
redis-cli -p 6380 SET greeting hello
redis-cli -p 6380 GET greeting
```

### 运行边界

活动缓存位于内存。RESP 成功响应可能早于 PostgreSQL 持久化，因此崩溃可能丢失已确认写入。队列和内存限制可能产生背压。这是有限的协议实现，不提供 RESP 身份认证，应限制在受信任网络，并将持久化进度与命令确认分别检查。
