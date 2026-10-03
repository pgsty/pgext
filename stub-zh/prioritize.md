## 用法

来源：

- [Official README](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/README.md)
- [SQL functions](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/prioritize.sql.in)
- [Control file](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/prioritize.control)

`prioritize` 提供 PostgreSQL 后端进程的操作系统优先级控制。它适合降低指定会话的调度优先级，不是 PostgreSQL 查询调度器。

### 查看和调整后端

```sql
CREATE EXTENSION prioritize;
SELECT get_backend_priority(pg_backend_pid());
SELECT set_backend_priority(pg_backend_pid(), 10);
```

任何用户都可以查询后端优先级。PostgreSQL 超级用户可以请求调整任意后端；其他用户只能调整使用同一数据库角色的后端。

### 调整同角色会话

```sql
SELECT set_backend_priority(pid, get_backend_priority(pid) + 5)
FROM pg_stat_activity
WHERE usename = CURRENT_USER;
```

增大 nice 数值会降低操作系统调度优先级。因此，这个例子降低这些后端的优先级，不会使它们执行得更快。

### 权限和限制

安装需要超级用户，无需预加载或重启。操作系统权限仍然生效：普通 PostgreSQL 进程通常不能通过降低 nice 数值来提高优先级。数据库超级用户权限不等于 root 权限。批量调整前，应检查平台调度策略和进程身份。控制文件使用 SQL 版本 1.0；发行版安装包版本 1.0.4 与之不同。
