## 用法

来源：

- [阿里云 RDS 在线 DDL 官方文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-online-ddl-extension-to-modify-column-data-type-online)
- [ApsaraDB RDS for PostgreSQL 官方概览](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_online_ddl` 是阿里云 ApsaraDB RDS for PostgreSQL 服务扩展，用于修改列数据类型，相比常规表重写可以缩短最终阻塞窗口。它把数据复制到临时表、建立索引、通过逻辑解码应用并发变更，最后短暂获取排他锁并交换表。

### 前置条件

服务商文档要求 RDS PostgreSQL 实例使用 PostgreSQL 12 或更高版本并满足最低小版本门槛，设置 `wal_level = logical`，目标表有主键或唯一约束，并使用权限足够的账户。安排操作前，应在当前服务控制台和文档中核对准确的引擎版本门槛。

### 核心流程

通过 RDS 支持的方式创建服务扩展，估算空间与运行时间，然后只把 `ALTER COLUMN ... TYPE ...` 子句提交给 `rds_online_ddl.alter_table`：

```sql
CREATE EXTENSION rds_online_ddl;

SELECT rds_online_ddl.alter_table(
  'public.orders',
  'ALTER COLUMN amount TYPE numeric(20,2)'
);
```

通过服务商提供的进度视图监控当前任务：

```sql
SELECT *
FROM rds_online_ddl.pg_stat_progress_online_ddl;
```

完成后，应验证列类型、约束、索引、行数、应用读写和查询计划，不要把函数返回当作端到端验证。

### 限制与运维风险

- 阿里云记录的是列数据类型变更，而不是任意在线 DDL；不要传入无关的 `ALTER TABLE` 操作。
- 复制期间应按大约两倍的表加索引占用预留空间，并监控实例存储、I/O、WAL 生成、复制延迟、长事务和逻辑解码保留量。
- 服务商不支持带外键的表和分区表。开始前应检查入向与出向约束，以及真实目标关系。
- 最终交换仍需要短暂排他锁。应安排维护窗口、控制锁等待行为，并清理可能延迟切换的长事务。
- 带 `USING` 的表达式受服务限制，不一定支持所有转换。应使用代表性数据演练精确子句，包括空值、默认值、生成值和转换失败。
- 按阿里云记录的流程，操作失败时由服务原子回滚；但仍应针对当前 RDS 引擎版本测试取消、故障转移、备份、副本和重试行为。
