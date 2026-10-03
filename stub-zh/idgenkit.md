## 用法

来源：

- [README.md](https://github.com/rahulbsw/idgenkit/blob/37126f591680b741305b54463d87336537e93d05/README.md)
- [postgres/idgenkit.control](https://github.com/rahulbsw/idgenkit/blob/37126f591680b741305b54463d87336537e93d05/postgres/idgenkit.control)
- [postgres/idgenkit--0.1.0.sql](https://github.com/rahulbsw/idgenkit/blob/37126f591680b741305b54463d87336537e93d05/postgres/idgenkit--0.1.0.sql)

`idgenkit` 0.1.0 在 PostgreSQL 14–18 上生成多种 ID。基本随机生成器无需预加载；相对 ID 与部分 Snowflake 配置有额外要求。

### 核心工作流

```sql
CREATE EXTENSION idgenkit;
SELECT ulid_generate(), uuidv7_generate(), nanoid_generate();
CREATE TABLE events (id uuid PRIMARY KEY DEFAULT ulid_generate_uuid(), body text);
```

### 函数

`ulid_generate_monotonic()` 与 `uuidv7_generate_monotonic()` 仅保证单个后端内递增，不保证全局递增。`ulid_to_uuid`、`ulid_from_uuid`、`ulid_timestamp` 与 `uuidv7_timestamp` 用于转换或解析 ID。`snowflake_generate`、`snowflake_timestamp`、`snowflake_machine_id`、`snowflake_sequence` 与 `snowflake_from_timestamp` 操作 64 位 ID。`relid_generate`、`relid_generate_monotonic`、`relid_tag` 与 `relid_timestamp` 提供带密钥标签与时间戳解析。

### 配置与边界

控制文件将扩展标记为可信且可重定位。带密钥的相对 ID 要求把 `idgenkit` 加入 `shared_preload_libraries`，重启后由超级用户配置至少 16 字节的 `idgenkit.relid_secret`。密钥应妥善保存且不进入版本控制；获准执行标签函数的角色能够检验猜测的键。30 位标签可能碰撞，不能用作授权边界。

Snowflake 使用共享状态。PostgreSQL 14–16 需要预加载并重启；17+ 可在首次使用时初始化，无需预加载。不同服务器应分配不同的 `idgenkit.machine_id`，并一致配置 `idgenkit.snowflake_epoch_ms`。文本 ULID 使用 `COLLATE "C"` 才能按时间排序。ID 会暴露时间信息，不能作为身份验证密钥。
