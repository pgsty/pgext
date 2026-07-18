## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/postgrespro/pg_grab_statement/blob/31273939223f2cf55c0eb5d28a30ba229c20d501/README.md)
- [1.0 版本 SQL 对象](https://github.com/postgrespro/pg_grab_statement/blob/31273939223f2cf55c0eb5d28a30ba229c20d501/pg_grab_statement--1.0.sql)
- [扩展 control 文件](https://github.com/postgrespro/pg_grab_statement/blob/31273939223f2cf55c0eb5d28a30ba229c20d501/pg_grab_statement.control)

`pg_grab_statement` 记录成功提交事务的详细语句。它挂接执行器开始与结束钩子，并直接写入非日志表 `grab.statement_log`；视图 `grab.statements` 解析用户与查询类型信息。

```sql
CREATE EXTENSION pg_grab_statement;
LOAD 'pg_grab_statement';
SELECT transaction, query_number, username, query_source
FROM grab.statements
ORDER BY transaction, query_number;
```

单会话可使用 `LOAD`；全实例捕获需要 `shared_preload_libraries`，角色或数据库默认值则可使用 `session_preload_libraries`。上游在纯 SELECT 基准中测得约 10–15% 开销，因此必须对真实负载进行压测。

记录包含查询文本、参数值、参数类型、用户 ID 与时间信息，可能收集密码和个人数据，必须限制访问与保留周期。该表是 unlogged 表，崩溃后可能被截断，因此不能作为审计账本。项目使用老旧 PostgreSQL 内部接口且没有当前兼容矩阵；应在精确服务器版本上验证钩子、预备语句、失败事务、崩溃行为与存储增长。
