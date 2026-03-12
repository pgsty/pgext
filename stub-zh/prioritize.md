

## 用法

> [prioritize: 获取和设置 PostgreSQL 后端进程的优先级](https://github.com/schmiddy/pg_prioritize)

`prioritize` 扩展为 PostgreSQL 后端进程暴露了 `getpriority()` 和 `setpriority()` 系统调用，允许你通过 SQL 对后端进程执行 `renice` 操作。

### 获取后端优先级

```sql
SELECT get_backend_priority(pg_backend_pid());
```

任何用户都可以查询任意后端的优先级。

### 设置后端优先级

```sql
SELECT set_backend_priority(pg_backend_pid(), 10);
```

超级用户可以设置任意后端的优先级。非特权用户只能调整同一角色的后端。

注意：优先级只能提高（更大的数值 = 更低的 OS 优先级）。只有 root 才能降低数值优先级，而 PostgreSQL 进程不应以 root 身份运行。

### 批量操作

```sql
-- 将当前用户所有后端的优先级提高 5
SELECT set_backend_priority(pid, get_backend_priority(pid) + 5)
  FROM pg_stat_activity
  WHERE usename = CURRENT_USER;
```
