## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_role_fkey_trigger_functions/pg_role_fkey_trigger_functions-1.0.4/README.md)
- [官方扩展控制文件 (pg_role_fkey_trigger_functions.control)](https://api.pgxn.org/src/pg_role_fkey_trigger_functions/pg_role_fkey_trigger_functions-1.0.4/pg_role_fkey_trigger_functions.control)
- [官方扩展 SQL (pg_role_fkey_trigger_functions--0.11.7.sql)](https://api.pgxn.org/src/pg_role_fkey_trigger_functions/pg_role_fkey_trigger_functions-1.0.4/sql/pg_role_fkey_trigger_functions--0.11.7.sql)

`pg_role_fkey_trigger_functions` PostgreSQL 扩展提供了一组触发器函数，用于帮助建立或维护引用 PostgreSQL 角色名称之列的参照完整性。它适用于相应的安全、审计或访问控制工作流。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_role_fkey_trigger_functions;

-- Using the `SET` command:
set pg_role_fkey_trigger_functions.trusted_tables TO '{pg_temp.evil_temp_tbl}';

-- Using the `set_config()` function:
select set_config(
    'pg_role_fkey_trigger_functions.trusted_tables',
    '{pg_temp.evil_temp_tbl}',
    false
);

-- Or, appending to instead of replacing the list of trusted tables:
select set_config(
    'pg_role_fkey_trigger_functions.trusted_tables',
    coalesce(
        current_setting('pg_role_fkey_trigger_functions.trusted_tables', true),
        '{}'
    )::text[] || 'pg_temp.evil_temp_tbl',
    false
);
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 重要对象

- `enforce_fkey_to_db_role()` 是扩展函数，返回 `trigger`。
- `grant_role_in_column1_to_role_in_column2()` 是扩展函数，返回 `trigger`。
- `maintain_referenced_role()` 是扩展函数，返回 `trigger`。
- `pg_role_fkey_trigger_functions_meta_pgxn()` 是扩展函数，返回 `jsonb`。
- `pg_role_fkey_trigger_functions_readme()` 是扩展函数，返回 `text`。
- `revoke_role_in_column1_from_role_in_column2()` 是扩展函数，返回 `trigger`。
- `test__pg_role_fkey_trigger_functions` 是扩展提供的过程。
- `test_dump_restore__maintain_referenced_role` 是扩展提供的过程。
- `test__customer` 是扩展安装或管理的表。
- `test__tbl` 是扩展安装或管理的表。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `1.0.4`。
- 控制文件将该扩展标记为可重定位。
- 控制文件不要求仅由超级用户安装。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
