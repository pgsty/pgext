## 用法

来源：

- [Official database.dev 包页面](https://database.dev/mansueli/function_vc)

`mansueli-function_vc` — 在 PostgreSQL 函数版本控制中使用此扩展。在管理或自动化上述数据库行为时，请使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "mansueli-function_vc";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `archive.save_function_history(function_name text, args text, return_type text, source_code text, schema_name text default 'public', lang_settings text default 'plpgsql')` 是一个扩展函数，返回 `void`。
- `archive.setup_function_history(schema_name text default 'public')` 是一个扩展函数，返回 `VOID`。
- `calculate_version()` 是一个扩展函数，返回 `TRIGGER`。
- `public.create_function_from_source(function_text text, schema_name text default 'public')` 是一个扩展函数，返回 `text`。
- `rollback_function(func_name text, version_no integer default 0, schema_n text default 'public')` 是一个扩展函数，返回 `text`。
- `archive.function_history` 是一个由扩展安装或管理的表。
- `archive` 是一个由扩展创建的模式。
- `before_insert_function_history` 是一个扩展定义的触发器。

### 要求与注意事项

- 该目录记录版本 `1.0.1`。
- 这是一个 database.dev/pg_tle 包；在发出引用的 `CREATE EXTENSION` 之前，请先注册或生成其包迁移。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
