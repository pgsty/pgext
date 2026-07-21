## 用法

来源：

- [官方v2.4 README](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/README.md)
- [v2.4版本发布日志](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/ChangeLog)
- [v2.4扩展控制文件](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/pg_dbms_errlog.control)
- [v2.4扩展SQL脚本](https://github.com/HexaCluster/pg_dbms_errlog/blob/v2.4/sql/pg_dbms_errlog--2.4.sql)

`pg_dbms_errlog` 为 PostgreSQL 提供了类似 Oracle 的 DML 错误日志功能。它会将 `INSERT`, `UPDATE`, 或 `DELETE` 失败时产生的错误排队，通过后台工作进程写入注册的 `ERR$_...` 表，并允许脚本在回滚到保存点后继续执行。这需要 `pg_statement_rollback` 权限或由调用方显式管理保存点。

### 启用扩展

将库添加到 `shared_preload_libraries`，确保 `max_worker_processes` 能容纳 `pg_dbms_errlog.max_workers` 加上固定的工作进程数，并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_dbms_errlog'
```

```sql
CREATE EXTENSION pg_dbms_errlog;
```

为每个 DML 目标创建并注册一个错误表：

```sql
CREATE TABLE raises (
    employee_id integer,
    salary integer CHECK (salary > 8000)
);

CALL dbms_errlog.create_error_log('raises');
-- Creates and registers public."ERR$_raises" by default.
```

### 错误发生后记录和继续

```sql
LOAD 'pg_statement_rollback';

SET pg_statement_rollback.enabled = on;
SET pg_dbms_errlog.enabled = on;
SET pg_dbms_errlog.query_tag = 'daily_load';
SET pg_dbms_errlog.reject_limit = 10;

BEGIN;
INSERT INTO raises VALUES (145, 15400);
INSERT INTO raises VALUES (161, 7700);  -- logged failure
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
INSERT INTO raises VALUES (175, 9680);
COMMIT;

SELECT * FROM "ERR$_raises";
```

错误表包含 `pg_err_number$`, `pg_err_mesg$`, `pg_err_optyp$`, `pg_err_tag$`, `pg_err_query$` 和 `pg_err_detail$`。

### API 和配置索引

- `dbms_errlog.create_error_log(dml_table_name, err_log_table_name, err_log_table_owner, err_log_table_space)`: 创建并注册一个错误表。
- `dbms_errlog.publish_queue(wait_for_completion)`: 请求工作进程处理排队的错误；默认情况下，`PUBLIC` 无法执行此操作。
- `dbms_errlog.queue_size()`: 报告排队中的错误数量。
- `pg_dbms_errlog.synchronous`: 默认为 `transaction`，可选值为 `query` 或 `off`。事务模式确保仅记录已提交事务的错误。
- `pg_dbms_errlog.reject_limit`: 事务范围内的错误限制；`-1` 表示无限制，`0` 不记录任何错误并回滚。
- `pg_dbms_errlog.no_client_error`: 抑制客户端错误消息的同时保留服务器日志；默认启用。
- `pg_dbms_errlog.frequency` 和 `pg_dbms_errlog.max_workers`: 异步工作进程的执行频率和并发数。

### 注意事项

- 调用方需要对目标表和错误表具有 DML 权限；创建错误表还需要上游描述的执行和注册表权限。
- `INSERT INTO ... SELECT ...` 是一个 PostgreSQL 语句，无法像 Oracle 那样仅保留成功行。
- 语法和其他解析时错误不会被记录。存储查询文本必须保持在 PostgreSQL 的 1 GB 值限制以下。
- 版本 `2.4` 没有更改 SQL API；它修复了工作进程关闭循环和动态后台工作进程崩溃的问题。
