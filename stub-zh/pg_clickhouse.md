## 用法

来源：[README](https://github.com/ClickHouse/pg_clickhouse/blob/main/README.md)、[reference](https://github.com/ClickHouse/pg_clickhouse/blob/main/doc/pg_clickhouse.md)、[tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/main/doc/tutorial.md)、[v0.3.2 release notes](https://github.com/ClickHouse/pg_clickhouse/releases/tag/v0.3.2)、[changelog](https://github.com/ClickHouse/pg_clickhouse/blob/main/CHANGELOG.md)

`pg_clickhouse` 通过 `clickhouse_fdw` foreign data wrapper，让 PostgreSQL 可以在 ClickHouse 上执行分析查询。上游文档说明支持 PostgreSQL 13+ 与 ClickHouse 23+；当前版本是 0.3.2。

### 连接 PostgreSQL 与 ClickHouse

```sql
CREATE EXTENSION pg_clickhouse;

CREATE SERVER taxi_srv
FOREIGN DATA WRAPPER clickhouse_fdw
OPTIONS (driver 'binary', host 'localhost', dbname 'taxi');

CREATE USER MAPPING FOR CURRENT_USER
SERVER taxi_srv
OPTIONS (user 'default');

CREATE SCHEMA taxi;
IMPORT FOREIGN SCHEMA taxi FROM SERVER taxi_srv INTO taxi;
```

上游文档列出的服务器选项：

- `driver`：必填，取值为 `binary` 或 `http`
- `host`
- `port`
- `dbname`
- `fetch_size`：HTTP streaming batch size；`0` 表示禁用 streaming
- `compression`：binary-driver compression，可取 `none`、`lz4` 或 `zstd`；v0.3.2 默认 `lz4`
- `secure`：显式 TLS 模式，可取 `on`、`off` 或 `auto`
- `min_tls_version`：binary 和 HTTP drivers 使用的最低 TLS protocol version

用户映射选项：

- `user`
- `password`

### 常见操作

```sql
ALTER EXTENSION pg_clickhouse UPDATE;
ALTER EXTENSION pg_clickhouse UPDATE TO '0.3';
SELECT pgch_version();
DROP SERVER taxi_srv CASCADE;
```

`IMPORT FOREIGN SCHEMA` 也支持 `LIMIT TO (...)` 和 `EXCEPT (...)`。reference 提醒：导入的 mixed-case 标识符会在 PostgreSQL 中加双引号，查询时也必须带引号。

### 查询与写入说明

`SELECT`、`EXPLAIN`、prepared statements、`INSERT` 和 `COPY` 都可作用于 `pg_clickhouse` foreign tables。使用 `EXPLAIN (VERBOSE)` 可以查看将发送到 ClickHouse 的远端 SQL。

```sql
EXPLAIN (VERBOSE)
SELECT node_id, count(*)
FROM logs
GROUP BY node_id;

INSERT INTO nodes(node_id, name, region, arch, os)
VALUES (9, 'west-node', 'us-west-2', 'amd64', 'Linux');
```

`COPY` 写入 foreign table 已有文档说明，但上游也指出当前仍通过 `INSERT` 语句实现，因为 FDW batch insertion 仍是未来工作。

### 版本与下推说明

- reference 区分 library version 和 extension version；`pgch_version()` 返回已加载的库版本。
- Release `v0.3.2` 对已有 SQL version `0.3` 而言只是二进制更新；从另一个 0.3 build 升级时不需要执行 `ALTER EXTENSION UPDATE`。
- Release `v0.3.2` 增加 `compression`、`secure` 和 `min_tls_version` server options，增加 `regexp_match()` 下推，并增加 PostgreSQL 19beta1 distribution support。
- Release `v0.3.2` 还修复 regular-expression flag pushdown，并在 regex argument 不是常量时避免下推 regex calls。
- Release `v0.3.1` 对已有 SQL version `0.3` 而言只是二进制更新；从 `v0.3.0` 升级时不需要执行 `ALTER EXTENSION UPDATE`。
- Release `v0.3.1` 将客户端库替换为 `ClickHouse/clickhouse-c`，支持流式读取 result blocks，并为 `SELECT` 和 `INSERT` 增加矩形多维数组支持。
- Release `v0.3.1` 还为 `pg_re2` 0.3.0 函数增加下推，例如 `re2extractallgroupshorizontal`、`re2extractallgroupsvertical`、`re2regexpquotemeta` 和 `re2splitbyregexp`，并修复 `UInt16` 到 PostgreSQL `int4` 的 cast。
- Release `v0.3.0` 使用 SQL version `0.3`；执行 `ALTER EXTENSION pg_clickhouse UPDATE TO '0.3'` 可应用其 SQL 层权限变更。
- Release `v0.3.0` 增加了 `re2` 函数、`soundex()`、双参数 `levenshtein()`、兼容的 `to_char(timestamp[tz], fmt)`、部分内置函数，以及 JSON/JSONB path operations 的下推。
- ClickHouse `JSON` 映射到 PostgreSQL `jsonb` 或 `json`；binary driver 的 `JSON` 映射要求 ClickHouse 24.10 或更新版本。
- `pg_clickhouse.pushdown_regex` 控制内置 PostgreSQL regex 下推。上游建议：如果正则处理需要直接下推，可以考虑使用 `re2` 扩展。

### 注意事项

- 在 0.3.0 中，`clickhouse_raw_query(text, text)` 不再对 `PUBLIC` 可执行；只应授予确实需要 ad-hoc ClickHouse 查询的角色。
- 上游将它定位为 analytics-first 扩展；轻量级 `DELETE` 和 `UPDATE` 支持仍在 roadmap 上。
- 完整示例请参考官方 tutorial：其中会创建 ClickHouse `taxi` 数据库，通过 `IMPORT FOREIGN SCHEMA` 导入，并查询生成的 foreign tables。
