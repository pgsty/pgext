


## 用法

> 来源：[pg_task upstream README](https://github.com/RekGRpth/pg_task)、[PGXN pg_task](https://pgxn.org/dist/pg_task/)。

`pg_task` 是一个后台工作进程调度器，用于在计划时间异步执行 SQL。上游文档说明它支持 PostgreSQL、Greenplum 和 Greengage。

在服务器启动时启用 worker，然后在拥有任务表的数据库中创建扩展：

```conf
shared_preload_libraries = 'pg_task'
```

```sql
CREATE EXTENSION pg_task;
```

### 调度任务

通过向配置的任务表插入行来调度工作。默认情况下，任务表是 `postgres` 数据库中的 `public.task`，除非通过 GUC 修改。

```sql
-- Run SQL immediately
INSERT INTO task (input) VALUES ('SELECT now()');

-- Run SQL after 5 minutes
INSERT INTO task (plan, input) VALUES (now() + '5 min'::interval, 'SELECT now()');

-- Run SQL at a specific time
INSERT INTO task (plan, input) VALUES ('2029-07-01 12:51:00', 'SELECT now()');

-- Repeat SQL every 5 minutes
INSERT INTO task (repeat, input) VALUES ('5 min', 'SELECT now()');

-- Exceptions are caught and written to the error column
INSERT INTO task (input) VALUES ('SELECT 1/0');

-- Limit concurrent tasks in a group.
-- max = 1 means one task at a time for this group.
INSERT INTO task ("group", max, input) VALUES ('billing', 1, 'SELECT refresh_billing_cache()');

-- Run SQL on a remote database
INSERT INTO task (input, remote) VALUES ('SELECT now()', 'user=user host=host');
```

### 任务表

任务表是面向用户可见的。上游说明用户可以为它添加列，也可以对它分区。

关键列：

| 名称 | 类型 | 默认值 | 说明 |
|------|------|---------|-------------|
| id | bigserial | autoincrement | 主键 |
| parent | bigint | pg_task.id | 父任务 ID |
| plan | timestamptz | statement_timestamp() | 计划开始时间 |
| start | timestamptz | | 实际开始时间 |
| stop | timestamptz | | 实际结束时间 |
| active | interval | 1 hour | 计划时间后任务仍可执行的活跃时段 |
| live | interval | 0 sec | 后台工作进程最长存活时间 |
| repeat | interval | 0 sec | 自动重复间隔 |
| timeout | interval | 0 sec | 任务运行允许的时间 |
| count | int | 0 | worker 退出前的最大任务数 |
| max | int | 0 | 分组内最大并发任务数；负值表示任务之间的暂停毫秒数 |
| pid | int | | 执行任务的进程 ID |
| state | enum | PLAN | PLAN、TAKE、WORK、DONE、STOP |
| delete | bool | true | 当 output 和 error 为空时自动删除 |
| drift | bool | false | 根据 stop 时间计算下一次重复 |
| header | bool | true | 在输出中显示列标题 |
| group | text | 'group' | 任务分组名称 |
| input | text | | 要执行的 SQL 命令 |
| output | text | | 接收到的结果 |
| error | text | | 捕获的错误 |
| remote | text | | 远程数据库连接字符串 |

### 配置

关键设置：

| 名称 | 类型 | 默认值 | 说明 |
|------|------|---------|-------------|
| pg_task.delete | bool | true | 自动删除已完成任务 |
| pg_task.drift | bool | false | 根据 stop 时间计算重复 |
| pg_task.header | bool | true | 在任务输出中显示列标题 |
| pg_task.string | bool | true | 输出中仅为字符串加引号 |
| pg_task.count | int | 0 | worker 退出前默认最多执行的任务数 |
| pg_task.fetch | int | 100 | 一次抓取的任务行数 |
| pg_task.limit | int | 1000 | 一次处理的任务行限制 |
| pg_task.max | int | 0 | 分组内默认最大并发任务数 |
| pg_task.run | int | 2147483647 | work 中最多同时执行的任务数 |
| pg_task.sleep | int | 1000 | 每 N 毫秒检查一次任务 |
| pg_task.active | interval | 1 hour | 计划时间后任务仍可执行的活跃时段 |
| pg_task.live | interval | 0 sec | worker 进程最长存活时间 |
| pg_task.repeat | interval | 0 sec | 默认重复间隔 |
| pg_task.timeout | interval | 0 sec | 默认任务超时 |
| pg_task.data | text | postgres | 任务表所在数据库名 |
| pg_task.user | text | postgres | 任务表所用用户名 |
| pg_task.schema | text | public | 任务表所在 schema |
| pg_task.table | text | task | 任务表名 |
| pg_task.json | json | `[{"data":"postgres"}]` | 多数据库配置 |
| pg_task.id | bigint | 0 | 当前任务 ID，只读会话设置 |

### 多数据库配置

通过 JSON 配置在多个数据库上运行任务：

```conf
pg_task.json = '[{"data":"database1"},{"data":"database2","user":"username2"},{"data":"database3","schema":"schema3"}]'
```

如果指定的数据库、用户、schema 或表不存在，`pg_task` 会创建它们。

本地元数据将该包标记为 `headless`，因此它依赖 `shared_preload_libraries`，而不是纯 SQL 的用户侧安装路径。把它用于关键作业前，请在目标 PostgreSQL 版本上验证后台调度行为。
