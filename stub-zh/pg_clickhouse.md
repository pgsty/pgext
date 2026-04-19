## 用法

来源：[README](https://github.com/ClickHouse/pg_clickhouse/blob/master/README.md)，[reference](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/pg_clickhouse.md)，[tutorial](https://github.com/ClickHouse/pg_clickhouse/blob/master/doc/tutorial.md)，[v0.2.0 release notes](https://github.com/ClickHouse/pg_clickhouse/releases/tag/v0.2.0)

`pg_clickhouse` 通过 `clickhouse_fdw` foreign data wrapper 在 PostgreSQL 中直接执行 ClickHouse 分析查询。上游文档说明它支持 PostgreSQL 13+ 与 ClickHouse 23+。

### 将 PostgreSQL 连接到 ClickHouse

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

上游记录的 server 选项包括：

- `driver`：必填，取值为 `binary` 或 `http`
- `host`
- `port`
- `dbname`
- `fetch_size`：HTTP 流式批大小，`0` 表示禁用流式

user mapping 选项包括：

- `user`
- `password`

### 常见操作

```sql
ALTER EXTENSION pg_clickhouse UPDATE;
ALTER EXTENSION pg_clickhouse UPDATE TO '0.2';
DROP SERVER taxi_srv CASCADE;
```

`IMPORT FOREIGN SCHEMA` 还支持 `LIMIT TO (...)` 与 `EXCEPT (...)`。参考文档特别提醒：导入的 mixed-case 标识符会在 PostgreSQL 中带双引号，查询时也必须显式加引号。

### 版本与下推说明

- 参考文档区分库版本和扩展版本；`pgch_version()` 是在 `v0.2.0` 中加入的。
- 仅补丁级别的发布可能只更新库文件，不需要执行 `ALTER EXTENSION`。
- `v0.2.0` 增加了数组、正则函数、`split_part()`、数组运算符与当前日期/时间表达式的更多下推能力，以及 `pg_clickhouse.pushdown_regex` 设置。

### 注意事项

- 上游将它定位为分析优先的扩展，路线图里更广泛的 DML 支持仍属于后续工作。
- 若需要完整示例，应按官方 tutorial 创建 ClickHouse `taxi` 数据库，通过 `IMPORT FOREIGN SCHEMA` 导入，再查询生成的 foreign tables。
