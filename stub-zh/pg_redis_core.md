## 用法

来源：

- [Official source](https://github.com/MegacoderKim/pg2redis/blob/v0.6.0/README.md)
- [Official source](https://github.com/MegacoderKim/pg2redis/releases/tag/v0.5.0)
- [Official source](https://github.com/MegacoderKim/pg2redis/releases/tag/v0.6.0)
- [Official source](https://github.com/MegacoderKim/pg2redis/blob/v0.6.0/sql/pg_redis_core--0.5.0--0.6.0.sql)

- [README 0.4.1](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/README.md)
- [Configuration and SQL API](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/docs/getting-started/configuration.md)
- [Control](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/pg_redis_core.control)
- [Worker initialization](https://github.com/MegacoderKim/pg2redis/blob/v0.4.1/src/lib.rs)
- [Release with generated SQL](https://github.com/MegacoderKim/pg2redis/releases/tag/v0.4.1)

`pg_redis_core` 0.6.0 随 pg2redis 项目分发，在 PostgreSQL 表上实现兼容 Redis 的命令层。已核对的发布提供 PostgreSQL 15、16 制品。SQL 接口无需后台监听器；通过 RESP 访问则需要预加载后台进程。

### SQL 流程

安装匹配的库，并以超级用户在目标数据库中创建扩展：

```sql
CREATE EXTENSION pg_redis_core;
SELECT redis_dispatcher('SET', 'hello', 'world', NULL);
SELECT redis_dispatcher('GET', 'hello', NULL, NULL);
SELECT key, convert_from(value, 'UTF8'), expires_at
FROM pg_redis_strings;
```

`redis_dispatcher` 返回按 RESP2 编码的文本。存储表包括 `pg_redis_strings`、`pg_redis_hashes`、`pg_redis_lists` 和 `pg_redis_zsets`。

### RESP 监听器

如需网络客户端访问，将库加入已有预加载配置并重启 PostgreSQL。此示例将监听器限制在本机：

```ini
shared_preload_libraries = 'pg_redis_core'
pg_redis.bind_address = '127.0.0.1'
pg_redis.port = 6379
pg_redis.database = 'postgres'
pg_redis.unlogged = off
```

扩展应创建在 `pg_redis.database` 指定的数据库中。上游默认监听地址是 `0.0.0.0`，启用前应检查网络访问范围。

```bash
redis-cli -h 127.0.0.1 -p 6379 SET hello world
redis-cli -h 127.0.0.1 -p 6379 GET hello
```

### 管理与限制

`pg_redis_evict_expired()` 执行一次有界的过期清理。`pg_redis_ensure_schema(false)` 确保存储对象存在。`pg_redis_set_unlogged(true)` 会使存储失去常规 WAL 持久性；需要崩溃恢复时应保留记录 WAL 的存储方式。

`pg_redis.cache_enabled`、`pg_redis.cache_ttl_ms`、`pg_redis.eviction_interval_ms` 控制缓存和过期清理。监听地址、端口及目标数据库是启动参数。

项目实现了特定的 Redis 命令子集，包括字符串、哈希、列表、有序集合、流和发布订阅；不能假定完全兼容 Redis。后台进程通过自己的 PostgreSQL 会话执行请求，因此 RESP 访问与 SQL 登录是不同的访问边界。此早期版本需要结合应用实际使用的客户端命令及持久性要求进行评估。

### 0.4 的认证与资源限制

`pg_redis.requirepass` 通过 `AUTH` 和 `HELLO` 启用 RESP 密码检查，空值表示禁用认证。密码应保存在受保护的服务器配置中。共享 RESP 凭据不会为每个客户端创建独立的 PostgreSQL 角色。

`pg_redis.lua_time_limit_ms` 限制沙箱 Lua 调用时间，`pg_redis.max_clients`、`pg_redis.max_frame_bytes` 和 `pg_redis.pubsub_queue_size` 分别限制连接、请求帧和订阅队列。过期管理通过 `pg_redis_key_expiry` 覆盖全部六种受支持数据类型。已有实例部署匹配文件后，必须执行随版本提供的 SQL 升级链。

### 0.6.0 的 TLS 与逻辑数据库

TLS 需显式启用。`pg_redis.tls_port` 默认为 0；设置该端口及 `pg_redis.tls_cert_file`、`pg_redis.tls_key_file` 后启用 TLS。TLS 监听不会关闭明文连接：仅允许 TLS 时须设置 `pg_redis.port = 0`。可选客户端证书使用 `pg_redis.tls_ca_cert_file` 和 `pg_redis.tls_auth_clients`（`no`、`optional`、`yes`）。证书通过 SIGHUP 重载，失败时保留可用配置。

`pg_redis.databases` 默认为 16，范围为 1–1024，可通过 SIGHUP 重载。SELECT 选择逻辑 Redis 数据库。键、过期时间及阻塞操作按数据库隔离，Pub/Sub 仍为全局。FLUSHDB 只清空所选数据库，FLUSHALL 则影响全部数据库。

安装匹配动态库和 SQL 文件，重启预加载工作进程，并执行 `ALTER EXTENSION pg_redis_core UPDATE TO '0.6.0';`。提供的 0.4.1→0.5.0→0.6.0 升级路径新增数据库列与复合键，已有记录保留在数据库 0。迁移前备份并测试现有客户端。
