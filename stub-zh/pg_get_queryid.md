## 用法

来源：

- [官方上游文档](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/README.md)
- [官方扩展 SQL](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/pg_get_queryid--1.0.sql)
- [官方实现](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/pg_get_queryid.c)

`pg_get_queryid` 按 PostgreSQL 后端 PID 返回 `pg_stat_statements` 最近生成或使用的查询标识符。它用于在活动视图还不暴露查询 ID 的 PostgreSQL 版本上关联 `pg_stat_activity` 与规范化语句。

### 预加载与安装

在 `pg_get_queryid` 之前加载 `pg_stat_statements`，重启 PostgreSQL，并在数据库中创建两个扩展：

```conf
shared_preload_libraries = 'pg_stat_statements,pg_get_queryid'
pg_get_queryid.track_utility = on
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_get_queryid;
```

`shared_preload_libraries` 中的顺序是上游要求的一部分。修改任一预加载项都需要重启服务器。

### 关联活动

`pg_get_queryid(integer)` 返回 `bigint`；该 PID 没有可用查询 ID 时返回 `0`：

```sql
SELECT a.pid,
       a.state,
       a.query AS active_query,
       pg_get_queryid(a.pid) AS queryid,
       s.query AS normalized_query
FROM pg_stat_activity AS a
LEFT JOIN pg_stat_statements AS s
  ON s.queryid = pg_get_queryid(a.pid)
WHERE a.query_start IS NOT NULL
  AND a.pid <> pg_backend_pid();
```

`pg_get_queryid.track_utility` 是默认启用、只能由超级用户设置的布尔参数，用于为工具语句生成标识符。预备执行会归入底层预备语句，嵌套跟踪遵循 `pg_stat_statements.track` 策略。

### 时序与兼容性

语句首次执行结束前不会出现在 `pg_stat_statements` 中；条目被淘汰或重置后，也要等下次执行结束才会重新出现。在查询速率极高时，`pg_stat_activity.query` 可能在扩展钩子更新所存标识符之前发生变化，因此连接结果可能瞬时配对到不同语句。应遵守正常的 `pg_stat_activity` 可见性规则，也不能把查询 ID 当作安全边界。

上游面向 PostgreSQL 9.6 及更高版本，但钩子代码使用较早的服务器 API。现代 PostgreSQL 已暴露 `pg_stat_activity.query_id`；在可用时应优先使用核心字段，并在构建前确认是否仍需要该扩展及其 ABI 是否兼容。
