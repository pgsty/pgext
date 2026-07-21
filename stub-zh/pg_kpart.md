## 用法

来源：

- [官方v1.0 README](https://github.com/HexaCluster/pg_kpart/blob/v1.0/README.md)
- [v1.0控制文件](https://github.com/HexaCluster/pg_kpart/blob/v1.0/pg_kpart.control)

`pg_kpart` 防止了对分表表进行全表扫描的意外查询，而这些查询没有有效的分区剪裁。其规划器钩子可以在执行前发出警告、告警或记录日志。功能单元是预加载的库；无需创建任何SQL对象，并且上游仅将 `CREATE EXTENSION` 描述为可选的系统目录注册。

### 启用与部署

为了集群范围内的强制执行，可以预加载库并重启PostgreSQL：

```conf
shared_preload_libraries = 'pg_kpart'
```

也可以在无需服务器重启的情况下，选择性地加载会话或数据库：

```conf
session_preload_libraries = 'pg_kpart'
```

在强制错误之前，以审计模式开始：

```sql
ALTER SYSTEM SET pg_kpart.message_level = 'warning';
SELECT pg_reload_conf();
```

一旦观察到的查询被理解，可以设置 `pg_kpart.message_level = 'error'`。

### 范围与行为

```sql
-- Check only these tables and their sub-partitions.
ALTER SYSTEM SET pg_kpart.blacklisted =
    'public.measurement, public.orders';

-- Or check all partitioned tables except selected hierarchies.
ALTER SYSTEM SET pg_kpart.whitelisted = 'public.audit_log';
SELECT pg_reload_conf();
```

```sql
-- Partition key is logdate.
SELECT * FROM measurement WHERE city_id = 5;              -- violation
SELECT * FROM measurement WHERE logdate = DATE '2026-07-01'; -- pruned, allowed
SELECT * FROM measurement WHERE logdate = $1;             -- runtime pruning, allowed
```

违反规则将使用SQLSTATE `FS001`，应用程序可以在 `message_level` 为 `error` 的情况下捕获这些错误。

### 配置索引与注意事项

- `pg_kpart.enabled`: 主控开关，默认值为 `on`。
- `pg_kpart.message_level`: 可以设置为 `error`, `warning`, `notice`, `log` 等其他PostgreSQL消息级别。
- `pg_kpart.min_partitions`: 检查的最小叶分区数量，默认值为 `2`。
- `pg_kpart.check_superuser`: 默认情况下，超级用户绕过检查。
- `pg_kpart.blacklisted`: 当非空时，仅检查命名层次结构，并忽略 `whitelisted`。
- `pg_kpart.whitelisted`: 在未设置黑名单的情况下，免于检查的层次结构。
- 如果谓词的范围仍然包括所有分区，则即使提到分区键也会被视为全表扫描并被拒绝。
- 规划器钩子也适用于 `UPDATE`, `DELETE` 和不带 `EXPLAIN` 的 `ANALYZE`。它依赖于PostgreSQL计划剪裁的结果，而不是对 `WHERE` 子句的文本检查。
- 上游v1.0在PostgreSQL 14及更高版本上进行了测试。
