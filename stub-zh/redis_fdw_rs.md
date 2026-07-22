## 用法

来源：

- [官方 README](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/README.md)
- [FDW 设置示例](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/examples/00_setup.sql)
- [选项与凭据加载](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/src/utils/helpers.rs)
- [FDW 选项验证器](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/src/core/validator.rs)
- [连接池实现](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/src/core/pool_manager.rs)
- [扩展控制文件](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/redis_fdw_rs.control)

`redis_fdw_rs` 0.6.0 是一个 Rust 外部数据包装器，可通过 PostgreSQL 外部表读写 Redis 字符串、哈希、列表、集合、有序集合与流。它支持 PostgreSQL 14 至 18、Redis Cluster、TLS、连接池、谓词下推、批量插入、TTL、模式导入以及部分连接下推。

### 创建包装器与表

安装扩展会创建处理器和验证器函数，但不会创建具名外部数据包装器或服务器。需要显式创建它们，再把每个表映射到一个 Redis 键或键模式：

```sql
CREATE EXTENSION redis_fdw_rs;

CREATE FOREIGN DATA WRAPPER redis_wrapper
HANDLER redis_fdw_handler
VALIDATOR redis_fdw_validator;

CREATE SERVER redis_server
FOREIGN DATA WRAPPER redis_wrapper
OPTIONS (host_port '127.0.0.1:6379');

CREATE FOREIGN TABLE user_profiles (
    field text,
    value text
)
SERVER redis_server
OPTIONS (
    database '0',
    table_type 'hash',
    table_key_prefix 'user:profiles'
);

INSERT INTO user_profiles VALUES
    ('name', 'Ada'),
    ('email', 'ada@example.com');

SELECT value FROM user_profiles WHERE field = 'email';
```

哈希谓词可以下推成 `HGET`。应使用 `EXPLAIN (VERBOSE)` 确认下推，而不是默认它会发生。列形状按位置解释：字符串有一个值列，哈希为字段/值，列表为元素或序号/元素，集合为成员，有序集合为成员/分数，流则为 ID 加字段/值列。

### Redis 类型与选项

- `table_type`：取 `string`、`hash`、`list`、`set`、`zset` 或 `stream` 之一；流不支持更新。
- `table_key_prefix`：精确键或 glob 模式。模式表增加一个首位键列，并可能需要 Redis `SCAN`。
- `database`：0 至 15 的 Redis 逻辑数据库；Redis Cluster 实际只使用数据库 0。
- `ttl`：以秒计的默认过期时间。可选 `ttl bigint` 列提供逐行值；`-1` 让键持久化。
- `batch_size` 和 `join_batch_size`：限制流水线插入及参数化连接的批量大小。
- `IMPORT FOREIGN SCHEMA`：抽样最多 10,000 个 Redis 键，按前缀分组并生成外部表定义。

FDW 为六种类型实现 `SELECT`、`INSERT`、`DELETE` 与大多数 `UPDATE`，但不支持流的 `UPDATE`。同服务器单列等值连接可有限下推；多键模式与流还有额外限制。设计写入路径前，应在官方矩阵中核对具体操作。

### TLS、凭据与连接池

在 `host_port` 中使用 `rediss://` 可要求验证 TLS 证书。文档中的 `#insecure` 后缀会关闭验证，只适用于隔离测试。

在复核的提交中，运行时会合并包装器、服务器与表选项，但读取 PostgreSQL 用户映射选项的代码已被注释。因此 `username` 与 `password` 必须作为服务器选项才能生效，具有目录读取权或能查看 DDL 的高权限用户可能看到它们。应限制所有权与可见性，优先使用权限受限的 Redis ACL 身份，并测试凭据轮换流程。每个后端的连接池缓存键能区分用户名或是否存在密码，却不包含密码值，所以修改密钥后仍可能复用旧连接池，直到连接或后端被回收。

### 一致性与破坏性操作

Redis 命令在远端立即生效，不会随 PostgreSQL 事务回滚。写入应设计为幂等，并预期语句错误、客户端取消、故障切换或重试后可能留下部分效果。

`TRUNCATE` 会解除链接精确的 Redis 键；用于模式表时则执行 `SCAN` 加批量 `UNLINK`，可能删除所有匹配键，而集群模式不支持多键 `TRUNCATE`。授予截断权限前，应复核 `table_key_prefix` 与 Redis ACL 范围。全量模式扫描、模式导入和下推的同服务器连接都可能大量消耗 Redis、网络或 PostgreSQL 内存，应使用生产规模的键基数测试并设置查询超时。
