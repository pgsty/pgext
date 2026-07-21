## 用法

来源：

- [官方v0.2.3 README](https://github.com/microsoft/pg_durable/blob/v0.2.3/README.md)
- [v0.2.3用户指南](https://github.com/microsoft/pg_durable/blob/v0.2.3/USER_GUIDE.md)
- [v0.2.3发布说明](https://github.com/microsoft/pg_durable/releases/tag/v0.2.3)
- [从v0.2.2到v0.2.3的升级SQL](https://github.com/microsoft/pg_durable/blob/v0.2.3/sql/pg_durable--0.2.2--0.2.3.sql)

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

### v0.2.3边界

- 新鲜安装v0.2.3将提供程序对象放置在`_duroxide`中；从0.2.2或更早版本升级后保留`duroxide`。`df.duroxide_schema()`报告当前活动的模式。
- 深度超过256级或节点数大于10,000的工作流被拒绝。条件查询返回空行时评估为假。
- 在执行`df.grant_usage()`后重新运行`ALTER EXTENSION ... UPDATE`，因为所有函数上的权限不会自动包含后来添加的函数。
- `{name}`变量替换是原始SQL文本替换；不要将不可信输入放置在这样的变量中。通过`$name`进行命名步骤结果替换执行SQL转义。
- `df.http()`可用性和出站策略是编译时特性。其限制不会对任意SQL或其他已安装的扩展进行沙盒化。
- 上游将该项目标记为预览版，发布的v0.2.3 Docker镜像用于评估和学习而非生产环境。
