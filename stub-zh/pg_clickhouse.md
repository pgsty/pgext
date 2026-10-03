## 用法

来源：

- [pg_clickhouse v0.11.0 README](https://github.com/ClickHouse/pg_clickhouse/blob/v0.11.0/README.md)
- [pg_clickhouse v0.11.0 参考文档](https://github.com/ClickHouse/pg_clickhouse/blob/v0.11.0/doc/pg_clickhouse.md)
- [pg_clickhouse v0.11.0 教程](https://github.com/ClickHouse/pg_clickhouse/blob/v0.11.0/doc/tutorial.md)
- [pg_clickhouse v0.11.0 变更日志](https://github.com/ClickHouse/pg_clickhouse/blob/v0.11.0/CHANGELOG.md)
- [pg_clickhouse v0.11.0 控制文件](https://github.com/ClickHouse/pg_clickhouse/blob/v0.11.0/pg_clickhouse.control)
- [0.10 to 0.11 upgrade SQL](https://github.com/ClickHouse/pg_clickhouse/blob/v0.11.0/sql/pg_clickhouse--0.10--0.11.sql)

`pg_clickhouse` 0.11.0 通过 `clickhouse_fdw` 外部数据封装器把 ClickHouse 表暴露给 PostgreSQL。上游面向 PostgreSQL 14 及以上版本与 ClickHouse 23.3 及以上版本；软件包可用性另行记录。正常使用无需预加载；`session_preload_libraries` 与 `shared_preload_libraries` 只是可选的连接启动优化。

### 连接 PostgreSQL 与 ClickHouse

```sql
CREATE EXTENSION pg_clickhouse;

CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (
  driver 'binary',
  host 'localhost',
  dbname 'taxi',
  compression 'lz4'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

必填的 `driver` 选项可取 `binary` 或 `http`。常用服务器选项包括 `host`、`port`、`dbname`、`compression`、`secure` 与 `min_tls_version`；用户映射接受 `user` 和 `password`。0.11 版本删除了 `fetch_size`，两个驱动均流式处理 Native 格式。

`IMPORT FOREIGN SCHEMA` 支持 `LIMIT TO (...)` 与 `EXCEPT (...)`。导入的混合大小写标识符会保留引号，引用时必须使用匹配的引号。

### 查询与写入外部表

```sql
EXPLAIN (VERBOSE)
SELECT node_id, count(*)
FROM taxi.logs
GROUP BY node_id;

INSERT INTO taxi.nodes(node_id, name)
VALUES (9, 'west-node');

COPY taxi.nodes(node_id, name) FROM STDIN;
```

`SELECT`、`EXPLAIN`、预备语句、`INSERT` 与 `COPY` 都可作用于外部表。在 0.10 版本中，二进制驱动以有界的 64 MiB 批次刷新写入，因此 `COPY` 已不再只是展开成逐行语句。使用 `EXPLAIN (VERBOSE)` 检查远端 SQL，并确认哪些过滤、连接、聚合与函数发生了下推。

### 直接查询与命令 API

0.10 版本新增了带类型的任意查询与命令接口：

```sql
GRANT EXECUTE ON FUNCTION clickhouse_query(text, text) TO analyst;
GRANT EXECUTE ON PROCEDURE clickhouse_perform(text, text) TO operator;

SELECT *
FROM clickhouse_query(
  'taxi_srv',
  'SELECT region, count() FROM taxi GROUP BY region'
) AS t(region text, n bigint);

CALL clickhouse_perform(
  'taxi_srv',
  'OPTIMIZE TABLE taxi.nodes FINAL'
);

SELECT clickhouse_server_version('taxi_srv');
```

`clickhouse_query(server, sql)` 按调用方提供的列定义返回行，而 `clickhouse_perform(server, sql)` 会丢弃结果。两者都能执行任意远端 SQL，因此 `EXECUTE` 已从 `PUBLIC` 撤销，只应按最小范围授权。0.11 已删除 `clickhouse_raw_query()`，须迁移到这两个接口。

### 下推与会话设置

0.10 版本扩展了聚合与函数下推，改善了本地分区和外部分区混合场景下的聚合执行，并修复了多处 PostgreSQL NULL 语义差异。子查询下推要求 ClickHouse 25.8 或以上版本；旧服务器会在本地计算这些子查询。

默认的 `pg_clickhouse.session_settings` 保持与 PostgreSQL 兼容的行为，其中包括 `join_use_nulls = 1`、`group_by_use_nulls = 1`、`final = 1` 与 `transform_null_in = 0`。覆盖它时，应保留工作负载所需的设置，尤其是安全下推 `IN` 所必需的 `transform_null_in = 0`。

### 升级与运维边界

```sql
ALTER EXTENSION pg_clickhouse UPDATE TO '0.11';
SELECT pgch_version();
```

扩展 SQL 版本为 `0.11`，`pgch_version()` 返回库版本 `0.11.0`。安装匹配文件后须执行 `ALTER EXTENSION`，从 SQL 目录中移除废弃函数。

把 `pg_clickhouse` 放入 `session_preload_libraries` 时，新会话会自动加载它；放入 `shared_preload_libraries` 时，更换动态库需要重启 PostgreSQL。与需要注册 postmaster 钩子的扩展不同，这两个设置都不是强制要求。

文档化的写入接口仍不包括轻量级 `UPDATE` 与 `DELETE`。应把直接远端 SQL 视为特权操作，使用贴近生产的数据验证 NULL 与类型相关的下推，并在依赖受版本约束的优化前核对 PostgreSQL 和 ClickHouse 版本。

0.11 默认拒绝来自 ClickHouse 的无效数据库编码文本或 JSON。服务器选项 `encoding_check` 可取 `fail`、`replace`、`remove` 或 `truncate`；后三者会改变返回数据，应按业务语义明确选择。导入类型还改进了大整数、时间间隔以及数组、Tuple、Map、Nested 的映射；已有外部表定义不会自动随推断规则变化，应逐列核对。
