

## 用法

> [pg_task: PostgreSQL 任务调度器](https://github.com/RekGRpth/pg_task)

`pg_task` 允许在后台异步地在任意指定时间执行任意 SQL 命令。支持 PostgreSQL、Greenplum 和 Greengage。

首先在 `postgresql.conf` 中添加：

```conf
shared_preload_libraries = 'pg_task'
```

然后通过向 `task` 表插入记录来调度任务：

```sql
-- 立即执行 SQL
INSERT INTO task (input) VALUES ('SELECT now()');

-- 5 分钟后执行 SQL
INSERT INTO task (plan, input) VALUES (now() + '5 min'::INTERVAL, 'SELECT now()');

-- 在指定时间执行 SQL
INSERT INTO task (plan, input) VALUES ('2029-07-01 12:51:00', 'SELECT now()');

-- 每 5 分钟重复执行 SQL
INSERT INTO task (repeat, input) VALUES ('5 min', 'SELECT now()');

-- 异常会被捕获并写入 error 列
INSERT INTO task (input) VALUES ('SELECT 1/0');

-- 限制同一组内的并发任务数
INSERT INTO task (group, max, input) VALUES ('group', 1, 'SELECT now()');

-- 在远程数据库上执行 SQL
INSERT INTO task (input, remote) VALUES ('SELECT now()', 'user=user host=host');
```


## 任务表列定义

| 列名 | 类型 | 默认值 | 说明 |
|------|------|---------|-------------|
| id | bigserial | 自增 | 主键 |
| parent | bigint | pg_task.id | 父任务 ID |
| plan | timestamptz | statement_timestamp() | 计划执行时间 |
| start | timestamptz | | 实际开始时间 |
| stop | timestamptz | | 实际结束时间 |
| active | interval | 1 hour | 计划时间后任务保持活跃的时段 |
| live | interval | 0 sec | 后台工作进程最大存活时间 |
| repeat | interval | 0 sec | 自动重复间隔 |
| timeout | interval | 0 sec | 任务执行允许的最大时间 |
| count | int | 0 | 工作进程退出前的最大任务数 |
| max | int | 0 | 同一组内最大并发任务数 |
| pid | int | | 执行任务的进程 ID |
| state | enum | PLAN | PLAN、TAKE、WORK、DONE、STOP |
| delete | bool | true | 当 output 和 error 为空时自动删除 |
| drift | bool | false | 根据结束时间计算下次重复时间 |
| header | bool | true | 在输出中显示列标题 |
| group | text | 'group' | 任务分组名称 |
| input | text | | 要执行的 SQL 命令 |
| output | text | | 接收到的结果 |
| error | text | | 捕获的错误 |
| remote | text | | 远程数据库连接字符串 |

你可以在此表上添加任意所需的列和/或创建分区。


## 配置参数（GUC）

关键设置：

| 参数 | 类型 | 默认值 | 说明 |
|------|------|---------|-------------|
| pg_task.data | text | postgres | 任务表所在数据库名 |
| pg_task.user | text | postgres | 任务表所用用户名 |
| pg_task.schema | text | public | 任务表所在模式名 |
| pg_task.table | text | task | 任务表名 |
| pg_task.sleep | int | 1000 | 每 N 毫秒检查一次任务 |
| pg_task.delete | bool | true | 自动删除已完成任务 |
| pg_task.drift | bool | false | 根据结束时间计算重复间隔 |
| pg_task.repeat | interval | 0 sec | 默认重复间隔 |
| pg_task.timeout | interval | 0 sec | 默认任务超时 |
| pg_task.max | int | 0 | 默认组内最大并发任务数 |
| pg_task.run | int | 2147483647 | 最大并发执行任务数 |
| pg_task.json | json | [{"data":"postgres"}] | 多数据库配置 |

### 多数据库配置

通过 JSON 配置在多个数据库上运行任务：

```conf
pg_task.json = '[{"data":"database1"},{"data":"database2","user":"username2"},{"data":"database3","schema":"schema3"}]'
```

如果指定的数据库、用户、模式或表不存在，`pg_task` 会自动创建它们。
