## 用法

来源：

- [官方上游 README](https://github.com/evoludigit/pggit/blob/6d1feef9b9fa215802b9243640e8183586dcdf6a/README.md)
- [官方扩展控制文件 (pggit.control)](https://github.com/evoludigit/pggit/blob/6d1feef9b9fa215802b9243640e8183586dcdf6a/pggit.control)
- [官方扩展 SQL (pggit--0.1.3.sql)](https://github.com/evoludigit/pggit/blob/6d1feef9b9fa215802b9243640e8183586dcdf6a/pggit--0.1.3.sql)

`pggit` — **PostgreSQL 中的 Git 类版本控制。** 分支、合并、比较和撤销数据库模式就像处理代码一样。在管理或自动化上述数据库行为时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pggit;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `pggit.abort_merge(p_merge_id uuid, p_reason text DEFAULT 'User aborted')` 是一个扩展函数，返回 `void`。
- `pggit.add_dependency(p_dependent_name TEXT, p_depends_on_name TEXT, p_dependency_type TEXT DEFAULT 'generic')` 是一个扩展函数，返回 `VOID`。
- `pggit.analyze_branch_for_pruning(p_branch_name TEXT)` 是一个扩展函数，返回 `TABLE`。
- `pggit.analyze_cqrs_dependencies(command_schema text DEFAULT 'command', query_schema text DEFAULT 'query')` 是一个扩展函数，返回 `TABLE`。
- `pggit.analyze_migration_intent(p_migration_content TEXT)` 是一个扩展函数，返回 `TABLE`。
- `pggit.analyze_migration_size_impact(p_migration_content TEXT)` 是一个扩展函数，返回 `TABLE`。
- `pggit.analyze_migration_with_ai(p_migration_id TEXT, p_migration_content TEXT, p_source_tool TEXT DEFAULT 'unknown')` 是一个扩展函数，返回 `TABLE`。
- `pggit.analyze_migration_with_ai_enhanced(p_migration_id TEXT, p_migration_content TEXT, p_source_tool TEXT DEFAULT 'unknown')` 是一个扩展函数，返回 `TABLE`。
- `pggit.analyze_query_performance()` 是一个扩展函数，返回 `jsonb`。
- `pggit.analyze_schema_change_frequency(p_branch_name text, p_days integer DEFAULT 30)` 是一个扩展函数，返回 `jsonb`。
- `pggit.apply_data_merge(p_merge_id UUID, p_source_branch TEXT, p_target_branch TEXT, p_resolution_strategy TEXT)` 是一个扩展函数，返回 `VOID`。
- `pggit.apply_migration(p_version TEXT)` 是一个扩展函数，返回 `VOID`。
- `pggit.apply_pruning_recommendation(p_recommendation_id INTEGER)` 是一个扩展函数，返回 `TEXT`。
- `pggit.apply_retention_policy(p_policy JSONB DEFAULT '{"full_days": 30, "incremental_days": 7}')` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.3`。
- 先安装并验证确认的扩展依赖项：`pgcrypto`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件标记该扩展为不受信任。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
