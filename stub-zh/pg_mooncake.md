## 用法

来源：

- [0.3.0 配置与教程](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/README.md)
- [0.3.0 SQL 过程](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/functions.rs)
- [依赖检查与 DuckDB 扩展加载](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/sql/bootstrap.sql)
- [扫描同步机制](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/duckdb_mooncake.rs)
- [0.3.0 验证与限制](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/docs/pg-mooncake-0.3.0-validation.md)

`pg_mooncake` 0.3.0 为 PostgreSQL 表创建只读列存镜像，用于分析查询。Moonlink 将已提交的行变更复制到 Arrow 缓冲区和 Parquet/Iceberg 存储，DuckDB 通过 `pg_duckdb` 执行扫描与聚合。此 PGSTY 版本面向 PostgreSQL 14–18，要求精确匹配的 `pg_duckdb` 1.2.0、DuckDB 2.0.0，以及配套打包的 DuckDB Mooncake 扩展。

### 配置与启用

由管理员配置逻辑复制，并按下列顺序将两个库加入已有预加载列表。扩展仓库参数应指向已安装的本地 DuckDB 扩展。以下示例采用 EL 上 PostgreSQL 17 的目录布局：

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
wal_level = logical
duckdb.allow_community_extensions = true
duckdb.allow_unsigned_extensions = true
mooncake.extension_repository = '/usr/pgsql-17/share/extension/pg_mooncake'
```

其他安装应使用对应 `pg_config --sharedir` 返回的目录，再接上 `/extension/pg_mooncake`。该 tag 在未设置此参数时会回退到社区仓库，不会自动选择打包目录。本地二进制要求允许加载未签名扩展。确保复制槽、WAL 发送进程和工作进程数量充足，然后重启 PostgreSQL。

```sql
CREATE EXTENSION pg_mooncake CASCADE;
```

创建扩展需要超级用户权限，并会检查已安装的依赖版本。创建镜像会打开独立数据库连接，因此源表 DDL 必须已经提交，连接角色还必须能够通过认证、读取源表并完成所需的复制与发布配置。下面的示例应在自动提交模式下运行。

### 写入源表，查询镜像

```sql
CREATE TABLE trades (
    id bigint PRIMARY KEY,
    symbol text,
    traded_at timestamp,
    price real
);
CALL mooncake.create_table('trades_analytics', 'trades');

INSERT INTO trades VALUES
    (1, 'AMZN', '2026-11-05 10:00:00', 207),
    (2, 'AMZN', '2026-11-05 10:05:00', 210),
    (3, 'AAPL', '2026-11-05 10:10:00', 203);
UPDATE trades SET price = 209 WHERE id = 1;
DELETE FROM trades WHERE id = 3;

SELECT symbol, avg(price) FROM trades_analytics GROUP BY symbol;
EXPLAIN SELECT avg(price) FROM trades_analytics WHERE symbol = 'AMZN';
```

辅助过程会复制现有行并启动逻辑复制。`INSERT`、`UPDATE` 和 `DELETE` 应发送到源表；镜像使用 `USING mooncake`，不能替代可写源表。已回滚的源表变更不会复制。查询可以通过支持的 DuckDB 执行计划组合镜像与普通行存表，应使用 `EXPLAIN` 检查是否采用 DuckDB 扫描。

该 tag 的扫描与快照同步使用当前后端的最后提交 LSN。不能据此认定其他所有会话的最新写入都已摄入，也不能假定远端源库与本地使用同一 WAL 位置。涉及跨会话或远端数据新鲜度要求时，应监控复制进度。

### 过程与维护

| 接口 | 用途 |
| --- | --- |
| `mooncake.create_table(dst, src, src_uri, table_config)` | 创建镜像；后两个参数默认为 NULL，远端 PostgreSQL 源连接的命名参数必须写作 `src_uri`。 |
| `mooncake.list_tables()` | 返回已登记的表、提交与刷盘 LSN，以及 Iceberg 仓库位置。 |
| `mooncake.create_snapshot(dst)` | 将缓冲数据按请求的同步位置刷入 Iceberg 快照。 |
| `mooncake.optimize_table(dst, mode)` | 请求压实，完整模式会压实数据与索引。 |
| `mooncake.load_files(dst, files)` | 0.3.0 尚未实现，会抛出 SQLSTATE `0A000`。 |

```sql
SELECT * FROM mooncake.list_tables();
CALL mooncake.create_snapshot('trades_analytics');
CALL mooncake.optimize_table('trades_analytics', 'full');
```

库内查询可以读取缓冲中的 Arrow 数据，外部 Iceberg 读取器只能看到已发布的快照。除镜像存储外，还应监控逻辑复制槽延迟及保留的 WAL。删除目标表会触发 Moonlink 清理，应检查登记记录和复制槽是否移除。在清理或迁移验证完成前，应保留源数据。

### 兼容性与升级边界

创建镜像时会处理无主键源表的复制标识，包括已有叶子分区。后续新增的分区在接收写入前也需要设置 `REPLICA IDENTITY FULL`。`TRUNCATE` 复制以及后续分区 DDL 的自动处理不在已支持或已验证范围内。不能假定镜像支持任意模式变更或 PostgreSQL 特有函数。

v0.3.0 tag 未提供从 0.2.0 原地升级的 SQL 路径。应在独立验证的新安装中，从保留的源表重建镜像。PostgreSQL 库与配套 DuckDB 扩展必须一起更新，为 DuckDB 1.4 构建的二进制不能用于 DuckDB 2.0。tag 中的验证覆盖本地存储和正常重启，这些结果不能证明 S3、崩溃恢复或压力场景的行为。
