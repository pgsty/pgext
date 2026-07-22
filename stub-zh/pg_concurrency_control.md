## 用法

来源：

- [阿里云使用文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-pg-concurrency-control-plug-in)
- [阿里云扩展支持矩阵](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`pg_concurrency_control` 1.0 是阿里云 ApsaraDB RDS for PostgreSQL 扩展；达到配置的并发上限时，它会按工作负载类别排队 SQL。它是防止资源耗尽的云厂商专用机制，不是可移植的社区扩展。当前官方使用页面将此流程限制在 RDS PostgreSQL 10 与 11。

### 启用与观察

在受支持的实例上创建扩展，并查看队列计数器：

```sql
CREATE EXTENSION pg_concurrency_control;

SELECT * FROM pg_concurrency_control_status();
```

状态函数返回 `autocommit_count`、`bigquery_count`、`query_count` 和 `transaction_count`。计数器为正表示该类别存在等待语句。

通过 RDS 参数管理流程配置并发参数。每个上限默认为 `0`，即禁用该类别；设置为 `1` 至 `1024` 会启用排队。

### 上限与超时

- `pg_concurrency_control.query_concurrency` 限制 SELECT 语句。
- `pg_concurrency_control.bigquery_concurrency` 限制显式标记为慢查询的语句。
- `pg_concurrency_control.transaction_concurrency` 限制事务块。
- `pg_concurrency_control.autocommit_concurrency` 限制自动提交 DML。
- `pg_concurrency_control.control_timeout` 与 `pg_concurrency_control.bigsql_control_timeout` 设置队列超时。
- `pg_concurrency_control.timeout_action` 与 `pg_concurrency_control.bigsql_timeout_action` 选择 `TCC_break`、`TCC_rollback` 或 `TCC_wait` 行为。

使用云厂商 hint 将语句标记为大查询类别：

```sql
/*+bigsql*/ SELECT pg_sleep(10);
```

### 运维边界

并发上限会改变所有匹配语句的延迟和失败行为。投入生产前应压力测试队列深度、应用超时、事务重试逻辑、连接池大小和每种超时动作。队列可以保护 CPU 或 I/O，但也可能把过载转移为连接池拥塞和客户端超时风暴。应将四个状态计数器与活动会话、延迟、错误和实例资源一起监控。

可用性和可编辑参数行为取决于准确的 RDS 引擎小版本。当前支持矩阵只在较早的受支持大版本中列出 1.0，因此升级大版本前必须核对目标实例。应按照阿里云的权限和维护模型配置与移除扩展；不要把其二进制文件复制到自托管 PostgreSQL。
