## 用法

来源：

- [官方扩展控制文件 (pg_json_validate.control)](https://github.com/jefbarn/pg_json_validate/blob/9804a7931c8f9b0a0dcb15c99d6a7d61488d13f5/pg_json_validate.control)
- [官方实现源代码](https://github.com/jefbarn/pg_json_validate/blob/9804a7931c8f9b0a0dcb15c99d6a7d61488d13f5/src/lib.rs)
- [官方 Rust 包清单](https://github.com/jefbarn/pg_json_validate/blob/9804a7931c8f9b0a0dcb15c99d6a7d61488d13f5/Cargo.toml)

`pg_json_validate` — JSON Schema 验证函数，用于 jsonb 值。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_json_validate;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `json_schema_is_valid` 是一个扩展函数。
- `json_schema_validate` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
