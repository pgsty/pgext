## 用法

来源：

- [PGXN 0.2.6 README](https://pgxn.org/dist/pg_durable/0.2.6/README.html)
- [0.2.6 用户指南](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/USER_GUIDE.md)
- [0.2.6 变更日志](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/CHANGELOG.md)
- [pg_durable 控制文件](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/pg_durable.control)
- [0.2.5 至 0.2.6 升级 SQL](https://api.pgxn.org/src/pg_durable/pg_durable-0.2.6/sql/pg_durable--0.2.5--0.2.6.sql)

`pg_durable` 在PostgreSQL中运行持久、容错的SQL工作流。一个工作流是一系列SQL步骤、定时器、信号、条件和并行分支组成的图，通过`df.start()`提交。执行状态会在PostgreSQL中进行检查点记录，因此在崩溃、重启或重试后不会重复已完成的步骤。

### 启用和授权访问

预加载工作进程，如果默认设置不合适，则选择其数据库和超级用户角色，然后重启PostgreSQL：

```conf
shared_preload_libraries = 'pg_durable'
pg_durable.database = 'postgres'
pg_durable.worker_role = 'postgres'
```

在`pg_durable.database`中创建扩展，并授予应用程序登录角色访问权限：

```sql
CREATE EXTENSION pg_durable;
SELECT df.grant_usage('app_role');
```

工作进程角色必须是超级用户，因为它会管理所有用户的实例并绕过行级安全。调用`df.start()`的角色必须具有`LOGIN`权限，因为工作流SQL是通过该捕获角色认证的连接执行的。

### 构建和运行一个工作流

```sql
SELECT df.start(
    'SELECT 100 AS amount' |=> 'total'
    ~> 'SELECT $total.amount * 2 AS doubled',
    'double-total'
);
```

`df.start()`返回实例ID。使用它来监控或控制运行：

```sql
SELECT df.status('a1b2c3d4');
SELECT df.result('a1b2c3d4');
SELECT * FROM df.instance_nodes('a1b2c3d4');
SELECT * FROM df.instance_executions('a1b2c3d4', 20);
SELECT df.cancel('a1b2c3d4', 'No longer needed');
```

### DSL索引

- `~>` 用于序列化步骤；`|=>` 为`$name`、`$name.column` 或 `$name.*` 替换命名结果。
- `&` / `df.join()` 等待并行分支；`|` / `df.race()` 保留第一个结果。
- `?>` 和 `!>` / `df.if()` 选择条件分支；`@>` / `df.loop()` 重复一个图。
- `df.sleep()`、`df.wait_for_schedule()` 和 `df.wait_for_signal()` 使等待持久化。
- `df.signal()`、`df.wait_for_completion()`、`df.explain()` 及实例检查函数操作正在运行或存储的实例。
- `df.setvar()`、`df.getvar()`、`df.unsetvar()` 和 `df.clearvars()` 管理在调用`df.start()`时捕获的用户变量。

### 0.2.6 版本边界

- 上游源码安装与发布镜像使用 `pgrx` 0.16.1，支持 PostgreSQL 17 与 18。扩展仍要求 `shared_preload_libraries`、重启以及超级用户工作角色。
- 经由 0.2.4 与 0.2.5 的升级包含会破坏重放的工作流变更。升级前应排空或取消运行中的 JOIN、RACE、循环与 `df.wait_for_schedule()` 工作；0.2.4 的 `df.nodes` 键迁移还会获取 `ACCESS EXCLUSIVE` 锁。
- `df.start(..., transaction_mode => 'new')` 会在调用者事务之外持久化独立启动。集群默认最多并发启动两个，由 `pg_durable.max_new_transaction_starts` 与 `pg_durable.new_transaction_start_timeout` 控制。
- 0.2.6 从左到右只解析一次变量替换，所以变量值引入的令牌形文本不会再次扫描。它仍是原始 SQL 替换；绝不能把不可信输入放进 `{name}` 变量。通过 `$name` 的命名步骤结果替换会执行 SQL 转义。
- 未公开的 `df.ensure_durofut(text)` 辅助函数已移除。升级前应删除或改写客户自有的依赖对象。
- 在 `ALTER EXTENSION ... UPDATE` 后重新运行 `df.grant_usage()`，因为对全部函数的授权不会自动覆盖后来新增的函数。
- `df.http()` 与 `df.http_multipart()` 的可用性和出站策略是编译时特性。其限制不会沙箱化任意 SQL 或其他已安装扩展。
- 项目仍处于 1.0 之前，上游发布的 Docker 镜像用于评估与学习，而不是生产。应阅读每个相邻版本的升级警告，不要假设未经测试的跨版本跳跃可安全重放。
