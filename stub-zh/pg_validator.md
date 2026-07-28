## 用法

来源：

- [官方上游 README](https://github.com/max-norin/pg_validator/blob/87931c07eb80fbf41ec999f81b745e320343acbe/README.md)
- [官方扩展控制文件 (pg_validator.control)](https://github.com/max-norin/pg_validator/blob/87931c07eb80fbf41ec999f81b745e320343acbe/dist/pg_validator.control)
- [官方扩展 SQL (pg_validator--1.0.sql)](https://github.com/max-norin/pg_validator/blob/87931c07eb80fbf41ec999f81b745e320343acbe/dist/pg_validator--1.0.sql)

`pg_validator` — 这是一个用于通过触发器验证数据的 PostgreSQL 扩展。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_validator;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `alpha_rule("value" ANYELEMENT)` 是一个扩展函数，返回 `BOOLEAN`。
- `array_is_unique("arr" ANYARRAY)` 是一个扩展函数，返回 `BOOLEAN`。
- `array_overlap_count("a" ANYARRAY, "b" ANYARRAY)` 是一个扩展函数，返回 `INT`。
- `array_unique("arr" ANYARRAY)` 是一个扩展函数，返回 `ANYARRAY`。
- `constraint_def_contained("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` 是一个扩展函数，返回 `BOOLEAN`。
- `constraint_def_contains("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` 是一个扩展函数，返回 `BOOLEAN`。
- `constraint_def_eq("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` 是一个扩展函数，返回 `BOOLEAN`。
- `constraint_def_neq("a" CONSTRAINT_DEF, "b" CONSTRAINT_DEF)` 是一个扩展函数，返回 `BOOLEAN`。
- `email_rule("value" ANYELEMENT)` 是一个扩展函数，返回 `BOOLEAN`。
- `exists_rule("relid" REGCLASS, "table_columns" TEXT[], "record" JSONB, "record_columns" TEXT[], "mode" FK_MODE = 'full', "where" TEXT = 'TRUE')` 是一个扩展函数，返回 `BOOLEAN`。
- `is_distinct_from("a" ANYELEMENT, "b" ANYELEMENT)` 是一个扩展函数，返回 `BOOLEAN`。
- `is_not_distinct_from("a" ANYELEMENT, "b" ANYELEMENT)` 是一个扩展函数，返回 `BOOLEAN`。
- `jsonb_array_append("json" JSONB, "path" TEXT[], "value" JSONB)` 是一个扩展函数，返回 `JSONB`。
- `jsonb_except("a" JSONB, "b" JSONB)` 是一个扩展函数，返回 `JSONB`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
