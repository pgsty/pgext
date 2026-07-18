## 用法

来源：

- [阿里云官方 pg_concurrency_control 指南](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-pg-concurrency-control-plug-in)
- [ApsaraDB RDS PostgreSQL 官方支持扩展列表](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`pg_concurrency_control` 是阿里云 ApsaraDB RDS PostgreSQL 的扩展，可按类别限制同时执行的 SQL，并将超额工作放入队列。它通过 `query_concurrency`、`autocommit_concurrency`、`transaction_concurrency` 和 `bigquery_concurrency`，分别限制 SELECT 查询、自动提交 DML、事务块和显式标记的慢查询。

### 启用查询并发限制

```sql
CREATE EXTENSION pg_concurrency_control;

SET pg_concurrency_control.query_concurrency = 10;

SELECT * FROM pg_concurrency_control_status();
```

每个并发限制默认值都是零，即不控制该类别。`pg_concurrency_control_status()` 报告自动提交、大查询、查询和事务队列中的当前数量。

### 超时与适用范围

队列等待限制由 `control_timeout` 和 `bigsql_control_timeout` 控制。文档给出的超时动作是 `TCC_break`、`TCC_rollback` 和 `TCC_wait`；应根据超时后是跳过语句、中止事务还是继续等待来选择。

### 注意事项

- 这是 ApsaraDB RDS PostgreSQL 的服务商专有扩展，并不是面向自管理 PostgreSQL 普遍发布的扩展。
- 专用使用指南当前要求 ApsaraDB RDS 实例运行 PostgreSQL 10 或 11。启用前应核对支持扩展表，并升级到最新次要引擎版本。
- 受支持的引擎组合列出了 1.0 版。服务商没有发布其源码、control 文件、软件许可证或可移植安装制品。
- 官方指南直接使用 `CREATE EXTENSION` 安装，没有说明该扩展需要预加载。
