## 用法

来源：

- [ColdFront v1.0.0-beta1 README](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/README.md)
- [ColdFront v1.0.0-beta1 发行版](https://github.com/pgEdge/coldfront/releases/tag/v1.0.0-beta1)
- [ColdFront 1.0 控制文件](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/extension/coldfront/coldfront.control)
- [ColdFront 1.0 扩展 SQL](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/extension/coldfront/coldfront--1.0.sql)
- [ColdFront 使用指南](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/docs/usage.md)
- [ColdFront 架构](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/docs/architecture.md)
- [ColdFront 分层模式架构](https://github.com/pgEdge/coldfront/blob/v1.0.0-beta1/docs/architecture_tiered.md)

`coldfront` 将 PostgreSQL 热分区和可写的 Apache Iceberg 冷数据透明地呈现为同一个 SQL 关系。它适合评估基于时间的冷热分层，或由 PostgreSQL 视图封装的纯 Iceberg 表。v1.0.0-beta1 是公开测试版：上游明确要求不要用于生产，因为接口、磁盘格式、行为和数据安全性仍可能发生变化。

### 配置运行时

ColdFront 依赖 `pg_duckdb` 在进程内执行 Iceberg I/O。两个库都必须按文档顺序预加载；修改此设置后需要重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_duckdb,coldfront'
coldfront.warehouse = 'wh'
coldfront.lakekeeper_endpoint = 'http://lakekeeper:8181/catalog'
```

v1.0.0-beta1 使用的并非任意原生 DuckDB 组合：其发行文档固定采用 PR 1025 中集成 DuckDB 1.5.3 的 `pg_duckdb`，以及打过补丁的 `duckdb-iceberg`。请使用该发行版匹配的构建或镜像，确保 Iceberg 提交和严格读取器互操作具备预期补丁。

重启后，在每个数据库中按相同依赖顺序创建扩展，然后保存对象存储凭据：

```sql
CREATE EXTENSION pg_duckdb;
CREATE EXTENSION coldfront;

SELECT coldfront.set_storage_secret(
  'access-key',
  'secret-key',
  'minio.example.com:9000',
  'us-east-1',
  'path',
  true
);
```

Lakekeeper 和对象存储是外部服务。使用托管关系前，应先初始化 Iceberg REST 目录与仓库。对于云端 AWS S3，请将端点传为 `NULL`，并传入存储桶的实际区域，让 DuckDB 使用虚拟主机式 HTTPS 寻址。不要把生产凭据写入迁移文件或示例。

### 运行分层表

分层模式从原生的 PostgreSQL 范围分区表开始。其主键必须覆盖分区键，归档器才能安全捕获并发变更：

```sql
CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY,
  ts timestamptz NOT NULL,
  status text,
  payload jsonb,
  PRIMARY KEY (id, ts)
) PARTITION BY RANGE (ts);
```

使用独立的 Go `archiver` 可执行程序注册并协调生命周期：

```bash
archiver register --config /etc/coldfront/config.yaml \
  --table events --period monthly \
  --hot-period "1 month" --retention "5 years"
archiver --config /etc/coldfront/config.yaml
```

这些是操作系统命令，不是扩展 SQL API。应通过 cron 或服务定时器运行 `archiver`，并对失败告警。它会预建分区，把原表转换为类似 `_events` 的热堆和透明的 `events` 视图，将早于 `hot_period` 的分区移动到 Iceberg，推进 `coldfront.archive_watermark`，并永久过期早于 `retention_period` 的数据。`retention_period` 必须大于 `hot_period`。

应用继续使用公开关系：

```sql
SELECT id, ts, status FROM events
WHERE ts >= '2026-07-01'
ORDER BY ts DESC;

SET coldfront.allow_mixed_writes = off;
UPDATE events
SET status = 'fixed'
WHERE ts >= '2026-07-01' AND id = 42;
```

设置 `coldfront.allow_mixed_writes = off` 后，无法明确归属某一层的 `UPDATE` 或 `DELETE` 会失败，而不是同时写入两层。当应用能够提供分区键谓词时，这样更安全。默认模式允许跨层写入，而双层提交并不具备崩溃安全性。

### 创建纯 Iceberg 表

解耦模式将所有行都存储在 Iceberg 中，不使用 `archiver` 或 PostgreSQL 热堆。先在 Lakekeeper 中预建 Iceberg 命名空间，再由扩展创建外部表、包装视图和注册表行：

```sql
SELECT coldfront.create_iceberg_table(
  p_schema  => 'public',
  p_table   => 'events_archive',
  p_columns => '[
    {"name":"id", "type":"bigint"},
    {"name":"ts", "type":"timestamptz"},
    {"name":"status", "type":"text"},
    {"name":"payload", "type":"jsonb"}
  ]'::jsonb
);

INSERT INTO events_archive VALUES
  (1, now(), 'new', '{"source":"demo"}');
