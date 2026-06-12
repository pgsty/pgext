来源：[pg_durable v0.2.2 README](https://github.com/microsoft/pg_durable/blob/v0.2.2/README.md)、[User Guide](https://github.com/microsoft/pg_durable/blob/v0.2.2/USER_GUIDE.md)、[control file](https://github.com/microsoft/pg_durable/blob/v0.2.2/pg_durable.control)、[GUC definitions](https://github.com/microsoft/pg_durable/blob/v0.2.2/src/lib.rs)。

## 用法

`pg_durable` 在 PostgreSQL 内运行持久、容错的 SQL 工作流。工作流由 SQL 字符串、函数和 DSL 操作符组成，并通过 `df.start()` 提交。状态会持久化在 PostgreSQL 中，因此崩溃或重启后已完成的步骤不会被重复执行。

`pg_durable` 必须通过 `shared_preload_libraries` 加载，然后重启 PostgreSQL。它的后台 worker 会连接到 `pg_durable.database` 指定的数据库，并以 `pg_durable.worker_role` 运行；上游默认值分别为 `postgres` 和 `azuresu`，且 worker role 必须是 superuser。

### 启用并授权访问

```sql
CREATE EXTENSION pg_durable;

SELECT df.grant_usage('app_role');
```

`CREATE EXTENSION` 不会把 usage 授给 `PUBLIC`。请为应用角色使用 `df.grant_usage()`，并在扩展升级后重新执行，确保新增加的函数也被覆盖。后台 worker 会在扩展创建后异步初始化；如果 `df.*` 调用报告 worker 尚未初始化，可以稍后重试。

### 启动并监控工作流

```sql
SELECT df.start('SELECT ''Hello, durable world!''', 'hello-job');

SELECT *
FROM df.list_instances();

SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

`df.start()` 返回一个 instance ID。使用该 ID 调用 `df.status()`、`df.result()`、`df.cancel()`、`df.signal()` 和 `df.explain()`。

### 组合 SQL 步骤

```sql
-- Run one step, name its result, then substitute it in the next step.
SELECT df.start(
  'SELECT 100 AS amount' |=> 'total'
  ~> 'SELECT $total * 2 AS doubled',
  'double-total'
);

-- Branch on a SQL condition.
SELECT df.start(
  'SELECT count(*) > 10 FROM orders'
    ?> 'SELECT ''high volume'''
    !> 'SELECT ''low volume''',
  'order-volume'
);

-- Run in parallel and wait for both branches.
SELECT df.start(
  'SELECT refresh_accounts()' & 'SELECT refresh_orders()',
  'parallel-refresh'
);
```

核心操作符包括表示顺序执行的 `~>`、为结果命名的 `|=>`、join 的 `&`、race 的 `|`、条件分支 `?>` 和 `!>`，以及循环 `@>`。

### 定时器、计划任务与信号

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
```

常用 DSL 函数包括 `df.sleep(seconds)`、`df.wait_for_schedule(cron)`、`df.wait_for_signal(name, timeout)`、`df.signal(id, name, data)`、`df.join()`、`df.race()`、`df.if()`、`df.loop()` 和 `df.explain()`。

### 配置与注意事项

- 必需 preload：把 `pg_durable` 加入 `shared_preload_libraries` 并重启 PostgreSQL。
- `pg_durable.database` 必须指向创建该扩展的数据库；工作流不会在其他数据库中被处理。
- `pg_durable.worker_role` 必须存在且为 superuser，因为 worker 会绕过 RLS 来管理所有用户的实例。
- 连接相关 GUC 包括 `pg_durable.max_management_connections`、`pg_durable.max_duroxide_connections`、`pg_durable.max_user_connections` 和 `pg_durable.execution_acquire_timeout`。
- `df.http()` 会执行出站 HTTP 工作，除非使用 `df.grant_usage(..., include_http => true)`，否则不会包含在标准授权中。
- 上游将 v0.2.2 标记为早期开发版本，尚不适合生产环境。
