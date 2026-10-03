## 用法

来源：

- [sql/pg_trickle--0.108.0--0.108.1.sql](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/sql/pg_trickle--0.108.0--0.108.1.sql)
- [Version 0.108.1 README](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/README.md)
- [SQL reference](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/SQL_REFERENCE.md)
- [Configuration](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/CONFIGURATION.md)
- [GUC catalog](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/GUC_CATALOG.md)
- [Upgrade guide](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/docs/UPGRADING.md)
- [Control file](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/pg_trickle.control)
- [0.107.0 to 0.108.0 migration](https://github.com/trickle-labs/pg-trickle/blob/v0.108.1/sql/pg_trickle--0.107.0--0.108.0.sql)

`pg_trickle` 0.108.1 在 PostgreSQL 18 上维护流式表：由 SQL 查询定义、可以常规查询的数据表；在支持时增量刷新，也可全量重算。流式表可依赖其他流式表，形成依赖图，同时支持在同一事务中维护结果。

### 启用扩展

将 `pg_trickle` 追加到 `shared_preload_libraries` 后重启 PostgreSQL，再以超级用户身份安装。应根据部署规模配置后台工作进程池；上游示例使用八个工作进程。

```ini
shared_preload_libraries = 'pg_trickle'
max_worker_processes = 8
```

```sql
CREATE EXTENSION pg_trickle;
```

`pg_trickle.cdc_mode` 默认为 `trigger`，事务内的变更捕获不需要逻辑 WAL 或复制槽。显式选择 `auto` 时先使用触发器，符合条件的源表随后可切换到带接收确认的 WAL 捕获。显式选择 `wal` 后，若准入检查或逻辑解码前置条件不满足，也会回退到触发器。不同方式的写入开销与运维要求不同。

### 创建并刷新流式表

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, region text, amount numeric);
SELECT pgtrickle.create_stream_table(
    name => 'regional_totals',
    query => 'SELECT region, SUM(amount) AS total, COUNT(*) AS cnt FROM orders GROUP BY region',
    schedule => '30s',
    refresh_mode => 'AUTO'
);
INSERT INTO orders VALUES (1, 'east', 10);
SELECT pgtrickle.refresh_stream_table('regional_totals');
SELECT * FROM regional_totals;
```

`initialize` 默认为 true，因此创建时会填充结果。`schedule` 接受时间间隔、`@hourly` 等 cron 表达式，或默认的 `calculated`，从下游依赖者继承刷新周期。`AUTO` 在可行时选择差分维护，也可回退为全量刷新。`DIFFERENTIAL` 会拒绝无法增量维护的查询；`FULL` 会清空并重新加载结果。

`IMMEDIATE` 在基表写入事务内使用语句级触发器，不使用 WAL 捕获，并拒绝实际生效的显式 WAL 请求。连接、聚合、子查询、递归查询等应遵循文档中的查询准入规则；支持某类语法不代表所有 SQL 表达式都能进行差分维护。

### 生命周期与监控

`pgtrickle.alter_stream_table` 修改定义或刷新策略，`pgtrickle.drop_stream_table` 删除托管表。恢复或管理员 DDL 导致捕获设施缺失时，`pgtrickle.repair_stream_table` 可修复设施并重置维护状态。应使用生命周期 API，不要直接写入托管流式表或在其上使用外键。

```sql
SELECT * FROM pgtrickle.pgt_status();
SELECT * FROM pgtrickle.health_check();
SELECT * FROM pgtrickle.dependency_tree();
SELECT * FROM pgtrickle.explain_st('regional_totals');
```

生命周期函数要求显式执行授权并检查所有权；全局管理操作仅限扩展所有者或超级用户。接受任意 SQL 的辅助函数保留调用者权限。捕获触发器会增加源表写入工作量，刷新失败可能导致变更缓冲积压，因此应监控健康状态与存储，不应将刷新周期视为严格的新鲜度保证。

### 外部协调与结果增量

`orchestration_mode` 可选择 `MANAGED` 调度或 `EXTERNAL` 协调，外部协调不能与即时维护同时使用。`pgtrickle.integration_capabilities` 公布可用接口契约；0.108.0 提供 Graph V1 的 1.2 版和 Delta V1 的 1.1 版。

`pgtrickle.output_delta_consumer_status` 报告消费者状态，`pgtrickle.validate_output_delta_consumer` 检查消费者能否恢复。`pgtrickle.request_output_delta_resnapshot`、`pgtrickle.begin_output_delta_resnapshot` 和 `pgtrickle.ack_output_delta_resnapshot` 管理基线重建。重新获取快照时，会校验数据库实例身份、输出契约摘要和行标识版本。推进外部投递前应遵循 SQL 参考中的准确签名与确认协议。

### 升级至 0.108.0

先安装新的库和扩展文件，再执行随版本提供的迁移：

```sql
ALTER EXTENSION pg_trickle UPDATE TO '0.108.1';
SELECT * FROM pgtrickle.output_delta_consumer_status();
```

升级保留消费者、游标、批次和类型化载荷，并新增快照重建校验。0.106.1 可通过随包提供的 0.107.0 迁移继续升级。恢复投递前应逐一验证消费者；出现 `INVALIDATED` 或 `RESNAPSHOT_REQUIRED` 时，必须创建并确认新的基线。0.105.2 及以前的部分版本存在文档列明的特定查询形态差分结果错误；升级不会自动修复已经物化的错误行。适用时应按升级指南进行比较和修复。

### 0.108.1 版本变化

本补丁修复上游 FULL 刷新或截断后下游 `IMMEDIATE` 表未及时更新的问题，支持恢复缺失的变更缓冲区，并修复源模式变更后的意外挂起及标量子查询增量刷新。同时增加 PostgreSQL 18.6 支持。安装匹配文件后执行 `ALTER EXTENSION pg_trickle UPDATE TO '0.108.1'`，并检查依赖流表结果及变更捕获状态。
