## 用法

来源：

- [官方 README](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/README.md)
- [官方用户指南](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/docs/USER_GUIDE.md)
- [扩展 control 文件](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/clickhousedb_fdw.control)

`clickhousedb_fdw` 通过 PostgreSQL 外部数据包装器 API 访问 ClickHouse 表。它支持远程读取和插入，并可下推聚合与连接，但并不是通用的读写兼容层：上游未提供 `UPDATE` 或 `DELETE` 支持。

### 核心流程

创建扩展前，先在每个 PostgreSQL 主机上安装并配置 ClickHouse ODBC 驱动。服务器选项用于指定 ClickHouse 数据库、ODBC 驱动和访问端点。

```sql
CREATE EXTENSION clickhousedb_fdw;

CREATE SERVER clickhouse_svr
  FOREIGN DATA WRAPPER clickhousedb_fdw
  OPTIONS (
    dbname 'default',
    driver '/usr/lib/libclickhouseodbc.so',
    host '127.0.0.1'
  );

CREATE USER MAPPING FOR CURRENT_USER
  SERVER clickhouse_svr;
```

定义列名和类型与远端 ClickHouse 表相符的外部表，然后通过它查询或插入数据。

```sql
CREATE FOREIGN TABLE events_ch (
  event_time timestamp,
  user_id bigint,
  payload text
)
SERVER clickhouse_svr
OPTIONS (table_name 'events');

SELECT user_id, count(*)
FROM events_ch
GROUP BY user_id;

INSERT INTO events_ch VALUES (clock_timestamp(), 42, 'created');
```

使用详细执行计划查看发送给 ClickHouse 的 SQL，并确认过滤、聚合或连接是否真正下推。

```sql
EXPLAIN (VERBOSE)
SELECT user_id, count(*)
FROM events_ch
WHERE event_time >= timestamp '2026-01-01'
GROUP BY user_id;
```

### 重要对象

- `clickhousedb_fdw`：扩展安装的外部数据包装器。
- 服务器选项 `dbname`、`driver` 和 `host`：分别选择 ClickHouse 数据库、ODBC 驱动库与端点。
- 外部表选项 `table_name`：把 PostgreSQL 外部表映射到 ClickHouse 表。
- 远程 `SELECT` 和 `INSERT`：已确认的 DML 接口；聚合与连接下推是规划器优化，并不保证适用于所有查询形态。

### 运维说明

所复核的上游版本将项目标为重度开发状态，并说明测试范围是 PostgreSQL 11–13。部署前应验证实际 PostgreSQL、ClickHouse、ODBC 驱动和类型转换组合。复杂连接和部分 ClickHouse 设置明确仍不完整。网络、ODBC 与远程执行错误会在 PostgreSQL 语句中发生，因此应设置合适的语句超时并测试故障恢复。不要假定其事务、约束、锁与回滚语义等同于本地 PostgreSQL 表。
