


## 用法

来源：

- [pg_durable README on main](https://github.com/microsoft/pg_durable/blob/main/README.md)
- [User Guide](https://github.com/microsoft/pg_durable/blob/main/USER_GUIDE.md)
- [control file](https://github.com/microsoft/pg_durable/blob/main/pg_durable.control)
- [0.2.3 to 0.2.4 upgrade SQL](https://github.com/microsoft/pg_durable/blob/main/sql/pg_durable--0.2.3--0.2.4.sql)
- [0.2.2 to 0.2.3 upgrade SQL](https://github.com/microsoft/pg_durable/blob/v0.2.3/sql/pg_durable--0.2.2--0.2.3.sql)

`pg_durable` 在 PostgreSQL 内运行持久、容错的 SQL 工作流。工作流是由 SQL 字符串、函数、定时器、信号、条件分支和并行分支组成的图，通过 `df.start()` 提交。扩展会把状态 checkpoint 到 PostgreSQL，因此崩溃、重启或步骤失败后，已完成步骤不会被重复执行。

当前 main 分支的包版本为 `0.2.4`；本次刷新观察到的最新 tag 是 `v0.2.3`。上游发布包和 Docker 镜像主要面向 PostgreSQL 17 与 18；本地 Pigsty 元数据可能另行验证更宽的 PostgreSQL 版本范围。

### 启用并授权访问

```sql
CREATE EXTENSION pg_durable;

SELECT df.grant_usage('app_role');
```

`pg_durable` 必须通过 `shared_preload_libraries` 加载，并重启 PostgreSQL。`CREATE EXTENSION` 后后台 worker 会异步初始化；如果 `df.*` 调用提示 worker 尚未初始化，稍后重试即可。`CREATE EXTENSION` 不会把 usage 授给 `PUBLIC`，因此需要为应用角色运行 `df.grant_usage()`，扩展升级后也应重新执行以覆盖新增函数。

worker 连接 `pg_durable.database` 指定的数据库，并以 `pg_durable.worker_role` 运行。worker role 必须是 superuser，因为它需要绕过 RLS 管理所有用户的 durable instances。

### 启动并监控工作流

```sql
SELECT df.start('SELECT ''Hello, durable world!''', 'hello-job');

SELECT *
FROM df.list_instances();

SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

`df.start()` 返回 instance ID。状态、结果、取消、信号、解释和等待辅助函数都使用 instance ID，而不是 label。`df.list_instances()` 有基础重载和分页重载；如果需要 `created_at`、`completed_at` 和 `next_cursor`，至少传入三个参数，不需要过滤的参数用 `NULL` 占位。

### 组合 SQL 步骤

```sql
-- 运行一个步骤，给结果命名，再在下一步替换使用。
SELECT df.start(
  'SELECT 100 AS amount' |=> 'total'
  ~> 'SELECT $total * 2 AS doubled',
  'double-total'
);

-- 在步骤之间传递多行结果集。
SELECT df.start(
  'SELECT id FROM documents WHERE processed = false LIMIT 100' |=> 'batch'
  ~> 'UPDATE documents SET processed = true WHERE id IN (SELECT id FROM $batch.*)',
  'process-documents'
);

-- 根据 SQL 条件分支。
SELECT df.start(
  'SELECT count(*) > 10 FROM orders'
    ?> 'SELECT ''high volume'''
    !> 'SELECT ''low volume''',
  'order-volume'
);
```

核心操作符包括表示顺序的 `~>`、为结果命名的 `|=>`、join 的 `&`、race 的 `|`、条件分支 `?>` / `!>`，以及循环 `@>`。结果替换支持 `$name`、`$name.column`、空值安全的 `$name?` / `$name.column?`，以及通过 `$name.*` 展开多行结果集。

### 定时器、计划任务、信号与变量

```sql
SELECT df.start(
  @> (
    df.wait_for_schedule('0 * * * *')
    ~> 'SELECT run_hourly_rollup()'
  ),
  'hourly-rollup'
);

SELECT df.start(
  'SELECT create_invoice()' |=> 'invoice'
  ~> df.wait_for_signal('approval', 86400)
  ~> 'SELECT finalize_invoice($invoice.id)',
  'invoice-approval'
);

SELECT df.setvar('api_url', 'https://example.internal');
SELECT df.getvar('api_url');
```

常用 DSL 函数包括 `df.sleep(seconds)`、`df.wait_for_schedule(cron)`、`df.wait_for_signal(name, timeout)`、`df.signal(id, name, data)`、`df.join()`、`df.join3()`、`df.race()`、`df.if()`、`df.if_rows()`、`df.loop()`、`df.break()`、`df.await_instance()`、`df.explain()` 和 durable variable 辅助函数。

### 监控与运行时状态

```sql
SELECT * FROM df.instance_nodes('a1b2c3d4');
SELECT * FROM df.instance_executions('a1b2c3d4', 20);

-- 需要直接授予管理权限；df.grant_usage() 不包含它。
SELECT * FROM df.metrics();
```

`df.instance_nodes()` 同时报告存储状态和推导状态。推导状态会为未被执行的分支或 race 失败分支显示 `skipped`，并在较新的循环迭代接管后，把旧迭代节点解释为 pending。这只是读取时解释，不改变工作流执行语义。

### 配置与注意事项

- 必需 preload：把 `pg_durable` 加入 `shared_preload_libraries` 并重启 PostgreSQL。
- `pg_durable.database` 必须指向创建该扩展的数据库；除非显式使用受支持的 database 参数，否则工作流不会在其他数据库中处理。
- v0.2.3 provider schema 迁移之后的新安装使用 `_duroxide`；从 <= 0.2.2 升级的安装保留旧 `duroxide` schema。`df.duroxide_schema()` 会报告当前安装应使用哪个 schema。
- `pg_durable.worker_role` 必须存在且为 superuser。
- 连接相关 GUC 包括 `pg_durable.max_management_connections`、`pg_durable.max_duroxide_connections`、`pg_durable.max_user_connections` 和 `pg_durable.execution_acquire_timeout`。
- `df.http()` 会执行出站 HTTP 工作。标准授权不包含 HTTP，除非使用 `df.grant_usage(..., include_http => true)`；release build 还可能通过 feature 限制 HTTP egress。
- 上游状态为 preview。发布的 Docker 镜像用于评估和学习，不建议生产使用。