SELECT * FROM events_archive;
```

`SELECT`、`INSERT`、`UPDATE` 和 `DELETE` 都以包装关系为目标；钩子会把数据修改语句重写到 Iceberg。对于这张表，PostgreSQL 只是 SQL 与计算前端，因此这些行不会进入其堆存储或普通 PostgreSQL 备份。

### 重要对象

- `coldfront.set_storage_secret(...)` 记录 S3 或 S3 兼容凭据，并生成持久化 DuckDB secret。`coldfront.set_storage_secret_azure(...)` 是 Azure ADLS Gen2 对应接口。
- `coldfront.create_iceberg_table(...)` 创建纯 Iceberg 表及其 PostgreSQL 包装层；它不是分层表的注册路径。
- `coldfront.grant_app_access(regrole)` 根据注册表为非超级用户授予托管视图运行时权限；它有意不授予管理函数或服务器文件角色。
- `coldfront.tiered_views` 注册透明关系，`coldfront.archive_watermark` 记录冷热分界，`coldfront.partition_config` 保存每张表的生命周期策略。
- `coldfront.warehouse` 和 `coldfront.lakekeeper_endpoint` 选择外部目录。`coldfront.allow_mixed_writes` 控制含糊的跨层 DML，`coldfront.local_pg_dsn` 启用文档所述的 PostgreSQL 到 Iceberg 流式路径。
- `archiver`、`partitioner` 和 `compactor` 是独立的 Go 程序。`archiver` 是分层迁移所必需的；`partitioner` 可以在没有冷层时管理 PostgreSQL 分区；`compactor` 执行 Iceberg 维护。安装 `coldfront` 不会调度其中任何程序。

### DDL 与 DML 边界

- 分层应用所见的关系是视图，其原生热表通常为 `_events`。普通写入不要绕过视图，也不要假定直接修改热表会同时影响 Iceberg。
- DDL 钩子会镜像针对已注册热表发出的受支持子集：`ALTER TABLE _events ADD/DROP COLUMN`、安全的 `ALTER COLUMN ... TYPE` 类型提升、列重命名以及表/视图重命名。不支持的类型或转换会被拒绝；每次模式迁移都应在两层数据副本上进行测试。
- `DROP TABLE _events`、删除托管视图以及 `TRUNCATE _events` 都会被阻止，因为单边操作会遗留或暴露冷数据。应先取消生命周期注册，再有意地分别拆除 PostgreSQL 与 Iceberg 两侧。
- 涉及冷层的 DML 不支持 `RETURNING`。分层自连接、对同一托管视图执行 `DELETE ... USING`，以及在子查询中再次引用该视图都会被拒绝。
- 正常的 `ROLLBACK` 会协调 PostgreSQL 与 DuckDB 事务，但如果后端在 Iceberg 快照操作与 PostgreSQL 提交之间崩溃，对象存储中可能留下孤儿文件。严格的单层谓词可降低双层风险；Iceberg 快照过期和孤儿文件维护仍是必需的运维工作。
- 可重复读隔离不会跨越同一 PostgreSQL 长事务内的多次 Iceberg 扫描。虽然可以读到自己在事务中的写入，但多次扫描之间可能观察到外部并发提交。

### 备份与恢复

逻辑 `pg_dump` 会包含 `coldfront.tiered_views`、`coldfront.archive_watermark` 和 `coldfront.partition_config`，但不会复制 Iceberg 数据、Lakekeeper 目录数据库、`coldfront.storage_secret` 中的对象存储凭据或 Bakery 协议的临时 claim 表。恢复 PostgreSQL 后，应重新执行 `coldfront.set_storage_secret(...)`；恢复的元数据随后会指向原来的外部 Iceberg 表。

应把 PostgreSQL 热层、Lakekeeper 目录数据库和对象存储桶作为相互独立但需要协调的系统进行备份与恢复。在同一恢复点保留兼容的 Iceberg 元数据与对象；只恢复 PostgreSQL 并不等于恢复冷数据，删除或重建外部仓库会让已恢复的注册表无法使用。

分层保留路径最终会在导出后销毁数据。缩短 `retention_period` 前必须备份，应演练外部服务不可用时的恢复，并确保对象存储生命周期策略不会删除仍被 Iceberg 快照引用的文件。

### 版本与兼容性边界

- v1.0.0-beta1 发行版包含 SQL 扩展版本 `1.0`；控制文件固定使用模式 `coldfront`，不可重定位，并要求预加载与重启服务器。
- 该标签发行版记录支持原生 PostgreSQL 16、17 和 18。应使用发行版验证过的精确依赖矩阵，不要假定其他 `pg_duckdb`、DuckDB 或 `duckdb-iceberg` 构建兼容。
- Iceberg 支持的列包括常见整数、浮点数、布尔值、时态类型、UUID、文本、带精度与标度限制的 numeric、`bytea`、`json`、`jsonb` 和 `interval`。数组、枚举、复合类型、范围、无限定 numeric、`inet` 与 `cidr` 会被拒绝。
- `jsonb` 和 `json` 在 Iceberg 中按字符串存储，并通过托管视图呈现为 `json`；`?` 与 `@>` 等仅适用于 JSONB 的操作符需要显式转换回 `jsonb`。
- 公开测试版状态是最重要的边界：不要把评估成功、上游基准测试或可用源码构建当作生产就绪的证据。
