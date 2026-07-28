## 用法

来源：

- [官方上游 README](https://github.com/erayack/pg-wal-budget/blob/945f29eea53a91fb297ce630e77f7512387f1e24/README.md)
- [官方扩展控制文件 (pg_wal_budget.control)](https://github.com/erayack/pg-wal-budget/blob/945f29eea53a91fb297ce630e77f7512387f1e24/pg_wal_budget.control)
- [官方扩展 SQL (pg_wal_budget--0.2.1--0.3.0.sql)](https://github.com/erayack/pg-wal-budget/blob/945f29eea53a91fb297ce630e77f7512387f1e24/sql/pg_wal_budget--0.2.1--0.3.0.sql)

`pg_wal_budget` — 一个使用 Rust 和 pgrx 实现的 PostgreSQL 17 扩展，用于观察、预测并可选地按策略范围强制执行 WAL 生成预算。当需要管理或自动化上述数据库行为时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_wal_budget;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pwb.clear_tenant()` 是一个扩展函数，返回 `void`。
- `pwb.counters()` 是一个扩展函数，返回 `table`。
- `pwb.create_policy(scope_kind text, scope_value text, wal_rate_bytes_per_sec bigint, wal_burst_bytes bigint, mode text default 'observe', priority integer default 100)` 是一个扩展函数，返回 `integer`。
- `pwb.disable_policy(policy_id integer)` 是一个扩展函数，返回 `void`。
- `pwb.flush_profiles()` 是一个扩展函数，返回 `void`。
- `pwb.policies()` 是一个扩展函数，返回 `setof`。
- `pwb.preload_status()` 是一个扩展函数，返回 `text`。
- `pwb.query_profiles()` 是一个扩展函数，返回 `table`。
- `pwb.recent_decisions(decision_limit integer default 100)` 是一个扩展函数，返回 `table`。
- `pwb.reset_profiles()` 是一个扩展函数，返回 `void`。
- `pwb.reset_stats()` 是一个扩展函数，返回 `void`。
- `pwb.scope_names()` 是一个扩展函数，返回 `table`。
- `pwb.scope_stats()` 是一个扩展函数，返回 `table`。
- `pwb.set_policy_mode(policy_id integer, mode text)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.3.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
