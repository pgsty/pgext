## 用法

来源：

- [官方上游 README](https://github.com/aadisanghani/pg_query_rules/blob/7c65bcd58c122cbf5b7b2b130d4cc391432c7ca4/README.md)
- [官方扩展控制文件 (pg_query_rules.control)](https://github.com/aadisanghani/pg_query_rules/blob/7c65bcd58c122cbf5b7b2b130d4cc391432c7ca4/pg_query_rules.control)
- [官方扩展 SQL (pg_query_rules--0.1.0.sql)](https://github.com/aadisanghani/pg_query_rules/blob/7c65bcd58c122cbf5b7b2b130d4cc391432c7ca4/pg_query_rules--0.1.0.sql)

`pg_query_rules` — pg_query_rules 是一个基于正则表达式的 PostgreSQL C 扩展，可以在运行时重写或阻止 SQL 查询。当需要管理或自动化上述数据库行为时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_query_rules;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `load_query_rules_to_runtime()` 是一个扩展函数，返回 `TEXT`。
- `pg_query_rules_updated_at()` 是一个扩展函数，返回 `TRIGGER`。
- `pgqr_test(sql_text TEXT)` 是一个扩展函数，返回 `TEXT`。
- `pgqr_version()` 是一个扩展函数，返回 `TEXT`。
- `pg_query_rules` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
