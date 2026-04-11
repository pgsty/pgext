
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION IF NOT EXISTS dblink;
> CREATE EXTENSION IF NOT EXISTS pg_variables;
> CREATE EXTENSION pgelog;
> SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');
> ```
>
> 来源：[README](https://github.com/anfiau/pgelog)

`pgelog` 通过 `dblink` 实现伪自治事务，把日志写入 PostgreSQL 表中。其核心目标是即使调用方的主事务回滚，日志记录也能保留下来。

## 前置条件

README 要求：

- PostgreSQL 11 或更高版本
- `dblink`
- `pg_variables`
- 本地免密码的 `dblink` 访问，通常通过 `pg_hba.conf` 中的 `peer` 本地规则实现

它还提醒，每个会话可能会为 `dblink` 额外打开 1 条连接，因此需要合理设置 `max_connections`。

## 对象

扩展会创建：

- `pgelog_params`，用于配置
- `pgelog_logs`，作为基础日志表
- `pgelog_vw_logs`，作为带时间信息的日志视图

日志表/视图会存储时间戳、日志类型、源函数、阶段、消息文本、事务 ID、SQLSTATE、SQLERRM 和连接名等字段。

## 基本记录

写入一条日志：

```sql
SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');
```

读取最新日志：

```sql
SELECT log_stamp, log_info
FROM pgelog_logs
ORDER BY log_stamp DESC
LIMIT 1;
```

## PL/pgSQL 异常日志

README 给出了一个更完整的 PL/pgSQL 示例：捕获异常、收集诊断信息、通过 `pgelog_to_log(...)` 写入 `FAIL` 日志，然后重新抛出异常。这是捕获可回滚失败日志的主要模式。

## 配置

配置参数通过以下函数管理：

```sql
SELECT pgelog_get_param('pgelog_ttl_minutes');
SELECT pgelog_set_param('pgelog_ttl_minutes', '2880');
```

README 通过 `pgelog_params` 表和辅助函数来说明 `pgelog_ttl_minutes` 等参数。
