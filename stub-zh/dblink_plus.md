## 用法

来源：

- [官方 README](https://github.com/ossc-db/dblink_plus/blob/cb0efffe7f9098e352ccf5da65738630cf40dc5a/README.md)
- [官方文档](https://ossc-db.github.io/dblink_plus/index.html)
- [官方扩展 SQL 模板](https://github.com/ossc-db/dblink_plus/blob/cb0efffe7f9098e352ccf5da65738630cf40dc5a/dblink_plus.sql.in)

`dblink_plus` 1.0.11 通过 connector plugin 在外部 PostgreSQL、Oracle、MySQL 或 SQLite 数据库上执行 SQL。与核心 `dblink` 不同，启用 `use_xa` 后，它可以用两阶段提交协调远程事务与本地事务。

### 核心流程

为已编译的 connector 创建 foreign-data-wrapper 定义、server 和 user mapping，然后打开命名连接：

```sql
CREATE EXTENSION dblink_plus;

CREATE FOREIGN DATA WRAPPER remote_postgres
  VALIDATOR dblink.postgres;

CREATE SERVER reporting_db
  FOREIGN DATA WRAPPER remote_postgres
  OPTIONS (dbname 'reporting');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER reporting_db
  OPTIONS (user 'report_reader');

BEGIN;
SELECT dblink.connect('reporting_conn', 'reporting_db');

SELECT *
FROM dblink.query(
  'reporting_conn',
  'SELECT id, name FROM customers ORDER BY id'
) AS t(id bigint, name text);

SELECT dblink.exec(
  'reporting_conn',
  'UPDATE job_state SET checked_at = CURRENT_TIMESTAMP WHERE id = 1'
);
COMMIT;

SELECT dblink.disconnect('reporting_conn');
```

由于 `dblink.query` 和 `dblink.call` 返回 `SETOF record`，调用方必须提供预期的列名和类型。传入 server 名而不是命名连接时，会创建在事务结束时关闭的匿名连接。

### 重要对象

- `dblink.connect(...)` 和 `dblink.disconnect(text)` 管理持久命名连接。
- `dblink.query(...)` 返回结果行，`dblink.exec(text, text)` 返回远程命令影响的行数。
- `dblink.open(...)`、`dblink.fetch(...)` 和 `dblink.close(...)` 提供事务范围内的 cursor 访问。
- `dblink.call(...)` 调用 stored function 并返回记录。
- `dblink.connections` 报告连接名、server OID、状态、`use_xa` 以及连接是否保持。
- `dblink.postgres`、`dblink.oracle`、`dblink.mysql` 和 `dblink.sqlite3` 是源码树中存在的 connector validator。

### 事务与配置边界

`dblink_plus.use_xa` 控制匿名连接的自动事务管理，默认为 on。`dblink.connect` 的显式 `use_xa` 参数优先；省略该参数时，即使 GUC 为 off，也会启动启用 XA 的连接。PostgreSQL 远端服务器必须把 `max_prepared_transactions` 设为大于零，才能使用两阶段提交。不能在事务块内执行的命令，例如 PostgreSQL `VACUUM` 或多种 Oracle DDL，必须通过 `use_xa = false` 的连接执行。

扩展无需预加载。connector 库及其客户端依赖在构建时选择，因此必须核验安装后的二进制实际包含 PostgreSQL、Oracle、MySQL 和 SQLite 中的哪些 connector。

### 运维说明

XA 路径使用 `dblink.atcommit` 中的 deferred delete trigger；它不能在只读事务或 hot standby 事务中运行，此时必须禁用 `use_xa`。大型远程结果会在本地 backend 中物化，可能消耗大量内存。Oracle 的 `max_value_len` 和 fetch size 会直接影响缓冲区大小，官方文档还列出了不支持的 Oracle LOB/raw 类型。对非超级用户，只授予 schema `dblink` 所需的 `USAGE`、函数 `EXECUTE`，以及文档要求的内部表 `DELETE` 权限。生产使用前应测试回滚、prepared transaction 恢复、日志、凭据存储与远程故障行为。
